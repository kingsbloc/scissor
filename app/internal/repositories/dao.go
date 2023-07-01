package repositories

import (
	"github.com/kingsbloc/scissor/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DAO interface {
}

type dao struct{}

func NewDAO(db *gorm.DB) DAO {
	DB = db
	return &dao{}
}

// Setup DB Connection
func InitDB() (*gorm.DB, error) {
	dsn := utils.GetEnvVar("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

// AutoMigration Setup
func AutoMigrate(models []interface{}, db *gorm.DB) {
	db.AutoMigrate(models...)
}
