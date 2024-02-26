package utils

import (
	"echo-test/components"

	"github.com/labstack/echo/v4"
)

type RedirectParams struct {
	headers map[string]string
	Message string
}

type DefaultRedirectHeaders struct {
	Url    string
	Swap   string
	Target string
}

func NewDefaultRedirectParams(params DefaultRedirectHeaders, meesage string) *RedirectParams {
	headers := make(map[string]string, 0)

	if params.Swap != "" {
		headers["HX-ReSwap"] = params.Swap
	}

	if params.Target != "" {
		headers["HX-Retarget"] = params.Target
	}

	if params.Url != "" {
		headers["HX-Push-Url"] = params.Url
	}

	return &RedirectParams{
		headers: headers,
		Message: meesage,
	}
}

func SendServerErrorNotification(c echo.Context) error {
	c.Response().Header().Add("HX-Reswap", "none")
	return components.Notification("end", "top", "alert-error", "Something went wrong").Render(c.Request().Context(), c.Response().Writer)
}

func SendSuccessNotification(c echo.Context, params RedirectParams) error {
	for key, value := range params.headers {
		c.Response().Header().Add(key, value)
	}
	return components.Notification("end", "top", "alert-success", params.Message).Render(c.Request().Context(), c.Response().Writer)
}
