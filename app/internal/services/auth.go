package services

import (
	"errors"

	"github.com/kingsbloc/scissor/internal/dto"
	"github.com/kingsbloc/scissor/internal/models"
	"github.com/kingsbloc/scissor/internal/repositories"
)

type AuthService interface {
	Signin(dto *dto.LoginDto) (*models.User, error)
}

type authService struct {
	dao repositories.DAO
}

func NewAuthService(dao repositories.DAO) AuthService {
	return &authService{dao: dao}
}

func (s *authService) Signin(dto *dto.LoginDto) (*models.User, error) {
	user := s.dao.NewUserQuery().FindByEmail(dto.Email)
	if user != nil {
		err := user.ComparePassword(dto.Password)
		if err != nil {
			return nil, errors.New("invalid login")
		}
		return user, nil
	}
	return nil, errors.New("invalid login")
}
