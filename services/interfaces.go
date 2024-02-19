package services

import (
	"echo-test/db"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type IAuthorService interface {
	Get(ctx echo.Context) ([]db.Author, error)
	GetOne(id int64, ctx echo.Context) (db.Author, error)
	Create(ctx echo.Context, params db.CreateAuthorParams) (db.Author, error)
	Delete(ctx echo.Context, id int64) error
	Update(ctx echo.Context, params db.UpdateAuthorParams) error
}

type ITestService interface {
	Get(ctx echo.Context) ([]db.ListTestsWithTeacherRow, error)
	GetOne(ctx echo.Context, id pgtype.UUID) (db.Test, error)
	Create(ctx echo.Context, params db.CreateTestParams) (db.Test, error)
	Delete(ctx echo.Context, id int64) error
	Update(ctx echo.Context, params db.UpdateTestParams) error
}

type IUserService interface {
	Get(ctx echo.Context) ([]db.ListTestsWithTeacherRow, error)
	GetOne(id int64, ctx echo.Context) (db.Test, error)
	Create(ctx echo.Context, params db.CreateTestParams) (db.User, error)
	Delete(ctx echo.Context, id int64) error
	Update(ctx echo.Context, params db.UpdateTestParams) error
	GetUserSuggestions(ctx echo.Context) ([]db.GetUserSuggestionsRow, error)
}
