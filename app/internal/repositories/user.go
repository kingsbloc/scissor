package repositories

import (
	"errors"

	"github.com/kingsbloc/scissor/internal/models"
	"gorm.io/gorm"
)

type UserQuery interface {
	Add(user *models.User) *gorm.DB
	FindByEmail(email string) *models.User
	FindByID(id uint) *models.User
}

type userQuery struct{}

func (s *userQuery) Add(user *models.User) *gorm.DB {
	result := DB.Create(&user)
	return result
}

func (s *userQuery) FindByEmail(email string) *models.User {
	var user models.User
	result := DB.Where(&models.User{Email: email}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func (s *userQuery) FindByID(id uint) *models.User {
	var user models.User
	result := DB.Where(&models.User{ID: id}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}
