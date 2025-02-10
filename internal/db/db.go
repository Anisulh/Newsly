package db

import (
	"Newsly/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	// Auto-migrate the schema to re-create the tables.
	return db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Like{},
		&models.Comment{},
		&models.ResearchPaper{},
		&models.SavedPaper{},
	)

}

func initializeCategories(db *gorm.DB) error {
	log.Print("Initializing categories")

	categories := []models.Category{
		{Key: "ml"},
		{Key: "quantum"},
		{Key: "neuroscience"},
		{Key: "genetics"},
		{Key: "renewables"},
		{Key: "astrophysics"},
		{Key: "robotics"},
		{Key: "biotech"},
		{Key: "blockchain"},
		{Key: "materials"},
		{Key: "medicine"},
		{Key: "social"},
		{Key: "engineering"},
		{Key: "cs"},
		{Key: "data_science"},
		{Key: "economics"},
	}
	for _, cat := range categories {
		// If a category with this key exists, FirstOrCreate does nothing.
		if err := db.Where("key = ?", cat.Key).FirstOrCreate(&cat).Error; err != nil {
			return err
		}
	}
	return nil
}

func setupDatabase(connectionString string) (*gorm.DB, error) {
	log.Print("Setting up database")
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	err = migrate(db)
	if err != nil {
		log.Printf("Error migrating database: %v", err)
		return nil, err
	}
	log.Print("Database migration successful")
	err = initializeCategories(db)
	if err != nil {
		log.Printf("Error initializing categories: %v", err)
		return nil, err
	}
	log.Print("Database initialization successful")
	return db, nil
}

func MustLoadDatabase(connectionString string) *gorm.DB {
	db, err := setupDatabase(connectionString)
	if err != nil {
		panic(err)
	}
	return db
}
