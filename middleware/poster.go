package middleware

import "github.com/gofiber/fiber/v2"

func UsePosterRouter(api fiber.Router) {
	api.Static("/poster", "/tmp")
}
