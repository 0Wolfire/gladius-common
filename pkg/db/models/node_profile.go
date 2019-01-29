package models

import (
	"time"
)

type NodeProfile struct {
	// Default Gorm Model Properties
	// Stating here to remove from JSON responses with `json:"-"`
	ID             uint       `json:"-" gorm:"primary_key"`
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"-"`
	DeletedAt      *time.Time `json:"-"`
	Name           string     `json:"name" gorm:"not null"`
	Email          string     `json:"email" gorm:"not null"`
	Bio            string     `json:"bio" gorm:"not null"`
	Location       string     `json:"location" gorm:"not null"`
	IPAddress      string     `json:"-" gorm:"not null"`
	EstimatedSpeed int        `json:"estimatedSpeed" gorm:"not null"`
	PoolAccepted   bool       `json:"-" gorm:"default:false"`
	NodeAccepted   bool       `json:"-" gorm:"default:false"`
	Pending        bool       `json:"pending" gorm:"default:true"`
	Approved       bool       `json:"approved" gorm:"default:false"`
	Wallet         string     `json:"wallet" gorm:"not null; unique"`
}

type NodeRequestPayload struct {
	EstimatedSpeed int    `json:"estimatedSpeed"`
	Wallet         string `json:"wallet"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Bio            string `json:"bio"`
	Location       string `json:"location"`
	IPAddress      string `json:"ipAddress"`
}