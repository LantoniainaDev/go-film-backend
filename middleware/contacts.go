package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetContact(c *fiber.Ctx) error {
	var err error

	file, err := os.Open("/ffmpeg.env")
	defer file.Close()

	c.Send(make([]byte, 2))
	return err
}
