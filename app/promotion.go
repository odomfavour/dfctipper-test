package app

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/ademuanthony/dfctipper/postgres/models"
	"github.com/dghubble/go-twitter/twitter"
	tb "gopkg.in/tucnak/telebot.v2"
)

func (a app) viewTweet(m *tb.Message) {
	ctx := context.Background()
	acc, err := a.currentUser(ctx, m)
	if err != nil {
		log.Error("a.currentUser", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	promotions, err := a.db.AvailablePromotion(ctx, acc.ID)
	if err != nil {
		log.Error("a.db.AvailablePromotion", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if len(promotions) == 0 {
		if _, err := a.b.Send(m.Sender,
			"No promotion to view at the moment, check back later", backToMyAccountMenu); err != nil {
			log.Error("viewTweet->send", err)
			return
		}
		return
	}

	for _, promotion := range promotions {
		message := fmt.Sprintf(`
		Twitter link: %s

		Possible earning: %d DFC
		`, promotion.TweetLink, int(promotion.RewardPerRetweet*40/100))

		if _, err = a.b.Send(m.Sender, message); err != nil {
			log.Error("a.b.Send", err)
		}
	}

	if _, err := a.b.Send(m.Sender, "Retweet and earn", backToMyAccountMenu); err != nil {
		log.Error("viewTweet->Send", err)
	}
}

func (a app) startCreatePromotion(m *tb.Message) {
	ctx := context.Background()
	if err := a.db.SetCurrentStep(ctx, m.Sender.ID, CREATEPROMOTION); err != nil {
		log.Error("a.startCreatePromotion->SetCurrentStep", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	message := `Send the new promotion in the the following format
	
	LINK_TO_TWEET|NUMBER_OF_RETWEET|AMOUNT_PER_RETWEET
	`
	if _, err := a.b.Send(m.Sender, message, backToMyAccountMenu); err != nil {
		log.Error("a.startCreatePromotion send", err)
		return
	}

	if err := a.db.SetCurrentStep(context.Background(), m.Sender.ID, CREATEPROMOTION); err != nil {
		log.Error("makeWithdrawal->SetCurrentStep", err)
		a.sendSystemErrorMsg(m, err)
		return
	}
}

func (a app) createPromotion(m *tb.Message) {
	component := strings.Split(m.Text, "|")

	invalidPromotionMessage := func() {
		message := "Invalid promotion format"
		if _, err := a.b.Send(m.Sender, message); err != nil {
			log.Error("a.startCreatePromotion send", err)
			return
		}

		message = `Send the new promotion in the the following format
	
	LINK_TO_TWEET|NUMBER_OF_RETWEET|AMOUNT_PER_RETWEET
	`
		if _, err := a.b.Send(m.Sender, message, backToMyAccountMenu); err != nil {
			log.Error("a.startCreatePromotion send", err)
			return
		}
	}

	if len(component) != 3 {
		invalidPromotionMessage()

		return
	}

	tweet := component[0]
	number, err := strconv.Atoi(component[1])
	if err != nil {
		invalidPromotionMessage()
		return
	}

	if number <= 0 {
		invalidPromotionMessage()
		return
	}

	amount, err := strconv.Atoi(component[2])
	if err != nil {
		invalidPromotionMessage()
		return
	}

	if number <= 0 {
		invalidPromotionMessage()
		return
	}

	total := number * amount

	ctx := context.Background()

	user, err := a.currentUser(ctx, m)
	if err != nil {
		log.Error("a.startCreatePromotion currentUser", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if user.Balance < int64(total) {
		if _, err := a.b.Send(m.Sender, "Insufficient Balance"); err != nil {
			log.Error("makeWithdrawal send")
		}
		return
	}

	if err := a.db.SetBalance(ctx, user.ID, user.Balance-int64(total)); err != nil {
		log.Error("a.startCreatePromotion CreatePromotion", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if err := a.db.CreatePromotion(ctx, tweet, number, amount, user.ID); err != nil {
		log.Error("a.startCreatePromotion CreatePromotion", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if err := a.db.SetCurrentStep(ctx, m.Sender.ID, NoStep); err != nil {
		log.Error("makeWithdrawal->SetCurrentStep", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if _, err := a.b.Send(m.Sender, "Promotion created"); err != nil {
		log.Error("createPromotion", err.Error())
	}

	go a.sendNewTweetNotification(tweet, amount)
}

func (a app) sendNewTweetNotification(tweet string, amount int) {
	ctx := context.Background()
	telegramInfo, err := a.db.ActiveUsersTelegram(ctx)
	if err != nil {
		log.Errorf("a.db.AllUserTelegram, %v", err)
		return
	}

	for i, user := range telegramInfo {

		message := fmt.Sprintf(`
		New Tweet Notificcation
		
		Twitter link: %s

		Possible earning: %d DFC
		`, tweet, int(amount*40/100))

		if _, err = a.b.Send(&tb.User{ID: user.TelegramID}, message); err != nil {
			log.Error("a.b.Send", err)
			if strings.Contains(err.Error(), "(401)") {
				if err = a.db.DeactivateByTelegramID(ctx, user.TelegramID); err != nil {
					log.Error("a.db.Activate", err)
				}
			}
		}

		if i%100 == 0 {
			log.Info("notified", i)
		}

	}
}

func (a app) processReward() {
	ctx := context.Background()

	uncompletedPromotions, err := a.db.UncompletedPromotion(ctx)
	if err != nil {
		log.Errorf("processReward->a.db.UncompletedPromotion %v", err)
		return
	}

	for _, promotion := range uncompletedPromotions {
		comps := strings.Split(promotion.TweetLink, "/status/")
		id, _ := strconv.Atoi(comps[1])

		log.Infof("Proceesing promotion %d", id)

		retweets, _, err := a.twitterClient.Statuses.Retweets(int64(id), &twitter.StatusRetweetsParams{
			Count: 100,
		})
		if err != nil {
			log.Error("a.twitterClient.Statuses.Retweets", err)
			continue
		}

		log.Infof("processing %d retweets for %d", len(retweets), id)

		for _, r := range retweets {
			if promotion.RetweetCount >= promotion.RewardCount {
				break
			}
			user, err := a.db.UserByTwitterID(ctx, r.User.ID)
			if err == sql.ErrNoRows {
				continue
			}

			if err != nil {
				log.Error("processReward->UserByTwitterID", err)
				continue
			}

			// no double earning
			if can, _ := a.db.CanEarn(ctx, promotion.ID, user.ID); !can {
				continue
			}

			if err := a.db.IncreaseRetweetCount(ctx, promotion.ID, promotion.RetweetCount+1); err != nil {
				log.Error("processReward->SetRetweetCount", err)
				continue
			}

			reward := int64(promotion.RewardPerRetweet*40) / 100

			if err := a.db.SaveReward(ctx, promotion.ID, user.ID, reward); err != nil {
				log.Error("sendReward->SaveReward", err)
				return
			}

			if err := a.sendReward(ctx, user, reward); err != nil {
				log.Error("processReward->sendReward", err)
				continue
			}

		}
	}
}

func (a app) sendReward(ctx context.Context, user *models.Account, reward int64) error {

	if err := a.db.SetBalance(ctx, user.ID, user.Balance+reward); err != nil {
		return err
	}

	message := fmt.Sprintf(`Hello %s
	
	You have earned %d DFC for retweeting
	
	Your account balance is %d
	
	Invite more friends and earn more`, user.FirstName, reward, user.Balance+reward)

	if _, err := a.b.Send(&tb.User{ID: user.TelegramID}, message); err != nil {
		log.Error("Send", err)
	}

	if user.ReferralID == "" {
		return nil
	}

	referral, err := a.db.GetUser(ctx, user.ReferralID)
	if err != nil {
		return err
	}

	if err := a.db.SetBalance(ctx, referral.ID, referral.Balance+reward); err != nil {
		return err
	}

	message = fmt.Sprintf(`Hello %s
	
	You have earned %d DFC from %s
	
	Your account balance is %d
	
	Invite more friends and earn more`, referral.FirstName, reward, user.FirstName, referral.Balance+reward)

	if _, err := a.b.Send(&tb.User{ID: referral.TelegramID}, message); err != nil {
		log.Error("Send", err)
	}

	return nil
}
