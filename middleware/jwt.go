package middleware

import (
	"context"
	"fmt"
	"groupinary/ent"
	"groupinary/ent/user"
	"groupinary/utils"
	"log"
	"net/http"
	"net/url"
	"time"

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

func parseIssuerURL(rawURL string) *url.URL {
	issuerURL, err := url.Parse(rawURL)
	if err != nil {
		log.Fatalf("Failed to parse the issuer URL: %v", err)
	}
	log.Printf("issuerURL: %s", issuerURL)
	return issuerURL
}

func setupJWTValidator(issuerURL *url.URL) (*validator.Validator, error) {
	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{"https://shrektionary.com/api", "4W01gsxupS4xoLLxbe8jdVVlGTFOKjd3"},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return &CustomClaims{}
		}),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to set up the JWT validator: %w", err)
	}
	return jwtValidator, nil
}

func handleValidationError(w http.ResponseWriter, err error) {
	log.Printf("Encountered error while validating JWT: %v", err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`{"message":"Failed to validate JWT."}`))
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func EnsureValidToken(client *ent.Client) func(next http.Handler) http.Handler {
	issuerURL := parseIssuerURL("https://dev-afmzazq3cr35ktpl.us.auth0.com/")
	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		handleValidationError(w, err)
	}

	jwtValidator, err := setupJWTValidator(issuerURL)
	if err != nil {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				errorHandler(w, r, err)
			})
		}
	}

	middleware := createJWTMiddleware(jwtValidator, errorHandler, client)

	return validateJWTAndHandleUser(client, middleware)
}

func createJWTMiddleware(jwtValidator *validator.Validator, errorHandler jwtmiddleware.ErrorHandler, client *ent.Client) *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)
}

func validateJWTAndHandleUser(client *ent.Client, middleware *jwtmiddleware.JWTMiddleware) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				claims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
				if !ok {
					http.Error(w, "failed to get validated claims", http.StatusInternalServerError)
					return
				}
				authID := claims.RegisteredClaims.Subject

				userId, err := checkUserExists(client, authID)

				if userId == nil && err != nil {
					userId, err = addUserToGraph(client, authID)
					if err != nil {
						log.Printf("Error adding user to graph: %v", err)
						http.Error(w, "Internal server error", http.StatusInternalServerError)
						return
					}
				}

				ctx := utils.AddUserIdToContext(r.Context(), *userId)

				r = r.WithContext(ctx)

				next.ServeHTTP(w, r)
			}))
	}
}

// Based on the auth string passed in, checks to see if that authID exists in the Graph
func checkUserExists(client *ent.Client, authID string) (*int, error) {
	user, err := client.User.Query().Where(user.AuthID(authID)).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("user does not exist: %w", err)
		}
		return nil, fmt.Errorf("querying user: %w", err)
	}
	return &user.ID, nil
}

// Adds the user to Graph with the AuthID if does not exist
func addUserToGraph(client *ent.Client, authID string) (*int, error) {
	user, err := client.User.Create().SetAuthID(authID).Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}
	log.Printf("finished addUserToGraph")
	return &user.ID, nil
}
