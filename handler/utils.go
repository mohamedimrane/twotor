package handler

import (
	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/mohamedimrane/twotor/data"
)

type HandlerWrapper struct {
	queries  *data.Queries
	validate *validator.Validate
}

func New(queries *data.Queries) *HandlerWrapper {
	return &HandlerWrapper{
		queries:  queries,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)

	for _, o := range options {
		o(componentHandler)
	}

	return adaptor.HTTPHandler(componentHandler)(c)
}

// Using HTMX requires setting the "HX-Redirect" response header to redirect the browser.
// Not doing that will result to the redirected page being rendered inside the HTMX target element.
func Redirect(c *fiber.Ctx, location string) error {
	c.Set("HX-Redirect", location)
	return nil
}

type ValidationError struct {
	ActualField string
	Field       string
	ActualTag   string
	Tag         string
	Error       string
	Value       interface{}
}

func Validate(v *validator.Validate, d interface{}) []ValidationError {
	var validationErrors []ValidationError

	errs := v.Struct(d)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var ve = ValidationError{
				ActualField: err.StructField(),
				Field:       err.Field(),
				ActualTag:   err.ActualTag(),
				Tag:         err.Tag(),
				Error:       err.Error(),
				Value:       err.Value(),
			}

			validationErrors = append(validationErrors, ve)
		}
	}

	return validationErrors
}

func userValidationErrorsToStrings(validationErrors []ValidationError) []string {
	var validationErrorsStrings []string
	for _, err := range validationErrors {
		switch err.ActualField {
		case "Username":
			switch err.ActualTag {
			case "required":
				str := "User name is required"
				validationErrorsStrings = append(validationErrorsStrings, str)
			case "min", "max":
				str := "User name must be 2 to 50 characters long"
				validationErrorsStrings = append(validationErrorsStrings, str)
			}

		case "Email":
			switch err.ActualTag {
			case "required":
				str := "Email is required"
				validationErrorsStrings = append(validationErrorsStrings, str)
			case "email":
				str := "Invalid email"
				validationErrorsStrings = append(validationErrorsStrings, str)
			}

		case "Password":
			switch err.ActualField {
			case "required":
				str := "Password is required"
				validationErrorsStrings = append(validationErrorsStrings, str)
			case "min", "max":
				str := "Password must be 8 to 32 characters long"
				validationErrorsStrings = append(validationErrorsStrings, str)
			}

		case "DisplayName":
			str := "Display name must be 1 to 50 characters long"
			validationErrorsStrings = append(validationErrorsStrings, str)

		case "Bio":
			str := "Bio cannot be longer than 200 characters"
			validationErrorsStrings = append(validationErrorsStrings, str)
		}
	}

	return validationErrorsStrings
}
