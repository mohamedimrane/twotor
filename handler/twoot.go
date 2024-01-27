package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohamedimrane/twotor/model"
	// parterr "github.com/mohamedimrane/twotor/views/partial/errors"
)

func (hw *HandlerWrapper) CreateTwoot(c *fiber.Ctx) error {
	// Parse twoot
	t := new(model.Twoot)
	if err := c.BodyParser(t); err != nil {
		return err
	}

	// Validate twoot
	errsStr := []string{}
	errs := Validate(hw.validate, t)
	if errs != nil {
		errsStr = append(errsStr, userValErrsToStrings(errs)...)
	}

	// return errors if needed
	if len(errsStr) != 0 {
		// return Render(c, parterr.CreateTwootErrors(errsStr))
		return nil
	}

	// Create twoot
	_, err := hw.queries.CreateTwoot(c.Context(), t.Contents)
	if err != nil {
		return nil
	}

	// Return success message
	return nil
}
