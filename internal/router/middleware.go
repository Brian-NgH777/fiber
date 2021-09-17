package router

import "github.com/gofiber/fiber/v2"

func MiddlewareAuth(c *fiber.Ctx) error {
	return c.Status(201).JSON("{}")
}