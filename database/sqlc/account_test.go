package database

import (
	"context"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	createdAccount, err := createRandomAccount()
	defer testQueries.DropAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotNil(t, createdAccount)

	require.NotZero(t, createdAccount.ID)
	require.NotZero(t, createdAccount.CreatedAt)
}

func TestShowAccount(t *testing.T) {
	randomAccount, err := createRandomAccount()
	defer testQueries.DropAccount(context.Background(), randomAccount.ID)

	account, err := testQueries.GetAccount(context.Background(), randomAccount.ID)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.Equal(t, account.ID, randomAccount.ID)
}

func createRandomAccount() (a Account, e error) {
	a, e = testQueries.CreateAccount(context.Background(), rand.Int63n(999999))
	return
}
