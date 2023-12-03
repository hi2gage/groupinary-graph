package middleware

import (
	"context"
	"groupinary/testutils"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/assert"
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
	// rr := httptest.NewRecorder()
	testCases := []struct {
		name          string
		expectedURL   string
		expectedError string
	}{
		{
			name:          "Correct URL",
			expectedURL:   "https://dev-afmzazq3cr35ktpl.us.auth0.com/",
			expectedError: "",
		},
		{
			name:          "Blank URL",
			expectedURL:   "",
			expectedError: "",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run("Valid Issuer URL", func(t *testing.T) {
			issuerURL := parseIssuerURL(tc.expectedURL)

			// Assert that the parsed URL matches the expected URL
			if issuerURL.String() != tc.expectedURL {
				t.Errorf("Expected issuer URL %s, got %s", tc.expectedURL, issuerURL.String())
			}
		})
	}
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

func TestCheckUserExists(t *testing.T) {
	fixturePaths := []string{
		"fixtures/users.yaml",
		"fixtures/groups.yaml",
	}

	client, db, err := testutils.OpenTest()
	if err != nil {
		t.Fatal(err)
	}
	// Register the cleanup function from testutils.
	t.Cleanup(func() {
		testutils.CleanupTestEnvironment(t, client)
	})

	expectedId := 50

	// Test the checkUserExists function
	testCases := []struct {
		name          string
		authID        string
		expected      *int
		expectedError string
	}{
		{
			name:          "User exists",
			authID:        "test_auth_id_1",
			expected:      &expectedId,
			expectedError: "",
		},
		{
			name:          "User does not exist",
			authID:        "nonexistentAuthID",
			expected:      nil,
			expectedError: "user does not exist: ent: user not found",
		},
		{
			name:          "Multiple Users with same authID",
			authID:        "test_auth_id_duplicate",
			expected:      nil,
			expectedError: "querying user: ent: user not singula",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			id, err := checkUserExists(client, tc.authID)

			if err != nil {
				assert.Error(t, err, "Expected error")
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain expected string")
				assert.Nil(t, id, "User should be nil when there is an error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NotNil(t, id, "User should not be nil when there is no error")

			}
			// if err != nil {
			// 	assert.Error(t, err, "Expected error")
			// } else {
			// 	assert.Equal(t, false, tc.expectedErr, "expectedErr should be empty // Got: %v", tc.expectedErr)
			// 	assert.NoError(t, err, "Unexpected error")
			// }

			// if tc.wantErr {
			// 	if err == nil {
			// 		t.Fatalf("Expected an error but got nil")
			// 	}
			// } else {
			// 	if err != nil {
			// 		t.Fatalf("Did not expect an error but got: %v", err)
			// 	}
			// }
			// if id != tc.expected {
			// 	t.Fatalf("Expected %d but got %d", tc.expected, id)
			// }
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
