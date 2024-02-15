package store

import (
	"context"
	"echo-test/db"
)

type AuthorStore struct {
	queries *db.Queries
}

func NewAuthorStore(queries *db.Queries) *AuthorStore {
	return &AuthorStore{
		queries: queries,
	}
}

func (as *AuthorStore) ListAuthors(ctx context.Context) ([]db.Author, error) {
	return as.queries.ListAuthors(ctx)
}

func (as *AuthorStore) GetAuthor(id int64, ctx context.Context) (db.Author, error) {
	return as.queries.GetAuthor(ctx, id)
}

func (as *AuthorStore) CreateAuthor(ctx context.Context, params db.CreateAuthorParams) (db.Author, error) {
	return as.queries.CreateAuthor(ctx, params)
}
