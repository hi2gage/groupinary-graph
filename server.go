package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"groupinary/ent"
	"groupinary/ent/migrate"
	"groupinary/graph/resolvers"
	"groupinary/middleware"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import the PostgreSQL driver package
)

const (
	defaultPort     = "8080"
	defaultCertFile = "cert.pem" // Path to your SSL/TLS certificate file
	defaultKeyFile  = "key.pem"  // Path to your SSL/TLS key file
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	environment := os.Getenv("ENVIRONMENT")
	log.Printf("Starting Server in: %s ENVIRONMENT", environment)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.Postgres, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	defer client.Close()

	// Seed the data

	ctx := context.Background()

	if err := client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	srv := handler.NewDefaultServer(resolvers.NewSchema(client))

	// Middleware for logging requests
	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			showLogs := false

			if showLogs {
				// Log the incoming request
				log.Printf("Received request: %s %s", r.Method, r.URL.Path)

				// Log the headers for preflight request or actual request
				if r.Method == http.MethodOptions {
					log.Println("Preflight Request Headers:")
				} else {
					log.Println("Actual Request Headers:")
				}

				// Print each header key-value pair
				for name, values := range r.Header {
					// Join multiple values for the same header with commas
					value := strings.Join(values, ", ")
					log.Printf("%s: %s", name, value)
				}
				log.Printf("==========================================================================================")
			}
			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}

	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set the Access-Control-Allow-Origin header
			// w.Header().Set("Access-Control-Allow-Origin", "https://studio.apollographql.com")
			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}

	issueURL := os.Getenv("ISSUERURL")
	audienceAPI := os.Getenv("AUDIENCE_API")
	audienceHash := os.Getenv("AUDIENCE_HASH")

	jwtENV := middleware.EnvJWTStruct{
		IssuerURL: issueURL,
		Audience:  []string{audienceAPI, audienceHash},
	}

	userTokenOperations := middleware.NewUserTokenOperator(client)
	authMiddleware := middleware.EnsureValidToken(userTokenOperations, jwtENV)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/viz", ent.ServeEntviz())
	http.Handle("/query", loggingMiddleware(corsHandler(authMiddleware(srv))))

	log.Printf("connect to https://localhost:%s/ for GraphQL playground", port)

	if environment == "dev" {
		log.Printf("started in TLS mode")
		log.Fatal(http.ListenAndServeTLS(":"+port, defaultCertFile, defaultKeyFile, nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}

}
