package resolvers

import (
	"context"
	"groupinary/ent/group"
	"groupinary/graph"
	"groupinary/testutils"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// Users

func TestUpdateUserName(t *testing.T) {
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

	// Create a query resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	lastName := "last Name"
	nickName := "nick Name"

	// Test cases table
	testCases := []struct {
		name          string
		firstName     string
		lastName      *string
		nickName      *string
		ctx           context.Context
		expectedError string
	}{
		{
			name:          "valid firstName, nil lastname, nil nickName",
			firstName:     "first name",
			lastName:      nil,
			nickName:      nil,
			ctx:           testutils.TestContext(2),
			expectedError: "",
		},
		{
			name:          "empty string firstName, nil lastname, nil nickName",
			firstName:     "",
			lastName:      nil,
			nickName:      nil,
			ctx:           testutils.TestContext(3),
			expectedError: "",
		},
		{
			name:          "valid firstName, valid lastname, nil nickName",
			firstName:     "first Name",
			lastName:      &lastName,
			nickName:      nil,
			ctx:           testutils.TestContext(4),
			expectedError: "",
		},
		{
			name:          "valid firstName, valid lastname, valid nickName",
			firstName:     "first Name",
			lastName:      &lastName,
			nickName:      &nickName,
			ctx:           testutils.TestContext(5),
			expectedError: "",
		},
		{
			name:          "User ID Not Added to Context",
			firstName:     "first Name",
			lastName:      nil,
			nickName:      nil,
			ctx:           context.Background(),
			expectedError: "could not retrieve user_id from context",
		},
		{
			name:          "User Not Found in Database",
			firstName:     "first Name",
			lastName:      nil,
			nickName:      nil,
			ctx:           testutils.TestContext(999999),
			expectedError: "ent: user not found",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			resultUser, err := resolver.UpdateUserName(tc.ctx, tc.firstName, tc.lastName, tc.nickName)

			if err != nil {
				assert.Error(t, err, "Expected error")
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain expected string")
				assert.Nil(t, resultUser, "User should be nil when there is an error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NotNil(t, resultUser, "User should not be nil when there is no error")

				// Check if firstName matches
				assert.Equal(t, tc.firstName, resultUser.FirstName, "User ID should match")

				// Check if lastName matches or is nil
				if tc.lastName != nil {
					assert.Equal(t, *tc.lastName, resultUser.LastName, "User lastName should match")
				} else {
					assert.Equal(t, "", resultUser.LastName, "User lastName should be an empty string // Got: %v", resultUser.LastName)
				}

				// Check if nickName matches or is nil
				if tc.nickName != nil {
					assert.Equal(t, *tc.nickName, resultUser.NickName, "User nickName should match")
				} else {
					assert.Equal(t, "", resultUser.NickName, "User nickName should be an empty string // Got: %v", resultUser.NickName)
				}
			}
		})
	}
}

// Groups

func TestCreateGroup(t *testing.T) {
	client, db, err := testutils.OpenTest()
	if err != nil {
		t.Fatal(err)
	}
	// Register the cleanup function from testutils.
	t.Cleanup(func() {
		testutils.CleanupTestEnvironment(t, client)
	})

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	userId := 1
	expectedDescription := "description1"
	expectedEmptyDescription := ""

	testCases := []struct {
		name             string
		ctx              context.Context
		nameInput        string
		descriptionInput *string
		expectedError    string
	}{
		{
			name:             "Happy Path",
			ctx:              testutils.TestContext(userId),
			nameInput:        "TestGroup",
			descriptionInput: nil,
			expectedError:    "",
		},
		{
			name:             "Happy Path with description",
			ctx:              testutils.TestContext(userId),
			nameInput:        "TestGroup",
			descriptionInput: &expectedDescription,
			expectedError:    "",
		},
		{
			name:             "Happy Path with (empty description)",
			ctx:              testutils.TestContext(userId),
			nameInput:        "TestGroup",
			descriptionInput: &expectedEmptyDescription,
			expectedError:    "",
		},
		{
			name:             "Invalid Input (empty name)",
			ctx:              testutils.TestContext(userId),
			nameInput:        "",
			descriptionInput: nil,
			expectedError:    "ent: validator failed for field \"Group.name\": value is less than the required length",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			array := []string{}
			testutils.LoadFixtures(db, array...)
			group, err := resolver.CreateGroup(tc.ctx, tc.nameInput, tc.descriptionInput)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, group, "Group should be nil on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NoError(t, err, "Unexpected error")
				assert.NotNil(t, group, "Group should not be nil")
				assert.Equal(t, tc.nameInput, group.Name, "Group name should match")

				// Check if the description matches or is nil
				if tc.descriptionInput != nil {
					assert.Equal(t, *tc.descriptionInput, group.Description, "Group description should match")
				} else {
					assert.Equal(t, "", group.Description, "Group description should be an empty string // Got: %v", group.Description)
				}
			}
		})
	}
}

