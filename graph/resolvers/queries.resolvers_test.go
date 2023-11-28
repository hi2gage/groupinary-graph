package resolvers

import (
	"context"
	"testing"

	"groupinary/testutils"
	"groupinary/utils"

	// "groupinary/ent/entc/integration/json/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCurrentUser(t *testing.T) {
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

	// Create a query resolver with the test client
	resolver := &queryResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	userId := 1 // This is inside of fixtures/users.yaml

	// Test cases table
	testCases := []struct {
		name        string
		ctx         context.Context
		expectedErr bool
	}{
		{
			name:        "User ID Not Added to Context",
			ctx:         context.Background(),
			expectedErr: true,
		},
		{
			name:        "User Not Found in Database",
			ctx:         utils.AddUserIdToContext(context.Background(), 999999),
			expectedErr: true,
		},
		{
			name:        "Happy Path",
			ctx:         utils.AddUserIdToContext(context.Background(), userId),
			expectedErr: false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			resultUser, err := resolver.CurrentUser(tc.ctx)

			if tc.expectedErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}

			if !tc.expectedErr {
				assert.NotNil(t, resultUser, "User should not be nil when there is no error")
				assert.Equal(t, userId, resultUser.ID, "User ID should match")
			}
		})
	}
}
