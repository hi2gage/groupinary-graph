package resolvers

import (
	"context"
	"groupinary/ent"
	"groupinary/graph"
	"groupinary/testutils"
	"groupinary/utils"
	"testing"

	"entgo.io/contrib/entgql"
	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	fixturePaths := []string{
		"fixtures/users.yaml",
		"fixtures/groups.yaml",
		"fixtures/user_groups.yaml",
		"fixtures/words.yaml",
		"fixtures/definitions.yaml",
		"fixtures/ent_types.yaml",
	}

	client, db, err := testutils.OpenTest()
	if err != nil {
		t.Fatal(err)
	}
	// Register the cleanup function from testutils.
	t.Cleanup(func() {
		testutils.CleanupTestEnvironment(t, client)
	})

	// Create a query queryResolver with the test client
	queryResolver := &queryResolver{
		Resolver: &Resolver{
			client: client.Debug(),
		},
	}

	// Create a mutation mutationResolver with the test client
	mutationResolver := &mutationResolver{
		Resolver: &Resolver{
			client: client.Debug(),
		},
	}
	stringTest := "test"
	user, _ := mutationResolver.Mutation().CreateGroup(context.Background(), stringTest, nil)

	nodeId := user.ID // This is inside of fixtures/users.yaml

	// Test cases table
	testCases := []struct {
		name          string
		ctx           context.Context
		nodeId        int
		expectedError string
	}{
		{
			name:          "Happy Path",
			ctx:           utils.AddUserIdToContext(context.Background(), 1),
			nodeId:        nodeId,
			expectedError: "",
		},
		{
			name:          "Node Not Found in Database",
			ctx:           utils.AddUserIdToContext(context.Background(), 999999),
			nodeId:        102,
			expectedError: "input: Could not resolve to a node with the global id of '102'",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)

			resultNode, err := queryResolver.Node(tc.ctx, tc.nodeId)

			if err != nil {
				assert.Error(t, err, "Expected error")
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain expected string")
				assert.Nil(t, resultNode, "User should be nil when there is an error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NotNil(t, resultNode, "Node should not be nil when there is no error")

				// Type assertion to get the concrete type
				var resultID int
				switch v := resultNode.(type) {
				case *ent.User:
					resultID = v.ID
					println("User")
				case *ent.Group:
					resultID = v.ID
					println("Group")
				case *ent.Word:
					resultID = v.ID
					println("Word")
				case *ent.Definition:
					resultID = v.ID
					print("Definition")
				default:
					t.Errorf("Unexpected type %T", v)
				}

				assert.Equal(t, nodeId, resultID, "ID should match")
			}
		})
	}
}

func TestNodes(t *testing.T) {
	fixturePaths := []string{
		"fixtures/users.yaml",
		"fixtures/groups.yaml",
		"fixtures/user_groups.yaml",
		"fixtures/words.yaml",
		"fixtures/definitions.yaml",
		"fixtures/ent_types.yaml",
	}

	client, db, err := testutils.OpenTest()
	if err != nil {
		t.Fatal(err)
	}
	// Register the cleanup function from testutils.
	t.Cleanup(func() {
		testutils.CleanupTestEnvironment(t, client)
	})

	// Create a query queryResolver with the test client
	queryResolver := &queryResolver{
		Resolver: &Resolver{
			client: client.Debug(),
		},
	}

	// Create a mutation mutationResolver with the test client
	mutationResolver := &mutationResolver{
		Resolver: &Resolver{
			client: client.Debug(),
		},
	}
	stringTest := "test"
	user, _ := mutationResolver.Mutation().CreateGroup(context.Background(), stringTest, nil)

	nodeId := user.ID // This is inside of fixtures/users.yaml

	// Test cases table
	testCases := []struct {
		name          string
		ctx           context.Context
		nodeId        int
		expectedError string
	}{
		{
			name:          "Happy Path",
			ctx:           utils.AddUserIdToContext(context.Background(), 1),
			nodeId:        nodeId,
			expectedError: "",
		},
		{
			name:          "Node Not Found in Database",
			ctx:           utils.AddUserIdToContext(context.Background(), 999999),
			nodeId:        102,
			expectedError: "input: Could not resolve to a node with the global id of '102'",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)

			resultNodes, err := queryResolver.Nodes(tc.ctx, []int{tc.nodeId})

			if err != nil {
				assert.Error(t, err, "Expected error")
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain expected string")
				assert.Nil(t, resultNodes, "User should be nil when there is an error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NotNil(t, resultNodes, "Node should not be nil when there is no error")

				// Type assertion to get the concrete type
				for _, node := range resultNodes {
					var resultID int
					switch v := node.(type) {
					case *ent.User:
						resultID = v.ID
						println("User")
					case *ent.Group:
						resultID = v.ID
						println("Group")
					case *ent.Word:
						resultID = v.ID
						println("Word")
					case *ent.Definition:
						resultID = v.ID
						print("Definition")
					default:
						t.Errorf("Unexpected type %T", v)
					}

					assert.Equal(t, tc.nodeId, resultID, "ID should match")
				}
			}
		})
	}
}

