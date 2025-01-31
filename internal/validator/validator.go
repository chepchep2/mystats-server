package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}

func ValidateRequest(c *fiber.Ctx, data interface{}) error {
	if err := c.BodyParser(data); err != nil {
		return err
	}
	return ValidateStruct(data)
}
