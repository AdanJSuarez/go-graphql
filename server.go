package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AdanJSuarez/go-graphql/graph"
	"github.com/AdanJSuarez/go-graphql/graph/generated"
	"github.com/AdanJSuarez/go-graphql/graph/model"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := model.Database{}
	db.Init()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: &db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/query for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
