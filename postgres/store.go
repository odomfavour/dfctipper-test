package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/ademuanthony/dfctipper/postgres/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (pg *PgDb) CreatUser(ctx context.Context, account *models.Account) error {
	return account.Insert(ctx, pg.db, boil.Infer())
}

func (pg *PgDb) ActivateByTelegramID(ctx context.Context, telegramID int64) error {
	colUp := models.M{
		models.AccountColumns.Active: 1,
	}
	_, err := models.Accounts(models.AccountWhere.TelegramID.EQ(telegramID)).UpdateAll(ctx, pg.db, colUp)
	return err
}

func (pg *PgDb) DeactivateByTelegramID(ctx context.Context, telegramID int64) error {
	colUp := models.M{
		models.AccountColumns.Active: 0,
	}
	_, err := models.Accounts(models.AccountWhere.TelegramID.EQ(telegramID)).UpdateAll(ctx, pg.db, colUp)
	return err
}

func (pg *PgDb) GetUser(ctx context.Context, id string) (*models.Account, error) {
	return models.Accounts(models.AccountWhere.ID.EQ(id)).One(ctx, pg.db)
}

func (pg *PgDb) ActiveUsersTelegram(ctx context.Context) (models.AccountSlice, error) {
	return models.Accounts(
		models.AccountWhere.Active.EQ(1),
		qm.Select(models.AccountColumns.ID, models.AccountColumns.TelegramID),
	).All(ctx, pg.db)
}

func (pg *PgDb) UserByTelegramID(ctx context.Context, telegramID int64) (*models.Account, error) {
	return models.Accounts(models.AccountWhere.TelegramID.EQ(int64(telegramID))).One(ctx, pg.db)
}

func (pg *PgDb) UserByTwitterID(ctx context.Context, twitterID int64) (*models.Account, error) {
	return models.Accounts(models.AccountWhere.TwitterID.EQ(int64(twitterID))).One(ctx, pg.db)
}

func (pg *PgDb) SetTwitterID(ctx context.Context, accID string, twitterID int64) error {
	colUp := models.M{
		models.AccountColumns.TwitterID: twitterID,
	}
	_, err := models.Accounts(models.AccountWhere.ID.EQ(accID)).UpdateAll(ctx, pg.db, colUp)
	return err
}

func (pg *PgDb) SetWalletAddress(ctx context.Context, telegramId int64, wallet string) error {
	colUp := models.M{
		models.AccountColumns.WalletAddress: wallet,
	}
	_, err := models.Accounts(models.AccountWhere.TelegramID.EQ(telegramId)).UpdateAll(ctx, pg.db, colUp)
	return err
}

func (pg *PgDb) SetCurrentStep(ctx context.Context, telegramID int64, step int) error {
	colUp := models.M{
		models.AccountColumns.CurrentStep: step,
	}
	_, err := models.Accounts(models.AccountWhere.TelegramID.EQ(telegramID)).UpdateAll(ctx, pg.db, colUp)
	return err
}

func (pg *PgDb) IncreaseDownlines(ctx context.Context, accID string) error {
	statement := "update account set downlines = downlines + 1, contest_downline = contest_downline + 1 where id = '" + accID + "'"
	_, err := pg.db.ExecContext(ctx, statement)
	return err
}

func (pg *PgDb) SetBalance(ctx context.Context, userID string, balance int64) error {
	colUp := models.M{
		models.AccountColumns.Balance: balance,
	}
	_, err := models.Accounts(models.AccountWhere.ID.EQ(userID)).UpdateAll(ctx, pg.db, colUp)
	return err
}

func (pg *PgDb) CurrentStep(ctx context.Context, telegramID int64) (int, error) {
	acc, err := models.Accounts(models.AccountWhere.TelegramID.EQ(telegramID)).One(ctx, pg.db)
	if err != nil {
		return -1, err
	}

	return acc.CurrentStep, nil
}

func (pg *PgDb) AllGetPendingWithdrawal(ctx context.Context) (models.WithdrawalSlice, error) {
	return models.Withdrawals(
		models.WithdrawalWhere.TXHash.EQ(""),
	).All(ctx, pg.db)
}

func (pg *PgDb) GetPendingWithdrawal(ctx context.Context, accID string) (*models.Withdrawal, error) {
	return models.Withdrawals(
		models.WithdrawalWhere.TXHash.EQ(""),
		models.WithdrawalWhere.UserID.EQ(accID)).One(ctx, pg.db)
}

func (pg *PgDb) MakeWithdrawalRequest(ctx context.Context, accID string, amount int64) error {
	req := models.Withdrawal{
		UserID: accID,
		Amount: amount,
		Date:   time.Now().UTC().Unix(),
	}

	return req.Insert(ctx, pg.db, boil.Infer())
}

func (pg *PgDb) UpdateTxHash(ctx context.Context, withID int, txHash string) error {
	colUp := models.M{
		models.WithdrawalColumns.TXHash: txHash,
	}
	_, err := models.Withdrawals(models.WithdrawalWhere.ID.EQ(withID)).UpdateAll(ctx, pg.db, colUp)
	return err
}

func (pg *PgDb) CreatePromotion(ctx context.Context, tweet string, number int, amount int, userID string) error {
	promotion := models.Promotion{
		CreatorID:        userID,
		CreatedAt:        time.Now().UTC().Unix(),
		TweetLink:        tweet,
		RewardCount:      number,
		RewardPerRetweet: amount,
	}

	return promotion.Insert(ctx, pg.db, boil.Infer())
}

func (pg *PgDb) IncreaseRetweetCount(ctx context.Context, promotionID, count int) error {
	statement := "update promotion set retweet_count = retweet_count + 1 where id = $1"
	_, err := pg.db.ExecContext(ctx, statement, promotionID)
	return err
}

func (pg *PgDb) SaveReward(ctx context.Context, promotionID int, userID string, reward int64) error {
	pReward := models.Reward{
		UserID: userID, PromotionID: promotionID,
		Date: time.Now().UTC().Unix(), Amount: int64(reward),
	}

	return pReward.Insert(ctx, pg.db, boil.Infer())
}

func (pg PgDb) UncompletedPromotion(ctx context.Context) (models.PromotionSlice, error) {
	return models.Promotions(
		qm.Where(fmt.Sprintf("%s < %s", models.PromotionColumns.RetweetCount, models.PromotionColumns.RewardCount)),
	).All(ctx, pg.db)
}

func (pg *PgDb) AvailablePromotion(ctx context.Context, accID string) (models.PromotionSlice, error) {
	promotions, err := models.Promotions(
		qm.Where(fmt.Sprintf("%s < %s", models.PromotionColumns.RetweetCount, models.PromotionColumns.RewardCount)),
	).All(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var availablePromotions models.PromotionSlice
	for _, promo := range promotions {
		exists, err := models.Rewards(
			models.RewardWhere.PromotionID.EQ(promo.ID),
			models.RewardWhere.UserID.EQ(accID),
		).Exists(ctx, pg.db)

		if err != nil {
			return nil, fmt.Errorf("models.Rewards -> %s", err.Error())
		}

		if exists {
			continue
		}

		availablePromotions = append(availablePromotions, promo)
	}

	return availablePromotions, nil
}

func (pg *PgDb) CanEarn(ctx context.Context, promotionID int, userID string) (bool, error) {
	exists, err := models.Rewards(
		models.RewardWhere.PromotionID.EQ(promotionID),
		models.RewardWhere.UserID.EQ(userID),
	).Exists(ctx, pg.db)

	if err != nil {
		return false, err
	}

	return !exists, nil
}
