package middleware

import (
	"groupinary/testutils"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestCheckUserExists(t *testing.T) {
	fixturePaths := []string{
		"fixtures/users.yaml",
	}

	client, db, err := testutils.OpenTest()
	if err != nil {
		t.Fatal(err)
	}
	// Register the cleanup function from testutils.
	t.Cleanup(func() {
		testutils.CleanupTestEnvironment(t, client)
	})

	operator := NewUserTokenOperator(client)

	expectedID := 60
	testCases := []struct {
		name          string
		authID        string
		expectedID    *int
		expectedError string
	}{
		{
			name:          "User 60",
			authID:        "test_auth_id_60",
			expectedID:    &expectedID,
			expectedError: "",
		},
		{
			name:          "User Does Not Exist",
			authID:        "nonexistent",
			expectedID:    nil,
			expectedError: "user does not exist",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)

			resultUserId, err := operator.CheckUserExists(tc.authID)

			if err != nil {
				assert.Error(t, err, "Expected error")
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain expected string")
				assert.Nil(t, resultUserId, "User should be nil when there is an error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NotNil(t, resultUserId, "User should not be nil when there is no error")

				// Check if expectedID matches
				assert.Equal(t, tc.expectedID, resultUserId, "User ID should match")
			}
		})
	}
}

func TestAddUserToGraph(t *testing.T) {
	fixturePaths := []string{
		"fixtures/users.yaml",
	}

	client, db, err := testutils.OpenTest()
	if err != nil {
		t.Fatal(err)
	}
	// Register the cleanup function from testutils.
	t.Cleanup(func() {
		testutils.CleanupTestEnvironment(t, client)
	})

	operator := NewUserTokenOperator(client)

	expectedID := 8589934593
	testCases := []struct {
		name          string
		authID        string
		expectedID    *int
		expectedError string
	}{
		{
			name:          "User 61",
			authID:        "test_auth_id_61",
			expectedID:    &expectedID,
			expectedError: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)

			resultUserId, err := operator.AddUserToGraph(tc.authID)

			if err != nil {
				assert.Error(t, err, "Expected error")
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain expected string")
				assert.Nil(t, resultUserId, "User should be nil when there is an error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NotNil(t, resultUserId, "User should not be nil when there is no error")

				// Check if expectedID matches
				assert.Equal(t, *tc.expectedID, *resultUserId, "User ID should match")
			}
		})
	}
}
