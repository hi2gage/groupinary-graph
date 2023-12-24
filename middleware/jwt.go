package middleware

import (
	"context"
	"fmt"
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

type EnvJWTStruct struct {
	IssuerURL string
	Audience  []string
}

func setupJWTValidator(env EnvJWTStruct) (*validator.Validator, error) {
	issuerURL, err := url.Parse(env.IssuerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the issuer URL: %w", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		env.Audience,
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

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	handleValidationError(w, err)
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func EnsureValidToken(client UserOperations, env EnvJWTStruct) func(next http.Handler) http.Handler {
	jwtValidator, err := setupJWTValidator(env)

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

func createJWTMiddleware(jwtValidator *validator.Validator, errorHandler jwtmiddleware.ErrorHandler, client UserOperations) *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)
}

func validateJWTAndHandleUser(client UserOperations, middleware *jwtmiddleware.JWTMiddleware) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				claims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
				if !ok {
					http.Error(w, "failed to get validated claims", http.StatusInternalServerError)
					return
				}
				authID := claims.RegisteredClaims.Subject

				userId, err := client.CheckUserExists(authID)

				if userId == nil && err != nil {
					userId, err = client.AddUserToGraph(authID)
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
