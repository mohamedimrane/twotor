package middleware

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mohamedimrane/twotor/data"
	"github.com/mohamedimrane/twotor/model"
)

type MiddlewareWrapper struct {
	queries *data.Queries
}

func New(queries *data.Queries) *MiddlewareWrapper {
	return &MiddlewareWrapper{
		queries: queries,
	}
}

func (mw *MiddlewareWrapper) Authenticated(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).Redirect("/")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).Redirect("/")
	}

	claimsId, err := strconv.Atoi(claims["sub"].(string))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).Redirect("/")
	}

	u, err := mw.queries.GetUser(c.Context(), int64(claimsId))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).Redirect("/")
	}

	user := model.User{
		Id:          uint(u.ID),
		Username:    u.Username,
		Email:       u.Email,
		Password:    u.Password,
		DisplayName: u.DisplayName.String,
		Bio:         u.Bio.String,
	}
	c.Locals("user", user)

	return c.Next()
}

func (mw *MiddlewareWrapper) Unauthenticated(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return c.Next()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Next()
	}

	claimsId, err := strconv.Atoi(claims["sub"].(string))
	if err != nil {
		return c.Next()
	}

	_, err = mw.queries.GetUser(c.Context(), int64(claimsId))
	if err != nil {
		return c.Next()
	}

	return c.Status(fiber.StatusPermanentRedirect).Redirect("/home")
}
