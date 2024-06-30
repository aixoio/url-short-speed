package router

import (
	"fmt"
	"os"

	"github.com/aixoio/url-short-speed/server/router/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartRouter() {
	app := fiber.New()

	app.Use(logger.New())

	app.Post("/api/link/new", handlers.NewLinkHandler)
	app.Get("/api/follow/:short", handlers.FollowLinkHandler)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
