package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(e echo.Context, t templ.Component) error {
	return t.Render(e.Request().Context(), e.Response())
}
