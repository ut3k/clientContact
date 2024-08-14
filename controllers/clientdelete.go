package controllers

import (
	"github.com/gofiber/fiber/v2"
	initializers "github.com/ut3k/clientContact/initialisers"
	"github.com/ut3k/clientContact/models"
)

func ClientDelete(c *fiber.Ctx) error {

	id := c.Params("id")

	initializers.ConnectToDataBase()

	var Client []models.Clients

	if err := initializers.DB.First(&Client, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Client not found")
	}

	return c.Render("clientdelete", fiber.Map{
		"Client":  Client,
		"BaseURL": c.BaseURL(),
	})

}

func ClientDeletePost(c *fiber.Ctx) error {

	id := c.Params("id")
	var Client []models.Clients

	initializers.ConnectToDataBase()

	err := initializers.DB.First(&Client, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Client not Found")
	}

	err = initializers.DB.Delete(Client, id).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Redirect("/tabview/")

}
