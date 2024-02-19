package helpers

import (
	"context"
	"echo-test/components"
	testComponents "echo-test/components/test"
	contextValues "echo-test/context"
	"echo-test/errors"
	"echo-test/services"
	"echo-test/types"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type StepOneForm struct {
	TestName    string `form:"test-name" validate:"required"`
	TeacherId   string `form:"teacher-id" validate:"required,uuid"`
	Duration    int    `form:"test-duration" validate:"required,gte=15,lte=180,numeric"`
	CurrentStep int    `form:"steps-loaded"`
}

func (s *StepOneForm) LteFieldError(fieldName string) string {
	panic("unimplemented")
}

type StepTwoForm struct {
	MaxScore         int `form:"max-score" validate:"required,gte=0,numeric"`
	MinRequiredScore int `form:"min-required-score" validate:"required,numeric,gte=0,ltefield=MaxScore"`
	CurrentStep      int `form:"steps-loaded"`
}

type CreateTestForm struct {
	TestName         string `form:"test-name" validate:"required"`
	TeacherId        string `form:"teacher-id" validate:"required,uuid"`
	Duration         int    `form:"test-duration" validate:"required,gte=15,lte=180,numeric"`
	MaxScore         int    `form:"max-score" validate:"required,gte=0,numeric"`
	MinRequiredScore int    `form:"min-required-score" validate:"required,numeric,gte=0,ltefield=MaxScore"`
}

func (s *StepTwoForm) LteFieldError(fieldName string) string {
	return "Cannot be greater than Maximal score"
}

type CreateTestHelper struct {
	testService services.ITestService
	userService services.IUserService
	steps       []string
}

func NewCreateTestHelper(ts services.ITestService, us services.IUserService) *CreateTestHelper {
	return &CreateTestHelper{
		testService: ts,
		userService: us,
		steps:       []string{"Main info", "Scoring requirements", "Questin definition", "Summary"},
	}
}

func (h *CreateTestHelper) PrepareStepOne(c echo.Context, errorsMap map[string]string) error {
	selectValues := make([]components.SelectValues, 0)
	userSuggestions, err := h.userService.GetUserSuggestions(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	for _, useruserSuggestion := range userSuggestions {
		uuid, err := ParseUUID(useruserSuggestion.Value)

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

	testComponents.CreateTestStepOne(testComponents.CreateTestStepOneProps{Erors: errorsMap}).Render(ctx, c.Response().Writer)
	return testComponents.Stepper(testComponents.StepperProps{CurrentStep: 1, Steps: h.steps, Oob: true}).Render(ctx, c.Response().Writer)

}

func (h *CreateTestHelper) HandleStepLoaded(c echo.Context, step int) error {
	return testComponents.Stepper(testComponents.StepperProps{CurrentStep: step, Steps: h.steps, Oob: true}).Render(c.Request().Context(), c.Response().Writer)
}

func (h *CreateTestHelper) PrepareStepTwo(c echo.Context, errorsMap map[string]string) error {
	testComponents.CreateTestStepScore(testComponents.CreateTestStepScoreProps{Erors: errorsMap}).Render(c.Request().Context(), c.Response().Writer)
	testComponents.Stepper(testComponents.StepperProps{CurrentStep: 2, Steps: h.steps, Oob: true}).Render(c.Request().Context(), c.Response().Writer)
	return nil
}

func (h *CreateTestHelper) PrepareStepThree(c echo.Context, errorsMap map[string]string) error {
	panic("haha")
}

func (h *CreateTestHelper) ValidateStepOne(c echo.Context) (bool, *map[string]string, error) {

	stepOneValues := &StepOneForm{}

	fmt.Printf("Current step: %d\n", stepOneValues.CurrentStep)

	return h.validateStep(c, stepOneValues)
}

func (h *CreateTestHelper) ValidateStepTwo(c echo.Context) (bool, *map[string]string, error) {
	stepTwoValues := &StepTwoForm{}

	fmt.Printf("Current step: %d\n", stepTwoValues.CurrentStep)

	return h.validateStep(c, stepTwoValues)
}

func (h *CreateTestHelper) validateStep(c echo.Context, step types.IStep) (bool, *map[string]string, error) {
	if err := c.Bind(step); err != nil {
		return false, nil, err
	}

	fmt.Printf("Form values: %v\n", step)

	errorsMap := make(map[string]string, 0)

	if err := c.Validate(step); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("%v\n", err.Tag())
			errorsMap[err.Field()] = errors.Message(err, step)
		}
	}

	if len(errorsMap) != 0 {
		return false, &errorsMap, nil
	}

	return true, nil, nil
}

func (h *CreateTestHelper) ValidateForm(c echo.Context) (*CreateTestForm, bool, error) {
	formValues := &CreateTestForm{}

	if err := c.Bind(formValues); err != nil {
		return nil, false, err
	}

	fmt.Printf("Final form values: %v\n", formValues)

	if err := c.Validate(formValues); err != nil {
		return nil, false, err
	}

	return formValues, true, nil
}
