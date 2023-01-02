package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiberfiber.Ctx) error {
	return c.Render("welcome", nil, "layouts/main")
}
