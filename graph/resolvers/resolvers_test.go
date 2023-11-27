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
			// Assuming the Resolver struct also has a client field
			client: client,
		},
	}

	testCreatorID := 1
	ctx := context.Background()

	testUser := &ent.User{
		ID:         testCreatorID,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		AuthID:     "test_auth_id",
		FirstName:  "Test",
		LastName:   "User",
		NickName:   "TestNick",
	}

	// Insert the test user into the database
	user, err := client.User.Create().
		SetCreateTime(testUser.CreateTime).
		SetUpdateTime(testUser.UpdateTime).
		SetAuthID(testUser.AuthID).
		SetFirstName(testUser.FirstName).
		SetLastName(testUser.LastName).
		SetNickName(testUser.NickName).
		Save(ctx)

	if err != nil {
		t.Fatal(err)
	}

	// Add user.ID into the context
	ctx = utils.AddUserIdToContext(ctx, user.ID)

	// Call the function and check the result
	resultUser, err := resolver.CurrentUser(ctx)
	assert.NoError(t, err, "Unexpected error")
	assert.NotNil(t, resultUser, "User should not be nil")
	compareUsers(t, testUser, resultUser)
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
