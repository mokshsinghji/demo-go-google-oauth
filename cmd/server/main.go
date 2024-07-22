package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mokshsinghji/go-google-oauth/internal/config"
	"github.com/mokshsinghji/go-google-oauth/internal/routes"
)

func main() {
	var app = fiber.New()

	config.InitEnvVariables()
	routes.RegisterRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
