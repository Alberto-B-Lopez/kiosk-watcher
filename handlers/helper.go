package handlers

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/lopez/kiosk-watcher/models"
)

func render(e echo.Context, t templ.Component) error {
	return t.Render(e.Request().Context(), e.Response())
}

func getUserBySession(c echo.Context) *User {
	sess, _ := cookieStore.Get(c.Request(), "session")
	if u, ok := sess.Values["user"]; ok {
		if user, ok := u.(*User); ok {
			return user
		} else {
			user = &User{
				Station: &models.Station{Name: ""},
				List:    []*models.Watcher{},
			}
		}
	}
	return nil
}

func updateStation(c echo.Context, station string) (*User, error) {
	sess, _ := cookieStore.Get(c.Request(), "session")
	if u, ok := sess.Values["user"]; ok {
		if user, ok := u.(*User); ok {
			user.Station.Name = station

			if err := sess.Save(c.Request(), c.Response()); err != nil {
				// Handle the error
				fmt.Println("Error saving session:", err)

				return nil, err
			}
			return user, nil
		}
	}
	return &User{}, nil
}

func updateList(c echo.Context, list []*models.Watcher) error {
	sess, _ := cookieStore.Get(c.Request(), "session")
	if u, ok := sess.Values["user"]; ok {
		if user, ok := u.(*User); ok {
			user.List = list
			if err := sess.Save(c.Request(), c.Response()); err != nil {
				// Handle the error
				fmt.Println("Error saving session:", err)
				return err
			}
		}
	}
	return nil
}
