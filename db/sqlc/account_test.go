package db

import (
	"context"
	"relinc/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func RandomTestAccount(t *testing.T) Account {
	user := CreateRandomUser(t)

	arg := CreateUserAccountParams{
		Owner:       user.ID,
		Balance:     util.RandomMoney(),
		Currency:    util.RandomCurrency(),
		AccountType: util.RandomAccountType(),
	}
	account, err := testQueries.CreateUserAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Owner, user.ID)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.AccountType, account.AccountType)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account

}
func TestCreateUserAccount(t *testing.T) {
	RandomTestAccount(t)
}
