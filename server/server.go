package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/brianlangdon/tada/cors"
	"github.com/brianlangdon/tada/db"
	"github.com/brianlangdon/tada/graph"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), clientOptions())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	http.Handle("/query", gqlHandler(db.New(client)))
	http.Handle("/playground",
		handler.Playground("GraphQL playground", "/query"),
	)
	http.Handle("/", http.FileServer(http.Dir("/dating")))
	err = http.ListenAndServe(":8080", nil)
	log.Println(err)
}

func gqlHandler(db db.DB) http.HandlerFunc {
	config := graph.Config{
		Resolvers: &graph.Resolver{DB: db},
	}
	gh := handler.GraphQL(graph.NewExecutableSchema(config))
	if os.Getenv("profile") != "prod" {
		gh = cors.Disable(gh)
	}
	return gh
}

func clientOptions() *options.ClientOptions {
	host := "db"
	if os.Getenv("profile") != "prod" {
		host = "localhost"
	}
	return options.Client().ApplyURI(
		"mongodb://" + host + ":27017",
	)
}
