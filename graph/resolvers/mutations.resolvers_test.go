package resolvers

import (
	"context"
	"groupinary/ent/enttest"
	"groupinary/ent/group"
	"groupinary/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Groups

func TestCreateGroup(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	userId := 1
	expectedDescription := "description"
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
			group, err := resolver.CreateGroup(tc.ctx, tc.nameInput, tc.descriptionInput)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, group, "Group should be nil on error")
			} else {
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

	client, err := testutils.OpenTest(fixturePaths...)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

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
			group, err := resolver.UpdateGroupName(tc.ctx, tc.id, tc.newName)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, group, "Group should be nil on error")
			} else {
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

	client, err := testutils.OpenTest(fixturePaths...)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

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
			result, err := resolver.DeleteGroup(tc.ctx, tc.id)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.False(t, result, "DeleteGroup should return false on error")
			} else {
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

	client, err := testutils.OpenTest(fixturePaths...)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

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
			expectedError:   "ent: definition not found",
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			word, err := resolver.AddRootWord(tc.ctx, tc.rootWordInput, tc.groupID, tc.definitionInput)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, word, "Word should be nil on error")
			} else {
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

	client, err := testutils.OpenTest(fixturePaths...)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

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
			word, err := resolver.UpdateWord(tc.ctx, tc.id, tc.newName)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, word, "Word should be nil on error")
			} else {
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

	client, err := testutils.OpenTest(fixturePaths...)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

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
			deleted, err := resolver.DeleteWord(tc.ctx, tc.id)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.False(t, deleted, "Deleted should be false on error")
			} else {
				assert.NoError(t, err, "Unexpected error")
				assert.True(t, deleted, "Deleted should be true")
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

	client, err := testutils.OpenTest(fixturePaths...)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

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
			definition, err := resolver.UpdateDefinition(tc.ctx, tc.id, tc.newName)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.Nil(t, definition, "Word should be nil on error")
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

	client, err := testutils.OpenTest(fixturePaths...)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

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
			deleted, err := resolver.DeleteDefinition(tc.ctx, tc.id)

			if err != nil {
				assert.Contains(t, err.Error(), tc.expectedError, "Error message should contain the expected substring")
				assert.False(t, deleted, "Deleted should be false on error")
			} else {
				assert.NoError(t, err, "Unexpected error")
				assert.True(t, deleted, "Deleted should be true")
			}
		})
	}
}
