package main

import (
	"github.com/buildtheui/DropMyFile/network"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!") 
	})

	network.PrintLanServerIpQr()

	app.Listen(":" + network.GetServerPort())
}