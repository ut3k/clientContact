package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	initializers "github.com/ut3k/clientContact/initialisers"
	"github.com/ut3k/clientContact/models"
)

func ClientUpdate(c *fiber.Ctx) error {

	id := c.Params("id")

	initializers.ConnectToDataBase()

	var Client []models.Clients

	if err := initializers.DB.First(&Client, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Client not found")
	}
	return c.Render("clientupdate", fiber.Map{
		"Client":  Client,
		"BaseURL": c.BaseURL(),
	})

}

func ClientUpdatePost(c *fiber.Ctx) error {

	id := c.Params("id")
	var Client models.Clients

	initializers.ConnectToDataBase()

	err := initializers.DB.First(&Client, id).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Client not Found")
	}
	Client.Name = c.FormValue("name")
	Client.Nip, err = strconv.Atoi(c.FormValue("nip"))
	Client.Phone, err = strconv.Atoi(c.FormValue("phone"))
	Client.ContactPersone = c.FormValue("contactpersone")
	Client.PersonePhone, err = strconv.Atoi(c.FormValue("personephone"))
	Client.WWW = c.FormValue("WWW")
	Client.FB = c.FormValue("FB")
	Client.Email = c.FormValue("email")
	Client.Adres = c.FormValue("adres")
	Client.AdresGoogleCODE = c.FormValue("adresgooglecode")
	Client.GoogleMapURL = c.FormValue("googlemapurl")
	Client.Status, err = strconv.Atoi(c.FormValue("status"))
	// Client.LastContact = c.FormValue("lastcontact")
	// Client.NextContact = c.FormValue("nextcontact")

	err = initializers.DB.Save(&Client).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Redirect("/clientupdate/" + id)

}
