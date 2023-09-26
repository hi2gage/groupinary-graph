package utils

import (
	"context"
	"errors"
)

func GetCreatorIDFromContext(ctx context.Context) (int, error) {
	creatorID, ok := ctx.Value("userID").(int)
	if !ok {
		return 0, errors.New("could not retrieve user_id from context")
	}
	return creatorID, nil
}
