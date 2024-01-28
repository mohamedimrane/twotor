package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohamedimrane/twotor/data"
	"github.com/mohamedimrane/twotor/model"
	"github.com/mohamedimrane/twotor/views/partial"
	parterr "github.com/mohamedimrane/twotor/views/partial/errors"
)

func (hw *HandlerWrapper) CreateTwoot(c *fiber.Ctx) error {
	// Parse twoot
	t := new(model.Twoot)
	if err := c.BodyParser(t); err != nil {
		return err
	}

	// Validate twoot
	errs := Validate(hw.validate, t)
	if errs != nil {
		return Render(c, parterr.CreateTwootErrors())
	}

	// Get user id
	u := c.Locals("user").(model.User)
	t.UserId = u.Id

	// Create twoot
	_, err := hw.queries.CreateTwoot(c.Context(), data.CreateTwootParams{
		Contents: t.Contents,
		UserID:   int64(t.UserId),
	})
	if err != nil {
		return nil
	}

	// Return success message
	return nil
}

func (hw *HandlerWrapper) ListTwoots(c *fiber.Ctx) error {
	twoots, err := hw.queries.ListTwoots(c.Context())
	if err != nil {
		return err
	}

	return Render(c, partial.TwootsList(twoots))
}
