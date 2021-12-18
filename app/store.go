package app

import (
	"context"

	"github.com/ademuanthony/dfctipper/postgres/models"
)

type Store interface {
	CreatUser(ctx context.Context, account *models.Account) error
	GetUser(ctx context.Context, id string) (*models.Account, error)
	UserByTelegramID(ctx context.Context, telegramID int64) (*models.Account, error)
	UserByTwitterID(ctx context.Context, twitterID int64) (*models.Account, error)
	SetTwitterID(ctx context.Context, accID string, twitterID int64) error
	SetWalletAddress(ctx context.Context, telegramId int64, wallet string) error
	SetBalance(ctx context.Context, userID string, balance int64) error

	SetCurrentStep(ctx context.Context, telegramID int64, step int) error
	CurrentStep(ctx context.Context, telegramID int64) (int, error)

	CreatePromotion(ctx context.Context, tweet string, number int, amount int, userID string) error
	AvailablePromotion(ctx context.Context, accID string) (models.PromotionSlice, error)
	UncompletedPromotion(ctx context.Context) (models.PromotionSlice, error)
	SetRewardCount(ctx context.Context, promotionID, count int) error
	SaveReward(ctx context.Context, promotionID int, userID string, reward int64) error

	GetPendingWithdrawal(ctx context.Context, accID string) (*models.Withdrawal, error)
	MakeWithdrawalRequest(ctx context.Context, accID string, amount int64) error
}
