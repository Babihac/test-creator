package handlers

import (
	questionPages "echo-test/pages/question"
	"echo-test/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type QuestionHandler struct {
	logger          *zerolog.Logger
	questionService services.IQuestionService
}

func NewQuestionHandler(logger *zerolog.Logger, questionService services.IQuestionService) *QuestionHandler {
	return &QuestionHandler{
		logger:          logger,
		questionService: questionService,
	}
}

func (q *QuestionHandler) Serve(e *echo.Echo) {
	group := e.Group("/question")

	group.GET("", q.index)
	group.GET("/new", q.new)
}

func (q *QuestionHandler) index(c echo.Context) error {
	questions, err := q.questionService.Get(c)

	if err != nil {
		return c.Redirect(http.StatusFound, "/question/not-found")
	}

	return questionPages.QuestionsPage(questionPages.QuestionsPageProps{Questions: questions}).Render(c.Request().Context(), c.Response().Writer)
}

func (q *QuestionHandler) new(c echo.Context) error {

	questionTypes, err := q.questionService.GetQuestionTypesSuggestions(c)

	if err != nil {
		return c.Redirect(http.StatusFound, "/500")
	}

	return questionPages.CreateQuestionPage(questionPages.CreateQuestionPageProps{QuestionTypes: questionTypes}).Render(c.Request().Context(), c.Response().Writer)
}
