package main

import (
	"fmt"
	conn "golang-db/conn"
)

func main() {
	fmt.Println("Starting tests...")

	db := conn.Init("mongodb://localhost:27017")

	client := db.connect()

	fmt.Printf("Client: %v", client)
}
