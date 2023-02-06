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

type GetUserAccountTxResult struct {
	User    User    `json:"user"`
	Account Account `json:"account"`
}
type GetUserAccountTxRequest struct {
	UserID int64 `json:"user_id"`
}

type ListUsersAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsersAccounts(ctx context.Context, arg GetUsersAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, getUsersAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.AccountType,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// Get user Account transaction body
func (store *Store) ListUsersAccountsTx(ctx context.Context, arg int64) (GetUserAccountTxResult, error) {
	var result GetUserAccountTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.GetUser(ctx, arg)

		if err != nil {
			return err
		}

		result.Account, err = q.GetUsertAccount(ctx, arg)
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

// Get user Account transaction body
func (store *Store) GetAccountTx(ctx context.Context, arg int64) (GetUserAccountTxResult, error) {
	var result GetUserAccountTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.GetUser(ctx, arg)

		if err != nil {
			return err
		}

		result.Account, err = q.GetUsertAccount(ctx, arg)
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
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
