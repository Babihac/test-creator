package services

import (
	"echo-test/db"
	"echo-test/store"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type AuthorService struct {
	log         *zerolog.Logger
	authorStore *store.AuthorStore
}

func NewAuthorService(log *zerolog.Logger, authorStore *store.AuthorStore) *AuthorService {
	return &AuthorService{
		log:         log,
		authorStore: authorStore,
	}
}

func (as *AuthorService) Get(ctx echo.Context) ([]db.Author, error) {
	return as.authorStore.ListAuthors(ctx.Request().Context())
}

func (as *AuthorService) GetOne(id int64, ctx echo.Context) (db.Author, error) {
	return as.authorStore.GetAuthor(id, ctx.Request().Context())
}

func (as *AuthorService) Create(ctx echo.Context, params db.CreateAuthorParams) (db.Author, error) {
	return as.authorStore.CreateAuthor(ctx.Request().Context(), params)
}

func (as *AuthorService) Delete(ctx echo.Context, id int64) error {
	return as.authorStore.DeleteAuthor(ctx.Request().Context(), id)
}

func (as *AuthorService) Update(ctx echo.Context, params db.UpdateAuthorParams) error {
	return as.authorStore.EditAuthor(ctx.Request().Context(), params)
}
