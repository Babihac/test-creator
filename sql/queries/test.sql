-- name: GetTest :one
SELECT * FROM test
where id = $1 LIMIT 1;

-- name: ListTests :many
SELECT * FROM test;

-- name: ListTestsWithTeacher :many
-- struct: TestWithTeacher
SELECT t.id, t.name, t.teacher_id, t.duration, t.max_points, t.date, CONCAT(u.first_name, ' ', u.last_name) as teacher_name, t.created_at
FROM "test" t
JOIN "user" u ON t.teacher_id = u.id;

-- name: CreateTest :one
INSERT INTO test (name, teacher_id, duration, max_points, date) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateTest :one
UPDATE test SET name = $1, teacher_id = $2, duration = $3, max_points = $4, date = $5 WHERE id = $6 RETURNING *;

-- name: DeleteTest :exec
DELETE FROM test WHERE id = $1;