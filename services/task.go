package services

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type Task struct {
	Id   int
	Name string
}

type TaskService struct {
	Log *zerolog.Logger
}

func NewTask(Log *zerolog.Logger) *TaskService {
	return &TaskService{
		Log: Log,
	}
}

func (ts *TaskService) Get(ctx echo.Context) ([]Task, error) {
	defaultTasks := make([]Task, 0)

	defaultTasks = append(defaultTasks, Task{
		Id:   1,
		Name: "First Tasker",
	})
	defaultTasks = append(defaultTasks, Task{
		Id:   2,
		Name: "Second Task",
	})
	defaultTasks = append(defaultTasks, Task{
		Id:   3,
		Name: "Third Task",
	})
	defaultTasks = append(defaultTasks, Task{
		Id:   4,
		Name: "Fourth Task",
	})

	return defaultTasks, nil
}
