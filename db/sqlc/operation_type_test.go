package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateOperationType(t *testing.T) {

	args := CreateOperationTypeParams{
		Multiplier:  1,
		Description: "COMPRA CREBITO",
	}

	operationType, err := testQueries.CreateOperationType(context.Background(), args)
	defer testQueries.DropOperationType(context.Background(), operationType.ID)
	require.NoError(t, err)
	require.NotNil(t, operationType)
	require.Equal(t, operationType.Description, args.Description)

	require.NotZero(t, operationType.ID)
	require.NotZero(t, operationType.CreatedAt)
	require.NotZero(t, operationType.UpdatedAt)

}

func TestShowOperationType(t *testing.T) {
	operationType, err := testQueries.GetOperationTypes(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, operationType)
	require.NotZero(t, operationType.ID)
}
