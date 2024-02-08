package handlers

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/lopez/kiosk-watcher/models"
)

type User struct {
	Station *models.Station
	List    []*models.Watcher
}

var (
	cookieStore = sessions.NewCookieStore([]byte(os.Getenv("SECRET"))) // Middleware
	db          models.Storage
)
