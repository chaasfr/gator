-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * from users where name = $1;

-- name: DeleteUsers :exec
TRUNCATE users;

-- name: GetUsers :many
SELECT * from users;