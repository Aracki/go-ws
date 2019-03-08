package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var client *mongo.Client

func init() {

	dbUrl := "mongodb://" + os.Getenv("host") + ":" + os.Getenv("port")

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	var err error
	client, err = mongo.Connect(ctx, dbUrl)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		log.Printf("connected to %s\n", dbUrl)
	}
}

func InsertNumber(num float32) error {

	if client == nil {
		return errors.New("cannot create mongo client")
	}

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
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

	if client == nil {
		return nil, errors.New("cannot create mongo client")
	}

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
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
