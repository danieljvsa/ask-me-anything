-- name: CreateUser :one
INSERT INTO users (
    username, password
) values (
    $1, $2
) RETURNING *;
