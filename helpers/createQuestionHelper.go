package helpers

import (
	"echo-test/components"
	questionComponents "echo-test/components/question"
	"echo-test/errors"
	"echo-test/services"
	"echo-test/utils"

	"github.com/labstack/echo/v4"
)

const OPEN_QUESTION = "49c95db4-5eb6-4968-968c-a22a453e1893"

type QuestionBody struct {
	QuestionName   string   `form:"question-name" validate:"required" json:"question-name"`
	QuestionType   string   `form:"question-type" validate:"required,uuid" json:"question-type"`
	QuestionPoints int      `form:"question-points" validate:"required,numeric,gt=0" json:"question-points"`
	QuestionText   string   `form:"question-text" validate:"required" json:"question-text"`
	AnswerTexts    []string `form:"answer[text]" validate:"required_with=answer[correct],dive,required" json:"answer-text"`
	AnswerCorrects []string `form:"answer[correct]" validate:"required_with=answer[text],dive,required,boolean" json:"answer-correct"`
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

func (h *CreateQuestionHelper) ValidateQuestion(c echo.Context) (*QuestionBody, *map[string]string, bool, error) {
	question := &QuestionBody{}
	errorsMap := make(map[string]string, 0)

	c.Request().ParseForm()

	if err := c.Bind(question); err != nil {
		return nil, nil, false, err
	}

	errors.ClearErrors(&errorsMap, c.Request())

	if err := c.Validate(question); err != nil {
		errors.PopulateErrorMap(&errorsMap, err, question)
		return nil, &errorsMap, false, nil
	}

	return question, nil, true, nil
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
	return questionComponents.FormBody(questionComponents.FormBodyProps{QuestionTypes: selectValues}).Render(c.Request().Context(), c.Response().Writer)
}
