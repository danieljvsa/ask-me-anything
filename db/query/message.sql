-- name: CreateMessage :one
INSERT INTO messages (
   message, user_id, parent_id, room_id
) values (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetMessage :one
SELECT * FROM messages
WHERE id=$1 LIMIT 1;

-- name: ListMessages :many
SELECT * FROM messages
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateMessage :one
UPDATE messages 
SET message = $2 
WHERE id=$1
RETURNING *;

-- name: UpdateParent :one
UPDATE messages 
SET parent_id = $2 
WHERE id=$1
RETURNING *;

-- name: UpdateLikes :one
UPDATE messages 
SET likes_count = $2 
WHERE id=$1
RETURNING *;

-- name: UpdateAnswered :one
UPDATE messages 
SET answered = $2 
WHERE id=$1
RETURNING *;

-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id=$1; 