func TestUpdateGroupName(t *testing.T) {
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

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	userId := 1      // This is inside of fixtures/users.yaml
	testGroupId := 1 // This is inside of fixtures/groups.yaml

	testCases := []struct {
		name          string
		ctx           context.Context
		id            int
		newName       string
		expectedError string
	}{
		{
			name:          "Happy Path",
			ctx:           testutils.TestContext(userId),
			id:            testGroupId,
			newName:       "UpdatedGroupName",
			expectedError: "",
		},
		{
			name:          "Invalid Input (empty name)",
			ctx:           testutils.TestContext(userId),
			id:            testGroupId,
			newName:       "",
			expectedError: "ent: validator failed for field \"Group.name\": value is less than the required length",
		},
		{
			name:          "Non-Existent Group",
			ctx:           testutils.TestContext(userId),
			id:            999, // Non-existent ID
			newName:       "UpdatedGroupName",
			expectedError: "ent: group not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			group, err := resolver.UpdateGroupName(tc.ctx, tc.id, tc.newName)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, group, "Group should be nil on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NoError(t, err, "Unexpected error")
				assert.NotNil(t, group, "Group should not be nil")
				assert.Equal(t, tc.newName, group.Name, "Group name should match the updated name")
			}
		})
	}
}

func TestDeleteGroup(t *testing.T) {
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

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client.Debug(),
		},
	}

	userId := 1      // This is inside of fixtures/users.yaml
	testGroupId := 1 // This is inside of fixtures/groups.yaml

	testCases := []struct {
		name          string
		ctx           context.Context
		id            int
		expectedError string
	}{
		{
			name:          "Happy Path",
			ctx:           testutils.TestContext(userId),
			id:            testGroupId,
			expectedError: "",
		},
		{
			name:          "Non-Existent Group",
			ctx:           testutils.TestContext(userId),
			id:            999, // Non-existent ID
			expectedError: "ent: group not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			result, err := resolver.DeleteGroup(tc.ctx, tc.id)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.False(t, result, "DeleteGroup should return false on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NoError(t, err, "Unexpected error")
				assert.True(t, result, "DeleteGroup should return true on success")

				deletedGroup, err := client.Group.Query().Where(group.ID(tc.id)).Only(tc.ctx)
				assert.Error(t, err, "Expecting an error when trying to query a deleted group")
				assert.Nil(t, deletedGroup, "Deleted group should be nil")
			}
		})
	}
}

// Words

