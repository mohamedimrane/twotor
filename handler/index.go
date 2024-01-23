package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohamedimrane/twotor/views"
)

func Index(c *fiber.Ctx) error {
	const isLoggedIn = false

	if isLoggedIn {
		return Render(c, views.Home())
	} else {
		return Render(c, views.UnloggedIndex())
	}
}
