package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohamedimrane/twotor/views"
)

func Index(c *fiber.Ctx) error {
	return Render(c, views.Home())
}
