package services

import "github.com/labstack/echo/v4"

type ITaskService interface {
	Get(ctx echo.Context) ([]Task, error)
}
