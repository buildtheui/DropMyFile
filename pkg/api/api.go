package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var App *fiber.App

func setupRoutes() {
	// Load extra needed files for the views like css or js
	App.Static("/assets", "./views/assets")

	
	App.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
}

func setUpApis() {
	api := App.Group("/api/v1")

	api.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")

		if err != nil {
			return err
		}

		c.SaveFile(file, "files/" + file.Filename)

		return c.Render("index", fiber.Map{})
	})
}

func RouterInit() *fiber.App {
	// Load templates
	var engine = html.New("./views", ".html")

	// Create a new Fiber app
	App = fiber.New(fiber.Config{
		Views: engine,
	})

	// Call setupRoutes to set up your routes
	setupRoutes()
	setUpApis()

	return App;
}
