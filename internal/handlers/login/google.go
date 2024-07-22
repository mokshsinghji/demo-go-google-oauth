package login

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mokshsinghji/go-google-oauth/internal/services"
)

func LoginWithGoogle(c *fiber.Ctx) error {
	err := services.RedirectToGoogleLogin(c)
	if err != nil {
		return err
	}
	return nil
}
