package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	fmt.Println("Go docker tut")
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://mongo")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Go docker handler")
		fmt.Fprintf(w, "Hello World, ", client.Ping(context.TODO(), nil))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
