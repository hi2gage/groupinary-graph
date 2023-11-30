package resolvers

import (
	"context"
	"testing"

	"groupinary/ent"
	"groupinary/testutils"

	// "groupinary/ent/entc/integration/json/ent/enttest"

	"entgo.io/contrib/entgql"
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
			ctx:         testutils.TestContext(999999),
			expectedErr: true,
		},
		{
			name:        "Happy Path",
			ctx:         testutils.TestContext(userId),
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

func TestDefinitionsConnections(t *testing.T) {
	fixturePaths := []string{
		"fixtures/users.yaml",
		"fixtures/groups.yaml",
		"fixtures/user_groups.yaml",
		"fixtures/words.yaml",
		"fixtures/definitions.yaml",
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

	userId := 33
	wordID := 33
	first := 10

	type expectedValues struct {
		count    int
		firstId  int
		pageInfo entgql.PageInfo[int]
	}

	testCases := []struct {
		name           string
		ctx            context.Context
		wordID         *int
		after          *entgql.Cursor[int]
		first          *int
		before         *entgql.Cursor[int]
		last           *int
		orderBy        *ent.DefinitionOrder
		where          *ent.DefinitionWhereInput
		expectedValues *expectedValues
		expectedError  string
	}{
		{
			name:    "Happy path",
			ctx:     testutils.TestContext(userId),
			wordID:  &wordID,
			after:   nil,
			first:   &first,
			before:  nil,
			last:    nil,
			orderBy: nil,
			expectedValues: &expectedValues{
				count:   2,
				firstId: 1,
				pageInfo: entgql.PageInfo[int]{
					HasNextPage:     false,
					HasPreviousPage: false,
					StartCursor: &entgql.Cursor[int]{
						ID:    1,
						Value: nil,
					},
					EndCursor: &entgql.Cursor[int]{
						ID:    100,
						Value: nil,
					},
				},
			},
			where: nil,

			expectedError: "",
		},
		// Add test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			conn, err := resolver.DefinitionsConnections(tc.ctx, tc.wordID, tc.after, tc.first, tc.before, tc.last, tc.orderBy, tc.where)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, conn, "Connection should be nil on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.NoError(t, err, "Unexpected error")
				assert.NotNil(t, conn, "Connection should not be nil")

				// expected values:
				assert.Equal(t, len(conn.Edges), tc.expectedValues.count)
				assert.Equal(t, conn.TotalCount, tc.expectedValues.count)
				assert.Equal(t, conn.Edges[0].Node.ID, tc.expectedValues.firstId, "first edge should match")

				// pageInfo
				assert.Equal(t, conn.PageInfo.HasNextPage, tc.expectedValues.pageInfo.HasNextPage)
				assert.Equal(t, conn.PageInfo.HasPreviousPage, tc.expectedValues.pageInfo.HasPreviousPage)
				assert.Equal(t, conn.PageInfo.StartCursor, tc.expectedValues.pageInfo.StartCursor)
				assert.Equal(t, conn.PageInfo.EndCursor, tc.expectedValues.pageInfo.EndCursor)
			}
		})
	}
}
