-- name: GetUser :one
SELECT id, phone, name, age FROM users
WHERE id = $1;
