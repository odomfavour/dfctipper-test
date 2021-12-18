// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Accounts", testAccounts)
	t.Run("Deposits", testDeposits)
	t.Run("Promotions", testPromotions)
	t.Run("Rewards", testRewards)
	t.Run("Withdrawals", testWithdrawals)
}

func TestDelete(t *testing.T) {
	t.Run("Accounts", testAccountsDelete)
	t.Run("Deposits", testDepositsDelete)
	t.Run("Promotions", testPromotionsDelete)
	t.Run("Rewards", testRewardsDelete)
	t.Run("Withdrawals", testWithdrawalsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsQueryDeleteAll)
	t.Run("Deposits", testDepositsQueryDeleteAll)
	t.Run("Promotions", testPromotionsQueryDeleteAll)
	t.Run("Rewards", testRewardsQueryDeleteAll)
	t.Run("Withdrawals", testWithdrawalsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceDeleteAll)
	t.Run("Deposits", testDepositsSliceDeleteAll)
	t.Run("Promotions", testPromotionsSliceDeleteAll)
	t.Run("Rewards", testRewardsSliceDeleteAll)
	t.Run("Withdrawals", testWithdrawalsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Accounts", testAccountsExists)
	t.Run("Deposits", testDepositsExists)
	t.Run("Promotions", testPromotionsExists)
	t.Run("Rewards", testRewardsExists)
	t.Run("Withdrawals", testWithdrawalsExists)
}

func TestFind(t *testing.T) {
	t.Run("Accounts", testAccountsFind)
	t.Run("Deposits", testDepositsFind)
	t.Run("Promotions", testPromotionsFind)
	t.Run("Rewards", testRewardsFind)
	t.Run("Withdrawals", testWithdrawalsFind)
}

func TestBind(t *testing.T) {
	t.Run("Accounts", testAccountsBind)
	t.Run("Deposits", testDepositsBind)
	t.Run("Promotions", testPromotionsBind)
	t.Run("Rewards", testRewardsBind)
	t.Run("Withdrawals", testWithdrawalsBind)
}

func TestOne(t *testing.T) {
	t.Run("Accounts", testAccountsOne)
	t.Run("Deposits", testDepositsOne)
	t.Run("Promotions", testPromotionsOne)
	t.Run("Rewards", testRewardsOne)
	t.Run("Withdrawals", testWithdrawalsOne)
}

func TestAll(t *testing.T) {
	t.Run("Accounts", testAccountsAll)
	t.Run("Deposits", testDepositsAll)
	t.Run("Promotions", testPromotionsAll)
	t.Run("Rewards", testRewardsAll)
	t.Run("Withdrawals", testWithdrawalsAll)
}

func TestCount(t *testing.T) {
	t.Run("Accounts", testAccountsCount)
	t.Run("Deposits", testDepositsCount)
	t.Run("Promotions", testPromotionsCount)
	t.Run("Rewards", testRewardsCount)
	t.Run("Withdrawals", testWithdrawalsCount)
}

func TestInsert(t *testing.T) {
	t.Run("Accounts", testAccountsInsert)
	t.Run("Accounts", testAccountsInsertWhitelist)
	t.Run("Deposits", testDepositsInsert)
	t.Run("Deposits", testDepositsInsertWhitelist)
	t.Run("Promotions", testPromotionsInsert)
	t.Run("Promotions", testPromotionsInsertWhitelist)
	t.Run("Rewards", testRewardsInsert)
	t.Run("Rewards", testRewardsInsertWhitelist)
	t.Run("Withdrawals", testWithdrawalsInsert)
	t.Run("Withdrawals", testWithdrawalsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("DepositToAccountUsingUser", testDepositToOneAccountUsingUser)
	t.Run("PromotionToAccountUsingCreator", testPromotionToOneAccountUsingCreator)
	t.Run("RewardToPromotionUsingPromotion", testRewardToOnePromotionUsingPromotion)
	t.Run("RewardToAccountUsingUser", testRewardToOneAccountUsingUser)
	t.Run("WithdrawalToAccountUsingUser", testWithdrawalToOneAccountUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AccountToUserDeposits", testAccountToManyUserDeposits)
	t.Run("AccountToCreatorPromotions", testAccountToManyCreatorPromotions)
	t.Run("AccountToUserRewards", testAccountToManyUserRewards)
	t.Run("AccountToUserWithdrawals", testAccountToManyUserWithdrawals)
	t.Run("PromotionToRewards", testPromotionToManyRewards)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("DepositToAccountUsingUserDeposits", testDepositToOneSetOpAccountUsingUser)
	t.Run("PromotionToAccountUsingCreatorPromotions", testPromotionToOneSetOpAccountUsingCreator)
	t.Run("RewardToPromotionUsingRewards", testRewardToOneSetOpPromotionUsingPromotion)
	t.Run("RewardToAccountUsingUserRewards", testRewardToOneSetOpAccountUsingUser)
	t.Run("WithdrawalToAccountUsingUserWithdrawals", testWithdrawalToOneSetOpAccountUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AccountToUserDeposits", testAccountToManyAddOpUserDeposits)
	t.Run("AccountToCreatorPromotions", testAccountToManyAddOpCreatorPromotions)
	t.Run("AccountToUserRewards", testAccountToManyAddOpUserRewards)
	t.Run("AccountToUserWithdrawals", testAccountToManyAddOpUserWithdrawals)
	t.Run("PromotionToRewards", testPromotionToManyAddOpRewards)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Accounts", testAccountsReload)
	t.Run("Deposits", testDepositsReload)
	t.Run("Promotions", testPromotionsReload)
	t.Run("Rewards", testRewardsReload)
	t.Run("Withdrawals", testWithdrawalsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Accounts", testAccountsReloadAll)
	t.Run("Deposits", testDepositsReloadAll)
	t.Run("Promotions", testPromotionsReloadAll)
	t.Run("Rewards", testRewardsReloadAll)
	t.Run("Withdrawals", testWithdrawalsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Accounts", testAccountsSelect)
	t.Run("Deposits", testDepositsSelect)
	t.Run("Promotions", testPromotionsSelect)
	t.Run("Rewards", testRewardsSelect)
	t.Run("Withdrawals", testWithdrawalsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Accounts", testAccountsUpdate)
	t.Run("Deposits", testDepositsUpdate)
	t.Run("Promotions", testPromotionsUpdate)
	t.Run("Rewards", testRewardsUpdate)
	t.Run("Withdrawals", testWithdrawalsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceUpdateAll)
	t.Run("Deposits", testDepositsSliceUpdateAll)
	t.Run("Promotions", testPromotionsSliceUpdateAll)
	t.Run("Rewards", testRewardsSliceUpdateAll)
	t.Run("Withdrawals", testWithdrawalsSliceUpdateAll)
}