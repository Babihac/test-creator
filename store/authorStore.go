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

func (as *AuthorStore) DeleteAuthor(ctx context.Context, id int64) error {
	return as.queries.DeleteAuthor(ctx, id)
}

func (as *AuthorStore) EditAuthor(ctx context.Context, params db.UpdateAuthorParams) error {
	return as.queries.UpdateAuthor(ctx, params)
}
