package controllers

import (
	"backend/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var DB *mongo.Database

func Connect(mongoDBUrl string, DBName string) error {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(mongoDBUrl))
	if err != nil {
		return nil
	}

	if err = client.Connect(context.TODO()); err != nil {
		return err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return err
	}
	fmt.Println("MongoDB connecté")

	DB = client.Database(DBName)
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := DB.Collection(collectionName)

	// s'assurer que les clées uniques restent uniques
	_, err := collection.Indexes().CreateMany(context.TODO(), models.GetIndexes())
	if err != nil {
		panic(err)
	}
	return collection
}
