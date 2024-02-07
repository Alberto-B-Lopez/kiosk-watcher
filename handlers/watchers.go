package handlers

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/lopez/kiosk-watcher/models"
	"github.com/lopez/kiosk-watcher/views"
)

func AddWatcher(c echo.Context) error {
	name := c.FormValue("name")
	fmt.Println("Adding watcher: ", name)

	for _, w := range list {
		if w.Name == name {
			w.Stn = station.Name
			return render(c, views.Home(w.Stn, list))
		}
	}

	nw := models.NewWatcher(name, station.Name)
	list = append(list, nw)
	return render(c, views.Home(nw.Stn, list))
}

func DeleteWatcher(c echo.Context) error {
	name := c.Param("id")
	fmt.Println("Deleting watcher: ", name)

	for i, w := range list {
		if w.Name == name {
			list = append(list[:i], list[i+1:]...)
			return render(c, views.Home(station.Name, list))
		}
	}

	return render(c, views.Home(station.Name, list))
}

func GetTime(c echo.Context) error {
	name := c.Param("id")

	for _, w := range list {
		if w.Name == name {
			w.Ticker = time.Since(w.StartTime)
			return render(c, views.TickerCheck(w))
		}
	}

	return render(c, views.Error())
}

func StartWatcher(c echo.Context) error {
	name := c.Param("id")
	for _, w := range list {
		if w.Name == name {
			if w.Ticker > 0 {
				w.IsRunning = true
				w.StartTime = time.Now().Add(-w.Ticker)
				return render(c, views.Watcher(w))
			}
			w.StartTime = time.Now()
			w.IsRunning = true

			return render(c, views.Watcher(w))
		}
	}
	return render(c, views.Error())
}

func StopWatcher(c echo.Context) error {
	name := c.Param("id")
	fmt.Println("Stopping count for: ", name)
	for _, w := range list {
		if w.Name == name {
			w.IsRunning = false

			return render(c, views.Watcher(w))
		}
	}
	return render(c, views.Error())
}

func ResetWatcher(c echo.Context) error {
	name := c.Param("id")
	fmt.Println("Resetting count for: ", name)
	for _, w := range list {
		if w.Name == name {
			w.IsRunning = false
			w.Ticker = 0

			return render(c, views.Watcher(w))
		}
	}
	return render(c, views.Error())
}

func Save(c echo.Context) error {
	name := c.Param("id")
	for _, w := range list {
		if w.Name == name {
			err := db.AddRow(w)
			if err != nil {
				fmt.Println("Error adding row: ", err)
			}
			w.IsRunning = false
			w.Ticker = 0
			fmt.Println("saved")

			return render(c, views.Watcher(w))
		}
	}
	return render(c, views.Error())
}
