package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransaction(t *testing.T) {

	args := CreateTransactionsParams{
		AccountID:       1,
		OperationTypeID: 2,
		Amount:          -25.5,
	}

	createdTransaction, err := testQueries.CreateTransactions(context.Background(), args)
	defer testQueries.DropTransaction(context.Background(), createdTransaction.ID)
	require.NoError(t, err)
	require.NotNil(t, createdTransaction)
	require.Equal(t, createdTransaction.AccountID, args.AccountID)
	require.Equal(t, createdTransaction.OperationTypeID, args.OperationTypeID)
	require.Equal(t, createdTransaction.Amount, args.Amount)

	require.NotZero(t, createdTransaction.ID)
	require.NotZero(t, createdTransaction.EventDate)
	require.NotZero(t, createdTransaction.UpdatedAt)

}

func TestShowTransaction(t *testing.T) {
	transaction, err := testQueries.GetTransactions(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, transaction)
	require.NotZero(t, transaction.ID)
}
