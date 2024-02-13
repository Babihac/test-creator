package handlers

import (
	"echo-test/components"
	"echo-test/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type TaskHandler struct {
	Log         *zerolog.Logger
	TaskService services.ITaskService
}

func NewTask(log *zerolog.Logger, ts services.ITaskService) *TaskHandler {
	return &TaskHandler{
		Log:         log,
		TaskService: ts,
	}
}

func (t *TaskHandler) Serve(e *echo.Echo) {
	group := e.Group("/tasks")

	group.GET("", t.get)

}

func (t *TaskHandler) get(c echo.Context) error {
	values, err := t.TaskService.Get(c)

	if err != nil {
		t.Log.Error().Msg("Error getting tasks")
		c.String(http.StatusInternalServerError, "failed")
		return err
	}

	comp := components.Tasks(values)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}
