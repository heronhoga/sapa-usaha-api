package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success", 
		"data": "Hello, World!",
	})
}