package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestCustomClaimsValidate(t *testing.T) {
	testCases := []struct {
		name        string
		claims      CustomClaims
		expectError bool
	}{
		{
			name:        "Valid CustomClaims",
			claims:      CustomClaims{Scope: "read write"},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Validate method.
			err := tc.claims.Validate(context.Background())

			// Check if the error matches the expectation.
			if (err != nil) != tc.expectError {
				t.Errorf("Expected error: %v, got error: %v", tc.expectError, err)
			}
		})
	}
}

func TestParseIssuerURL(t *testing.T) {
	t.Run("Valid Issuer URL", func(t *testing.T) {
		expectedURL := "https://dev-afmzazq3cr35ktpl.us.auth0.com/"
		issuerURL := parseIssuerURL()

		// Assert that the parsed URL matches the expected URL
		if issuerURL.String() != expectedURL {
			t.Errorf("Expected issuer URL %s, got %s", expectedURL, issuerURL.String())
		}
	})
}

func TestHandleValidationError(t *testing.T) {
	t.Run("Handles Validation Error", func(t *testing.T) {
		// Create a mock response writer
		w := httptest.NewRecorder()

		// Simulate a validation error
		err := someValidationErrorFunction()

		// Call the function to handle the validation error
		handleValidationError(w, err)

		// Check the response status code
		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
		}

		// Check the response body
		expectedResponseBody := `{"message":"Failed to validate JWT."}`
		if w.Body.String() != expectedResponseBody {
			t.Errorf("Expected response body %s, got %s", expectedResponseBody, w.Body.String())
		}

		// Check the response content type header
		if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
			t.Errorf("Expected Content-Type header 'application/json', got '%s'", contentType)
		}
	})
}

// Function to simulate a validation error for testing purposes
func someValidationErrorFunction() error {
	// Replace this with actual logic that might produce a validation error
	return nil
}

func TestSetupJWTValidator(t *testing.T) {
	testCases := []struct {
		name           string
		issuerURL      *url.URL
		expectedError  bool
		expectedNilJWT bool
	}{
		{
			name:           "Valid JWT Validator",
			issuerURL:      mustParseURL("https://mock-auth0-issuer.com/"),
			expectedError:  false,
			expectedNilJWT: false,
		},
		{
			name:           "Error Creating JWT Validator",
			issuerURL:      mustParseURL(""), // Invalid URL to trigger an error
			expectedError:  true,
			expectedNilJWT: true,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jwtValidator, err := setupJWTValidator(tc.issuerURL)

			// Check for expected error
			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got error: %v", tc.expectedError, err)
			}

			// Check for expected nil JWT validator
			if (jwtValidator == nil) != tc.expectedNilJWT {
				t.Errorf("Expected nil JWT validator: %v, got: %v", tc.expectedNilJWT, jwtValidator)
			}
		})
	}
}

// Helper function to parse URL and panic on error (for test simplicity)
func mustParseURL(rawURL string) *url.URL {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return parsedURL
}
