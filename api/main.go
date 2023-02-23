package main

import (
	"log"
	"os"
	"url-shortenr/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		return
	}

	app.Use(logger.New())

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}
