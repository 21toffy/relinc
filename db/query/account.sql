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

-- name: GetUsertAccountByAccountId :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetUsersAccounts :many
SELECT * FROM accounts
ORDER BY owner
LIMIT $1
OFFSET $2;

-- name: GetUsersAccountsByUserEmail :many
SELECT accounts.*, users.email_address
FROM accounts
INNER JOIN users ON accounts.owner = users.id
WHERE users.email_address = $1 ;




-- name: GetUserAccountForUpdateAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;



-- name: UpdateUserAccount :one
UPDATE accounts
SET balance = balance + $2
WHERE id = $1
RETURNING *;


-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;


-- name: GetUserAccountWithUserFields :one
SELECT accounts.*, users.email_address, users.id as user_id
FROM accounts
INNER JOIN users ON accounts.owner = users.id
WHERE accounts.id = $1;