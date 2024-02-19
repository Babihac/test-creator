// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: question_type.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createQuestionType = `-- name: CreateQuestionType :one
INSERT INTO question_type (type) VALUES ($1) RETURNING id, type
`

func (q *Queries) CreateQuestionType(ctx context.Context, type_ string) (QuestionType, error) {
	row := q.db.QueryRow(ctx, createQuestionType, type_)
	var i QuestionType
	err := row.Scan(&i.ID, &i.Type)
	return i, err
}

const deleteQuestionType = `-- name: DeleteQuestionType :exec
DELETE FROM  question_type WHERE id = $1
`

func (q *Queries) DeleteQuestionType(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteQuestionType, id)
	return err
}

const getQuestionType = `-- name: GetQuestionType :one
SELECT id, type FROM question_type
where id = $1 LIMIT 1
`

func (q *Queries) GetQuestionType(ctx context.Context, id pgtype.UUID) (QuestionType, error) {
	row := q.db.QueryRow(ctx, getQuestionType, id)
	var i QuestionType
	err := row.Scan(&i.ID, &i.Type)
	return i, err
}

const getQuestionTypeSuggestions = `-- name: GetQuestionTypeSuggestions :many
SELECT id::text as value, type as label FROM question_type
`

type GetQuestionTypeSuggestionsRow struct {
	Value string
	Label string
}

func (q *Queries) GetQuestionTypeSuggestions(ctx context.Context) ([]GetQuestionTypeSuggestionsRow, error) {
	rows, err := q.db.Query(ctx, getQuestionTypeSuggestions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetQuestionTypeSuggestionsRow
	for rows.Next() {
		var i GetQuestionTypeSuggestionsRow
		if err := rows.Scan(&i.Value, &i.Label); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listQuestionTypes = `-- name: ListQuestionTypes :many
SELECT id, type FROM question_type
`

func (q *Queries) ListQuestionTypes(ctx context.Context) ([]QuestionType, error) {
	rows, err := q.db.Query(ctx, listQuestionTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []QuestionType
	for rows.Next() {
		var i QuestionType
		if err := rows.Scan(&i.ID, &i.Type); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateQuestionType = `-- name: UpdateQuestionType :one
UPDATE question_type SET type = $1 WHERE id = $2 RETURNING id, type
`

type UpdateQuestionTypeParams struct {
	Type string
	ID   pgtype.UUID
}

func (q *Queries) UpdateQuestionType(ctx context.Context, arg UpdateQuestionTypeParams) (QuestionType, error) {
	row := q.db.QueryRow(ctx, updateQuestionType, arg.Type, arg.ID)
	var i QuestionType
	err := row.Scan(&i.ID, &i.Type)
	return i, err
}
