package main

import (
	"log"
	"net/http"
	"os"
	"test/SkyScraper/graph"
	"test/SkyScraper/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("Database access", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for interactive access", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
