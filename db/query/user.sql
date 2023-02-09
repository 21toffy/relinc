-- name: CreateUser :one
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
RETURNING *;


-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;


-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users SET first_name = $2, last_name = $3, email_address = $4, phone_number = $5, username = $6, dob = $7, address = $8
WHERE id = $1
RETURNING first_name, last_name, email_address, phone_number, username, dob, address;
