-- name: GetSession :one
SELECT * FROM Sessions
WHERE id = $1 LIMIT 1;

-- name: CreateSession :one
INSERT INTO Sessions (
    expires_at, user_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM Sessions
WHERE id = $1;

-- name: DeleteExpiredSessions :exec
DELETE FROM Sessions
WHERE expires_at <= now();
