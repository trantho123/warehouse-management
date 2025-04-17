-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password,
    role_id,
    created_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
    username = COALESCE(sqlc.narg(username),username),
    email = COALESCE(sqlc.narg(email), email),
    password = COALESCE(sqlc.narg(password), password),
    role_id = COALESCE(sqlc.narg(role_id), role_id),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

