package handlers

import (
	"fmt"
	"math/rand"

	"github.com/aixoio/url-short-speed/server/db"
	"github.com/gofiber/fiber/v2"
)

type link_req struct {
	Long string `json:"long"`
}

func NewLinkHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var dat link_req
	c.BodyParser(&dat)

	d := db.DB

	link := db.Link{
		LongLink:  dat.Long,
		ShortLink: generateShortLink(),
	}

	d.Create(&link)

	return c.Status(fiber.StatusCreated).JSON(map[string]string{
		"short": fmt.Sprintf("/api/follow/%s", link.ShortLink),
	})
}

func generateShortLink() string {
	const letters = "QWERTYUIOPASDFGHJKLZXCVBNMmnbvcxzasdfghjklpoiuytrewq1234567890"
	out := ""
	for i := 0; i < 8; i++ {
		out += string(letters[rand.Intn(len(letters))])
	}

	return out
}
