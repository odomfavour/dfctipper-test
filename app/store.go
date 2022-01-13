package app

import (
	"context"

	"github.com/ademuanthony/dfctipper/postgres/models"
)

type Store interface {
	CreatUser(ctx context.Context, account *models.Account) error
	ActivateByTelegramID(ctx context.Context, accID int64) error
	DeactivateByTelegramID(ctx context.Context, accID int64) error
	GetUser(ctx context.Context, id string) (*models.Account, error)
	UserByTelegramID(ctx context.Context, telegramID int64) (*models.Account, error)
	UserByTwitterID(ctx context.Context, twitterID int64) (*models.Account, error)
	ActiveUsersTelegram(ctx context.Context) (models.AccountSlice, error)
	SetTwitterID(ctx context.Context, accID string, twitterID int64) error
	SetWalletAddress(ctx context.Context, telegramId int64, wallet string) error
	SetBalance(ctx context.Context, userID string, balance int64) error
	IncreaseDownlines(ctx context.Context, accID string) error

	SetCurrentStep(ctx context.Context, telegramID int64, step int) error
	CurrentStep(ctx context.Context, telegramID int64) (int, error)

	CreatePromotion(ctx context.Context, tweet string, number int, amount int, userID string) error
	AvailablePromotion(ctx context.Context, accID string) (models.PromotionSlice, error)
	UncompletedPromotion(ctx context.Context) (models.PromotionSlice, error)
	SetRetweetCount(ctx context.Context, promotionID, count int) error
	CanEarn(ctx context.Context, promotionID int, userID string) (bool, error)
	SaveReward(ctx context.Context, promotionID int, userID string, reward int64) error

	GetPendingWithdrawal(ctx context.Context, accID string) (*models.Withdrawal, error)
	AllGetPendingWithdrawal(ctx context.Context) (models.WithdrawalSlice, error)
	MakeWithdrawalRequest(ctx context.Context, accID string, amount int64) error
	UpdateTxHash(ctx context.Context, withID int, txHash string) error
}
