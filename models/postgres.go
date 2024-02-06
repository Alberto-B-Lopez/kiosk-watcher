package models

import (
	"database/sql"
	"fmt"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresdb() (*PostgresStorage, error) {
	connStr := "user=postgres dbname=postgres password=mysecretpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (p *PostgresStorage) Init() error {
	return p.createTable()
}

func (p *PostgresStorage) dropAndCreateTable() error {
	// Drop the table if it exists
	_, err := p.db.Exec("DROP TABLE IF EXISTS kiosk")
	if err != nil {
		return err
	}

	// Create the table
	err = p.createTable()
	return err
}

func (p *PostgresStorage) createTable() error {
	fmt.Println("We in createTable func")
	query := ` create table if not exists kiosk (
      id serial primary key,
      stn varchar(3),
      name varchar(255),
      duration varchar(255) ,
      bag_tag_printed bool,
      created_at timestamp default current_timestamp
    )`

	_, err := p.db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("looks like we done creating table")

	return err
}

func (p *PostgresStorage) AddRow(w *Watcher) error {
	fmt.Println("We in Add func")
	query := `insert into kiosk
          (stn, name, duration, bag_tag_printed)
          values ($1, $2, $3, $4)`
	_, err := p.db.Exec(query, w.Stn, w.Name, w.Ticker.String(), w.BagTagPrinted)
	if err != nil {
		return err
	}
	fmt.Println("looks like we done")
	return err
}
