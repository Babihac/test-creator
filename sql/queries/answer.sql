-- name: GetAnswer :one
SELECT * from answer
WHERE id = $1
LIMIT 1;

-- name: ListAnswers :many
SELECT * from answer;

-- name: CreateAnswer :one
INSERT INTO answer (value, question_id, correct) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateAnswer :one
UPDATE answer SET value = $1, question_id = $2, correct = $3 WHERE id = $4 RETURNING *;

-- name: DeleteAnswer :exec
DELETE FROM answer WHERE id = $1;