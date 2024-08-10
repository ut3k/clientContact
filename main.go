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
)

func init() {
	initializers.ConnectToDataBase()
	initializers.SyncDB()
	initializers.PopulateClients()
}

func main() {
	engine := html.New("./views", ".html")
	// Setup app

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

	//middleware (cokolwiek to jest)
	app.Use(recover.New())
	app.Use(logger.New())

	// Setup routes
	app.Get("/", controllers.ClientListView)
	app.Get("/tabview/", controllers.ClientTableView)

	// setup
	app.Static("/", "./public")

	app.Use(handlers.NotFound)
	log.Fatal(app.Listen(":3000"))

}
