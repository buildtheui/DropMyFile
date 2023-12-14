package api

import (
	"fmt"

	"github.com/buildtheui/DropMyFile/pkg/global"
	"github.com/buildtheui/DropMyFile/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var App *fiber.App

func setupRoutes() {
	// Load extra needed files for the views like css or js
	App.Static("/assets", "./views/assets")

	
	App.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{ "Session": global.GSession })
	})
}

func setUpApis() {
	api := App.Group("/api/v1")

	api.Post("/upload", func(c *fiber.Ctx) error {		

		form, err := c.MultipartForm()

		if err != nil {
			return err
		}

		files := form.File["files"]

		for _, file := range files {

			err := c.SaveFile(file, fmt.Sprintf("./files/%s", utils.RenameFileToUnique(file.Filename)))

			if err != nil {
				return err
			}
		}

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

	// Only LAN users can access with the correct session printed in console
	App.Use(global.ValidateSession)

	// Call setupRoutes to set up your routes
	setupRoutes()
	setUpApis()

	return App;
}
