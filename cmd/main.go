package main

import (
	"log"
	"os"

	"github.com/lopez/kiosk-watcher/handlers"
	"github.com/lopez/kiosk-watcher/models"
)

func main() {
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	pass := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")

	db, err := models.NewPostgresdb(user, dbname, pass, host)
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
