package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohamedimrane/twotor/model"
	"github.com/mohamedimrane/twotor/views"
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
		return Render(c, views.CreateAccountErrors(validationErrorsStrings))
	}

	// Create user

	// Authenticate user
	// Return success message

	return nil
}
