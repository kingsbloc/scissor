package models

import (
	"time"

	"github.com/kingsbloc/scissor/internal/config"
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

func (c *Shorten) ToBase() Shorten {
	return Shorten{
		ID:          c.ID,
		UserID:      c.UserID,
		Url:         c.Url,
		CustomShort: config.New().Server.ServerUrl + "/" + c.CustomShort,
		Exp:         c.Exp,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		DeletedAt:   c.DeletedAt,
	}
}
