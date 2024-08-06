package models

import "gorm.io/gorm"

type Clients struct {
	gorm.Model
	Title string
	Body  string
}
