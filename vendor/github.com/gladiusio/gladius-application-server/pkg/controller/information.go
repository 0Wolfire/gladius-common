package controller

import (
	"github.com/gladiusio/gladius-application-server/pkg/db/models"
	"github.com/jinzhu/gorm"
)

func PoolInformation(db *gorm.DB) (models.PoolInformation, error) {
	var profile models.PoolInformation

	err := db.First(&profile).Error

	return profile, err
}