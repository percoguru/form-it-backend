package user

import "github.com/gofiber/fiber/v2"

func GetUsers(c *fiber.Ctx) error {
	return c.SendString("Get User")
}
