package login

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mokshsinghji/go-google-oauth/internal/services"
)

func WithGoogle(c *fiber.Ctx) error {
	err := services.RedirectToGoogleLogin(c)
	if err != nil {
		return err
	}
	return nil
}

func GoogleCallback(c *fiber.Ctx) error {
	err := services.GoogleLoginCallback(c)
	if err != nil {
		return err
	}

	return nil
}
