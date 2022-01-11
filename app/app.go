package app

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ademuanthony/dfctipper/postgres/models"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	WELCOME_BONUS     = 50000
	MINIMUMWITHDRAWAL = 500000000
)

type app struct {
	db            Store
	twitterClient *twitter.Client
	b             *tb.Bot
	client        *ethclient.Client
	config        BlockchainConfig
}

func Start(ctx context.Context, db Store, twitterClient *twitter.Client, b *tb.Bot) error {

	app := &app{db: db, twitterClient: twitterClient, b: b}

	buildMenuItems(b)

	b.Handle("/start", app.startHandler)
	b.Handle("/ajah", app.startCreatePromotion)
	b.Handle(tb.OnText, app.textHandler)
	b.Handle(&btnMyAccount, app.myAccountMenu)
	b.Handle(&btnBackToMenu, app.sendMainMenu)
	b.Handle(&btnBackToMyAccount, app.myAccountMenu)
	b.Handle(&btnAccountBalance, app.accountBalance)
	b.Handle(&btnWallet, app.viewWallet)
	b.Handle(&btnWithdraw, app.withdrawal)
	b.Handle(&btnReferralLink, app.referralLink)
	b.Handle(&btnStartEarning, app.viewTweet)

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

func (a app) askforTwitter(m *tb.Message) {

	ctx := context.Background()

	refTelegramId, _ := strconv.Atoi(m.Payload)
	referrer, err := a.db.UserByTelegramID(ctx, int64(refTelegramId))
	if err != nil && err != sql.ErrNoRows {
		log.Error("a.db.UserByTelegramID", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if err = a.db.SetCurrentStep(ctx, m.Sender.ID, ConnectTwitter); err != nil {
		log.Error("a.db.SetCurrentStep", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	message := "Let's connect your twitter account. Please enter your Twitter username"
	if _, err = a.b.Send(m.Sender, message); err != nil {
		log.Error("a.b.Send", err)
		return
	}

	message = `Hello %s
	
	You have a new referral, %s
	
	You will earn 100%% of all his earnings
	
	Invite more people to increase your earnings`

	if referrer != nil {
		if _, err = a.b.Send(&tb.User{ID: referrer.TelegramID}, fmt.Sprintf(message, referrer.FirstName, m.Sender.FirstName)); err != nil {
			log.Error("a.b.Send", err)
			return
		}
	}
}

func (a app) connectTwitter(m *tb.Message) {
	ctx := context.Background()
	acc, err := a.currentUser(ctx, m)
	if err != nil {
		log.Error("connectTwitter", "a.currentUser", err)
		return
	}
	username := strings.Trim(m.Text, "@")
	twitterAcc, _, err := a.twitterClient.Users.Lookup(
		&twitter.UserLookupParams{
			ScreenName: []string{username},
		},
	)

	if err != nil {
		log.Error("connectTwitter", "a.twitterClient.Users.Lookup", err)
		return
	}

	if len(twitterAcc) == 0 {
		message := "Invalid Twitter username. Please try again"
		if _, err := a.b.Send(m.Sender, message); err != nil {
			log.Error("a.b.Send", err)
			return
		}
	}

	if _, err = a.db.UserByTwitterID(ctx, twitterAcc[0].ID); err == nil {
		message := "Another account is using this twitter handle"
		if _, err := a.b.Send(m.Sender, message); err != nil {
			log.Error("a.b.Send", err)
			return
		}
	}

	if err := a.db.SetTwitterID(ctx, acc.ID, twitterAcc[0].ID); err != nil {
		log.Error("SetTwitterID", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if err = a.db.SetCurrentStep(ctx, m.Sender.ID, NoStep); err != nil {
		log.Error("connectTwitter", "setCurrentStep", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	message := "Congratulations! Your Twitter account has been linked successfully"
	if _, err = a.b.Send(m.Sender, message); err != nil {
		log.Error("a.b.Send", err)
	}

	if err = a.db.IncreaseDownlines(ctx, acc.ReferralID); err != nil {
		log.Error("a.db.IncreaseDownlines", err)
	}

	a.sendMainMenu(m)
}
