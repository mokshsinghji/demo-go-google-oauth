package handlers

import "github.com/gofiber/fiber/v2"

func HomePage(c *fiber.Ctx) error {
	err := c.JSON(fiber.Map{
		"message": "Hello World! Go to /login/google to authenticate",
	})

	if err != nil {
		return err
	}

	return nil
}
