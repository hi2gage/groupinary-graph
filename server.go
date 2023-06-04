package main

import (
	"graphql-api/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

const (
	defaultPort     = "8080"
	defaultCertFile = "cert.pem" // Path to your SSL/TLS certificate file
	defaultKeyFile  = "key.pem"  // Path to your SSL/TLS key file
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"https://studio.apollographql.com"}, // or "*" for wildcard
		AllowCredentials: true,
	}

	corsHandler := cors.New(corsOptions).Handler(srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", corsHandler)

	log.Printf("connect to https://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServeTLS(":"+port, defaultCertFile, defaultKeyFile, nil))
}
