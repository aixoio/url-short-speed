package router

import (
	"fmt"
	"os"

	"github.com/aixoio/url-short-speed/server/router/handlers"
	"github.com/gofiber/fiber/v2"
)

func StartRouter() {
	app := fiber.New()

	app.Post("/api/link/new", handlers.NewLinkHandler)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
