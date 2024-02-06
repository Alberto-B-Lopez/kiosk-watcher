package handlers

import "github.com/lopez/stopwatch-delta/models"

var (
	list    = []*models.Watcher{}
	station = models.NewStation("")
	db      models.Storage
)
