package app

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

func (a app) accountBalance(m *tb.Message) {
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
	ctx := context.Background()
	user, err := a.currentUser(ctx, m)

	if err != nil {
		log.Error("viewWallet->currentUser", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	address := user.WalletAddress
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

	a.myAccountMenu(m)
}

func (a app) withdrawal(m *tb.Message) {
	ctx := context.Background()
	user, err := a.currentUser(ctx, m)
	if err != nil {
		log.Error("makeWithdrawal->currentUser", err)
		a.sendSystemErrorMsg(m, err)
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
	amount, _ := strconv.Atoi(m.Text)
	if amount <= 0 {
		if _, err := a.b.Send(m.Sender, "Invalid amount. Amount must be a positive number with a period(.)"); err != nil {
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
		if _, err := a.b.Send(m.Sender, fmt.Sprintf("Invalid amount. Amount must be greater than %d", MINIMUMWITHDRAWAL)); err != nil {
			log.Error("makeWithdrawal")
		}
	}

	if user.Balance < int64(amount) {
		if _, err := a.b.Send(m.Sender, "Insufficient Balance"); err != nil {
			log.Error("makeWithdrawal")
		}
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
}

func (a app) referralLink(m *tb.Message) {
	ctx := context.Background()
	user, err := a.currentUser(ctx, m)
	if err != nil {
		log.Error("referralLink->currentUser", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	message := fmt.Sprintf(`🎁 Defitipper give away on Telegram 🎁
	🎈 Earn billions of DFC performing simple task 🎈
	
	🆓 Get free DFC tokens in bot
	
	Claim Now👇
	 http://t.me/dfctippingbot?start=%d`, user.TelegramID)

	if _, err := a.b.Send(m.Sender, message); err != nil {
		log.Error("referralLink->send", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	message = `♻️ Send this banner to your friends and get 100% of all their earnings
	📛 Note: Fake Referral Not Pay`

	if _, err := a.b.Send(m.Sender, message); err != nil {
		log.Error("referralLink->Send", err)
		a.sendSystemErrorMsg(m, err)
	}
}
