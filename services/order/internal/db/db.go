package db

import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var OrdersCollection *mongo.Collection

func InitDB(uri string, dbName string) {
    var err error
    Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }

    OrdersCollection = Client.Database(dbName).Collection("orders")
    log.Println("Connected to MongoDB database")
}
