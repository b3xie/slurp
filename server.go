package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/b3xie/slurp/graph"
	model "github.com/b3xie/slurp/graph/model"
	"github.com/b3xie/slurp/internal"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	var user model.NewUser
	user.Username = "TEST"
	user.Secret = "TESTPASS"
	ctx := context.TODO()
	internal.CreateUser(ctx, user)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
