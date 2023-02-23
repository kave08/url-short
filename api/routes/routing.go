package routes

import (
	"github.com/gofiber/fiber/v2"
	"url-shortenr/service"
	)

func SetupRoutes(app *fiber.App) {

	app.Get("/:url",service.ResolverURL)
	app.Post("/api/v1", service.ShortenURL)
}
