package main

import (
	"log"

	"github.com/buildtheui/DropMyFile/network"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	// Load templates
	engine := html.New("./views", ".html")


	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Load extra needed files for the views like css or js
	app.Static("/assets", "./views/assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	network.PrintLanServerIpQr()

	log.Fatal(app.Listen(":" + network.GetServerPort()))
}