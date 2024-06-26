package router

import (
	"echo-test/db"
	"echo-test/handlers"
	"echo-test/services"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func Init(
	e *echo.Echo,
	logger *zerolog.Logger,
	authorService services.IAuthorService,
	testService services.ITestService,
	userService services.IUserService,
	questionService services.IQuestionService,
	answerService services.IAnswerService,
	db *pgxpool.Pool,
	queries *db.Queries,
) {
	authorHandler := handlers.NewAuthor(logger, authorService)
	testHandler := handlers.NewTestHandler(logger, testService, userService, questionService)
	questionHandler := handlers.NewQuestionHandler(logger, questionService, answerService, db, queries)
	authorHandler.Serve(e)
	testHandler.Serve(e)
	questionHandler.Serve(e)
}
