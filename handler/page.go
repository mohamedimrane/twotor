package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/mohamedimrane/twotor/model"
	"github.com/mohamedimrane/twotor/views/page"
)

func (*HandlerWrapper) Index(c *fiber.Ctx) error {
	return Render(c, page.UnloggedIndex())
}

func (*HandlerWrapper) Home(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)

	log.Printf("%+v", user)

	return Render(c, page.Home(user))
}
