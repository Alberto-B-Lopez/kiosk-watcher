package main

import (
	"log"

	"github.com/lopez/stopwatch-delta/handlers"
	"github.com/lopez/stopwatch-delta/models"
)

func main() {
	db, err := models.NewPostgresdb()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Init()
	if err != nil {
		log.Fatal(err)
	}

	server := handlers.NewServer(db)
	log.Fatal(server.Start())
}
