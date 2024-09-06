package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	initializers "github.com/ut3k/clientContact/initialisers"
	"github.com/ut3k/clientContact/models"
)

func ClientTableView(c *fiber.Ctx) error {

	initializers.ConnectToDataBase()

	var Clients []models.Clients
	pageStr := c.Query("page", "1")
	page, err := strconv.Atoi(pageStr)

	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	sortColumn := c.Query("sort", "id")
	sortOrder := c.Query("order", "asc")

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc"
	}

	filterName := c.Query("name", "")

	query := initializers.DB.Offset(offset).Limit(pageSize)

	query = query.Order(sortColumn + " " + sortOrder)

	if filterName != "" {
		query = query.Where("name LIKE ?", "%"+filterName+"%h")
	}

	result := query.Find(&Clients)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Databes Query Error")
	}

	prev := page - 1
	if prev < 1 {
		prev = 1
	}

	return c.Render("tablelist", fiber.Map{
		"Clients":    Clients,
		"BaseURL":    c.BaseURL(),
		"Current":    page,
		"Next":       page + 1,
		"Prev":       page,
		"SortColumn": sortColumn,
		"SortOrder":  sortOrder,
		"FilterName": filterName,
	})

}
