package main

import (
	config "backend/configs"
	"backend/controllers"
	"backend/middleware"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("backend démaré")

	// --------initialisation---------
	config.Setup()
	// -----------------------
	err := controllers.Connect(config.Getenv("MONGO_URI"), config.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	// l'API
	api := app.Group("/api")

	middleware.FilmRouter(api, "filmStreamTest")
	middleware.CompileRouter(app, "serie", "serie")
	middleware.CompileRouter(api, "anime", "anime")
	middleware.UsePosterRouter(api)

	//le font
	app.Static("/", "/build")

	app.Listen(":3000")
}
