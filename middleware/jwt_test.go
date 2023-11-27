package middleware

import (
	"context"
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
