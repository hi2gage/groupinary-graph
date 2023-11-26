package utils

import (
	"context"
	"testing"
)

func TestGetCreatorIDFromContext_TableDriven(t *testing.T) {
	testCases := []struct {
		name          string
		ctx           context.Context
		expectedID    int
		expectedError string
	}{
		{
			name:          "Valid User ID",
			ctx:           AddUserIdToContext(context.Background(), 123),
			expectedID:    123,
			expectedError: "",
		},
		{
			name:          "No User ID",
			ctx:           context.Background(),
			expectedID:    0,
			expectedError: "could not retrieve user_id from context",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultID, err := GetCreatorIDFromContext(testCase.ctx)

			// Check if the result matches the expected values
			if resultID != testCase.expectedID {
				t.Errorf("Expected user ID %d, got %d", testCase.expectedID, resultID)
			}

			if err == nil && testCase.expectedError != "" {
				t.Errorf("Expected error: %s, got nil", testCase.expectedError)
			} else if err != nil && err.Error() != testCase.expectedError {
				t.Errorf("Expected error: %s, got: %v", testCase.expectedError, err)
			}
		})
	}
}
