package app

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

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

	a.myAccountMenu(m)
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

		amount := with.Amount
		dfcAmount, err := a.convertDollarToDfc(ctx, amount)
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

		if _, err := a.b.Send(&tb.User{ID: account.TelegramID}, fmt.Sprintf(message, account.FirstName, amount, txHash)); err != nil {
			log.Error("a.b.Send", err)
		}

		if err := a.db.UpdateTxHash(ctx, with.ID, txHash); err != nil {
			log.Error("a.db.UpdateTxHash", err)
		}

		log.Info("Withdrawal proccessed", txHash)
	}
}
