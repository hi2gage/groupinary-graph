package utils

import (
	"context"
	"errors"
)

// UserContextKey is a key used to store and retrieve user IDs in context.Context.
type UserContextKey struct{}

func AddUserIdToContext(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, UserContextKey{}, userID)
}

func GetCreatorIDFromContext(ctx context.Context) (int, error) {
	creatorID, ok := ctx.Value(UserContextKey{}).(int)
	if !ok {
		return 0, errors.New("could not retrieve user_id from context")
	}
	return creatorID, nil
}
