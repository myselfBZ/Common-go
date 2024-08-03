package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)





type MongoStore struct{
    db *mongo.Database
}


func ConnectDB() *mongo.Client {
    
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        panic(err)
    }

    return client 
}
