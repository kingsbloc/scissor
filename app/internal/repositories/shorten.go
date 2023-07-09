package repositories

import (
	"github.com/kingsbloc/scissor/internal/dto"
	"github.com/kingsbloc/scissor/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ShortenQuery interface {
	Add(dto dto.AddShortenDto, userId uint) *gorm.DB
	ListByUserID(id uint) []models.Shorten
	Delete(id uint, userId uint) (models.Shorten, *gorm.DB)
}

type shortenQuery struct{}

func (s *shortenQuery) Add(dto dto.AddShortenDto, userId uint) *gorm.DB {
	result := DB.Create(&models.Shorten{
		UserID:      userId,
		Url:         dto.Url,
		CustomShort: dto.CustomUrl,
		Exp:         dto.Exp,
	})
	return result

}

func (s *shortenQuery) Delete(id uint, userId uint) (models.Shorten, *gorm.DB) {
	var shorten models.Shorten
	result := DB.Clauses(clause.Returning{}).Where(&models.Shorten{UserID: userId, ID: id}).Delete(&shorten)
	return shorten, result
}

func (s *shortenQuery) ListByUserID(id uint) []models.Shorten {
	var result []models.Shorten
	DB.Where(&models.Shorten{UserID: id}).Find(&result)
	return result
}
