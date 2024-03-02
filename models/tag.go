package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TempTag struct {
	ID   uuid.UUID
	Name string
}

type Tag struct {
	ID                  uuid.UUID `gorm:"type:uuid;primary_key;"`
	Text                string
	Icon                string
	Colour              string
	IconColour          string
	IconSecondaryColour string
	Questions           []Question `gorm:"many2many:question_tags;"`
}

func (u *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
