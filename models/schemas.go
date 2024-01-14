package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Schema interface{}

type Film struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	Title    string   `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	Url      string   `json:"url,omitempty" bson:"url,omitempty" validate:"required"`
	Poster   string   `json:"poster,omitempty" bson:"poster,omitempty"`
	Cover    string   `json:"cover,omitempty" bson:"cover,omitempty"`
	Synopsis string   `json:"synopsis,omitempty" bson:"synopsis,omitempty" validate:"required"`
	Price    int      `json:"price,omitempty" bson:"price,omitempty"`
	Source   string   `json:"source,omitempty" bson:"source,omitempty" validate:"required"`
	Category []string `json:"genres,omitempty" bson:"genres,omitempty"`
}

type Compile struct {
	Film     `bson:",inline"`
	Episodes []string `json:"episodes" bson:"episodes,omitempty"`
}

func (c *Compile) UnmarshalBSON(data []byte) error {
	// pour eviter les boucles infinies
	type Alias Compile
	var filmClone Film

	var clone Alias

	// fmt.Println("unMarshaling")
	// deserialise les proprietes de Compile
	err := bson.Unmarshal(data, &clone)
	if err != nil {
		return err
	}

	// deserialise les proprietes incororées depuis film par Compile
	err = bson.Unmarshal(data, &filmClone)
	if err != nil {
		return err
	}
	c.Film = filmClone

	// injecte les proprietes de Compiles
	c.Episodes = clone.Episodes
	return nil
}

func (f *Film) Validate() error {
	//validation du film
	f.Title = validateTitle(f.Title)

	// validation de l'URL
	f.Url = validateUrl(f.Url)

	// validation du poster
	f.Poster = slash(f.Poster)

	// validation du cover
	if f.Cover == "" {
		f.Cover = f.Poster
	}

	// validation de la source
	if err := validaterSource(f.Source, false); err != nil {
		return err
	}

	return nil
}

func (f *Compile) Validate() error {
	//validation de la compile
	f.Title = validateTitle(f.Title)

	// validation de l'URL
	f.Url = validateUrl(f.Url)

	// validation du poster
	f.Poster = slash(f.Poster)

	// validation du cover
	if f.Cover == "" {
		f.Cover = f.Poster
	}

	// validation de la source
	if err := validaterSource(f.Source, true); err != nil {
		return err
	}

	return nil
}

func GetIndexes() mongo.IndexModel {
	indexOption := options.Index().SetUnique(true)
	return mongo.IndexModel{
		Keys:    bson.M{"title": 2},
		Options: indexOption,
	}
}
