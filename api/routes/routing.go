package routes

import (
	"url-shortenr/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/:url", handler.ResolverURL)
	app.Post("/api/v1", handler.ShortenURL)
}
