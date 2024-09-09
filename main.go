package main

import (
	"html/template"
	"os"
	"path/filepath"

	"log"

	// "fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/recipes/template-asset-bundling/handlers"
	"github.com/gofiber/template/html/v2"

	"github.com/ut3k/clientContact/controllers"
	initializers "github.com/ut3k/clientContact/initialisers"
	"github.com/ut3k/clientContact/utils"
)

func init() {
	initializers.ConnectToDataBase()
	initializers.SyncDB()
	initializers.PopulateClients()
}

func main() {
	engine := html.New("./views", ".html")
	// Setup app
	engine.AddFunc("ShortDate", utils.ShortDate)
	engine.Reload(true)
	// Disable on production

	engine.AddFunc("getCssAsset", func(name string) (res template.HTML) {
		filepath.Walk("public", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == name {
				res = template.HTML("<link rel=\"stylesheet\" href=\"/" + path + "\">")
			}
			return nil
		})
		return
	})

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	// inicjalizacja utils

	//middleware (cokolwiek to jest)
	app.Use(recover.New())
	app.Use(logger.New())

	// Setup routes
	app.Get("/", controllers.ClientListView)
	app.Get("/tabview/", controllers.ClientTableView)
	app.Get("/client/:id", controllers.ClientSingle)

	// Client ADD
	app.Get("/clientadd", controllers.ClientAdd)
	app.Post("/clientadd", controllers.ClientAddPost)

	// client Update
	app.Get("/clientupdate/:id", controllers.ClientUpdate)
	app.Post("/clientupdate/:id", controllers.ClientUpdatePost)

	// Client Delete
	app.Get("/clientdelete/:id", controllers.ClientDelete)
	app.Post("/clientdelete/:id", controllers.ClientDeletePost)

	// setup
	app.Static("/", "./public")

	app.Use(handlers.NotFound)
	log.Fatal(app.Listen(":3000"))

}
