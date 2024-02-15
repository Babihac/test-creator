package handlers

import (
	authors "echo-test/components"
	components "echo-test/components/author/index"
	newAuthor "echo-test/components/author/new"
	"echo-test/db"
	"echo-test/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type CreateAuthorBody struct {
	Name string `form:"name" validate:"required"`
	Bio  string `form:"bio"`
}

type AuthorHandler struct {
	log           *zerolog.Logger
	authorService services.IAuthorService
}

func NewAuthor(log *zerolog.Logger, authorService services.IAuthorService) *AuthorHandler {
	return &AuthorHandler{
		log:           log,
		authorService: authorService,
	}
}

func (a *AuthorHandler) Serve(e *echo.Echo) {
	group := e.Group("/author")

	group.GET("", a.get)
	group.GET("/new", a.newAuthor)
	group.GET("/:id", a.authorDetail)

	group.POST("", a.createAuthor)

}

func (a *AuthorHandler) get(c echo.Context) error {
	values, err := a.authorService.Get(c)

	if err != nil {
		a.log.Error().Msg("Error getting authors")
		c.String(http.StatusInternalServerError, err.Error())
		fmt.Println("ddoddodo")
		return err
	}

	comp := authors.Authors(values)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

func (a *AuthorHandler) authorDetail(c echo.Context) error {
	id := c.Param("id")
	num, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		a.log.Error().Msg(fmt.Sprintf("Author not found: %s\n", id))
		c.String(http.StatusInternalServerError, err.Error())
		return err
	}

	author, err := a.authorService.GetOne(num, c)

	if err != nil {
		a.log.Error().Msg(fmt.Sprintf("Error getting author with id: %s\n", id))
		c.String(http.StatusInternalServerError, err.Error())
		return err
	}

	comp := components.Author(author)

	return comp.Render(c.Request().Context(), c.Response().Writer)

}

func (a *AuthorHandler) createAuthor(c echo.Context) error {
	body := &CreateAuthorBody{}

	c.Request().ParseForm()
	if err := c.Bind(body); err != nil {
		a.log.Error().Msg("Error creating new author\n")
		return c.Redirect(http.StatusBadRequest, "/author")
	}
	errors := make(map[string]string)
	if err := c.Validate(body); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors[strings.ToLower(err.Field())] = err.Tag()
		}
		return newAuthor.NewAuthor(body.Name, body.Bio, errors).Render(c.Request().Context(), c.Response().Writer)
	}

	author, error := a.authorService.Create(c, db.CreateAuthorParams{
		Name: body.Name,
		Bio:  pgtype.Text{String: body.Bio, Valid: body.Name != ""},
	})

	if error != nil {
		a.log.Error().Msg("Error creating new author\n")
		return newAuthor.NewAuthor(body.Name, body.Bio, map[string]string{}).Render(c.Request().Context(), c.Response().Writer)
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("/author/%d", author.ID))
}

func (a *AuthorHandler) newAuthor(c echo.Context) error {
	return newAuthor.NewAuthor("", "", map[string]string{}).Render(c.Request().Context(), c.Response().Writer)
}
