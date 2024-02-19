package services

import (
	"echo-test/db"

	"github.com/jackc/pgx/v5/pgtype"
	echo "github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type TestService struct {
	logger  *zerolog.Logger
	queries *db.Queries
}

// Create implements ITestService.
func (t *TestService) Create(ctx echo.Context, params db.CreateTestParams) (db.Test, error) {
	return t.queries.CreateTest(ctx.Request().Context(), params)
}

// Delete implements ITestService.
func (*TestService) Delete(ctx echo.Context, id int64) error {
	panic("unimplemented")
}

// Get implements ITestService.
func (t *TestService) Get(ctx echo.Context) ([]db.ListTestsWithTeacherRow, error) {
	return t.queries.ListTestsWithTeacher(ctx.Request().Context())
}

// GetOne implements ITestService.
func (t *TestService) GetOne(ctx echo.Context, id pgtype.UUID) (db.Test, error) {
	return t.queries.GetTest(ctx.Request().Context(), id)
}

// Update implements ITestService.
func (*TestService) Update(ctx echo.Context, params db.UpdateTestParams) error {
	panic("unimplemented")
}

func NewTestService(logger *zerolog.Logger, queries *db.Queries) *TestService {
	return &TestService{
		queries: queries,
		logger:  logger,
	}
}
