package models

import (
	"time"

	"gorm.io/gorm"
)

type Clients struct {
	gorm.Model
	ID              uint
	Name            string
	Nip             int
	Phone           int
	ContactPersone  string
	PersonePhone    int
	WWW             string
	FB              string
	Email           string
	Adres           string
	AdresGoogleCODE string
	GoogleMapURL    string
	Status          int
	LastContact     time.Time
	NextContact     time.Time
}
