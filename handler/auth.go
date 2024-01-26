package handler

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mohamedimrane/twotor/data"
	"github.com/mohamedimrane/twotor/model"
	"github.com/mohamedimrane/twotor/views/page"
	parterr "github.com/mohamedimrane/twotor/views/partial/errors"
)

func (*HandlerWrapper) Index(c *fiber.Ctx) error {
	return Render(c, page.UnloggedIndex())
}

func (*HandlerWrapper) Home(c *fiber.Ctx) error {
	log.Printf("%+v", c.Locals("user").(model.User))
	return Render(c, page.Home())
}

func (*HandlerWrapper) SignUp(c *fiber.Ctx) error {
	return Render(c, page.CreateAccount())
}

func (hw *HandlerWrapper) CreateUser(c *fiber.Ctx) error {
	// Parse user
	u := new(model.User)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	// Validate user
	validationErrors := Validate(hw.validate, u)
	if validationErrors != nil {
		validationErrorsStrings := userValidationErrorsToStrings(validationErrors)
		return Render(c, parterr.CreateAccountErrors(validationErrorsStrings))
	}

	// Create user
	udb, err := hw.queries.CreateUser(c.Context(), data.CreateUserParams{
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(udb.ID)),
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	cookie := &fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(cookie)

	// Return success message
	return Redirect(c, "/home")
}
