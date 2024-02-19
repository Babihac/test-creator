-- name: GetQuestionType :one
SELECT * FROM question_type
where id = $1 LIMIT 1;

-- name: GetQuestionTypeSuggestions :many
SELECT id::text as value, type as label FROM question_type;

-- name: ListQuestionTypes :many
SELECT * FROM question_type;

-- name: CreateQuestionType :one
INSERT INTO question_type (type) VALUES ($1) RETURNING *;

-- name: UpdateQuestionType :one
UPDATE question_type SET type = $1 WHERE id = $2 RETURNING *;

-- name: DeleteQuestionType :exec
DELETE FROM  question_type WHERE id = $1;