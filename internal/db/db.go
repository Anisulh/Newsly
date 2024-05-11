package db

import (
	"Newsly/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func migrate(db *gorm.DB) error {
	// Migrate the schema
	return db.AutoMigrate(&models.Keyword{}, &models.User{}, &models.Content{} )
}

func setupDatabase(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = migrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MustLoadDatabase(connectionString string) *gorm.DB {
	db, err := setupDatabase(connectionString)
	if err != nil {
		panic(err)
	}
	return db
}
