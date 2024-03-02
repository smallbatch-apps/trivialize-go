package models

import "github.com/google/uuid"

type Answer struct {
	CustomModel
	Text       string
	Points     int
	Sort       int
	QuestionID uuid.UUID `json:"question_id"`
	Question   Question
}
