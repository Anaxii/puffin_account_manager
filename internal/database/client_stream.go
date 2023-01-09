package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

func iterateChangeStream(clientChanges chan primitive.ObjectID, routineCtx context.Context, waitGroup sync.WaitGroup, stream *mongo.ChangeStream) {
	defer stream.Close(routineCtx)
	defer waitGroup.Done()
	for stream.Next(routineCtx) {
		var data bson.M
		if err := stream.Decode(&data); err != nil {
			panic(err)
		}
		id := data["documentKey"].(primitive.M)["_id"].(primitive.ObjectID)
		clientChanges <- id
	}
}

func (d *Database) ClientStream(clientChanges chan primitive.ObjectID) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(d.DBURI))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	database := client.Database("puffin_clients")
	episodesCollection := database.Collection("settings")

	var waitGroup sync.WaitGroup

	episodesStream, err := episodesCollection.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		panic(err)
	}
	waitGroup.Add(1)
	routineCtx, _ := context.WithCancel(context.Background())
	go iterateChangeStream(clientChanges, routineCtx, waitGroup, episodesStream)

	waitGroup.Wait()
}
