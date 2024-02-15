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
	name = strings.ToUpper(name)

	for _, w := range user.List {
		if w.Name == name {
			w.Stn = user.Station.Name
			return render(c, views.Home(w.Stn, user.List))
		}
	}

	nw := models.NewWatcher(name, user.Station.Name)
	user.List = append(user.List, nw)
	err := updateList(c, user.List)
	if err != nil {
		return render(c, views.Error())
	}

	return render(c, views.Home(nw.Stn, user.List))
}

func DeleteWatcher(c echo.Context) error {
	user := getUserBySession(c)
	name := c.Param("id")

	for i, w := range user.List {
		if w.Name == name {
			user.List = append(user.List[:i], user.List[i+1:]...)
			err := updateList(c, user.List)
			if err != nil {
				return render(c, views.Error())
			}

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
			err := updateList(c, user.List)
			if err != nil {
				return render(c, views.Error())
			}

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
				err := updateList(c, user.List)
				if err != nil {
					return render(c, views.Error())
				}

				return render(c, views.WatcherState(w))
			}
			w.StartTime = time.Now()
			w.CurrTime = time.Now()
			w.IsRunning = true
			err := updateList(c, user.List)
			if err != nil {
				return render(c, views.Error())
			}

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
			err := updateList(c, user.List)
			if err != nil {
				return render(c, views.Error())
			}

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
			err := updateList(c, user.List)
			if err != nil {
				return render(c, views.Error())
			}

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
			w.EndTime = time.Now()
			err := db.AddRow(w)
			if err != nil {
				fmt.Println("Error adding row: ", err)
			}

			if w.IsRunning {
				w.IsRunning = false
			}

			w.Ticker = 0
			err = updateList(c, user.List)
			if err != nil {
				return render(c, views.Error())
			}

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
			if c.FormValue("bagtag") == "on" {
				w.BagTagPrinted = true
				err := updateList(c, user.List)
				if err != nil {
					return render(c, views.Error())
				}
			} else {
				w.BagTagPrinted = false
				err := updateList(c, user.List)
				if err != nil {
					return render(c, views.Error())
				}
			}

			return render(c, views.Watcher(w))
		}
	}
	return render(c, views.Error())
}

func BoardingPassToggle(c echo.Context) error {
	user := getUserBySession(c)
	name := c.Param("id")
	for _, w := range user.List {
		if w.Name == name {
			if c.FormValue("boardingpass") == "on" {
				w.BoardingPassPrinted = true
				err := updateList(c, user.List)
				if err != nil {
					return render(c, views.Error())
				}
			} else {
				w.BoardingPassPrinted = false
				err := updateList(c, user.List)
				if err != nil {
					return render(c, views.Error())
				}
			}

			return render(c, views.Watcher(w))
		}
	}
	return render(c, views.Error())
}
