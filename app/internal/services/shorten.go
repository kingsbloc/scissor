package services

import (
	"github.com/kingsbloc/scissor/internal/repositories"
)

type ShortenService interface{}

type shortenService struct {
	dao repositories.DAO
}

func NewShortenService(dao repositories.DAO) ShortenService {
	return &shortenService{dao: dao}
}
