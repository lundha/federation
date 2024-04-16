//go:generate go run ../../../testdata/gqlgen.go
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"example.com/federation-demo/users/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/gqlerror"
)

const defaultPort = "4002"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := graph.NewResolver()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log.Printf("users - recovered error: %v", err)
		return gqlerror.Errorf("internal server error")
	})

	var handle http.Handler = srv
	handle = graph.UserLoaderMiddleware(resolver.Repo, handle)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handle)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
