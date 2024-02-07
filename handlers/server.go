package handlers

import (
	"github.com/labstack/echo/v4"
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

	// GET Routes
	app.GET("/", ShowIndex)
	app.GET("/get-time/:id", GetTime)
	app.GET("/start-watcher/:id", StartWatcher)
	app.GET("/stop-watcher/:id", StopWatcher)
	app.GET("/reset-watcher/:id", ResetWatcher)
	// POST Routes
	app.POST("/set-station", ShowHome)
	app.POST("/delete/:id", DeleteWatcher)
	app.POST("/add", AddWatcher)
	app.POST("/save/:id", Save)

	return app.Start("localhost:8080")
}
