package middleware

import (
	"context"
	"fmt"
	"groupinary/ent"
	"groupinary/ent/user"
	"log"
)

// UserOperations defines the interface for user-related operations.
type UserOperations interface {
	CheckUserExists(authID string) (*int, error)
	AddUserToGraph(authID string) (*int, error)
}

// UserTokenOperator is an implementation of UserOperations using ent.Client.
type UserTokenOperator struct {
	Client *ent.Client
}

// NewUserTokenOperator creates a new RealUserOperations instance.
func NewUserTokenOperator(client *ent.Client) *UserTokenOperator {
	return &UserTokenOperator{
		Client: client,
	}
}

// CheckUserExists checks if the user with the given authID exists in the Graph.
func (u *UserTokenOperator) CheckUserExists(authID string) (*int, error) {
	user, err := u.Client.User.Query().Where(user.AuthID(authID)).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("user does not exist: %w", err)
		}
		return nil, fmt.Errorf("querying user: %w", err)
	}
	return &user.ID, nil
}

// AddUserToGraph adds the user to the Graph with the given AuthID if it does not exist.
func (u *UserTokenOperator) AddUserToGraph(authID string) (*int, error) {
	user, err := u.Client.User.Create().SetAuthID(authID).Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}
	log.Printf("finished AddUserToGraph")
	return &user.ID, nil
}
