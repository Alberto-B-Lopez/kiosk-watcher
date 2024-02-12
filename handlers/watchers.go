package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/lopez/kiosk-watcher/models"
	"github.com/lopez/kiosk-watcher/views"
)

func AddWatcher(c echo.Context) error {
	user := getUserBySession(c)
	name := c.FormValue("name")
	name = strings.TrimSpace(name)
	name = strings.Join(strings.Fields(name), "")

	for _, w := range user.List {
		if w.Name == name {
			w.Stn = user.Station.Name
			return render(c, views.Home(w.Stn, user.List))
		}
	}

	nw := models.NewWatcher(name, user.Station.Name)
	user.List = append(user.List, nw)
	updateList(c, user.List)

	return render(c, views.Home(nw.Stn, user.List))
}

func DeleteWatcher(c echo.Context) error {
	user := getUserBySession(c)
	name := c.Param("id")

	for i, w := range user.List {
		if w.Name == name {
			user.List = append(user.List[:i], user.List[i+1:]...)
			updateList(c, user.List)
			return render(c, views.RenderWatchers(user.List))
		}
	}

	return render(c, views.Home(user.Station.Name, user.List))
}

func GetTime(c echo.Context) error {
	user := getUserBySession(c)
	name := c.Param("id")

	for _, w := range user.List {
		if w.Name == name {
			w.Ticker = time.Since(w.CurrTime)
			updateList(c, user.List)
			return render(c, views.TimerCheck(w))
		}
	}

	return render(c, views.Error())
}

func StartWatcher(c echo.Context) error {
	user := getUserBySession(c)
	name := c.Param("id")
	for _, w := range user.List {
		if w.Name == name {
			if w.Ticker > 0 {
				w.IsRunning = true
				w.CurrTime = time.Now().Add(-w.Ticker)
				updateList(c, user.List)
				return render(c, views.WatcherState(w))
			}
			w.StartTime = time.Now()
			w.CurrTime = time.Now()
			w.IsRunning = true
			updateList(c, user.List)

			return render(c, views.WatcherState(w))
		}
	}
	return render(c, views.Error())
}

func StopWatcher(c echo.Context) error {
	user := getUserBySession(c)
	name := c.Param("id")
	for _, w := range user.List {
		if w.Name == name {
			w.IsRunning = false
			w.EndTime = time.Now()
			updateList(c, user.List)

			return render(c, views.WatcherState(w))
		}
	}
	return render(c, views.Error())
}

func ResetWatcher(c echo.Context) error {
	user := getUserBySession(c)
	name := c.Param("id")
	for _, w := range user.List {
		if w.Name == name {
			w.IsRunning = false
			w.Ticker = 0
			updateList(c, user.List)

			return render(c, views.WatcherState(w))
		}
	}
	return render(c, views.Error())
}

func Save(c echo.Context) error {
	user := getUserBySession(c)
	name := c.Param("id")
	for _, w := range user.List {
		if w.Name == name {
			err := db.AddRow(w)
			if err != nil {
				fmt.Println("Error adding row: ", err)
			}

			if w.IsRunning {
				w.IsRunning = false
				w.EndTime = time.Now()
			}
			w.Ticker = 0
			fmt.Println("saved")
			updateList(c, user.List)

			return render(c, views.Watcher(w))
		}
	}
	return render(c, views.Error())
}

func BagTagToggle(c echo.Context) error {
	user := getUserBySession(c)
	name := c.Param("id")
	for _, w := range user.List {
		if w.Name == name {
			if c.FormValue("myCheckbox") == "on" {
				w.BagTagPrinted = true
				updateList(c, user.List)
			} else {
				fmt.Println(w.BagTagPrinted)
				w.BagTagPrinted = false
				updateList(c, user.List)
			}

			return render(c, views.Watcher(w))
		}
	}
	return render(c, views.Error())
}
