// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, email FROM "user"
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
	)
	return i, err
}

const getUserSuggestions = `-- name: GetUserSuggestions :many
SELECT id as "value", concat(first_name, ' ', last_name) as "label" FROM "user"
`

type GetUserSuggestionsRow struct {
	Value pgtype.UUID
	Label interface{}
}

func (q *Queries) GetUserSuggestions(ctx context.Context) ([]GetUserSuggestionsRow, error) {
	rows, err := q.db.Query(ctx, getUserSuggestions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserSuggestionsRow
	for rows.Next() {
		var i GetUserSuggestionsRow
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

const listUsers = `-- name: ListUsers :many
SELECT id, first_name, last_name, email FROM "user"
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
