package repositories

import (
	"context"
	"time"

	"github.com/cassiusbessa/create-service/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateService(db string, service entities.Service) error {
	collection := Repo.Client.Database(db).Collection("company")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	service.Id = primitive.NewObjectID()
	update := bson.M{"$push": bson.M{"services": service}}
	_, err := collection.UpdateOne(ctx, bson.D{}, update)
	if err != nil {
		cancel()
		return err
	}
	return nil
}
