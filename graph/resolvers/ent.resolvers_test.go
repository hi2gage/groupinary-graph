package resolvers

import (
	"context"
	"groupinary/ent"
	"groupinary/testutils"
	"groupinary/utils"
	"testing"

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
