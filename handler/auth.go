package handler

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mohamedimrane/twotor/data"
	"github.com/mohamedimrane/twotor/model"
	"github.com/mohamedimrane/twotor/views/page"
	parterr "github.com/mohamedimrane/twotor/views/partial/errors"
	"golang.org/x/crypto/bcrypt"
)

func (*HandlerWrapper) SignUp(c *fiber.Ctx) error {
	return Render(c, page.CreateAccount())
}

func (*HandlerWrapper) LogIn(c *fiber.Ctx) error {
	return Render(c, page.LogInToAccount())
}

func (*HandlerWrapper) LogOut(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return c.Redirect("/")
}

func (hw *HandlerWrapper) CreateUser(c *fiber.Ctx) error {
	// Parse user
	u := new(model.User)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	// Validate user
	errsStr := []string{}

	// verify email availability
	if _, err := hw.queries.GetUserByEmail(c.Context(), u.Email); err == nil {
		errsStr = append(errsStr, "Email already taken")
	}

	// verify username availability
	if _, err := hw.queries.GetUserByUsername(c.Context(), u.Username); err == nil {
		errsStr = append(errsStr, "Username already taken")
	}

	// validate input
	errs := Validate(hw.validate, u)
	if errs != nil {
		errsStr = append(errsStr, userValErrsToStrings(errs)...)
	}

	// return errors if needed
	if len(errsStr) != 0 {
		return Render(c, parterr.CreateAccountErrors(errsStr))
	}

	// Create user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}

	udb, err := hw.queries.CreateUser(c.Context(), data.CreateUserParams{
		Username:    u.Username,
		Email:       u.Email,
		Password:    string(hashedPassword),
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

func (hw *HandlerWrapper) LogInUser(c *fiber.Ctx) error {
	// Parse user
	u := new(model.User)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	// Get user form database
	udb, err := hw.queries.GetUserByEmail(c.Context(), u.Email)
	if err != nil {
		return Render(c, parterr.LogInToAccountErrors())
	}

	// Verify user
	if err := bcrypt.CompareHashAndPassword([]byte(udb.Password), []byte(u.Password)); err != nil {
		return Render(c, parterr.LogInToAccountErrors())
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
