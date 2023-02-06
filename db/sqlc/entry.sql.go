// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: entry.sql

package db

import (
	"context"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (
  account_id, 
  amount
) VALUES (
  $1, $2
)
RETURNING id, account_id, amount, created_at
`

type CreateEntryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.AccountID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1
`

func (q *Queries) DeleteEntry(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntry, id)
	return err
}

const getEntry = `-- name: GetEntry :one
SELECT id, account_id, amount, created_at FROM entries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntrys = `-- name: ListEntrys :many
SELECT id, account_id, amount, created_at FROM entries
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListEntrysParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEntrys(ctx context.Context, arg ListEntrysParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntrys, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entry{}
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
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

const updateEntry = `-- name: UpdateEntry :exec
UPDATE entries SET account_id = $2, amount = $3
WHERE id = $1
RETURNING account_id, amount
`

type UpdateEntryParams struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

type UpdateEntryRow struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) UpdateEntry(ctx context.Context, arg UpdateEntryParams) error {
	_, err := q.db.ExecContext(ctx, updateEntry, arg.ID, arg.AccountID, arg.Amount)
	return err
}
