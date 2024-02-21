package handlers

import (
	"echo-test/components"
	authorComponents "echo-test/components/author"
	"echo-test/db"
	"echo-test/errors"
	"echo-test/pages"
	"echo-test/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type CreateAuthorBody struct {
	Name string `form:"name" validate:"required"`
	Bio  string `form:"bio" validate:"min=5"`
}

// LteFieldError implements types.IStep.
func (*CreateAuthorBody) LteFieldError(fieldName string) string {
	panic("unimplemented")
}

type UpdateAuthorBody struct {
	Name string `form:"name" validate:"required"`
	Bio  string `form:"bio" validate:"min=5"`
	ID   int64  `param:"id" validate:"required"`
}

// LteFieldError implements types.IStep.
func (*UpdateAuthorBody) LteFieldError(fieldName string) string {
	panic("unimplemented")
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
	group.GET("/edit/:id", a.showEditAuthor)
	group.GET("/:id", a.authorDetail)

	group.POST("", a.createAuthor)

	group.PATCH("/:id", a.editAuthor)

	group.DELETE("/:id", a.deleteAuthor)

}

func (a *AuthorHandler) get(c echo.Context) error {
	values, err := a.authorService.Get(c)

	if err != nil {
		a.log.Error().Msg("Error getting authors")
		c.String(http.StatusInternalServerError, err.Error())
		return err
	}

	comp := authorComponents.Authors(values)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

func (a *AuthorHandler) authorDetail(c echo.Context) error {
	fmt.Println(c.Request().Header)

	id := c.Param("id")
	num, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		a.log.Error().Msg(fmt.Sprintf("Author not found: %s\n", id))
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Cannot find user with ID %s", id))
	}

	author, err := a.authorService.GetOne(num, c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Cannot find user with ID %s", id))
	}

	comp := pages.Author(author)

	return comp.Render(c.Request().Context(), c.Response().Writer)

}

func (a *AuthorHandler) showEditAuthor(c echo.Context) error {
	id := c.Param("id")
	num, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		a.log.Error().Msg(fmt.Sprintf("Author not found: %s\n", id))
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Cannot find user with ID %s", id))
	}

	author, err := a.authorService.GetOne(num, c)

	if err != nil {
		a.log.Error().Msg(fmt.Sprintf("Error getting author with id: %s\n", id))
		c.String(http.StatusInternalServerError, err.Error())
		return err
	}

	return pages.EditAuthorPage(author.Name, author.Bio.String, num, map[string]string{}).Render(c.Request().Context(), c.Response().Writer)
}

func (a *AuthorHandler) createAuthor(c echo.Context) error {
	body := &CreateAuthorBody{}

	c.Request().ParseForm()
	if err := c.Bind(body); err != nil {
		a.log.Error().Msg("Error creating new author\n")
		return c.Redirect(http.StatusBadRequest, "/author")
	}
	errorsMap := make(map[string]string)
	if err := c.Validate(body); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Tag(), err.Param())
			errorsMap[strings.ToLower(err.Field())] = errors.Message(err, body)
		}
		c.Response().Status = http.StatusBadRequest
		return authorComponents.NewAuthor(body.Name, body.Bio, errorsMap).Render(c.Request().Context(), c.Response().Writer)
	}

	author, error := a.authorService.Create(c, db.CreateAuthorParams{
		Name: body.Name,
		Bio:  pgtype.Text{String: body.Bio, Valid: body.Name != ""},
	})

	if error != nil {
		a.log.Error().Msg("Error creating new author\n")
		return authorComponents.NewAuthor(body.Name, body.Bio, map[string]string{}).Render(c.Request().Context(), c.Response().Writer)
	}

	return c.Redirect(http.StatusFound, fmt.Sprintf("/author/%d", author.ID))
}

func (a *AuthorHandler) editAuthor(c echo.Context) error {
	body := &UpdateAuthorBody{}
	fmt.Printf("id params: %s\n", c.Param("id"))

	if err := c.Bind(body); err != nil {
		fmt.Println(err)
		a.log.Error().Msg("Error creating new author\n")
		return c.Redirect(http.StatusBadRequest, "/author")
	}

	errorsMap := make(map[string]string)
	if err := c.Validate(body); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Tag(), err.Param(), err.Field())
			errorsMap[strings.ToLower(err.Field())] = errors.Message(err, body)
		}
		return authorComponents.EditAuthor(body.Name, body.Bio, body.ID, errorsMap).Render(c.Request().Context(), c.Response().Writer)
	}

	err := a.authorService.Update(c, db.UpdateAuthorParams{
		ID:   body.ID,
		Name: body.Name,
		Bio:  pgtype.Text{String: body.Bio, Valid: body.Name != ""},
	})

	if err != nil {
		a.log.Error().Msg("Error updating new author\n")
		return authorComponents.EditAuthor(body.Name, body.Bio, body.ID, map[string]string{}).Render(c.Request().Context(), c.Response().Writer)
	}

	c.Response().Header().Add("HX-Push-Url", fmt.Sprintf("/author/%d", body.ID))

	authorComponents.Author(db.Author{ID: body.ID, Name: body.Name, Bio: pgtype.Text{String: body.Bio, Valid: body.Name != ""}}).Render(c.Request().Context(), c.Response().Writer)
	components.Notification("end", "top", "alert-success", fmt.Sprintf("Author %s updated successfuly", body.Name)).Render(c.Request().Context(), c.Response().Writer)
	return nil
}

func (a *AuthorHandler) newAuthor(c echo.Context) error {
	return authorComponents.NewAuthor("", "", map[string]string{}).Render(c.Request().Context(), c.Response().Writer)
}

func (a *AuthorHandler) deleteAuthor(c echo.Context) error {
	id := c.Param("id")

	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.Response().Status = http.StatusBadRequest
		c.Response().Header().Add("HX-Reswap", "none")
		components.Notification("end", "top", "alert-error", "W").Render(c.Request().Context(), c.Response().Writer)
		return err
	}

	author, err := a.authorService.GetOne(parsedId, c)

	if err != nil {
		c.Response().Status = http.StatusBadRequest
		c.Response().Header().Add("HX-Reswap", "none")
		components.Notification("end", "top", "alert-error", "Author with this ID does not exist").Render(c.Request().Context(), c.Response().Writer)
		return err
	}

	err = a.authorService.Delete(c, parsedId)

	if err != nil {
		c.Response().Status = http.StatusInternalServerError
		c.Response().Header().Add("HX-Reswap", "none")
		components.Notification("end", "top", "alert-error", "Something went wrong").Render(c.Request().Context(), c.Response().Writer)
		return err
	}

	components.Notification("end", "top", "alert-success", fmt.Sprintf("Author %s deleted successfuly", author.Name)).Render(c.Request().Context(), c.Response().Writer)

	return nil
}
