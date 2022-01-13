package app

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/ademuanthony/dfctipper/postgres/models"
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
	b             *tb.Bot
	client        *ethclient.Client
	config        BlockchainConfig
}

func Start(ctx context.Context, db Store, twitterClient *twitter.Client,
	client *ethclient.Client, cfg BlockchainConfig, b *tb.Bot) error {

	app := &app{db: db, twitterClient: twitterClient, b: b, client: client, config: cfg}

	buildMenuItems(b)

	b.Handle("/start", app.startHandler)
	b.Handle("/ajah", app.startCreatePromotion)
	b.Handle(tb.OnText, app.wrapHandler(app.textHandler))
	b.Handle(&btnMyAccount, app.wrapHandler(app.myAccountMenu))
	b.Handle(&btnBackToMenu, app.wrapHandler(app.sendMainMenu))
	b.Handle(&btnBackToMyAccount, app.wrapHandler(app.myAccountMenu))
	b.Handle(&btnAccountBalance, app.wrapHandler(app.accountBalance))
	b.Handle(&btnWallet, app.wrapHandler(app.viewWallet))
	b.Handle(&btnWithdraw, app.wrapHandler(app.withdrawal))
	b.Handle(&btnReferralLink, app.wrapHandler(app.referralLink))
	b.Handle(&btnStartEarning, app.wrapHandler(app.viewTweet))

	go func() {
		for {
			app.processReward()
			time.Sleep(30 * time.Second)
		}
	}()

	go func() {
		for {
			app.processWithdrawals()
			time.Sleep(30 * time.Second)
		}
	}()

	return nil
}

func (a app) startHandler(m *tb.Message) {

	ctx := context.Background()

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

type handlerFunc func (*tb.Message)

func (a app) wrapHandler(f handlerFunc) interface{} {
	return func (m *tb.Message)  {
		ctx := context.Background()
		if err := a.db.ActivateByTelegramID(ctx, m.Sender.ID); err != nil {
			log.Error("a.db.ActivateByTelegramID", err)
		}
		f(m)
	}
}