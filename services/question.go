package services

import (
	"echo-test/db"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type QuestionService struct {
	logger  *zerolog.Logger
	queries *db.Queries
}

func NewQuestionService(logger *zerolog.Logger, queries *db.Queries) *QuestionService {
	return &QuestionService{
		logger:  logger,
		queries: queries,
	}
}

func (q *QuestionService) Get(ctx echo.Context) ([]db.Question, error) {
	return q.queries.ListQuestions(ctx.Request().Context())
}

func (q *QuestionService) GetQuestionTypesSuggestions(ctx echo.Context) ([]db.GetQuestionTypeSuggestionsRow, error) {
	return q.queries.GetQuestionTypeSuggestions(ctx.Request().Context())
}
