package repository

import (
	"context"
	"time"

	"github.com/cassiusbessa/create-service/entity"
	"github.com/cassiusbessa/create-service/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func MongoConnection() (*mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	return client, cancel
}

func CreateService(db string, service entity.Service) error {
	collection := client.Database(db).Collection("company")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	update := bson.M{"$push": bson.M{"services": service}}
	_, err := collection.UpdateOne(ctx, bson.D{}, update)
	if err != nil {
		cancel()
		return err
	}
	defer cancel()
	return nil
}

func SaveError(db string, customErr errors.CustomError) {
	collection := client.Database(db).Collection("errors")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.InsertOne(ctx, customErr)
	if err != nil {
		cancel()
	}
	defer cancel()
}
