package main

import (
	router "echo-test/route"
	"echo-test/services"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	e := echo.New()

	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Failed to open log file")
	}

	logger := zerolog.New(file)

	taskService := services.NewTask(&logger)

	e.Static("/", "assets")

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	router.Init(e, &logger, taskService)

	e.Logger.Fatal(e.Start(":8000"))
}
