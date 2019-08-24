package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rkania3/go-nosql-manager/conn"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("Starting tests...")

	db := conn.Init{
		"mongodb://localhost:27017",
	}

	client := db.Connect()

	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
}
