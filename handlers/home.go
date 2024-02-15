package handlers

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/lopez/kiosk-watcher/models"
	"github.com/lopez/kiosk-watcher/views"
)

func ShowIndex(c echo.Context) error {
	sess, _ := cookieStore.Get(c.Request(), "session")

	// Check if "user" key exists in sess.Values
	if _, ok := sess.Values["user"]; ok {
		return ShowHome(c)
	} else {
		// "user" key does not exist, create it and set a default value
		defaultUser := &User{
			Station: &models.Station{Name: ""},
			List:    []*models.Watcher{},
		}
		sess.Values["user"] = defaultUser
		// Save the session after setting the default value
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			// Handle the error
			fmt.Println("Error saving session:", err)

			return err
		}
	}

	return render(c, views.Setup())
}

func ShowHome(c echo.Context) error {
	user := getUserBySession(c)
	if user.Station.Name == "" {
		user.Station.Name = strings.ToUpper(c.FormValue("name"))
		for _, w := range user.List {
			w.Stn = user.Station.Name
		}
		updateStation(c, user.Station.Name)
	}

	return render(c, views.Home(user.Station.Name, user.List))
}

func ShowEdit(c echo.Context) error {
	user := getUserBySession(c)
	user.Station.Name = ""
	for _, w := range user.List {
		w.Stn = user.Station.Name
	}
	updateStation(c, user.Station.Name)
	return render(c, views.Setup())
}
