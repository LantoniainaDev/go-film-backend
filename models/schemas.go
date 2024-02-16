package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Schema interface{}

type Film struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	Title    string    `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	Url      string    `json:"url,omitempty" bson:"url,omitempty" validate:"required"`
	Poster   string    `json:"poster,omitempty" bson:"poster,omitempty"`
	Cover    string    `json:"cover,omitempty" bson:"cover,omitempty"`
	Synopsis string    `json:"synopsis,omitempty" bson:"synopsis,omitempty" validate:"required"`
	Price    int       `json:"price,omitempty" bson:"price,omitempty"`
	Source   string    `json:"source,omitempty" bson:"source,omitempty" validate:"required"`
	Category []string  `json:"genres,omitempty" bson:"genres,omitempty"`
	Date     time.Time `json:"published" bson:"published"`
}

type Compile struct {
	Film     `bson:",inline"`
	Episodes []string `json:"episodes" bson:"episodes,omitempty"`
}

type Modif struct {
	PosterImg string `json:"posterImage,omitempty" bson:"posterImage,omitempty"`
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

// staitc methods
func GetIndexes() []mongo.IndexModel {
	indexOption := options.Index().SetUnique(true)
	return []mongo.IndexModel{
		{
			Keys:    bson.M{"title": 2},
			Options: indexOption,
		},
		{
			Keys:    bson.M{"url": 3},
			Options: options.Index().SetUnique(true),
		},
	}
}
