package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/gladiusio/gladius-application-server/pkg/db/models"
)

func Initialize(db *gorm.DB) (*gorm.DB, error) {
	var err error

	if db == nil {
		db, err = gorm.Open("sqlite3", "test.db")
	}

	if err != nil {
		return db, err
	}

	// Migrate the schemas
	db.AutoMigrate(&models.PoolInformation{})
	db.AutoMigrate(&models.NodeProfile{})

	return db, err
}