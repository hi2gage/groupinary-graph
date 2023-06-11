package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"shrektionary_api/ent"
	"shrektionary_api/ent/migrate"
	"shrektionary_api/graph"
	"shrektionary_api/middleware"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq" // Import the PostgreSQL driver package

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.Postgres, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	// Seed the data

	defer client.Close()
	ctx := context.Background()

	if err := client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	if err := seedData(ctx, client); err != nil {
		log.Fatal("seeding data", err)
	}

	srv := handler.NewDefaultServer(graph.NewSchema(client))

	// Middleware for logging requests
	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}

	// corsOptions := cors.Options{
	// 	AllowedOrigins:   []string{"studio.apollographql.com"}, // or "*" for wildcard
	// 	AllowCredentials: true,
	// }
	// corsHandler := cors.New(corsOptions)
	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set the Access-Control-Allow-Origin header
			// w.Header().Set("Access-Control-Allow-Origin", "https://studio.apollographql.com")
			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}

	authMiddleware := middleware.EnsureValidToken()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", loggingMiddleware(corsHandler(authMiddleware(srv))))

	log.Printf("connect to https://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServeTLS(":"+port, defaultCertFile, defaultKeyFile, nil))
}

func seedData(ctx context.Context, client *ent.Client) error {
	// Reset the database schema
	if err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		return err
	}

	// Delete records from the "Word" table
	_, errDelete := client.Word.
		Delete().
		Exec(ctx)
	if errDelete != nil {
		return errDelete
	}

	// Create definitions
	definitions_shrekt := []*ent.Definition{
		client.Definition.
			Create().
			SetDescription("Definition 1 for Shrek").
			SaveX(ctx),
		client.Definition.
			Create().
			SetDescription("Definition 2 for Shrek").
			SaveX(ctx),
		client.Definition.
			Create().
			SetDescription("Definition 3 for Shrek").
			SaveX(ctx),
	}

	definitions_bot := []*ent.Definition{
		client.Definition.
			Create().
			SetDescription("Definition 1 for Bot").
			SaveX(ctx),
		client.Definition.
			Create().
			SetDescription("Definition 2 for Bot").
			SaveX(ctx),
		client.Definition.
			Create().
			SetDescription("Definition 3 for Bot").
			SaveX(ctx),
	}

	// Create words and associate definitions
	words := []*ent.WordCreate{
		client.Word.
			Create().
			SetDescription("Shrek").
			AddDefinitions(
				definitions_shrekt[0],
				definitions_shrekt[1],
				definitions_shrekt[2],
			),
		client.Word.
			Create().
			SetDescription("Bot").
			AddDefinitions(
				definitions_bot[0],
				definitions_bot[1],
				definitions_bot[2],
			),
	}

	_, err := client.Word.CreateBulk(words...).Save(ctx)
	return err
}
