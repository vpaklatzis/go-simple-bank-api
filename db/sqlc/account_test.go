package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vpaklatzis/go-simple-bank/util"
)

func createRandomAccount(t *testing.T) Account {
	param := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, param.Owner, account.Owner)
	require.Equal(t, param.Balance, account.Balance)
	require.Equal(t, param.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.UpdatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accountA := createRandomAccount(t)
	accountB, err := testQueries.GetAccount(context.Background(), accountA.ID)
	require.NoError(t, err)
	require.NotEmpty(t, accountB)

	require.Equal(t, accountA.ID, accountB.ID)
	require.Equal(t, accountA.Owner, accountB.Owner)
	require.Equal(t, accountA.Balance, accountB.Balance)
	require.Equal(t, accountA.Currency, accountB.Currency)
	require.WithinDuration(t, accountA.CreatedAt, accountB.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	accountA := createRandomAccount(t)
	params := UpdateAccountParams{
		ID:      accountA.ID,
		Balance: util.RandomBalance(),
	}
	accountB, err := testQueries.UpdateAccount(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, accountB)

	require.Equal(t, accountA.ID, accountB.ID)
	require.Equal(t, accountA.Owner, accountB.Owner)
	require.Equal(t, params.Balance, accountB.Balance)
	require.Equal(t, accountA.Currency, accountB.Currency)
	require.WithinDuration(t, accountA.CreatedAt, accountB.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	accountA := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), accountA.ID)
	require.NoError(t, err)

	accountB, err := testQueries.GetAccount(context.Background(), accountA.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountB)
}

func TestListAccounts(t *testing.T) {
	for i := 1; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
