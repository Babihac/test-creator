package helpers

import (
	"echo-test/components"
	questionComponents "echo-test/components/question"
	"echo-test/errors"
	"echo-test/services"
	"echo-test/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

const OPEN_QUESTION = "49c95db4-5eb6-4968-968c-a22a453e1893"

type QuestionBody struct {
	QuestionName   string   `form:"question-name" validate:"required"`
	QuestionType   string   `form:"question-type" validate:"required,uuid"`
	QuestionPoints int      `form:"question-points" validate:"required,numeric,gt=0"`
	QuestionText   string   `form:"question-text" validate:"required"`
	AnswerTexts    []string `form:"answer[text]" validate:"dive,required_with=answer[correct],gt=0"`
	AnswerCorrects []string `form:"answer[correct]" validate:"dive,boolean,required_with=answer[text],gt=0"`
}

func (q *QuestionBody) LteFieldError(fieldName string) string {
	panic("unimplemented")
}

type CreateQuestionHelper struct {
	questionService services.IQuestionService
}

func NewCreateQuestionHelper(qs services.IQuestionService) *CreateQuestionHelper {
	return &CreateQuestionHelper{
		questionService: qs,
	}
}

func (h *CreateQuestionHelper) ValidateQuestion(c echo.Context) (*QuestionBody, *map[string]string, error) {
	question := &QuestionBody{}
	errorsMap := make(map[string]string, 0)

	if err := c.Bind(question); err != nil {
		return nil, nil, err
	}

	if err := c.Validate(question); err != nil {
		errors.PopulateErrorMap(&errorsMap, err, question)
	}

	if len(errorsMap) != 0 {
		fmt.Printf("%v\n", errorsMap)
		return nil, &errorsMap, nil
	}

	return question, nil, nil
}

func (h *CreateQuestionHelper) PrepareQuestionForm(c echo.Context, errorsMap map[string]string) error {
	selectValues := make([]components.SelectValues, 0)
	questionTypes, err := h.questionService.GetQuestionTypesSuggestions(c)

	if err != nil {
		return utils.SendServerErrorNotification(c)
	}

	for _, questionType := range questionTypes {
		selectValues = append(selectValues, components.SelectValues{Value: questionType.Value, Label: questionType.Label})
	}
	return questionComponents.FormBody(questionComponents.FormBodyProps{QuestionTypes: selectValues, Errors: errorsMap}).Render(c.Request().Context(), c.Response().Writer)
}
