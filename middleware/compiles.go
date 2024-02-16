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

func CompileRouter(app fiber.Router, collectionName string, root string) error {
	FilmCollection := controllers.GetCollection(collectionName)
	validate := validator.New()

	// read all
	app.Get(fmt.Sprintf("/%s", root), func(c *fiber.Ctx) error {
		var films []models.Compile
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
		var film models.Compile

		var id string = c.Params("id")

		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}

		ctx := context.TODO()
		result := FilmCollection.FindOne(ctx, bson.M{"_id": objectId})

		result.Decode(&film)

		c.JSON(film)
		return nil
	})

	// create
	app.Post(fmt.Sprintf("/%s", root), func(c *fiber.Ctx) error {
		var film models.Compile

		if err := c.BodyParser(&film); err != nil {
			return err
		}

		// validation du schema
		if err := film.Validate(); err != nil {
			return err
		}

		// process du schema avant la modif de la base de donnée

		if err := validate.Struct(&film); err != nil {
			c.SendStatus(http.StatusNotImplemented)
			return err
		}

		if err := film.Validate(); err != nil {
			return err
		}

		if err := film.PreSave(root); err != nil {
			return err
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
		var film models.Compile

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

		if err := film.PreSave(root); err != nil {
			return err
		}

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
	videoCompileRouter(app, collectionName, root)
	return nil
}
