package models

import (
	"time"

	"github.com/google/uuid"
)

type EventStatus int

const (
	Draft EventStatus = iota + 1
	Published
	Completed
	Cancelled
	Archived
)

type Event struct {
	CustomModel
	Description string
	Status      int
	StartTime   time.Time
	Location    string
	SeriesID    uuid.UUID `json:"series_id"`
	Series      Series
}
