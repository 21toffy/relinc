package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransfertx(t *testing.T) {
	store := NewStore(testDB)

	account1 := RandomTestAccount(t)
	account2 := RandomTestAccount(t)

	n := 5

	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//check account entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		//check account entries
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check accounts

		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount.ID)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount.ID)
		require.Equal(t, account2.ID, toAccount.ID)

		//check account
		dif1 := account1.Balance - fromAccount.Balance
		dif2 := toAccount.Balance - account2.Balance

		require.Equal(t, dif1, dif2)
		require.True(t, dif1 >= dif2)
		require.True(t, dif1%amount == 0)
		k := int(dif1 / amount)

		require.True(t, k >= 1 && k <= n)
		existed[k] = true

	}

	// check the final updated balance

	updatedAccount1, err := testQueries.GetUsertAccount(context.Background(), account1.Owner)

	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetUsertAccount(context.Background(), account2.Owner)
	require.NoError(t, err)

	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)

}
