package services

import (
	"echo-test/db"

	"github.com/jackc/pgx/v5/pgtype"
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

func (q *QuestionService) GetOne(ctx echo.Context, id pgtype.UUID) (db.Question, error) {
	return q.queries.GetQuestion(ctx.Request().Context(), id)
}

func (q *QuestionService) GetQuestionTypesSuggestions(ctx echo.Context) ([]db.GetQuestionTypeSuggestionsRow, error) {
	return q.queries.GetQuestionTypeSuggestions(ctx.Request().Context())
}

func (q *QuestionService) Create(ctx echo.Context, params db.CreateQuestionParams) (db.Question, error) {
	return q.queries.CreateQuestion(ctx.Request().Context(), params)
}

func (q *QuestionService) CreateTx(ctx echo.Context, params db.CreateQuestionParams, trx *db.Queries) (db.Question, error) {
	return trx.CreateQuestion(ctx.Request().Context(), params)
}
