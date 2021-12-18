package app

import (
	"context"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	menu            = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	btnStartEarning = menu.Text("💰 Start Earning")
	btnMyAccount    = menu.Text("👤 My Account")
	btnReferralLink = menu.Text("🔗 Referral Link")

	myAccountMenu       = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	backToMyAccountMenu = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	btnAccountBalance   = myAccountMenu.Text("🏦 Balance")
	btnWallet           = myAccountMenu.Text("💳️ Wallet")
	btnWithdraw         = myAccountMenu.Text("💰 Withdraw")
	btnBackToMyAccount  = myAccountMenu.Text("⬅️ Back to My Account")

	btnBackToMenu = menu.Text("⬅️ Back to Menu")
)

func buildMenuItems(b *tb.Bot) {
	menu.Reply(
		menu.Row(btnStartEarning, btnMyAccount, btnReferralLink),
	)

	myAccountMenu.Reply(
		myAccountMenu.Row(btnAccountBalance, btnWallet, btnWithdraw),
		myAccountMenu.Row(btnBackToMenu),
	)

	backToMyAccountMenu.Reply(
		backToMyAccountMenu.Row(btnBackToMyAccount),
	)
}

func (a app) sendMainMenu(m *tb.Message) {
	if _, err := a.b.Send(m.Sender, "Pick an item from the menu to continue", menu); err != nil {
		log.Error("a.b.Send", err)
	}
}

func (a app) myAccountMenu(m *tb.Message) {
	ctx := context.Background()

	if err := a.db.SetCurrentStep(ctx, m.Sender.ID, NoStep); err != nil {
		log.Error("viewWallet->SetCurrentStep", err)
		a.sendSystemErrorMsg(m, err)
		return
	}

	if _, err := a.b.Send(m.Sender, "Manage your account below", myAccountMenu); err != nil {
		log.Error("a.b.Send", err)
	}
}
