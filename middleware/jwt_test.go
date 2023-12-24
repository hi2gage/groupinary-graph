package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

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
		jwtENV         EnvJWTStruct
		expectedError  string
		expectedNilJWT bool
	}{
		{
			name: "Valid JWT Validator",
			jwtENV: EnvJWTStruct{
				IssuerURL: "https://mock-auth0-issuer.com/",
				Audience: []string{
					"https://shrektionary.com/api",
					"4W01gsxupS4xoLLxbe8jdVVlGTFOKjd3",
				},
			},
			expectedError:  "",
			expectedNilJWT: false,
		},
		{
			name: "Empty String URL",
			jwtENV: EnvJWTStruct{
				IssuerURL: "",
				Audience:  []string{},
			},
			expectedError:  "failed to set up the JWT validator: issuer url is required but was empty",
			expectedNilJWT: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jwtValidator, err := setupJWTValidator(tc.jwtENV)

			if err != nil {
				assert.Error(t, err, "Expected error")
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain expected string")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
				assert.Nil(t, jwtValidator, "jwtValidator should be nil when there is no error // Got: %v", jwtValidator)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NotNil(t, jwtValidator, "jwtValidator should not be nil when there is no error")
			}
		})
	}
}

// // MockUserOperations is a mock implementation of UserOperations for testing purposes.
// type MockUserOperations struct {
// 	mock.Mock
// }

// func (m *MockUserOperations) CheckUserExists(authID string) (*int, error) {
// 	args := m.Called(authID)
// 	value := args.Int(0)
// 	return &value, args.Error(1)
// }

// func (m *MockUserOperations) AddUserToGraph(authID string) (*int, error) {
// 	args := m.Called(authID)
// 	value := args.Int(0)
// 	return &value, args.Error(1)
// }

// func TestEnsureValidToken_ErrorDuringSetup(t *testing.T) {
// 	mockClient := new(MockUserOperations)

// 	// mockValidator := &validator.Validator{} // You can use a mock validator if needed

// 	// Simulate an error during setup
// 	mockEnv := EnvJWTStruct{
// 		IssuerURL: "https://dev-afmzazq3cr35ktpl.us.auth0.com/",
// 		Audience: []string{
// 			"https://shrektionary.com/api",
// 			"4W01gsxupS4xoLLxbe8jdVVlGTFOKjd3",
// 		},
// 	}

// 	// Call the EnsureValidToken function
// 	middleware := EnsureValidToken(mockClient, mockEnv)

// 	// Create a test request and response
// 	req := httptest.NewRequest("GET", "/query", nil)
// 	w := httptest.NewRecorder()

// 	// Call the middleware with the test request and response
// 	// middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(w, req)

// 	mockClient.On("CheckUserExists", mock.Anything).Return(0, nil).Once()
// 	mockClient.On("AddUserToGraph", mock.Anything).Return(0, nil).Once()

// 	// Call the middleware with the test request and response
// 	middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Your test logic here
// 	})).ServeHTTP(w, req)

// 	// Assert that the expectations were met
// 	mockClient.AssertExpectations(t)
// }
