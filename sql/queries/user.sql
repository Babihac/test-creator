-- name: GetUser :one
SELECT * FROM "user"
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user";

-- name: GetUserSuggestions :many
SELECT id as "value", concat(first_name, ' ', last_name) as "label" FROM "user";