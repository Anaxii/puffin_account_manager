package database

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"puffin_account_manager/pkg/global"
	"time"
)

type Database struct {
	DBURI string
}

func (d *Database) GetClients() ([]global.ClientSettings, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(d.DBURI))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:CheckIfExists"}).Error("Failed to connect to mongodb client")
		return []global.ClientSettings{}, err
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("puffin_clients").Collection("settings")
	cur, err := requestsCollection.Find(context.TODO(), bson.D{{"status", "active"}})
	if err != nil {
		log.Error(err)
	}

	var results []global.ClientSettings
	for cur.Next(context.TODO()) {
		var result global.ClientSettings
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, result)
	}

	return results, nil
}

func (d *Database) GetClientUsers(clientUUID int) ([]global.ClientUsers, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(d.DBURI))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:CheckIfExists"}).Error("Failed to connect to mongodb client")
		return []global.ClientUsers{}, err
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("puffin_clients").Collection("users")
	cur, err := requestsCollection.Find(context.TODO(), bson.D{{"client", clientUUID}})
	if err != nil {
		log.Error(err)
	}

	var results []global.ClientUsers
	for cur.Next(context.TODO()) {
		var result global.ClientUsers
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, result)
	}

	return results, nil
}
