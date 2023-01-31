package db

import (
	"context"
	"time"
)

type CreateAccountTxParams struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	EmailAddress   string    `json:"email_address"`
	PhoneNumber    string    `json:"phone_number"`
	Username       string    `json:"username"`
	Dob            time.Time `json:"dob"`
	Address        string    `json:"address"`
	ProfilePicture string    `json:"profile_picture"`
	Gender         string    `json:"gender"`
	Balance        int64     `json:"balance"`
	Currency       string    `json:"currency"`
	AccountType    string    `json:"account_type"`
}
type CreateAccountTxResult struct {
	User    User    `json:"user"`
	Account Account `json:"account"`
}

// TransferTx performs a money transfer from one account to another
// it will create transfer record, add account entries and update account balance within a single db transaction
func (store *Store) CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (CreateAccountTxResult, error) {
	var result CreateAccountTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, CreateUserParams{
			FirstName:      arg.FirstName,
			LastName:       arg.LastName,
			EmailAddress:   arg.EmailAddress,
			PhoneNumber:    arg.PhoneNumber,
			Username:       arg.Username,
			Dob:            arg.Dob,
			Address:        arg.Address,
			ProfilePicture: arg.ProfilePicture,
			Gender:         arg.Gender,
		})

		if err != nil {
			return err
		}

		result.Account, err = q.CreateUserAccount(ctx, CreateUserAccountParams{
			Owner:       result.User.ID,
			Balance:     arg.Balance,
			Currency:    arg.Currency,
			AccountType: arg.AccountType,
		})
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
