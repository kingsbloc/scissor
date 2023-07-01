package repositories

import (
	"github.com/kingsbloc/scissor/internal/models"
	"gorm.io/gorm"
)

type UserQuery interface {
	Add(user *models.User) *gorm.DB
}

type userQuery struct{}

func (s *userQuery) Add(user *models.User) *gorm.DB {
	result := DB.Create(&user)
	return result
}
