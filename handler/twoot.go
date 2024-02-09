package handler

import (
	"database/sql"

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

	pti := t.TwootId
	if pti != 0 {
		_, err := hw.queries.GetTwootById(c.Context(), int64(pti))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).Send([]byte("no no no"))
		}
	}

	// Get user id
	u := c.Locals("user").(model.User)
	t.UserId = u.Id

	// Create twoot
	var ptidb sql.NullInt64
	if pti == 0 {
		ptidb = sql.NullInt64{Valid: false}
	} else {
		ptidb = sql.NullInt64{
			Int64: int64(pti),
			Valid: true,
		}
	}

	_, err := hw.queries.CreateTwoot(c.Context(), data.CreateTwootParams{
		Contents: t.Contents,
		UserID:   int64(t.UserId),
		TwootID:  ptidb,
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
