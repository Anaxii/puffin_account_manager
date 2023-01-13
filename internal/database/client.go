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
			log.Error(err)
			continue
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
			log.Error(err)
		}

		results = append(results, result)
	}

	return results, nil
}

func (d *Database) GetUserClients(user string) (map[int]bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(d.DBURI))
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(), "file": "Database:CheckIfExists"}).Error("Failed to connect to mongodb client")
		return map[int]bool{}, err
	}
	defer client.Disconnect(ctx)

	requestsCollection := client.Database("puffin_clients").Collection("users")
	cur, err := requestsCollection.Find(context.TODO(), bson.D{{"user", user}})
	if err != nil {
		log.Error(err)
	}

	clients := map[int]bool{}
	for cur.Next(context.TODO()) {
		var result global.ClientUsers
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		clients[result.Client] = true
	}

	return clients, nil
}

func (d *Database) AddClientUser(c global.ClientUsers) error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(d.DBURI))
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	clientSettings := client.Database("puffin_clients").Collection("users")
	_, err = clientSettings.InsertOne(context.TODO(), c)

	return err
}

func (d *Database) DeleteUser(user string, id int) error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(d.DBURI))
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	clientSettings := client.Database("puffin_clients").Collection("users")
	_, err = clientSettings.DeleteOne(context.TODO(), bson.D{{"user", user},  {"client", id}})
	return err
}

func (d *Database) GetUser(u string) (global.Account, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(d.DBURI))
	if err != nil {
		return global.Account{}, err
	}
	defer client.Disconnect(ctx)

	clientSettings := client.Database("puffin").Collection("account_requests")
	res := clientSettings.FindOne(context.TODO(), bson.D{{"wallet_address", u}}, nil)
	var result global.Account
	err = res.Decode(&result)
	if err != nil {
		return global.Account{}, err
	}
	return result, err
}