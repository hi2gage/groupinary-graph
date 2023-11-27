package resolvers

import (
	"context"
	"groupinary/ent/enttest"
	"groupinary/ent/group"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGroup(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	expectedDescription := "description"
	expectedEmptyDescription := "description"

	testCases := []struct {
		name             string
		ctx              context.Context
		nameInput        string
		descriptionInput *string
		expectedError    string
	}{
		{
			name:             "Happy Path",
			ctx:              context.Background(),
			nameInput:        "TestGroup",
			descriptionInput: nil,
			expectedError:    "",
		},
		{
			name:             "Happy Path with description",
			ctx:              context.Background(),
			nameInput:        "TestGroup",
			descriptionInput: &expectedDescription,
			expectedError:    "",
		},
		{
			name:             "Happy Path with (empty description)",
			ctx:              context.Background(),
			nameInput:        "TestGroup",
			descriptionInput: &expectedEmptyDescription,
			expectedError:    "",
		},
		{
			name:             "Invalid Input (empty name)",
			ctx:              context.Background(),
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
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	// Create a test group to update
	testGroup, err := resolver.CreateGroup(context.Background(), "TestGroup", nil)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name          string
		ctx           context.Context
		id            int
		newName       string
		expectedError string
	}{
		{
			name:          "Happy Path",
			ctx:           context.Background(),
			id:            testGroup.ID,
			newName:       "UpdatedGroupName",
			expectedError: "",
		},
		{
			name:          "Invalid Input (empty name)",
			ctx:           context.Background(),
			id:            testGroup.ID,
			newName:       "",
			expectedError: "ent: validator failed for field \"Group.name\": value is less than the required length",
		},
		{
			name:          "Non-Existent Group",
			ctx:           context.Background(),
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
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create a mutation resolver with the test client
	resolver := &mutationResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	// Create a test group to delete
	testGroup, err := resolver.CreateGroup(context.Background(), "TestGroup", nil)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name          string
		ctx           context.Context
		id            int
		expectedError string
	}{
		{
			name:          "Happy Path",
			ctx:           context.Background(),
			id:            testGroup.ID,
			expectedError: "",
		},
		{
			name:          "Non-Existent Group",
			ctx:           context.Background(),
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
