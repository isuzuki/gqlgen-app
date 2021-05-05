package main

import (
	"app/graph"
	"app/graph/generated"
	"app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	initDB()

	log.Printf("connect to http://localhost:%s/graphql for GraphQL API", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initDB() {
	db, err := gorm.Open("postgres", "postgres://gqlgen_user:gqlgen_password@db:5432/gqlgen_db?sslmode=disable")
	if err != nil {
		log.Fatalln("connection failed.", err)
	}

	db.LogMode(true)
	db.AutoMigrate(&models.Item{})

	defer db.Close()
}