func TestDefinitions(t *testing.T) {
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
				count:   3,
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
			conn, err := resolver.Definitions(tc.ctx, tc.after, tc.first, tc.before, tc.last, tc.orderBy, tc.where)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, conn, "Connection should be nil on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NoError(t, err, "Unexpected error")
				assert.NotNil(t, conn, "Connection should not be nil")

				// expected values:
				assert.Equal(t, tc.expectedValues.count, len(conn.Edges))
				assert.Equal(t, tc.expectedValues.count, conn.TotalCount)
				assert.Equal(t, tc.expectedValues.firstId, conn.Edges[0].Node.ID, "first edge should match")

				// pageInfo
				assert.Equal(t, tc.expectedValues.pageInfo.HasNextPage, conn.PageInfo.HasNextPage)
				assert.Equal(t, tc.expectedValues.pageInfo.HasPreviousPage, conn.PageInfo.HasPreviousPage)
				assert.Equal(t, tc.expectedValues.pageInfo.StartCursor, conn.PageInfo.StartCursor)
				assert.Equal(t, tc.expectedValues.pageInfo.EndCursor, conn.PageInfo.EndCursor)
			}
		})
	}
}

func TestGroups(t *testing.T) {
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

	userId := 33

	testCases := []struct {
		name          string
		ctx           context.Context
		expectedID    int
		expectedCount int
		expectedError string
	}{
		{
			name:          "Happy path",
			ctx:           testutils.TestContext(userId),
			expectedID:    1,
			expectedCount: 5,
			expectedError: "",
		},
		// Add test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			groups, err := resolver.Groups(tc.ctx)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, groups, "Connection should be nil on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NoError(t, err, "Unexpected error")
				assert.NotNil(t, groups, "Connection should not be nil")

				// expected values:
				assert.Equal(t, tc.expectedID, groups[0].ID)
				assert.Equal(t, tc.expectedCount, len(groups))
			}
		})
	}
}

func TestUsers(t *testing.T) {
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

	userId := 33

	testCases := []struct {
		name          string
		ctx           context.Context
		expectedID    int
		expectedCount int
		expectedError string
	}{
		{
			name:          "Happy path",
			ctx:           testutils.TestContext(userId),
			expectedID:    1,
			expectedCount: 9,
			expectedError: "",
		},
		// Add test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			users, err := resolver.Users(tc.ctx)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, users, "Connection should be nil on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NoError(t, err, "Unexpected error")
				assert.NotNil(t, users, "Connection should not be nil")

				// expected values:
				assert.Equal(t, tc.expectedID, users[0].ID)
				assert.Equal(t, tc.expectedCount, len(users))
			}
		})
	}
}

func TestWords(t *testing.T) {
	fixturePaths := []string{
		"fixtures/users.yaml",
		"fixtures/groups.yaml",
		"fixtures/user_groups.yaml",
		"fixtures/words.yaml",
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
		orderBy        *ent.WordOrder
		where          *ent.WordWhereInput
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
				count:   6,
				firstId: 1,
				pageInfo: entgql.PageInfo[int]{
					HasNextPage:     false,
					HasPreviousPage: false,
					StartCursor: &entgql.Cursor[int]{
						ID:    1,
						Value: nil,
					},
					EndCursor: &entgql.Cursor[int]{
						ID:    50,
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
			conn, err := resolver.Words(tc.ctx, tc.after, tc.first, tc.before, tc.last, tc.orderBy, tc.where)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, conn, "Connection should be nil on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NoError(t, err, "Unexpected error")
				assert.NotNil(t, conn, "Connection should not be nil")

				// expected values:
				assert.Equal(t, tc.expectedValues.count, len(conn.Edges))
				assert.Equal(t, tc.expectedValues.count, conn.TotalCount)
				assert.Equal(t, tc.expectedValues.firstId, conn.Edges[0].Node.ID, "first edge should match")

				// pageInfo
				assert.Equal(t, tc.expectedValues.pageInfo.HasNextPage, conn.PageInfo.HasNextPage)
				assert.Equal(t, tc.expectedValues.pageInfo.HasPreviousPage, conn.PageInfo.HasPreviousPage)
				assert.Equal(t, tc.expectedValues.pageInfo.StartCursor, conn.PageInfo.StartCursor)
				assert.Equal(t, tc.expectedValues.pageInfo.EndCursor, conn.PageInfo.EndCursor)
			}
		})
	}
}

// Mutation

func TestQuery(t *testing.T) {
	// Create a new instance of Resolver
	resolver := &Resolver{} // You may need to initialize fields or dependencies if required

	// Call the Mutation method
	queryResolver := resolver.Query()

	// Assert that the returned object is not nil
	if queryResolver == nil {
		t.Error("Expected non-nil QueryResolver, got nil")
	}

	//lint:ignore S1040 ignoring type assertion warning because QueryResolver is explicitly declared
	// Assert that the returned object is of the correct type (graph.QueryResolver)
	_, isMutationResolver := queryResolver.(graph.QueryResolver)
	if !isMutationResolver {
		t.Error("Expected QueryResolver to implement graph.QueryResolver interface")
	}
}
