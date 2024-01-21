package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohamedimrane/twotor/views"
)

func Hello(c *fiber.Ctx) error {
	return Render(c, views.Hello("Imrane"))
}

func List(c *fiber.Ctx) error {
	return Render(c, views.List())
}
