package routes

import (
	"url-shortenr/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/:url",handler.ResolverURL)
	// e.Post("/api/v1", handler.ShortenURL)
}
