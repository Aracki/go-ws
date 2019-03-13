package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"github.com/pkg/errors"
)

const (
	dbNAME = "test"
)

func configDB(ctx context.Context) (*mongo.Client, error) {
	uri := fmt.Sprintf(`mongodb://%s`,
		os.Getenv("mongo_host"),
	)
	client, err := mongo.NewClient(uri)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to mongo: %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("mongo client couldn't connect with background context: %v", err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "ping mongodb failed")
	}
	log.Println("mongodb client connected:", uri)

	return client, nil
}

func InsertNumber(num float32) error {

	ctx := context.Background()
	dbClient, err := configDB(ctx)
	if err != nil {
		return err
	}
	defer dbClient.Disconnect(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	res, err := dbClient.Database(dbNAME).Collection("numbers").InsertOne(ctx, bson.M{"name": "pi", "value": num})
	if err != nil {
		return err
	}

	fmt.Println("inserted: ", res.InsertedID)
	return nil
}

func GetAllValues() (values []interface{}, err error) {

	ctx := context.Background()
	dbClient, err := configDB(ctx)
	if err != nil {
		return nil, err
	}
	defer dbClient.Disconnect(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	cur, err := dbClient.Database(dbNAME).Collection("numbers").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var res bson.M
		err := cur.Decode(&res)
		if err != nil {
			return nil, err
		}
		values = append(values, bson.M(res)["value"])
	}

	return values, nil
}
