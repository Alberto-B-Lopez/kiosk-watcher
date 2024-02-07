package handlers

import "github.com/lopez/kiosk-watcher/models"

var (
	list    = []*models.Watcher{}
	station = models.NewStation("")
	db      models.Storage
)
