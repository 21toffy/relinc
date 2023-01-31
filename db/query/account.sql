-- name: CreateUserAccount :one
INSERT INTO accounts (
    owner,
    balance,
    currency,
    account_type
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetUsertAccount :one
SELECT * FROM accounts
WHERE owner = $1 LIMIT 1;

-- name: GetUsersAccounts :many
SELECT * FROM accounts
ORDER BY owner
LIMIT $1
OFFSET $2;


-- name: GetUserAccountForUpdateAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- name: UpdateUserAccount :one
UPDATE accounts SET balance = $2
WHERE id = $1
RETURNING *;
