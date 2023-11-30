package resolvers

// func TestNode(t *testing.T) {
// 	fixturePaths := []string{
// 		"fixtures/users.yaml",
// 		"fixtures/ent_types.yaml",
// 	}

// 	client, db, err := testutils.OpenTest()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// Register the cleanup function from testutils.
// 	t.Cleanup(func() {
// 		testutils.CleanupTestEnvironment(t, client)
// 	})

// 	// Create a query resolver with the test client
// 	resolver := &queryResolver{
// 		Resolver: &Resolver{
// 			client: client,
// 		},
// 	}

// 	nodeId := 1 // This is inside of fixtures/users.yaml

// 	// Test cases table
// 	testCases := []struct {
// 		name          string
// 		ctx           context.Context
// 		nodeId        int
// 		expectedError string
// 	}{
// 		{
// 			name:          "Happy Path",
// 			ctx:           utils.AddUserIdToContext(context.Background(), nodeId),
// 			nodeId:        1,
// 			expectedError: "Got: no such table: ent_types",
// 		},
// 		{
// 			name:          "Node Not Found in Database",
// 			ctx:           utils.AddUserIdToContext(context.Background(), 999999),
// 			nodeId:        999999,
// 			expectedError: "Got: no such table: ent_types",
// 		},
// 	}

// 	// Run test cases
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			testutils.LoadFixtures(db, fixturePaths...)
// 			resultNode, err := resolver.Node(tc.ctx, tc.nodeId)

// 			if err != nil {
// 				assert.Error(t, err, "Expected error")
// 				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain expected string")
// 				assert.Nil(t, resultNode, "User should be nil when there is an error")
// 				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
// 			} else {
// 				assert.NotNil(t, resultNode, "Node should not be nil when there is no error")

// 				// Type assertion to get the concrete type
// 				var resultID int
// 				switch v := resultNode.(type) {
// 				case *ent.User:
// 					resultID = v.ID
// 				case *ent.Group:
// 					resultID = v.ID
// 				case *ent.Word:
// 					resultID = v.ID
// 				case *ent.Definition:
// 					resultID = v.ID
// 				default:
// 					t.Errorf("Unexpected type %T", v)
// 				}

// 				assert.Equal(t, nodeId, resultID, "ID should match")
// 			}
// 		})
// 	}
// }
