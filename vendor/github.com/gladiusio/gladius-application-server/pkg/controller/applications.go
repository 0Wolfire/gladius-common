package controller

import (
	"errors"
	"github.com/gladiusio/gladius-application-server/pkg/db/models"
	"github.com/jinzhu/gorm"
	)

// temp
func TempDBCalls() {
	//db, _ := Initialize(nil)
	//
	////Temp for testing
	//request := models.NodeRequestPayload{
	//	EstimatedSpeed: 100,
	//	Wallet:         "0x97543295ABC235DDD",
	//	Name:           "Name",
	//	Email:          "email@fds.com",
	//	Bio:            "bio",
	//	Location:       "location",
	//	IPAddress:      "0.0.0.0",
	//}
	//
	//NodeApplyToPool(db, request)
	//
	//requestUpdate := models.NodeRequestPayload{
	//	Wallet:   "0x97543295ABC235DDD",
	//	Name:     "Name Updated",
	//	Email:    "email@fds.com Updated",
	//	Bio:      "bio Updated",
	//	Location: "location Updated",
	//}
	//
	//_, err := NodeUpdateProfile(db, requestUpdate)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Pool Accepts Application
	//PoolApplicationStatus(db, "0x97543295ABC235DDD", false)
	//// Node Denies Application
	//NodeApplicationStatus(db, "0x97543295ABC235DDD", false)
	//
	//// Pool Accepts Application
	//PoolApplicationStatus(db, "0x97543295ABC235DDD", true)
	//// Node Denies Application
	//NodeApplicationStatus(db, "0x97543295ABC235DDD", true)
	//
	//poolInfo := models.PoolInformation{
	//	Name:     "Gladius Pool",
	//	Address:  "124.232.83.8",
	//	Bio:      "Gladius Testing Pool",
	//	Location: "Washington D.C.",
	//	Rating:   5,
	//	Wallet:   "0x96585865865",
	//	Public:   true,
	//}
	//
	//PoolCreateUpdateData(db, poolInfo)
}

func PoolCreateUpdateData(db *gorm.DB, poolInfo models.PoolInformation) {
	var pool models.PoolInformation

	db.Model(&pool).FirstOrCreate(&pool)
	db.Model(&pool).Updates(&poolInfo)
}

func NodeApplyToPool(db *gorm.DB, payload models.NodeRequestPayload) (models.NodeProfile, error) {
	profile := models.CreateApplication(&payload)
	err := db.Model(&profile).Where("wallet like ?", payload.Wallet).FirstOrCreate(&profile).Error

	return profile, err
}

func NodeUpdateProfile(db *gorm.DB, payload models.NodeRequestPayload) (models.NodeProfile, error) {
	profile, err := NodeProfile(db, payload.Wallet)
	if err != nil {
		return profile, err
	}

	db.Model(&profile).Updates(
		models.NodeProfile{
			Name:     payload.Name,
			Bio:      payload.Bio,
			Email:    payload.Email,
			Location: payload.Location,
		},
	)

	return profile, nil
}

func NodeProfile(db *gorm.DB, wallet string) (models.NodeProfile, error) {
	var profile models.NodeProfile

	if err := db.Model(&profile).Where("lower(wallet) like lower(?)", wallet).First(&profile).Error; err != nil {
		return models.NodeProfile{}, errors.New("NodeProfile() profile not found for given wallet address")
	}

	return profile, nil
}

type FullProfile struct {
	NodeProfile     models.NodeProfile     `json:"profile"`
	PoolInformation models.PoolInformation `json:"pool"`
}

func NodePoolApplication(db *gorm.DB, wallet string) (FullProfile, error) {
	var poolInformation models.PoolInformation
	var fullProfile FullProfile

	profile, err := NodeProfile(db, wallet)
	if err != nil {
		return fullProfile, err
	}

	poolInformation, err = PoolInformation(db)

	if err != nil {
		return fullProfile, err
	}

	return FullProfile{
		NodeProfile:profile,
		PoolInformation:poolInformation,
	}, err
}

func PoolApplicationStatus(db *gorm.DB, wallet string, accepted bool) {
	profile, _ := NodeProfile(db, wallet)
	db.Save(&profile)
}

func NodeApplicationStatus(db *gorm.DB, wallet string, accepted bool) {
	profile, _ := NodeProfile(db, wallet)
	db.Save(&profile)
}
