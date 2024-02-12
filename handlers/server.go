package handlers

import (
	"encoding/gob"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"github.com/lopez/kiosk-watcher/models"
)

type Server struct {
	store models.Storage
}

func NewServer(store models.Storage) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) Start() error {
	app := echo.New()
	db = s.store

	// Static files
	app.Static("/css", "css")

	cookieStore.Options.Secure = false
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.MaxAge = 2 * 60 * 60
	gob.Register(&User{})

	app.Use(session.Middleware(cookieStore))

	// app.Use(middleware.Logger())

	// // CORS
	// app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{"*"},
	// })

	// GET Routes
	app.GET("/", ShowIndex)
	app.GET("/edit", ShowEdit)
	app.GET("/get-time/:id", GetTime)
	app.GET("/start-watcher/:id", StartWatcher)
	app.GET("/stop-watcher/:id", StopWatcher)
	app.GET("/reset-watcher/:id", ResetWatcher)
	// POST Routes
	app.POST("/set-station", ShowHome)
	app.POST("/delete/:id", DeleteWatcher)
	app.POST("/add", AddWatcher)
	app.POST("/save/:id", Save)
  app.POST("/bag-tag/:id", BagTagToggle)

	return app.Start("0.0.0.0:10000")
}
