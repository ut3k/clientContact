package controllers

import (
	"github.com/gofiber/fiber/v2"
	initializers "github.com/ut3k/clientContact/initialisers"
	"github.com/ut3k/clientContact/models"
)

func ClientAdd(c *fiber.Ctx) error {

	return c.Render("clientadd", fiber.Map{
		"BaseURL": c.BaseURL(),
	})

}

func ClientAddPost(c *fiber.Ctx) error {

	id := c.Params("id")

	initializers.ConnectToDataBase()

	var Client []models.Clients

	if err := initializers.DB.First(&Client, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Client not found")
	}

	return c.Render("clientsingle", fiber.Map{
		"Client":  Client,
		"BaseURL": c.BaseURL(),
	})

}
