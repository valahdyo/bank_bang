package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/valahdyo/bank_bang/utils"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: utils.RandomOwner(),
		Balance: utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	getAccount, err:= testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, getAccount)

	require.Equal(t, account1.ID, getAccount.ID)
	require.Equal(t, account1.Owner, getAccount.Owner)
	require.Equal(t, account1.Balance, getAccount.Balance)
	require.Equal(t, account1.Currency, getAccount.Currency)
	require.WithinDuration(t, account1.CreatedAt, getAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID: account1.ID,
		Balance: utils.RandomMoney(),
	}

	getAccount, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account1)

	require.Equal(t, account1.ID, getAccount.ID)
	require.Equal(t, account1.Owner, getAccount.Owner)
	require.Equal(t, arg.Balance, getAccount.Balance)
	require.Equal(t, account1.Currency, getAccount.Currency)
	require.WithinDuration(t, account1.CreatedAt, getAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)

	getAccount, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, getAccount)
}

func TestListAccount(t *testing.T) {
	for i := 0; i<10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)
	
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}