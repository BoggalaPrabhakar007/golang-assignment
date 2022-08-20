package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ConnString is for connecting to mongo database
const ConnString string = "mongodb://mongoadmin:secret@localhost:27017"

var client *mongo.Client

//createConnection creates the mongodb client object.
func createConnection() error {
	var err error
	clientOptions := options.Client().ApplyURI(ConnString)
	//connecting to mongo database and getting the mongo client
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	return nil
}

//GetConnectionClient returns the mongoDb client.
func GetConnectionClient() (*mongo.Client, error) {
	var err error
	//if connection with mongo database is not there then create the new connection else return the existing client
	if client == nil {
		err = createConnection()
	}
	return client, err
}

//DisconnectConnection will disconnect the connection with the MongoDB repo. Once the connection is disconnected
func DisconnectConnection() error {
	var err error
	if client != nil {
		err = client.Disconnect(context.Background())
	}
	return err
}