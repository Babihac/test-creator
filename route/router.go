package router

import (
	"echo-test/handlers"
	"echo-test/services"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func Init(e *echo.Echo, logger *zerolog.Logger, taskService services.ITaskService) {
	taskHandler := handlers.NewTask(logger, taskService)
	taskHandler.Serve(e)
}
