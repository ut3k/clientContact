package models

import (
	"time"

	"gorm.io/gorm"
)

type Clients struct {
	gorm.Model
	ID              uint
	Name            string
	Nip             int `gorm:"default:NULL"`
	Phone           int
	ContactPersone  string `gorm:"default:NULL"`
	PersonePhone    int    `gorm:"default:NULL"`
	WWW             string `gorm:"default:NULL"`
	FB              string `gorm:"default:NULL"`
	Email           string `gorm:"default:NULL"`
	Adres           string `gorm:"default:NULL"`
	AdresGoogleCODE string `gorm:"unique"`
	GoogleMapURL    string
	Status          int `gorm:"default:0"`
	LastContact     time.Time
	NextContact     time.Time
}
