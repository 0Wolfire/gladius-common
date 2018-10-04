package db

import (
	"github.com/gladiusio/gladius-application-server/pkg/db/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite for now, SQL later on
)

func Initialize() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schemas
	db.AutoMigrate(&models.PoolInformation{})
	db.AutoMigrate(&models.NodeProfile{})
}
