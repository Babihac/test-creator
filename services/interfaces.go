package services

import (
	"echo-test/db"

	"github.com/labstack/echo/v4"
)

type ITaskService interface {
	Get(ctx echo.Context) ([]Task, error)
}

type IAuthorService interface {
	Get(ctx echo.Context) ([]db.Author, error)
	GetOne(id int64, ctx echo.Context) (db.Author, error)
	Create(ctx echo.Context, params db.CreateAuthorParams) (db.Author, error)
}
