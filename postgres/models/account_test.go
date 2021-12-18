// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testAccounts(t *testing.T) {
	t.Parallel()

	query := Accounts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAccountsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Accounts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AccountSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AccountExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Account exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AccountExists to return true, but got false.")
	}
}

func testAccountsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	accountFound, err := FindAccount(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if accountFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAccountsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Accounts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAccountsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Accounts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAccountsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountOne := &Account{}
	accountTwo := &Account{}
	if err = randomize.Struct(seed, accountOne, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	if err = randomize.Struct(seed, accountTwo, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = accountOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = accountTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Accounts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAccountsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	accountOne := &Account{}
	accountTwo := &Account{}
	if err = randomize.Struct(seed, accountOne, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	if err = randomize.Struct(seed, accountTwo, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = accountOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = accountTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testAccountsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAccountsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(accountColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAccountToManyUserDeposits(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c Deposit

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, depositDBTypes, false, depositColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, depositDBTypes, false, depositColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.UserID = a.ID
	c.UserID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserDeposits().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.UserID == b.UserID {
			bFound = true
		}
		if v.UserID == c.UserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AccountSlice{&a}
	if err = a.L.LoadUserDeposits(ctx, tx, false, (*[]*Account)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserDeposits); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserDeposits = nil
	if err = a.L.LoadUserDeposits(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserDeposits); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAccountToManyCreatorPromotions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c Promotion

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, promotionDBTypes, false, promotionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, promotionDBTypes, false, promotionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.CreatorID = a.ID
	c.CreatorID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.CreatorPromotions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.CreatorID == b.CreatorID {
			bFound = true
		}
		if v.CreatorID == c.CreatorID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AccountSlice{&a}
	if err = a.L.LoadCreatorPromotions(ctx, tx, false, (*[]*Account)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CreatorPromotions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CreatorPromotions = nil
	if err = a.L.LoadCreatorPromotions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CreatorPromotions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAccountToManyUserRewards(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c Reward

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, rewardDBTypes, false, rewardColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, rewardDBTypes, false, rewardColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.UserID = a.ID
	c.UserID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserRewards().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.UserID == b.UserID {
			bFound = true
		}
		if v.UserID == c.UserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AccountSlice{&a}
	if err = a.L.LoadUserRewards(ctx, tx, false, (*[]*Account)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserRewards); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserRewards = nil
	if err = a.L.LoadUserRewards(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserRewards); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAccountToManyUserWithdrawals(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c Withdrawal

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, withdrawalDBTypes, false, withdrawalColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, withdrawalDBTypes, false, withdrawalColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.UserID = a.ID
	c.UserID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserWithdrawals().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.UserID == b.UserID {
			bFound = true
		}
		if v.UserID == c.UserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AccountSlice{&a}
	if err = a.L.LoadUserWithdrawals(ctx, tx, false, (*[]*Account)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserWithdrawals); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserWithdrawals = nil
	if err = a.L.LoadUserWithdrawals(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserWithdrawals); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAccountToManyAddOpUserDeposits(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c, d, e Deposit

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Deposit{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, depositDBTypes, false, strmangle.SetComplement(depositPrimaryKeyColumns, depositColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Deposit{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserDeposits(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.UserID {
			t.Error("foreign key was wrong value", a.ID, first.UserID)
		}
		if a.ID != second.UserID {
			t.Error("foreign key was wrong value", a.ID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserDeposits[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserDeposits[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserDeposits().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAccountToManyAddOpCreatorPromotions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c, d, e Promotion

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Promotion{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, promotionDBTypes, false, strmangle.SetComplement(promotionPrimaryKeyColumns, promotionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Promotion{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCreatorPromotions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CreatorID {
			t.Error("foreign key was wrong value", a.ID, first.CreatorID)
		}
		if a.ID != second.CreatorID {
			t.Error("foreign key was wrong value", a.ID, second.CreatorID)
		}

		if first.R.Creator != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Creator != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.CreatorPromotions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CreatorPromotions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CreatorPromotions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAccountToManyAddOpUserRewards(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c, d, e Reward

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Reward{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, rewardDBTypes, false, strmangle.SetComplement(rewardPrimaryKeyColumns, rewardColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Reward{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserRewards(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.UserID {
			t.Error("foreign key was wrong value", a.ID, first.UserID)
		}
		if a.ID != second.UserID {
			t.Error("foreign key was wrong value", a.ID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserRewards[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserRewards[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserRewards().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAccountToManyAddOpUserWithdrawals(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Account
	var b, c, d, e Withdrawal

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Withdrawal{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, withdrawalDBTypes, false, strmangle.SetComplement(withdrawalPrimaryKeyColumns, withdrawalColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Withdrawal{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserWithdrawals(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.UserID {
			t.Error("foreign key was wrong value", a.ID, first.UserID)
		}
		if a.ID != second.UserID {
			t.Error("foreign key was wrong value", a.ID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserWithdrawals[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserWithdrawals[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserWithdrawals().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAccountsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAccountsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AccountSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAccountsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Accounts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	accountDBTypes = map[string]string{`ID`: `character varying`, `ReferralID`: `character varying`, `Balance`: `bigint`, `Username`: `character varying`, `FirstName`: `character varying`, `LastName`: `character varying`, `TwitterID`: `bigint`, `TelegramID`: `bigint`, `JoinAt`: `bigint`, `WalletAddress`: `character varying`, `DepositAddress`: `character varying`, `CurrentStep`: `integer`}
	_              = bytes.MinRead
)

func testAccountsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(accountAllColumns) == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, accountDBTypes, true, accountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAccountsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(accountAllColumns) == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, accountDBTypes, true, accountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(accountAllColumns, accountPrimaryKeyColumns) {
		fields = accountAllColumns
	} else {
		fields = strmangle.SetComplement(
			accountAllColumns,
			accountPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := AccountSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAccountsUpsert(t *testing.T) {
	t.Parallel()

	if len(accountAllColumns) == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Account{}
	if err = randomize.Struct(seed, &o, accountDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Account: %s", err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, accountDBTypes, false, accountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Account: %s", err)
	}

	count, err = Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}