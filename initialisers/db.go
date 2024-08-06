package initializers

import (
	"fmt"

	"github.com/ut3k/clientContact/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDataBase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data/clientContact.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to DataBase")

	}
	// return db, nil

}

func SyncDB() {
	DB.AutoMigrate(&models.Clients{})
}
