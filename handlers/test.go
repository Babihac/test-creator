package handlers

import (
	"context"
	"echo-test/components"
	contextValues "echo-test/context"
	testPages "echo-test/pages/test"
	"echo-test/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type TestHandler struct {
	logger      *zerolog.Logger
	testService services.ITestService
	userService services.IUserService
}

func NewTestHandler(logger *zerolog.Logger, testService services.ITestService, userService services.IUserService) *TestHandler {
	return &TestHandler{
		logger:      logger,
		testService: testService,
		userService: userService,
	}
}

func (t *TestHandler) Serve(echo *echo.Echo) {
	group := echo.Group("/test")

	group.GET("", t.Index)
	group.GET("/new", t.new)
}

func (t *TestHandler) Index(c echo.Context) error {
	tests, err := t.testService.Get(c)

	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/500")
	}

	return testPages.TestsPage(tests).Render(c.Request().Context(), c.Response().Writer)
}

func (t TestHandler) new(c echo.Context) error {
	selectValues := make([]components.SelectValues, 0)
	userSuggestions, err := t.userService.GetUserSuggestions(c)

	if err != nil {
		c.Redirect(http.StatusInternalServerError, "/500")
	}

	for _, useruserSuggestion := range userSuggestions {
		selectValues = append(selectValues, components.SelectValues{
			Label: useruserSuggestion.Label,
			Value: string(useruserSuggestion.Value.Bytes[:]),
		})
	}

	key := contextValues.USER_SUGGESTIONS
	ctx := context.WithValue(c.Request().Context(), key, selectValues)

	return testPages.CreateTestPage().Render(ctx, c.Response().Writer)
}
