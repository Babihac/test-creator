package handlers

import (
	"echo-test/components"
	questionComponents "echo-test/components/question"
	"echo-test/db"
	"echo-test/helpers"
	"echo-test/pages"
	questionPages "echo-test/pages/question"
	"echo-test/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type QuestionHandler struct {
	logger               *zerolog.Logger
	questionService      services.IQuestionService
	answerService        services.IAnswerService
	createQuestionHelper *helpers.CreateQuestionHelper
	db                   *pgxpool.Pool
}

func NewQuestionHandler(logger *zerolog.Logger, questionService services.IQuestionService, answerService services.IAnswerService, db *pgxpool.Pool) *QuestionHandler {
	return &QuestionHandler{
		logger:               logger,
		questionService:      questionService,
		createQuestionHelper: helpers.NewCreateQuestionHelper(questionService),
		answerService:        answerService,
		db:                   db,
	}
}

func (q *QuestionHandler) Serve(e *echo.Echo) {
	group := e.Group("/question")

	group.GET("", q.index)
	group.GET("/new", q.new)
	group.GET("/add-answer", q.addAnswer)
	group.GET("/:id", q.Detail)

	group.POST("", q.createQuestion)
}

func (q *QuestionHandler) index(c echo.Context) error {
	questions, err := q.questionService.Get(c)

	if err != nil {
		return c.Redirect(http.StatusFound, "/question/not-found")
	}

	return questionPages.QuestionsPage(questionPages.QuestionsPageProps{Questions: questions}).Render(c.Request().Context(), c.Response().Writer)
}

func (q *QuestionHandler) Detail(c echo.Context) error {
	id := c.Param("id")

	uuid := helpers.StringToUUID(id)

	question, error := q.questionService.GetOne(c, uuid)

	if error != nil {
		return pages.NotFound(pages.NotFoundProps{Message: "This question was not found."}).Render(c.Request().Context(), c.Response().Writer)
	}

	return questionPages.QuestionDetail(questionPages.QuestionPageProps{Question: question}).Render(c.Request().Context(), c.Response().Writer)
}

func (q *QuestionHandler) new(c echo.Context) error {
	selectValues := make([]components.SelectValues, 0)
	questionTypes, err := q.questionService.GetQuestionTypesSuggestions(c)

	if err != nil {
		return c.Redirect(http.StatusFound, "/500")
	}

	for _, questionType := range questionTypes {
		selectValues = append(selectValues, components.SelectValues{Value: questionType.Value, Label: questionType.Label})
	}

	return questionPages.CreateQuestionPage(questionPages.CreateQuestionPageProps{QuestionTypes: selectValues}).Render(c.Request().Context(), c.Response().Writer)
}

func (q *QuestionHandler) addAnswer(c echo.Context) error {
	return questionComponents.AnswerDefinitionRow().Render(c.Request().Context(), c.Response().Writer)
}

func (q *QuestionHandler) createQuestion(c echo.Context) error {
	quetionBody, errorsMap, err := q.createQuestionHelper.ValidateQuestion(c)

	if err != nil {
		return helpers.SendServerErrorNotification(c)
	}

	if errorsMap != nil {
		return q.createQuestionHelper.PrepareQuestionForm(c, *errorsMap)
	}

	fmt.Println(quetionBody)

	question, err := q.questionService.Create(c, db.CreateQuestionParams{
		QuestionType: helpers.StringToUUID(quetionBody.QuestionType),
		Points:       int32(quetionBody.QuestionPoints),
		Name:         quetionBody.QuestionName,
		QuestionText: quetionBody.QuestionText,
	})

	for i, answerText := range quetionBody.AnswerTexts {
		correct, _ := strconv.ParseBool(quetionBody.AnswerCorrects[i]) // already validated

		//tx, err := q.db.Begin(c.Request().Context())

		answer, err := q.answerService.Create(c, db.CreateAnswerParams{
			Value:      answerText,
			Correct:    correct,
			QuestionID: question.ID,
		})

		if err != nil {
			return helpers.SendServerErrorNotification(c)
		}

		fmt.Printf("New answer: %v", answer)
	}

	if err != nil {
		return helpers.SendServerErrorNotification(c)
	}
	uuid, _ := helpers.ParseUUID(question.ID)

	redirectParams := helpers.NewDefaultRedirectParams(helpers.DefaultRedirectHeaders{Url: fmt.Sprintf("/question/%s", uuid), Swap: "OuterHTML", Target: "#create-question-component"},
		fmt.Sprintf("Question %s created successfuly", quetionBody.QuestionName))

	helpers.SendSuccessNotification(c, *redirectParams)
	return questionComponents.QuestionDetail(questionComponents.QuestioNDetailProps{Question: question}).Render(c.Request().Context(), c.Response().Writer)
}