func TestCreateWord(t *testing.T) {
	fixturePaths := []string{
		"fixtures/users.yaml",
		"fixtures/groups.yaml",
		"fixtures/user_groups.yaml",
	}

	client, db, err := testutils.OpenTest()
	if err != nil {
		t.Fatal(err)
	}
	// Register the cleanup function from testutils.
	t.Cleanup(func() {
		testutils.CleanupTestEnvironment(t, client)
	})

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	userId := 1      // This is inside of fixtures/users.yaml
	testGroupId := 1 // This is inside of fixtures/groups.yaml

	expectedDescription := "description"
	expectedEmptyDescription := ""

	testCases := []struct {
		name            string
		ctx             context.Context
		rootWordInput   string
		groupID         int
		definitionInput *string
		expectedError   string
	}{
		{
			name:            "Happy Path",
			ctx:             testutils.TestContext(userId),
			rootWordInput:   "Test Root Word",
			groupID:         testGroupId,
			definitionInput: nil,
			expectedError:   "",
		},
		{
			name:            "Happy Path with defintion",
			ctx:             testutils.TestContext(userId),
			rootWordInput:   "Test Root Word",
			groupID:         testGroupId,
			definitionInput: &expectedDescription,
			expectedError:   "",
		},
		{
			name:            "Invalid Input (empty defintion)",
			ctx:             testutils.TestContext(userId),
			rootWordInput:   "Test Root Word",
			groupID:         testGroupId,
			definitionInput: &expectedEmptyDescription,
			expectedError:   "ent: validator failed for field \"Definition.description\": value is less than the required length",
		},
		{
			name:            "Invalid Input (empty name)",
			ctx:             testutils.TestContext(userId),
			rootWordInput:   "",
			groupID:         testGroupId,
			definitionInput: nil,
			expectedError:   "ent: validator failed for field \"Word.description\": value is less than the required length",
		},
		{
			name:            "User Not Found in Database",
			ctx:             context.Background(),
			rootWordInput:   "Test Root Word",
			groupID:         testGroupId,
			definitionInput: nil,
			expectedError:   "could not retrieve user_id from context",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			word, err := resolver.AddRootWord(tc.ctx, tc.rootWordInput, tc.groupID, tc.definitionInput)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, word, "Word should be nil on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NoError(t, err, "Unexpected error")
				assert.NotNil(t, word, "Word should not be nil")
				assert.Equal(t, tc.rootWordInput, word.Description, "Word name should match")

				// Check if the definitionInput matches or is nil
				if tc.definitionInput != nil {
					definition, err := word.QueryDefinitions().First(tc.ctx)
					assert.NoError(t, err, "Error should be nil")
					assert.NotNil(t, definition, "Definition should not be nil")
					assert.Equal(t, *tc.definitionInput, definition.Description, "Definition description should match")
				} else {
					definition, err := word.QueryDefinitions().First(tc.ctx)
					assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
					assert.Nil(t, definition, "Definition should be nil")
				}
			}
		})
	}
}

func TestUpdateWordName(t *testing.T) {
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

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	userId := 1      // This is inside of fixtures/users.yaml
	testGroupId := 1 // This is inside of fixtures/groups.yaml

	testCases := []struct {
		name          string
		ctx           context.Context
		id            int
		newName       string
		expectedError string
	}{
		{
			name:          "Happy Path",
			ctx:           testutils.TestContext(userId),
			id:            testGroupId,
			newName:       "UpdatedGroupName",
			expectedError: "",
		},
		{
			name:          "Invalid Input (empty name)",
			ctx:           testutils.TestContext(userId),
			id:            testGroupId,
			newName:       "",
			expectedError: "ent: validator failed for field \"Word.description\": value is less than the required length",
		},
		{
			name:          "Non-Existent Word",
			ctx:           testutils.TestContext(userId),
			id:            999, // Non-existent ID
			newName:       "UpdatedWordName",
			expectedError: "ent: word not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			word, err := resolver.UpdateWord(tc.ctx, tc.id, tc.newName)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, word, "Word should be nil on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NoError(t, err, "Unexpected error")
				assert.NotNil(t, word, "Word should not be nil")
				assert.Equal(t, tc.newName, word.Description, "Word name should match the updated name")
			}
		})
	}
}

func TestDeleteWord(t *testing.T) {
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

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	userId := 1     // This is inside of fixtures/users.yaml
	testWordId := 1 // This is inside of fixtures/words.yaml

	testCases := []struct {
		name          string
		ctx           context.Context
		id            int
		expectedError string
	}{
		{
			name:          "Happy Path",
			ctx:           testutils.TestContext(userId),
			id:            testWordId,
			expectedError: "",
		},
		{
			name:          "Non-Existent Word",
			ctx:           testutils.TestContext(userId),
			id:            999, // Non-existent ID
			expectedError: "ent: word not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			deleted, err := resolver.DeleteWord(tc.ctx, tc.id)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.False(t, deleted, "Deleted should be false on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NoError(t, err, "Unexpected error")
				assert.True(t, deleted, "Deleted should be true")
			}
		})
	}
}

