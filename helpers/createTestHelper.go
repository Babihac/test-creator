package helpers

import (
	"context"
	"echo-test/components"
	testComponents "echo-test/components/test"
	contextValues "echo-test/context"
	"echo-test/errors"
	"echo-test/services"
	"echo-test/types"
	"echo-test/utils"

	"github.com/labstack/echo/v4"
)

type StepOneForm struct {
	TestName    string `form:"test-name" validate:"required" json:"test-name"`
	TeacherId   string `form:"teacher-id" validate:"required,uuid" json:"teacher-id"`
	Duration    int    `form:"test-duration" validate:"required,gte=15,lte=180,numeric" json:"duration"`
	CurrentStep int    `form:"steps-loaded"`
}

func (s *StepOneForm) LteFieldError(fieldName string) string {
	panic("unimplemented")
}

type StepTwoForm struct {
	MaxScore         int `form:"max-score" validate:"required,gte=0,numeric" json:"max-score"`
	MinRequiredScore int `form:"min-required-score" validate:"required,numeric,gte=0,ltefield=MaxScore" json:"min-required-score"`
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
	testService     services.ITestService
	userService     services.IUserService
	questionService services.IQuestionService
	steps           []string
}

func NewCreateTestHelper(ts services.ITestService, us services.IUserService, qs services.IQuestionService) *CreateTestHelper {
	return &CreateTestHelper{
		testService:     ts,
		userService:     us,
		questionService: qs,
		steps:           []string{"Main info", "Scoring", "Questions", "Summary"},
	}
}

func (h *CreateTestHelper) PrepareStepOne(c echo.Context, errorsMap map[string]string) error {
	selectValues := make([]components.SelectValues, 0)
	userSuggestions, err := h.userService.GetUserSuggestions(c)

	if err != nil {
		return utils.SendServerErrorNotification(c)
	}

	for _, useruserSuggestion := range userSuggestions {

		selectValues = append(selectValues, components.SelectValues{
			Label: useruserSuggestion.Label,
			Value: useruserSuggestion.Value,
		})
	}

	key := contextValues.USER_SUGGESTIONS
	ctx := context.WithValue(c.Request().Context(), key, selectValues)

	c.Response().Header().Add("HX-Retarget", "#create-test-form-step-1")

	testComponents.CreateTestStepOne(testComponents.CreateTestStepOneProps{}).Render(ctx, c.Response().Writer)
	return testComponents.Stepper(testComponents.StepperProps{CurrentStep: 1, Steps: h.steps, Oob: true}).Render(ctx, c.Response().Writer)

}

func (h *CreateTestHelper) HandleStepLoaded(c echo.Context, step int) error {
	return testComponents.Stepper(testComponents.StepperProps{CurrentStep: step, Steps: h.steps, Oob: true}).Render(c.Request().Context(), c.Response().Writer)
}

func (h *CreateTestHelper) PrepareStepTwo(c echo.Context, errorsMap map[string]string) error {
	c.Response().Header().Add("HX-Retarget", "#create-test-form-step-2")
	testComponents.CreateTestStepScore(testComponents.CreateTestStepScoreProps{Erors: errorsMap}).Render(c.Request().Context(), c.Response().Writer)
	testComponents.Stepper(testComponents.StepperProps{CurrentStep: 2, Steps: h.steps, Oob: true}).Render(c.Request().Context(), c.Response().Writer)
	return nil
}

func (h *CreateTestHelper) PrepareStepThree(c echo.Context, errorsMap map[string]string) error {

	questions, err := h.questionService.GetQuestionSuggestions(c)

	if err != nil {
		return utils.SendServerErrorNotification(c)
	}

	c.Response().Header().Add("HX-Retarget", "#create-test-form-step-3")
	testComponents.StepQuestions(errorsMap, questions).Render(c.Request().Context(), c.Response().Writer)
	testComponents.Stepper(testComponents.StepperProps{CurrentStep: 3, Steps: h.steps, Oob: true}).Render(c.Request().Context(), c.Response().Writer)
	return nil
}

func (h *CreateTestHelper) PrepareStepFour(c echo.Context) error {
	c.Response().Header().Add("HX-Retarget", "#create-test-form-step-4")
	testComponents.StepSummary().Render(c.Request().Context(), c.Response().Writer)
	testComponents.Stepper(testComponents.StepperProps{CurrentStep: 4, Steps: h.steps, Oob: true}).Render(c.Request().Context(), c.Response().Writer)

	return nil
}

func (h *CreateTestHelper) ValidateStepOne(c echo.Context) (bool, *map[string]string, error) {

	stepOneValues := &StepOneForm{}

	c.Request().ParseForm()

	return h.validateStep(c, stepOneValues)
}

func (h *CreateTestHelper) ValidateStepTwo(c echo.Context) (bool, *map[string]string, error) {
	stepTwoValues := &StepTwoForm{}

	return h.validateStep(c, stepTwoValues)
}

func (h *CreateTestHelper) validateStep(c echo.Context, form types.IForm) (bool, *map[string]string, error) {
	if err := c.Bind(form); err != nil {
		return false, nil, err
	}

	errorsMap := make(map[string]string, 0)

	if err := c.Validate(form); err != nil {
		errors.PopulateErrorMap(&errorsMap, err, form)
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

	if err := c.Validate(formValues); err != nil {
		return nil, false, err
	}

	return formValues, true, nil
}

func (h *CreateTestHelper) InputErrors(c echo.Context, errors map[string]string) error {
	return components.InputErrors(components.InputErrorsProps{Errors: errors}).Render(c.Request().Context(), c.Response().Writer)
}
