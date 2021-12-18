package app

import (
	"context"
	"encoding/json"

	"github.com/ademuanthony/dfctipper/postgres/models"
	tb "gopkg.in/tucnak/telebot.v2"
)

func (a app) deleteMessage(messageID int, chatID int64) error {
	var payload = struct {
		ChatID    int64 `json:"chat_id"`
		MessageID int   `json:"message_id"`
	}{
		ChatID:    chatID,
		MessageID: messageID,
	}

	_, err := a.b.Raw("deleteMessage", payload)

	return err
}

func (a app) isAnAdmin(userID int, group string) (bool, error) {
	// return true, nil
	var payload = struct {
		ChatID string `json:"chat_id"`
		UserID int    `json:"user_id"`
	}{
		ChatID: group,
		UserID: userID,
	}

	data, err := a.b.Raw("getChatMember", payload)
	if err != nil {
		log.Error("getChatMember", err)
		return false, err
	}

	var resp struct {
		Result struct {
			Status string `json:"status"`
			User   struct {
				IsBot bool `json:"is_bot"`
			} `json:"user"`
		} `json:"result"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		log.Error("getChatMember", err)
		return false, err
	}
	return resp.Result.Status == "creator" || resp.Result.Status == "administrator", nil
}

func (a app) sendSystemErrorMsg(m *tb.Message, err error) {
	log.Error(err)
	if _, err := a.b.Send(m.Sender, "😧 I am having some internal issues. 🙏 Please try again later!"); err != nil {
		log.Error(err)
	}
}

func (a app) currentUser(ctx context.Context, m *tb.Message) (*models.Account, error) {
	return a.db.UserByTelegramID(ctx, m.Sender.ID)
}
