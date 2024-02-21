-- name: GetQuestion :one
SELECT * FROM question 
WHERE id = $1 LIMIT 1;

-- name: ListQuestions :many
SELECT * FROM question;

-- name: ListQuestionsByQuestionType :many
SELECT * FROM question 
WHERE question_type = $1;

-- name: CreateQuestion :one
INSERT INTO question (
  question_type, points, name, question_text
) VALUES ( 
  $1, $2, $3,$4
) RETURNING *;

-- name: UpdateQuestion :one
UPDATE question SET 
  question_type = $1, points = $2, name = $3
WHERE id = $4
RETURNING *;

-- name: DeleteQuestion :exec
DELETE FROM question 
WHERE id = $1;
