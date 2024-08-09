package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	initializers "github.com/ut3k/clientContact/initialisers"
	"github.com/ut3k/clientContact/models"
)

func ClientListView(c *fiber.Ctx) error {

	initializers.ConnectToDataBase()

	var Clients []models.Clients
	pageStr := c.Query("page", "1")
	page, _ := strconv.Atoi(pageStr)

	if page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	initializers.DB.Offset(offset).Limit(pageSize).Find(&Clients)

	return c.Render("index", fiber.Map{
		"Clients": Clients,
		"BaseURL": c.BaseURL(),
	})

}
