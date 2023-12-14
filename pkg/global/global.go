package global

import (
	"strings"

	"github.com/buildtheui/DropMyFile/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

var GSession = utils.GenerateRandomString(6)
var AllowPaths = []string{"/assets/"}

func IsvalidSession(session string) bool {
	return GSession == session
}

func ValidateSession(c *fiber.Ctx) error {
	path := c.Path()

	for _, allowPath :=range(AllowPaths) {
		if strings.HasPrefix(path, allowPath) {
			return c.Next();
		}
	}

	if !IsvalidSession(c.Query("s")) {
		return c.Status(401).SendString("Access denied")
	}

	return c.Next();
}