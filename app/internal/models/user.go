package models

import (
	"time"

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
