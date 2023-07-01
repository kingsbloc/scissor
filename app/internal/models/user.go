package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"index:;primaryKey;autoIncrement:true" json:"id"`
	Email     string         `gorm:"unique;not null;type:varchar(200)" json:"email"`
	Password  string         `gorm:"not null;type:varchar(200)" json:"-"`
	Name      string         `gorm:"not null;type:varchar(200)" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" swaggertype:"primitive,string"`
}

func (user *User) HashPassword(password string) error {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(byte)
	return nil
}

func (user *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
