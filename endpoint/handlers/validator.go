package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateSpotParam(c *fiber.Ctx) error {
	req := new(params)
	if err := c.QueryParser(req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	c.Locals("spot", req)
	return c.Next()
}
