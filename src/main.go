package main

import (
	"app/graph"
	"app/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/graphql for GraphQL API", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
