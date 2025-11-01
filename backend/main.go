package backend

import (
	config "backend/configs"
	"backend/controllers"
	"backend/middleware"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start() {
	fmt.Println("backend démaré")

	// --------initialisation---------
	config.Setup()
	// -----------------------
	// wd, errGtwd := os.Getwd()
	// fmt.Println("wd: ", wd)
	// if errGtwd != nil {
	// 	panic(errGtwd)
	// }
	// if err := models.Thumble(fmt.Sprintf("%s\\%s", wd, "vid.mp4"), "serie/popol/output.jpeg"); err != nil {
	// 	panic(err)
	// }

	err := controllers.Connect(config.Getenv("MONGO_URI"), config.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))

	// l'API

	api := app.Group("/api")

	middleware.FilmRouter(api, config.Getenv("FILM_BD"))
	middleware.CompileRouter(api, config.Getenv("SERIE_BD"), "serie")
	middleware.CompileRouter(api, config.Getenv("ANIME_BD"), "anime")
	middleware.UsePosterRouter(api)

	api.Get("/contacts", middleware.GetContact)

	//le font
	app.Static("/", config.Getenv("STATIC"))

	port := config.Getenv("PORT")
	app.Listen(":" + port)
}
