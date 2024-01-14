package middleware

import (
	"backend/controllers"
	"backend/models"
	"context"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func videoFilmRouter(api fiber.Router, collectionName string) {
	root := filmRoot
	FilmCollection := controllers.GetCollection(collectionName)

	// video
	api.Get(fmt.Sprintf("/video/%s/:url", root), func(c *fiber.Ctx) error {
		var film models.Film
		url := "/" + c.Params("url")
		singleResult := FilmCollection.FindOne(context.TODO(), bson.M{"url": url})

		//si aucun resultat
		if err := singleResult.Decode(&film); err != nil {
			c.SendStatus(http.StatusNotFound)
			return nil
		}

		err := c.SendFile(film.Source, false)
		if err != nil {
			return err
		}

		return nil
	})
}

func videoCompileRouter(api fiber.Router, collectionName string, pathLink string) {
	root := pathLink
	FilmCollection := controllers.GetCollection(collectionName)

	// video
	api.Get(fmt.Sprintf("/video/%s/:url/:ep", root), func(c *fiber.Ctx) error {
		var film models.Compile
		url := "/" + c.Params("url")
		episodeId, err := strconv.Atoi(c.Params("ep"))
		if err != nil {
			c.Status(http.StatusBadRequest).SendString("episode introuvable")
			return nil
		}
		singleResult := FilmCollection.FindOne(context.TODO(), bson.M{"url": url})

		//si aucun resultat
		if err := singleResult.Decode(&film); err != nil {
			c.SendStatus(http.StatusNotFound)
			return err
		}

		episode := film.Episodes[episodeId%len(film.Episodes)]
		if err := c.SendFile(path.Join(film.Source, episode), false); err != nil {
			return err
		}

		return nil
	})
}
