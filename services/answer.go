package services

import (
	"echo-test/db"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type AnswerService struct {
	logger  *zerolog.Logger
	queries *db.Queries
}

func NewAnswerService(logger *zerolog.Logger, queries *db.Queries) *AnswerService {
	return &AnswerService{
		logger:  logger,
		queries: queries,
	}
}

func (a *AnswerService) Get(ctx echo.Context) ([]db.Answer, error) {
	return a.queries.ListAnswers(ctx.Request().Context())
}

func (a *AnswerService) Create(ctx echo.Context, params db.CreateAnswerParams) (db.Answer, error) {
	return a.queries.CreateAnswer(ctx.Request().Context(), params)
}