func TestConnectWords(t *testing.T) {
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

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	userId := 1       // This is inside of fixtures/users.yaml
	parentWordId := 1 // This is inside of fixtures/words.yaml
	childWordId := 2  // This is inside of fixtures/words.yaml

	testCases := []struct {
		name          string
		ctx           context.Context
		parentWordId  int
		childWordId   int
		expectedError string
	}{
		{
			name:          "Happy Path",
			ctx:           testutils.TestContext(userId),
			parentWordId:  parentWordId,
			childWordId:   childWordId,
			expectedError: "",
		},
		{
			name:          "Non-Existent Word",
			ctx:           testutils.TestContext(userId),
			parentWordId:  parentWordId, // Non-existent ID
			childWordId:   999,          // Non-existent ID
			expectedError: "ent: constraint failed: add m2m edge for table word_descendants: FOREIGN KEY constraint failed",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			word, err := resolver.ConnectWords(tc.ctx, tc.parentWordId, tc.childWordId)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, word, "Word should be nil on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NoError(t, err, "Unexpected error")
				assert.NotNil(t, word, "Word should not be nil")
				assert.Equal(t, tc.childWordId, word.ID, "Word id should match the updated name")

				// Check if the return parent word is corect
				parentWord, err := word.QueryParents().First(tc.ctx)
				assert.NoError(t, err, "Error should be nil")
				assert.NotNil(t, parentWord, "Definition should not be nil")
				assert.Equal(t, tc.parentWordId, parentWord.ID, "Definition description should match")

			}
		})
	}
}

// Definitions

func TestUpdateDefinitionName(t *testing.T) {
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

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	userId := 1           // This is inside of fixtures/users.yaml
	testDefinitionId := 1 // This is inside of fixtures/words.yaml

	testCases := []struct {
		name          string
		ctx           context.Context
		id            int
		newName       string
		expectedError string
	}{
		{
			name:          "Happy Path",
			ctx:           testutils.TestContext(userId),
			id:            testDefinitionId,
			newName:       "UpdatedGroupName",
			expectedError: "",
		},
		{
			name:          "Invalid Input (empty name)",
			ctx:           testutils.TestContext(userId),
			id:            testDefinitionId,
			newName:       "",
			expectedError: "ent: validator failed for field \"Definition.description\": value is less than the required length",
		},
		{
			name:          "Non-Existent Word",
			ctx:           testutils.TestContext(userId),
			id:            999, // Non-existent ID
			newName:       "UpdatedWordName",
			expectedError: "ent: definition not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			definition, err := resolver.UpdateDefinition(tc.ctx, tc.id, tc.newName)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, definition, "Word should be nil on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.NoError(t, err, "Unexpected error")
				assert.NotNil(t, definition, "Word should not be nil")
				assert.Equal(t, tc.newName, definition.Description, "Word name should match the updated name")
			}
		})
	}
}

func TestDeleteDefinition(t *testing.T) {
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

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	userId := 1           // This is inside of fixtures/users.yaml
	testDefinitionId := 1 // This is inside of fixtures/words.yaml

	testCases := []struct {
		name          string
		ctx           context.Context
		id            int
		expectedError string
	}{
		{
			name:          "Happy Path",
			ctx:           testutils.TestContext(userId),
			id:            testDefinitionId,
			expectedError: "",
		},
		{
			name:          "Non-Existent Definition",
			ctx:           testutils.TestContext(userId),
			id:            999, // Non-existent ID
			expectedError: "ent: definition not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils.LoadFixtures(db, fixturePaths...)
			deleted, err := resolver.DeleteDefinition(tc.ctx, tc.id)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.False(t, deleted, "Deleted should be false on error")
				assert.NotEqual(t, "", tc.expectedError, "expectedError should not be an empty string // Got: %v", err)
			} else {
				assert.Equal(t, "", tc.expectedError, "expectedError should be empty // Got: %v", tc.expectedError)
				assert.NoError(t, err, "Unexpected error")
				assert.True(t, deleted, "Deleted should be true")
			}
		})
	}
}

// Mutation

func TestMutation(t *testing.T) {
	// Create a new instance of Resolver
	resolver := &Resolver{} // You may need to initialize fields or dependencies if required

	// Call the Mutation method
	mutationResolver := resolver.Mutation()

	// Assert that the returned object is not nil
	if mutationResolver == nil {
		t.Error("Expected non-nil MutationResolver, got nil")
	}

	//lint:ignore S1040 ignoring type assertion warning because MutationResolver is explicitly declared
	// Assert that the returned object is of the correct type (graph.MutationResolver)
	_, isMutationResolver := mutationResolver.(graph.MutationResolver)
	if !isMutationResolver {
		t.Error("Expected MutationResolver to implement graph.MutationResolver interface")
	}
}
