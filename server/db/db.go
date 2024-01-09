package db

import (
	"github.com/Anisulh/content_personalization/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(utils.DBConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
