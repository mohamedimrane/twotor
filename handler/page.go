package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/mohamedimrane/twotor/model"
	"github.com/mohamedimrane/twotor/views/page"
)

func (*HandlerWrapper) Index(c *fiber.Ctx) error {
	return Render(c, page.UnloggedIndex())
}

func (*HandlerWrapper) Home(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)

	return Render(c, page.Home(user))
}

func (hw *HandlerWrapper) Twoot(c *fiber.Ctx) error {
	idstr := c.Params("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.Status(fiber.StatusNotFound).Send([]byte(err.Error()))
	}

	trow, err := hw.queries.GetTwootById(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).Send([]byte(err.Error()))
	}

	user := c.Locals("user").(model.User)

	return Render(c, page.Twoot(user, trow))
}
