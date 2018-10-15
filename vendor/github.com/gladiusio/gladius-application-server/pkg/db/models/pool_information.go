package models

import "time"

type PoolInformation struct {
	// Default Gorm Model Properties
	// Stating here to remove from JSON responses with `json:"-"`
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`

	OnlyRow   int    `json:"-" gorm:"unique;not null;default:1"`
	Address   string `json:"address" gorm:"not null"`
	Name      string `json:"name" gorm:"not null"`
	Bio       string `json:"bio"`
	Location  string `json:"location" gorm:"not null"`
	Rating    int    `json:"rating"`
	NodeCount int    `json:"nodeCount" gorm:"not null;default:0"`
	Wallet    string `json:"wallet" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Url       string `json:"url" gorm:"-"`
	Public    bool   `json:"public" gorm:"not null;default:false"`
}

func (PoolInformation) TableName() string {
	return "pool_information"
}
