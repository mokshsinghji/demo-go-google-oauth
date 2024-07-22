package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mokshsinghji/go-google-oauth/internal/handlers"
	"github.com/mokshsinghji/go-google-oauth/internal/handlers/login"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", handlers.HomePage)
	app.Get("/login/google", login.LoginWithGoogle)
}
