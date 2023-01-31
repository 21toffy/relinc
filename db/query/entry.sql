-- name: CreateEntry :one
INSERT INTO entries (
  account_id, 
  amount
) VALUES (
  $1, $2
)
RETURNING *;


-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;


-- name: ListEntrys :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;

-- name: UpdateEntry :exec
UPDATE entries SET account_id = $2, amount = $3
WHERE id = $1
RETURNING account_id, amount;
