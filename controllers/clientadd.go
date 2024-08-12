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

	var Client []models.Clients

	err := c.BodyParser(Client)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	initializers.ConnectToDataBase()

	err = initializers.DB.Create(Client).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Redirect("clientadd")

}
