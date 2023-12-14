package global

import (
	"github.com/buildtheui/DropMyFile/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

var GSession = utils.GenerateRandomString(6)

func IsvalidSession(session string) bool {
	return GSession == session
}

func ValidateSession(c *fiber.Ctx) error {
	if !IsvalidSession(c.Query("s")) {
		return c.Status(401).SendString("Access denied")
	}

	return c.Next();
}