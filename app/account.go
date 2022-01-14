package app

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	tb "gopkg.in/tucnak/telebot.v2"
)

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

	if otherUser, err := a.db.UserByTwitterID(ctx, twitterAcc[0].ID); err == nil {
		if otherUser.ID != acc.ID {
			message := "Another account is using this twitter handle"
			if _, err := a.b.Send(m.Sender, message); err != nil {
				log.Error("a.b.Send", err)
			}
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

	if err = a.db.ActivateByTelegramID(ctx, acc.TelegramID); err != nil {
		log.Error("a.db.Activate", err)
	}

	if err = a.db.IncreaseDownlines(ctx, acc.ReferralID); err != nil {
		log.Error("a.db.IncreaseDownlines", err)
	}

	a.sendMainMenu(m)
}

func (a app) ensureAccount(m *tb.Message) bool {
	ctx := context.Background()
	user, err := a.currentUser(ctx, m)
	if err != nil {
		a.startHandler(m)
		return false
	}

	if user.TwitterID == 0 {
		a.askforTwitter(m)
		return false
	}

	return true
}

func (a app) accountBalance(m *tb.Message) {
	if !a.ensureAccount(m) {
		return
	}
	ctx := context.Background()
	user, err := a.currentUser(ctx, m)
	if err != nil {
		log.Error("accountBalance->currentUser", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	message := fmt.Sprintf(`
	⚖️ Balance
	💰 Your Balance: %d DFC
	
	🏃‍♂️ Invite your friends and get 100%% of all their earnings`, user.Balance)

	if _, err := a.b.Send(m.Sender, message, backToMyAccountMenu); err != nil {
		fmt.Println(err)
	}
}

func (a app) viewWallet(m *tb.Message) {
	if !a.ensureAccount(m) {
		return
	}
	ctx := context.Background()
	user, err := a.currentUser(ctx, m)

	if err != nil {
		log.Error("viewWallet->currentUser", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	address := user.WalletAddress
	if address == "" {
		address = "empty"
	}
	if _, err := a.b.Send(m.Sender, fmt.Sprintf("Your wallet address is %s", address)); err != nil {
		log.Error("viewWallet", err)
	}

	if err = a.db.SetCurrentStep(ctx, m.Sender.ID, SETWALLET); err != nil {
		log.Error("viewWallet->SetCurrentStep", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if _, err := a.b.Send(m.Sender, "Send a BEP20 DFC address to change your wallet address", backToMyAccountMenu); err != nil {
		log.Error("viewWallet", err)
	}
}

func (a app) setWalletMsg(m *tb.Message) {
	if !a.ensureAccount(m) {
		return
	}
	if !strings.HasPrefix(strings.ToLower(m.Text), "0x") {
		if _, err := a.b.Send(m.Sender, "Please send a valid DFC wallet"); err != nil {
			a.sendSystemErrorMsg(m, err)
			return
		}
	}

	ctx := context.Background()
	if err := a.db.SetWalletAddress(ctx, m.Sender.ID, m.Text); err != nil {
		a.sendSystemErrorMsg(m, err)
		return
	}
	if _, err := a.b.Send(m.Sender, `✅ Wallet updated successfully!`, backToMyAccountMenu); err != nil {
		log.Error(err)
		return
	}

	if err := a.db.SetCurrentStep(ctx, m.Sender.ID, NoStep); err != nil {
		log.Error("viewWallet->SetCurrentStep", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	a.mySettingMenu(m)
}

func (a app) withdrawal(m *tb.Message) {
	if !a.ensureAccount(m) {
		return
	}
	ctx := context.Background()
	user, err := a.currentUser(ctx, m)
	if err != nil {
		log.Error("makeWithdrawal->currentUser", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if user.WalletAddress == "" {
		message := "Cannot proccess withdrwal"
		if _, err := a.b.Send(&tb.User{ID: user.TelegramID}, message); err != nil {
			log.Error(err)
		}
		a.viewWallet(m)
		return
	}

	pendingWidrawal, err := a.db.GetPendingWithdrawal(ctx, user.ID)
	if err != nil && err != sql.ErrNoRows {
		log.Error("makeWithdrawal->GetPendingWithdrawal", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if pendingWidrawal != nil {
		message := fmt.Sprintf("You have a pending withdrawal request of %d DFC", pendingWidrawal.Amount)
		if _, err := a.b.Send(m.Sender, message, backToMyAccountMenu); err != nil {
			log.Error("makeWithdrawal->Send", err)
		}
		return
	}

	if err := a.db.SetCurrentStep(ctx, m.Sender.ID, MAKEWITHDRAW); err != nil {
		log.Error("makeWithdrawal->SetCurrentStep", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if _, err = a.b.Send(m.Sender, "How much do you want to withdraw", backToMyAccountMenu); err != nil {
		log.Error("makeWithdrawal->Send", err)
		return
	}
}

func (a app) makeWithdrawal(m *tb.Message) {
	if !a.ensureAccount(m) {
		return
	}
	amount, _ := strconv.Atoi(m.Text)
	if amount <= 0 {
		msg := "Invalid amount. Amount must be a positive number with a period(.)"
		if _, err := a.b.Send(m.Sender, msg, backToMyAccountMenu); err != nil {
			log.Error("makeWithdrawal")
		}
	}

	ctx := context.Background()
	user, err := a.currentUser(ctx, m)
	if err != nil {
		log.Error("makeWithdrawal->currentUser", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if amount < MINIMUMWITHDRAWAL {
		msg := fmt.Sprintf("Invalid amount. Amount must be greater than %d", MINIMUMWITHDRAWAL)
		if _, err := a.b.Send(m.Sender, msg, backToMyAccountMenu); err != nil {
			log.Error("makeWithdrawal")
		}
		return
	}

	if user.Balance < int64(amount) {
		if _, err := a.b.Send(m.Sender, "Insufficient Balance", backToMyAccountMenu); err != nil {
			log.Error("makeWithdrawal")
		}
		return
	}

	if err := a.db.SetBalance(ctx, user.ID, user.Balance-int64(amount)); err != nil {
		log.Error("SetBalance", err)
		return
	}

	if err := a.db.MakeWithdrawalRequest(ctx, user.ID, int64(amount)); err != nil {
		log.Error("makeWithdrawal->currentUser", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if err := a.db.SetCurrentStep(ctx, m.Sender.ID, NoStep); err != nil {
		log.Error("makeWithdrawal->SetCurrentStep", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if _, err = a.b.Send(m.Sender, "Your withdrawal request has been placed. Thank you"); err != nil {
		log.Error("makeWithdrawal->send", err)
		return
	}
	a.myAccountMenu(m)
}

func (a app) referralLink(m *tb.Message) {
	if !a.ensureAccount(m) {
		return
	}
	ctx := context.Background()
	user, err := a.currentUser(ctx, m)
	if err != nil {
		log.Error("referralLink->currentUser", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	message := fmt.Sprintf(`🎁 DFC give away on Telegram 🎁
	🎈 Earn billions of DFC performing simple task 🎈
	
	🆓 Get free DFC tokens in bot
	
	Claim Now👇
	 http://t.me/dfctippingbot?start=%d`, user.TelegramID)

	if _, err := a.b.Send(m.Sender, message); err != nil {
		log.Error("referralLink->send", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	message = `♻️ Share this link to invite your friends to earn more DFC for FREE
	📛 Note: Fake referrals won't get paid
	🎁 Number of referrals: %d
	`

	if _, err := a.b.Send(m.Sender, fmt.Sprintf(message, user.Downlines)); err != nil {
		log.Error("referralLink->Send", err)
		a.sendSystemErrorMsg(m, err)
	}
}

func (a app) processWithdrawals() {
	ctx := context.Background()
	pendingWithdrawal, err := a.db.AllGetPendingWithdrawal(ctx)
	if err != nil {
		log.Error("a.db.AllGetPendingWithdrawal", err)
		return
	}

	for _, with := range pendingWithdrawal {
		log.Info("Proccessing withdrawal ID", with.ID)
		account, err := a.db.GetUser(ctx, with.UserID)
		if err != nil {
			log.Error("FindAccount", err)
			continue
		}

		if account.WalletAddress == "" {
			continue
			// message := "Unable to proccess your withdrawal request. Please set your wallet address from the Account menu"
			// if _, err := a.b.Send(&tb.User{ID: account.TelegramID}, message, myAccountMenu); err != nil {
			// 	log.Error(err)
			// }
		}

		amount := big.NewInt(with.Amount)
		dfcAmount := amount.Mul(amount, big.NewInt(1e8))
		if err != nil {
			log.Errorf("processPaymentQueue->convertClubDollarToBnb %v", err)
			continue
		}

		txHash, err := a.transferDfc(ctx, a.config.MasterAddressKey, account.WalletAddress, dfcAmount)
		if err != nil {
			log.Errorf("processPaymentQueue->m.transfer %v - %v", err, dfcAmount)
			continue
		}

		message := `Hello %s

		Your withdrawal of %d DFC has been processed
		
		https://bscscan.com/tx/%s`

		if _, err := a.b.Send(&tb.User{ID: account.TelegramID}, fmt.Sprintf(message, account.FirstName, with.Amount, txHash)); err != nil {
			log.Error("a.b.Send", err)
		}

		if err := a.db.UpdateTxHash(ctx, with.ID, txHash); err != nil {
			log.Error("a.db.UpdateTxHash", err)
		}

		log.Info("Withdrawal proccessed", txHash)
	}
}
