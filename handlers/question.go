package handlers

import (
	"echo-test/components"
	questionComponents "echo-test/components/question"
	"echo-test/db"
	"echo-test/helpers"
	"echo-test/pages"
	questionPages "echo-test/pages/question"
	"echo-test/services"
	"echo-test/utils"
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
	queries              *db.Queries
}

func NewQuestionHandler(logger *zerolog.Logger, questionService services.IQuestionService, answerService services.IAnswerService, db *pgxpool.Pool, queries *db.Queries) *QuestionHandler {
	return &QuestionHandler{
		logger:               logger,
		questionService:      questionService,
		createQuestionHelper: helpers.NewCreateQuestionHelper(questionService),
		answerService:        answerService,
		db:                   db,
		queries:              queries,
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

	if c.Request().Header.Get("Hx-Request") == "true" {
		return questionComponents.Questions().Render(c.Request().Context(), c.Response().Writer)
	}

	return questionPages.QuestionsPage(questionPages.QuestionsPageProps{Questions: questions}).Render(c.Request().Context(), c.Response().Writer)
}

func (q *QuestionHandler) Detail(c echo.Context) error {
	id := c.Param("id")

	uuid := utils.StringToUUID(id)

	question, error := q.questionService.GetOne(c, uuid)

	if error != nil {
		return pages.NotFound(pages.NotFoundProps{Message: "This question was not found."}).Render(c.Request().Context(), c.Response().Writer)
	}

	if c.Request().Header.Get("Hx-Request") == "true" {
		return questionComponents.QuestionDetail(questionComponents.QuestioNDetailProps{Question: question}).Render(c.Request().Context(), c.Response().Writer)
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

	if c.Request().Header.Get("Hx-Request") == "true" {
		return questionComponents.CreateQuestion(questionComponents.CreateQuestionProps{QuestionTypes: selectValues}).Render(c.Request().Context(), c.Response().Writer)
	}

	return questionPages.CreateQuestionPage(questionPages.CreateQuestionPageProps{QuestionTypes: selectValues}).Render(c.Request().Context(), c.Response().Writer)
}

func (q *QuestionHandler) addAnswer(c echo.Context) error {
	index := c.QueryParam("questions-count")

	num, err := strconv.ParseInt(index, 10, 32)

	if err != nil {
		return utils.SendServerErrorNotification(c)
	}

	return questionComponents.AnswerDefinitionRow(num).Render(c.Request().Context(), c.Response().Writer)
}

func (q *QuestionHandler) createQuestion(c echo.Context) error {
	quetionBody, errorsMap, ok, err := q.createQuestionHelper.ValidateQuestion(c)

	if err != nil {
		return utils.SendServerErrorNotification(c)
	}

	if !ok {
		return components.InputErrors(components.InputErrorsProps{Errors: *errorsMap}).Render(c.Request().Context(), c.Response().Writer)
	}

	tx, err := q.db.Begin(c.Request().Context())
	defer tx.Rollback(c.Request().Context())

	if err != nil {
		return utils.SendServerErrorNotification(c)
	}

	qtx := q.queries.WithTx(tx)

	question, err := q.questionService.CreateTx(c, db.CreateQuestionParams{
		QuestionType: utils.StringToUUID(quetionBody.QuestionType),
		Points:       int32(quetionBody.QuestionPoints),
		Name:         quetionBody.QuestionName,
		QuestionText: quetionBody.QuestionText,
	}, qtx)

	if err != nil {
		return utils.SendServerErrorNotification(c)
	}

	for i, answerText := range quetionBody.AnswerTexts {
		correct, _ := strconv.ParseBool(quetionBody.AnswerCorrects[i]) // already validated

		_, err := q.answerService.CreateTx(c, db.CreateAnswerParams{
			Value:      answerText,
			Correct:    correct,
			QuestionID: question.ID,
		}, qtx)

		if err != nil {
			return utils.SendServerErrorNotification(c)
		}

	}

	tx.Commit(c.Request().Context())

	uuid, _ := utils.ParseUUID(question.ID)

	redirectParams := utils.NewDefaultRedirectParams(utils.DefaultRedirectHeaders{Url: fmt.Sprintf("/question/%s", uuid), Swap: "OuterHTML", Target: "#create-question-component"},
		fmt.Sprintf("Question %s created successfuly", quetionBody.QuestionName))

	utils.SendSuccessNotification(c, *redirectParams)
	return questionComponents.QuestionDetail(questionComponents.QuestioNDetailProps{Question: question}).Render(c.Request().Context(), c.Response().Writer)
}
