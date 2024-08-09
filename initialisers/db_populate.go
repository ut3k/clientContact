package initializers

import (
	"fmt"

	"github.com/ut3k/clientContact/models"
)

func PopulateCity() {
	ConnectToDataBase()
	ClientsDummy := []*models.Clients{}

	DB.Create(ClientsDummy)
}
