-- name: CreateRoom :one 
INSERT INTO rooms (
    user_id
) values (
    $1
) RETURNING *;

-- name: GetRoom :one
SELECT * FROM rooms
WHERE id=$1 LIMIT 1;

-- name: ListRooms :many
SELECT * FROM rooms
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateRoom :one
UPDATE rooms 
SET user_id = $2 
WHERE id=$1
RETURNING *;

-- name: DeleteRoom :exec
DELETE FROM rooms
WHERE id=$1; 