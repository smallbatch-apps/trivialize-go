package models

import "github.com/google/uuid"

type SeriesStatus int

const (
	New SeriesStatus = iota + 1
	Active
	Finalised
)

type Series struct {
	CustomModel
	Name         string
	Description  string
	PromoImage   string
	CompanyID    uuid.UUID `json:"company_id"`
	SeriesStatus SeriesStatus
	Company      Company
	Events       []Event
}
