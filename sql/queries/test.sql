-- name: GetTest :one
SELECT * FROM test
where id = $1 LIMIT 1;

-- name: ListTests :many
SELECT * FROM test;

-- name: ListTestsWithTeacher :many
-- struct: TestWithTeacher
SELECT t.id, t.name, t.teacher_id, t.duration, t.max_points, CONCAT(u.first_name, ' ', u.last_name) as teacher_name, t.created_at
FROM "test" t
JOIN "user" u ON t.teacher_id = u.id;

-- name: CreateTest :one
INSERT INTO test (name, teacher_id, duration, max_points) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateTest :one
UPDATE test SET name = $1, teacher_id = $2, duration = $3, max_points = $4 WHERE id = $5 RETURNING *;

-- name: DeleteTest :exec
DELETE FROM test WHERE id = $1;