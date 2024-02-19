package handlers

import (
	"context"
	"echo-test/components"
	testComponents "echo-test/components/test"
	contextValues "echo-test/context"
	"echo-test/db"
	"echo-test/helpers"
	testPages "echo-test/pages/test"
	"echo-test/services"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type TestHandler struct {
	logger           *zerolog.Logger
	testService      services.ITestService
	userService      services.IUserService
	createTestHelper *helpers.CreateTestHelper
}

func NewTestHandler(logger *zerolog.Logger, testService services.ITestService, userService services.IUserService) *TestHandler {
	return &TestHandler{
		logger:           logger,
		testService:      testService,
		userService:      userService,
		createTestHelper: helpers.NewCreateTestHelper(testService, userService),
	}
}

func (t *TestHandler) Serve(echo *echo.Echo) {
	group := echo.Group("/test")

	group.GET("", t.Index)
	group.GET("/new", t.new)
	group.GET("/not-found", t.NotFound)

	group.POST("/new", t.createTest)

	group.GET("/new/step1", t.createTestStepOne)
	group.GET("/new/step2", t.createTestStepTwo)
	group.GET("/new/step3", t.createTestStepThree)
	group.GET("/new/step4", t.new)

	group.GET("/:id", t.Detail)

	group.POST("/new/step1", t.createTestStepOne)
	group.POST("/new/step2", t.createTestStepTwo)
	group.POST("/new/step3", t.createTestStepThree)
	group.POST("/new/step4", t.new)

}

func (t *TestHandler) Index(c echo.Context) error {
	tests, err := t.testService.Get(c)

	if err != nil {
		return c.Redirect(http.StatusFound, "/test/not-found")
	}

	return testPages.TestsPage(tests).Render(c.Request().Context(), c.Response().Writer)
}

func (t *TestHandler) NotFound(c echo.Context) error {
	id := c.Param("id")

	return testComponents.NotFound(testComponents.NotFoundProps{Id: id}).Render(c.Request().Context(), c.Response().Writer)
}

func (t *TestHandler) Detail(c echo.Context) error {
	id := c.Param("id")

	uuid := helpers.StringToUUID(id)

	test, err := t.testService.GetOne(c, uuid)

	if err != nil {
		return c.Redirect(http.StatusFound, "/test/not-found")
	}

	return testPages.TestDetailPage(testComponents.TestDetailProps{Name: test.Name}).Render(c.Request().Context(), c.Response().Writer)
}

func (t *TestHandler) new(c echo.Context) error {
	selectValues := make([]components.SelectValues, 0)
	userSuggestions, err := t.userService.GetUserSuggestions(c)

	if err != nil {
		c.Redirect(http.StatusFound, "/test/not-found")
	}

	for _, useruserSuggestion := range userSuggestions {
		uuid, err := helpers.ParseUUID(useruserSuggestion.Value)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error parsing UUID value")
		}

		selectValues = append(selectValues, components.SelectValues{
			Label: useruserSuggestion.Label,
			Value: uuid,
		})
	}

	key := contextValues.USER_SUGGESTIONS
	ctx := context.WithValue(c.Request().Context(), key, selectValues)

	var defaultTeacherId string

	if len(selectValues) > 0 {
		defaultTeacherId = selectValues[0].Value
	}

	return testPages.CreateTestPage(testComponents.CreateTestProps{DefaultTeacherId: defaultTeacherId}).Render(ctx, c.Response().Writer)
}

func (t *TestHandler) createTest(c echo.Context) error {
	values, ok, err := t.createTestHelper.ValidateForm(c)

	if !ok || err != nil {
		fmt.Printf("%t, %v\n", ok, err)
		c.Response().Header().Add("HX-Reswap", "none")
		return components.Notification("end", "top", "alert-error", "Something went wrong").Render(c.Request().Context(), c.Response().Writer)
	}

	newTest, err := t.testService.Create(c, db.CreateTestParams{
		Name:      values.TestName,
		TeacherID: helpers.StringToUUID(values.TeacherId),
		Duration: pgtype.Interval{
			Microseconds: time.Duration.Microseconds(time.Minute * time.Duration(values.Duration)),
			Valid:        time.Duration.Microseconds(time.Minute*time.Duration(values.Duration)) != 0,
		},
		MaxPoints: int32(values.MaxScore),
	})

	if err != nil {
		c.Response().Header().Add("HX-Reswap", "none")
		return components.Notification("end", "top", "alert-error", "Something went wrong").Render(c.Request().Context(), c.Response().Writer)
	}
	uuid, _ := helpers.ParseUUID(newTest.ID)
	c.Response().Header().Add("HX-Push-Url", fmt.Sprintf("/test/%s", uuid))
	c.Response().Header().Add("HX-Reswap", "OuterHTML")
	c.Response().Header().Add("HX-Retarget", "#create-test-component")

	components.Notification("end", "top", "alert-success", fmt.Sprintf("Test %s created successfuly", newTest.Name)).Render(c.Request().Context(), c.Response().Writer)
	return testComponents.TestDetail(testComponents.TestDetailProps{Name: newTest.Name}).Render(c.Request().Context(), c.Response().Writer)

}

func (t *TestHandler) createTestStepOne(c echo.Context) error {
	if c.Request().Header.Get("Hx-Request") != "true" {
		return c.Redirect(http.StatusFound, "/test/new")
	}

	c.Response().Header().Add("HX-Push-Url", "/test/new/step1")
	return t.createTestHelper.PrepareStepOne(c, map[string]string{})
}

func (t *TestHandler) createTestStepTwo(c echo.Context) error {
	if c.Request().Header.Get("Hx-Request") != "true" {
		return c.Redirect(http.StatusFound, "/test/new")
	}

	ok, errorsMap, err := t.createTestHelper.ValidateStepOne(c)

	if err != nil {
		c.Response().Header().Add("HX-Reswap", "none")
		return components.Notification("end", "top", "alert-error", "Something went wrong").Render(c.Request().Context(), c.Response().Writer)
	}

	if !ok {
		return t.createTestHelper.PrepareStepOne(c, *errorsMap)
	}

	c.Response().Header().Add("HX-Push-Url", "/test/new/step2")
	return t.createTestHelper.PrepareStepTwo(c, map[string]string{})
}

func (t *TestHandler) createTestStepThree(c echo.Context) error {
	if c.Request().Header.Get("Hx-Request") != "true" {
		return c.Redirect(http.StatusFound, "/test/new")
	}

	ok, errorsMap, err := t.createTestHelper.ValidateStepTwo(c)

	if err != nil {
		c.Response().Header().Add("HX-Reswap", "none")
		return components.Notification("end", "top", "alert-error", "Something went wrong").Render(c.Request().Context(), c.Response().Writer)
	}

	if !ok {
		return t.createTestHelper.PrepareStepTwo(c, *errorsMap)
	}

	return t.createTestHelper.PrepareStepThree(c, map[string]string{})
}
