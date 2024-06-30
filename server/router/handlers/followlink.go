package handlers

import (
	"github.com/aixoio/url-short-speed/server/db"
	"github.com/gofiber/fiber/v2"
)

func FollowLinkHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	short := c.Params("short")

	link := db.Link{
		ShortLink: short,
	}

	db.DB.Where(&link).First(&link)

	return c.Redirect(link.LongLink, fiber.StatusMovedPermanently)
}
