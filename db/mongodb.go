package db

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func InsertNumber(num float32) error {

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		return err
	}

	collection := client.Database("testing").Collection("numbers")

	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": num})
	if err != nil {
		return err
	}
	id := res.InsertedID

	fmt.Println("inserted: ", id)
	return nil
}

func GetAllValues() (values []interface{}, err error) {

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		return nil, err
	}

	collection := client.Database("testing").Collection("numbers")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var res bson.M
		err := cur.Decode(&res)
		if err != nil {
			return nil, err
		} else {
			values = append(values, bson.M(res)["value"])
		}
	}

	return values, nil
}
