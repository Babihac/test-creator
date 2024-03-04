package services

import (
	"echo-test/db"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type IAuthorService interface {
	Get(ctx echo.Context) ([]db.Author, error)
	GetOne(id int64, ctx echo.Context) (db.Author, error)
	Create(ctx echo.Context, params db.CreateAuthorParams) (db.Author, error)
	Delete(ctx echo.Context, id int64) error
	Update(ctx echo.Context, params db.UpdateAuthorParams) error
}

type ITestService interface {
	Get(ctx echo.Context) ([]db.ListTestsWithTeacherRow, error)
	GetOne(ctx echo.Context, id pgtype.UUID) (db.Test, error)
	Create(ctx echo.Context, params db.CreateTestParams) (db.Test, error)
	Delete(ctx echo.Context, id int64) error
	Update(ctx echo.Context, params db.UpdateTestParams) error
}

type IUserService interface {
	Get(ctx echo.Context) ([]db.ListTestsWithTeacherRow, error)
	GetOne(id int64, ctx echo.Context) (db.Test, error)
	Create(ctx echo.Context, params db.CreateTestParams) (db.User, error)
	Delete(ctx echo.Context, id int64) error
	Update(ctx echo.Context, params db.UpdateTestParams) error
	GetUserSuggestions(ctx echo.Context) ([]db.GetUserSuggestionsRow, error)
}

type IQuestionService interface {
	Get(ctx echo.Context) ([]db.Question, error)
	GetQuestionSuggestions(ctx echo.Context) ([]db.ListQuestionsWithTypeRow, error)
	GetQuestionTypesSuggestions(ctx echo.Context) ([]db.GetQuestionTypeSuggestionsRow, error)
	Create(ctx echo.Context, params db.CreateQuestionParams) (db.Question, error)
	CreateTx(ctx echo.Context, params db.CreateQuestionParams, trx *db.Queries) (db.Question, error)
	GetOne(ctx echo.Context, id pgtype.UUID) (db.Question, error)
}

type IAnswerService interface {
	Get(ctx echo.Context) ([]db.Answer, error)
	Create(ctx echo.Context, params db.CreateAnswerParams) (db.Answer, error)
	CreateTx(ctx echo.Context, params db.CreateAnswerParams, trx *db.Queries) (db.Answer, error)
}
