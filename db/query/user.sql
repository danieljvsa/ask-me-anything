-- name: CreateUser :one
INSERT INTO users (
    username, password, email
) values (
    $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id=$1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUsername :one
UPDATE users 
SET username = $2 
WHERE id=$1
RETURNING *;

-- name: UpdateEmail :one
UPDATE users 
SET email = $2 
WHERE id=$1
RETURNING *;

-- name: UpdatePassword :one
UPDATE users 
SET password = $2 
WHERE id=$1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id=$1; 