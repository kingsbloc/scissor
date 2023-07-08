package services

import (
	"errors"
	"strconv"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/kingsbloc/scissor/internal/dto"
	"github.com/kingsbloc/scissor/internal/models"
	"github.com/kingsbloc/scissor/internal/repositories"
	"github.com/kingsbloc/scissor/internal/utils"
	"gorm.io/gorm"
)

type UserService interface {
	NewUser(dto *dto.AddUserDto) (bool, error)
	AddNewShorten(dto dto.AddShortenDto, id string) *gorm.DB
}

type userService struct {
	dao repositories.DAO
}

func NewUserService(dao repositories.DAO) UserService {
	return &userService{dao: dao}
}

func (s *userService) NewUser(dto *dto.AddUserDto) (bool, error) {
	var user models.User

	user.Name = dto.Name
	user.Email = dto.Email
	user.HashPassword(dto.Password)

	result := s.dao.NewUserQuery().Add(&user)

	if result.Error != nil {
		var perr *pgconn.PgError
		if errors.As(result.Error, &perr) {
			return false, utils.HandleSqlError(result.Error)
		}
		return false, result.Error
	}
	return true, nil
}

func (s *userService) AddNewShorten(dto dto.AddShortenDto, id string) *gorm.DB {
	userId, _ := strconv.Atoi(id)
	return s.dao.NewShortenQuery().Add(dto, uint(userId))
}
