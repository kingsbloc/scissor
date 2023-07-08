package services

import (
	"math/rand"
	"strings"

	"github.com/kingsbloc/scissor/internal/config"
	"github.com/kingsbloc/scissor/internal/repositories"
	"github.com/kingsbloc/scissor/internal/utils"
)

type ShortenService interface {
	GenerateNewShort() string
	EnforceHTTP(url string) string
	CheckDomainError(url string) bool
}

type shortenService struct {
	dao repositories.DAO
}

func NewShortenService(dao repositories.DAO) ShortenService {
	return &shortenService{dao: dao}
}

func (s *shortenService) GenerateNewShort() string {
	return utils.EncodeBase62(rand.Uint64())
}

func (s *shortenService) EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func (s *shortenService) CheckDomainError(url string) bool {
	domain := config.New().Server.ServerUrl

	if url == domain {
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	return newURL != domain
}
