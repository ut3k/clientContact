package main

import (
	"html/template"
	"os"
	"path/filepath"

	"log"

	// "fmt"
	"game_politico/controllers"
	"game_politico/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/recipes/template-asset-bundling/handlers"
	"github.com/gofiber/template/html/v2"
)
