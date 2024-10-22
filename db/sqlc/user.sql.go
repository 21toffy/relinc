// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: user.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  first_name, 
  last_name,
  email_address,
  phone_number,
  username,
  dob,
  password,
  address,
  profile_picture,
  gender
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING id, first_name, last_name, email_address, phone_number, username, password, password_changed_at, dob, address, profile_picture, gender, created_at
`

type CreateUserParams struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	EmailAddress   string    `json:"email_address"`
	PhoneNumber    string    `json:"phone_number"`
	Username       string    `json:"username"`
	Dob            time.Time `json:"dob"`
	Password       string    `json:"password"`
	Address        string    `json:"address"`
	ProfilePicture string    `json:"profile_picture"`
	Gender         string    `json:"gender"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.EmailAddress,
		arg.PhoneNumber,
		arg.Username,
		arg.Dob,
		arg.Password,
		arg.Address,
		arg.ProfilePicture,
		arg.Gender,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.EmailAddress,
		&i.PhoneNumber,
		&i.Username,
		&i.Password,
		&i.PasswordChangedAt,
		&i.Dob,
		&i.Address,
		&i.ProfilePicture,
		&i.Gender,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, email_address, phone_number, username, password, password_changed_at, dob, address, profile_picture, gender, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.EmailAddress,
		&i.PhoneNumber,
		&i.Username,
		&i.Password,
		&i.PasswordChangedAt,
		&i.Dob,
		&i.Address,
		&i.ProfilePicture,
		&i.Gender,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email_address, phone_number, username, password, password_changed_at, dob, address, profile_picture, gender, created_at FROM users
WHERE email_address = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, emailAddress string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, emailAddress)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.EmailAddress,
		&i.PhoneNumber,
		&i.Username,
		&i.Password,
		&i.PasswordChangedAt,
		&i.Dob,
		&i.Address,
		&i.ProfilePicture,
		&i.Gender,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, first_name, last_name, email_address, phone_number, username, password, password_changed_at, dob, address, profile_picture, gender, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.EmailAddress,
		&i.PhoneNumber,
		&i.Username,
		&i.Password,
		&i.PasswordChangedAt,
		&i.Dob,
		&i.Address,
		&i.ProfilePicture,
		&i.Gender,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, first_name, last_name, email_address, phone_number, username, password, password_changed_at, dob, address, profile_picture, gender, created_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.EmailAddress,
			&i.PhoneNumber,
			&i.Username,
			&i.Password,
			&i.PasswordChangedAt,
			&i.Dob,
			&i.Address,
			&i.ProfilePicture,
			&i.Gender,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users SET first_name = $2, last_name = $3, email_address = $4, phone_number = $5, username = $6, dob = $7, address = $8
WHERE id = $1
RETURNING first_name, last_name, email_address, phone_number, username, dob, address
`

type UpdateUserParams struct {
	ID           int64     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	EmailAddress string    `json:"email_address"`
	PhoneNumber  string    `json:"phone_number"`
	Username     string    `json:"username"`
	Dob          time.Time `json:"dob"`
	Address      string    `json:"address"`
}

type UpdateUserRow struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	EmailAddress string    `json:"email_address"`
	PhoneNumber  string    `json:"phone_number"`
	Username     string    `json:"username"`
	Dob          time.Time `json:"dob"`
	Address      string    `json:"address"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.EmailAddress,
		arg.PhoneNumber,
		arg.Username,
		arg.Dob,
		arg.Address,
	)
	var i UpdateUserRow
	err := row.Scan(
		&i.FirstName,
		&i.LastName,
		&i.EmailAddress,
		&i.PhoneNumber,
		&i.Username,
		&i.Dob,
		&i.Address,
	)
	return i, err
}
