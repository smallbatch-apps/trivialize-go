package models

import "github.com/google/uuid"

type QuestionType int

const (
	Simple QuestionType = iota + 1
	MultipleChoice
	TrueFalse
	MultipleCorrect
)

type Question struct {
	CustomModel
	Text      string
	Type      QuestionType
	CompanyID uuid.UUID `json:"company_id"`
	Company   Company
	Documents []Document
	Tags      []Tag `gorm:"many2many:question_tags;"`
}
