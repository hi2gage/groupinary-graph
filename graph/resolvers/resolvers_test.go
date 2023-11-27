package resolvers

import (
	"context"
	"testing"
	"time"

	"groupinary/ent"
	"groupinary/ent/enttest"
	"groupinary/utils"

	// "groupinary/ent/entc/integration/json/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCurrentUser(t *testing.T) {
	// Create a new ent.Client for testing
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	// Create a query resolver with the test client
	resolver := &queryResolver{
		Resolver: &Resolver{
			client: client,
		},
	}

	// Define the base user for testing
	baseUser := &ent.User{
		ID:         123,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		AuthID:     "test_auth_id",
		FirstName:  "Test",
		LastName:   "User",
		NickName:   "TestNick",
	}

	// Insert the test user into the database
	testUser, err := client.User.Create().
		SetCreateTime(baseUser.CreateTime).
		SetUpdateTime(baseUser.UpdateTime).
		SetAuthID(baseUser.AuthID).
		SetFirstName(baseUser.FirstName).
		SetLastName(baseUser.LastName).
		SetNickName(baseUser.NickName).
		Save(context.Background())

	if err != nil {
		t.Fatal(err)
	}

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
			ctx:         utils.AddUserIdToContext(context.Background(), testUser.ID),
			expectedErr: false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultUser, err := resolver.CurrentUser(tc.ctx)

			if tc.expectedErr {
				assert.Error(t, err, "Expected error")
			} else {
				assert.NoError(t, err, "Unexpected error")
			}

			if !tc.expectedErr {
				assert.NotNil(t, resultUser, "User should not be nil when there is no error")
				compareUsers(t, testUser, resultUser)
			}
		})
	}
}

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

func compareUsers(t *testing.T, expected, actual *ent.User) {
	assert.Equal(t, expected.ID, actual.ID, "User ID should match")
	assert.True(t, expected.CreateTime.Equal(actual.CreateTime.UTC()), "User CreateTime should match")
	assert.True(t, expected.UpdateTime.Equal(actual.UpdateTime.UTC()), "User UpdateTime should match")
	assert.Equal(t, expected.AuthID, actual.AuthID, "User AuthID should match")
	assert.Equal(t, expected.FirstName, actual.FirstName, "User FirstName should match")
	assert.Equal(t, expected.LastName, actual.LastName, "User LastName should match")
	assert.Equal(t, expected.NickName, actual.NickName, "User NickName should match")
}
