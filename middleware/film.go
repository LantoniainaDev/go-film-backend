package middleware

import (
	"backend/controllers"
	"backend/models"
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var filmRoot string = "film"

func FilmRouter(app fiber.Router, collectionName string) error {
	root := filmRoot
	FilmCollection := controllers.GetCollection(collectionName)
	validate := validator.New()

	// read all
	app.Get(fmt.Sprintf("/%s", root), func(c *fiber.Ctx) error {
		var films []models.Film
		cursor, err := FilmCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			return err
		}

		cursor.All(context.TODO(), &films)
		c.JSON(films)
		return nil
	})

	// read
	app.Get(fmt.Sprintf("/%s/:id", root), func(c *fiber.Ctx) error {
		var film models.Film

		var id string = c.Params("id")

		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}

		ctx := context.TODO()
		result := FilmCollection.FindOne(ctx, bson.M{"_id": objectId})

		if err := result.Decode(&film); err != nil {
			return err
		}

		c.JSON(film)
		return nil
	})

	// create
	app.Post(fmt.Sprintf("/%s", root), func(c *fiber.Ctx) error {
		var film models.Film

		if err := c.BodyParser(&film); err != nil {
			return err
		}

		// validation du schema
		if err := film.Validate(); err != nil {
			return err
		}

		if err := validate.Struct(&film); err != nil {
			c.SendStatus(http.StatusNotImplemented)
			return err
		}

		// process du schema avant la modif de la base de donnée
		if err := film.PreSave(); err != nil {
			return nil
		}

		film.Id = primitive.NewObjectID()

		if _, err := FilmCollection.InsertOne(context.TODO(), film); err != nil {
			return err
		}

		c.JSON(film)

		return nil
	})

	// delete
	app.Delete(fmt.Sprintf("/%s/:id", root), func(c *fiber.Ctx) error {
		id := c.Params("id")
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}

		res, err := FilmCollection.DeleteOne(context.TODO(), bson.M{
			"_id": objId,
		})
		if err != nil {
			return err
		}

		c.JSON(res.DeletedCount)

		return nil
	})

	// update
	app.Patch(fmt.Sprintf("/%s/:id", root), func(c *fiber.Ctx) error {
		var film models.Film

		id, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return err
		}

		res := FilmCollection.FindOne(context.TODO(), bson.M{"_id": id}, nil)
		res.Decode(&film)

		if err := c.BodyParser(&film); err != nil {
			return err
		}

		if err := film.Validate(); err != nil {
			return err
		}

		// process du schema avant la modif de la base de donnée
		var modif models.Modif

		c.BodyParser(&modif)
		shouldSkip := modif.PosterImg != ""

		fmt.Println("skiped? ", shouldSkip)

		if err := film.PreSave(shouldSkip); err != nil {
			return nil
		}

		// modification du poster
		if modif.PosterImg != "" {
			if err = CopyImage(modif.PosterImg, "tmp", "film", film.Poster); err != nil {
				return err
			}
		}

		// ---------------------

		updateBson := bson.M{"$set": film}
		if err != nil {
			return nil
		}

		_, err = FilmCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, updateBson)
		if err != nil {
			return err
		}

		c.Status(http.StatusCreated).JSON(film)

		return nil
	})

	// video
	videoFilmRouter(app, collectionName)

	return nil
}
