package handler

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohamedimrane/twotor/data"
	"github.com/mohamedimrane/twotor/model"
	views "github.com/mohamedimrane/twotor/views/page"
	errpartial "github.com/mohamedimrane/twotor/views/partial/errors"
)

func (*HandlerWrapper) Index(c *fiber.Ctx) error {
	const isLoggedIn = false

	if isLoggedIn {
		return Render(c, views.Home())
	} else {
		return Render(c, views.UnloggedIndex())
	}
}

func (*HandlerWrapper) SignUp(c *fiber.Ctx) error {
	return Render(c, views.CreateAccount())
}

func (hw *HandlerWrapper) CreateUser(c *fiber.Ctx) error {
	// Parse user
	u := new(model.User)
	if err := c.BodyParser(u); err != nil {
		log.Println(err)
		return err
	}

	// Validate user
	validationErrors := Validate(hw.validate, u)
	if validationErrors != nil {
		validationErrorsStrings := userValidationErrorsToStrings(validationErrors)
		return Render(c, errpartial.CreateAccountErrors(validationErrorsStrings))
	}

	// Create user
	_, err := hw.queries.CreateUser(c.Context(), data.CreateUserParams{
		Username:    u.Username,
		Email:       u.Email,
		Password:    u.Password,
		DisplayName: sql.NullString{String: u.DisplayName},
		Bio:         sql.NullString{String: u.Bio},
	})
	if err != nil {
		return err
	}

	// Authenticate user
	// Return success message

	return c.Redirect("/")
}
