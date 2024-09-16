-- name: CreateMessage :one
INSERT INTO messages (
   message, user_id, parent_id, room_id
) values (
    $1, $2, $3, $4
) RETURNING *;
