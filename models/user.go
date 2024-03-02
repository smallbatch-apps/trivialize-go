package models

import (
	"github.com/smallbatch-apps/trivialize-go/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	CustomModel
	Name      string
	Email     string `gorm:"index:idx_email_unique,unique"`
	Password  string `json:"-"`
	CompanyId uint   `json:"company_id"`
	Company   Company
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err := u.CustomModel.BeforeCreate(tx); err != nil {
		return err
	}
	u.Password, err = HashPassword(u.Password)
	return
}

func (u *User) FindUserByEmail(email string) error {
	return database.Database.Db.Where("email = ?", email).First(&u).Error
}

func (u *User) FindUserById(db *gorm.DB, id string) error {
	return database.Database.Db.Where("id = ?", id).First(&u).Error
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
