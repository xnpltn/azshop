-- name: CreateUser :exec
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3);

-- name: GetUserByID :one
SELECT id, name, email, verified FROM users WHERE id=$1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email=$1;



