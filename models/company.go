package models

type Company struct {
	CustomModel
	Name      string
	Users     []User
	Series    []Series
	Questions []Question
}
