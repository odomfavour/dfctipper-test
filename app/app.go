package app

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/ademuanthony/dfctipper/postgres/models"
	"github.com/ademuanthony/dfctipper/web"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	WELCOME_BONUS     = 50000
	MINIMUMWITHDRAWAL = 200000000
)

type app struct {
	db            Store
	twitterClient *twitter.Client
	server        *web.Server
	b             *tb.Bot
	client        *ethclient.Client
	config        BlockchainConfig

	EnableWeb     bool
	EnableTwitter bool

	MgDomain string
	MgKey    string
}

func Start(ctx context.Context, server *web.Server, db Store, twitterClient *twitter.Client,
	client *ethclient.Client, cfg BlockchainConfig, b *tb.Bot, mgDomain, mgKey string, enableWeb, enableTwitter bool) error {

	app := &app{
		db: db, twitterClient: twitterClient,
		server: server, b: b, client: client, config: cfg,
		MgDomain:      mgDomain,
		MgKey:         mgKey,
		EnableWeb:     enableWeb,
		EnableTwitter: enableTwitter,
	}

	app.initBot()

	if err := app.initWeb(); err != nil {
		return err
	}

	return nil
}

func (a *app) initBot() {
	buildMenuItems(a.b)

	a.b.Handle("/start", a.startHandler)
	a.b.Handle("/ajah", a.startCreatePromotion)
	a.b.Handle(tb.OnText, a.wrapHandler(a.textHandler))
	a.b.Handle(&btnMyAccount, a.wrapHandler(a.myAccountMenu))
	a.b.Handle(&btnSettings, a.wrapHandler(a.mySettingMenu))
	a.b.Handle(&btnBackToMenu, a.wrapHandler(a.sendMainMenu))
	a.b.Handle(&btnBackToMyAccount, a.wrapHandler(a.myAccountMenu))
	a.b.Handle(&btnBackToMySetting, a.wrapHandler(a.mySettingMenu))
	a.b.Handle(&btnAccountBalance, a.wrapHandler(a.accountBalance))
	a.b.Handle(&btnWithdraw, a.wrapHandler(a.withdrawal))
	a.b.Handle(&btnReferralLink, a.wrapHandler(a.referralLink))
	a.b.Handle(&btnStartEarning, a.wrapHandler(a.viewTweet))

	a.b.Handle(&btnWallet, a.wrapHandler(a.viewWallet))
	a.b.Handle(&btnTwitter, a.wrapHandler(a.askforTwitter))

	if a.EnableTwitter {
		go func() {
			for {
				a.processReward()
				time.Sleep(30 * time.Second)
			}
		}()
	}

	go func() {
		for {
			a.processWithdrawals()
			time.Sleep(30 * time.Second)
		}
	}()
}

func (a *app) initWeb() error {
	if err := a.server.Templates.AddTemplate("home"); err != nil {
		return fmt.Errorf("AddTemplate: %v", err)
	}

	if err := a.server.Templates.AddTemplate("advertiser"); err != nil {
		return fmt.Errorf("AddTemplate: %v", err)
	}

	if err := a.server.Templates.AddTemplate("advertiser-thankyou"); err != nil {
		return fmt.Errorf("AddTemplate: %v", err)
	}

	log.Info("adding web routes")
	a.server.AddRoute("/", web.GET, a.homePage)
	a.server.AddRoute("/advertiser", web.GET, a.advertiser)
	a.server.AddRoute("/contactpostback", web.POST, a.contactPostBack)
	a.server.AddRoute("/thankyou", web.GET, a.advertiserThankyou)

	return nil
}

func (a app) startHandler(m *tb.Message) {

	ctx := context.Background()

	if acc, err := a.db.UserByTelegramID(ctx, m.Sender.ID); err == nil {
		if acc.TwitterID > 0 {
			a.sendMainMenu(m)
		} else {
			a.askforTwitter(m)
		}
		return
	}

	refTelegramId, _ := strconv.Atoi(m.Payload)
	refName := "None"
	refId := ""
	referrer, err := a.db.UserByTelegramID(ctx, int64(refTelegramId))
	if err == nil {
		refName = referrer.FirstName
		refId = referrer.ID
	}

	welcomeMessage := fmt.Sprintf(`🎈 Welcome to Defitipper 🎈
	🎁 Welcome Gift: %d DFC
	
	🏃‍♂️ You have been invited by this user (%s)
	
	📛 You will have the oppurtunity to earn money for performing simple task like retweeting.

	✅ Invite your friends to increase your earnings.
	`, WELCOME_BONUS, refName)

	if _, err := a.b.Send(m.Sender, welcomeMessage); err != nil {
		log.Errorf("Cannot send message - %v", err)
		return
	}

	// create an account and then ask for twitter handle
	_, err = a.db.UserByTelegramID(ctx, m.Sender.ID)
	if err == nil {
		a.sendMainMenu(m)
		return
	}
	if err != sql.ErrNoRows {
		log.Error("startHandler->a.db.UserByTelegramID", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	newAcc := &models.Account{
		ID:         uuid.NewString(),
		TelegramID: m.Sender.ID,
		ReferralID: refId,
		Balance:    WELCOME_BONUS,
		FirstName:  m.Sender.FirstName,
		LastName:   m.Sender.LastName,
		Username:   m.Sender.Username,
		JoinAt:     time.Now().UTC().Unix(),
	}

	if err = a.db.CreatUser(ctx, newAcc); err != nil {
		log.Error("a.db.CreatUser", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	a.askforTwitter(m)
}

func (a app) textHandler(m *tb.Message) {
	ctx := context.Background()
	currentStep, err := a.db.CurrentStep(ctx, m.Sender.ID)
	if err != nil {
		log.Error("textHandler", "currentStep", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	switch currentStep {
	case ConnectTwitter:
		a.connectTwitter(m)
	case SETWALLET:
		a.setWalletMsg(m)
	case MAKEWITHDRAW:
		a.makeWithdrawal(m)
	case CREATEPROMOTION:
		a.createPromotion(m)
	default:
		a.b.Send(m.Sender, "Please click on any of the items on the menu for your interaction")
	}
}

type handlerFunc func(*tb.Message)

func (a app) wrapHandler(f handlerFunc) interface{} {
	return func(m *tb.Message) {
		ctx := context.Background()
		if err := a.db.ActivateByTelegramID(ctx, m.Sender.ID); err != nil {
			log.Error("a.db.ActivateByTelegramID", err)
		}
		f(m)
	}
}
