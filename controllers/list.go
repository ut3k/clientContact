package controllers

import (
	"github.com/gofiber/fiber/v2"
	initializers "github.com/ut3k/clientContact/initialisers"
	"github.com/ut3k/clientContact/models"
)

func ClientListView(c *fiber.Ctx) error {

	initializers.ConnectToDataBase()

	var Cities []models.Clients
	initializers.DB.Find(&Cities)

	return c.Render("index", fiber.Map{
		"Cities":  Cities,
		"BaseURL": c.BaseURL(),
	})

}
