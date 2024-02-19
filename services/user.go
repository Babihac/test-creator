package services

import (
	"echo-test/db"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type UserService struct {
	logger  *zerolog.Logger
	queries *db.Queries
}

// Create implements IUserService.
func (*UserService) Create(ctx echo.Context, params db.CreateTestParams) (db.User, error) {
	panic("unimplemented")
}

// Delete implements IUserService.
func (*UserService) Delete(ctx echo.Context, id int64) error {
	panic("unimplemented")
}

// Get implements IUserService.
func (*UserService) Get(ctx echo.Context) ([]db.ListTestsWithTeacherRow, error) {
	panic("unimplemented")
}

// GetOne implements IUserService.
func (*UserService) GetOne(id int64, ctx echo.Context) (db.Test, error) {
	panic("unimplemented")
}

// Update implements IUserService.
func (*UserService) Update(ctx echo.Context, params db.UpdateTestParams) error {
	panic("unimplemented")
}

func NewUserService(logger *zerolog.Logger, queries *db.Queries) *UserService {
	return &UserService{
		logger:  logger,
		queries: queries,
	}
}

func (u *UserService) GetUserSuggestions(ctx echo.Context) ([]db.GetUserSuggestionsRow, error) {
	return u.queries.GetUserSuggestions(ctx.Request().Context())
}
