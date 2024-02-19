package main

import (
	"context"
	"echo-test/db"
	router "echo-test/route"
	"echo-test/services"
	"echo-test/store"
	customValidator "echo-test/validator"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func main() {
	e := echo.New()

	dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:golang_fun@localhost:5432/fun")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database :( ): %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	queries := db.New(dbpool)

	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Failed to open log file")
	}

	logger := zerolog.New(file)

	authorStore := store.NewAuthorStore(queries)

	authorService := services.NewAuthorService(&logger, authorStore)
	testService := services.NewTestService(&logger, queries)
	userService := services.NewUserService(&logger, queries)
	questionService := services.NewQuestionService(&logger, queries)

	e.Static("/", "assets")

	e.Validator = &customValidator.CustomValidator{
		Validator: validator.New(),
	}

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

	router.Init(e, &logger, authorService, testService, userService, questionService)

	e.Logger.Fatal(e.Start(":8000"))
}
