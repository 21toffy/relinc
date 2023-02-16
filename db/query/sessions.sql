-- name: CreateSession :one
INSERT INTO sessions (
  id, 
  username,
  refresh_token,
  email_address,
  user_agent,
  client_ip,
  is_blocked,
  expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;


-- name: GetSession :one
SELECT * FROM sessions
WHERE id = $1 LIMIT 1;