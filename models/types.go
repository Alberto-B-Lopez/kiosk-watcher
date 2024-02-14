package models

import (
	"fmt"
	"time"
)

type Storage interface {
	AddRow(*Watcher) error
}

type Station struct {
	Name string
}

func NewStation(name string) *Station {
	return &Station{
		Name: name,
	}
}

type Watcher struct {
	ID                  int
	Stn                 string
	Name                string
	Ticker              time.Duration
	CurrTime            time.Time
	StartTime           time.Time
	EndTime             time.Time
	IsRunning           bool
	CurrnetTime         time.Time
	BagTagPrinted       bool
	BoardingPassPrinted bool
}

func NewWatcher(name string, station string) *Watcher {
	return &Watcher{
		Name:          name,
		Stn:           station,
		Ticker:        time.Duration(0),
		IsRunning:     false,
		BagTagPrinted: false,
	}
}

func (w *Watcher) FormatDuration() string {
	// Extract minutes and seconds
	minutes := int(w.Ticker.Minutes())
	seconds := int(w.Ticker.Seconds()) % 60

	// Format the result as a string
	if minutes > 0 {
		return fmt.Sprintf("%dm%ds", minutes, seconds)
	}
	return fmt.Sprintf("%ds", seconds)
}
