package conn

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Struct for future DB connection params
type Init struct {
	URI string
}

// Struct for mongoDB collections
type collection struct {
	db         string
	collection string
}

func (data *Init) Connect() *mongo.Client {
	// Client options
	clientOptions := options.Client().ApplyURI(data.URI)

	// Connect to mongoDB
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Set timeout for connection
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	// Check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

// func close(client *mongo.Client) bool {
// 	err = client.Disconnect(context.TODO())

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connection to MongoDB closed!")

// 	return true
// }

func SetCollection(data *collection, client *mongo.Client) bool {
	collection := client.Database(data.db).Collection(data.collection)

	fmt.Printf("Set collection in %s to %s", data.db, data.collection)
	fmt.Println(collection)

	return true
}
