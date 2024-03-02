package models

import "github.com/google/uuid"

type Document struct {
	CustomModel
	Title      string
	Location   string
	TempFile   string
	FileName   string
	Public     bool
	Status     int
	QuestionID uuid.UUID `json:"question_id"`
	Question   Question
}
