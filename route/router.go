package router

import (
	"echo-test/handlers"
	"echo-test/services"

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
) {
	authorHandler := handlers.NewAuthor(logger, authorService)
	testHandler := handlers.NewTestHandler(logger, testService, userService)
	questionHandler := handlers.NewQuestionHandler(logger, questionService)
	authorHandler.Serve(e)
	testHandler.Serve(e)
	questionHandler.Serve(e)
}
