package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"server/config"
	"server/database/models"
)

func newDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Settings.GetPostgresDSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.EndorsedCoinInfo{})
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}

var Database = newDatabase()
