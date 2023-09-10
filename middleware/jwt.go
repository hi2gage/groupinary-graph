package middleware

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"shrektionary_api/ent"
	"shrektionary_api/ent/user"
	"time"

	"entgo.io/ent/dialect"
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func EnsureValidToken() func(next http.Handler) http.Handler {
	issuerURL, err := url.Parse("https://dev-afmzazq3cr35ktpl.us.auth0.com/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer URL: %v", err)
	}
	log.Printf("issuerURL: %s", issuerURL)

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{"https://shrektionary.com/api", "UJuMkLZ3wOERGYV4icU0hpIldfcZ07sW"},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the JWT validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				claims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
				if !ok {
					http.Error(w, "failed to get validated claims", http.StatusInternalServerError)
					return
				}
				userID := claims.RegisteredClaims.Subject
				exists := checkUserExists(userID)
				if !exists {
					addUserToGraph(userID)
				}
				next.ServeHTTP(w, r)
			}))
	}
}

func checkUserExists(userID string) bool {
	client, err := ent.Open(dialect.Postgres, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	defer client.Close()

	exists, err := client.User.Query().Where(user.AuthID(userID)).Exist(context.Background())
	if err != nil {
		log.Fatal("querying user", err)
	}
	return exists
}

func addUserToGraph(userID string) {
	client, err := ent.Open(dialect.Postgres, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	defer client.Close()

	_, err = client.User.Create().SetAuthID(userID).Save(context.Background())

	if err != nil {
		log.Fatal("querying user", err)
	}
	log.Printf("finished addUserToGraph")
}
