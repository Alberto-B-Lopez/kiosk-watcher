package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/lopez/kiosk-watcher/views"
)

func ShowIndex(c echo.Context) error {
	return render(c, views.Index())
}

func ShowHome(c echo.Context) error {
	station.Name = c.FormValue("name")
	return render(c, views.Home(station.Name, list))
}
