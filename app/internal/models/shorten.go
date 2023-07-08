package models

import (
	"time"

	"gorm.io/gorm"
)

type Shorten struct {
	ID          uint           `gorm:"index:;primaryKey;autoIncrement:true" json:"id"`
	UserID      uint           `gorm:"index:;primaryKey;" json:"user_id"`
	Url         string         `gorm:"not null;type:varchar(2000)" json:"url"`
	CustomShort string         `gorm:"index;not null;type:varchar(2000)" json:"custom_short"`
	Exp         int32          `gorm:"not null;" json:"exp"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" swaggertype:"primitive,string"`
}
