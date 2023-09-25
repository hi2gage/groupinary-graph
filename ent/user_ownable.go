package ent

import (
	"context"
	"errors"
)

type UserOwnable interface {
	SetCreatorID(ctx context.Context) error
	getCreatorID() *int
	ValidateCreateInput() error
}

// Words
func (input *CreateDefinitionInput) getCreatorID() *int {
	return input.CreatorID
}

func (input *CreateDefinitionInput) ValidateCreateInput() error {
	return validateCreateInput(input)
}

func (input *CreateWordInput) SetCreatorID(ctx context.Context) error {
	creatorID, err := getCreatorIDFromContext(ctx)
	if err != nil {
		return err
	}

	input.CreatorID = &creatorID
	return nil
}

// Definitions
func (input *CreateWordInput) ValidateCreateInput() error {
	return validateCreateInput(input)
}

func (input *CreateWordInput) getCreatorID() *int {
	return input.CreatorID
}

func (input *CreateDefinitionInput) SetCreatorID(ctx context.Context) error {
	creatorID, err := getCreatorIDFromContext(ctx)
	if err != nil {
		return err
	}

	input.CreatorID = &creatorID
	return nil
}

// Reusable UserOwnable Functions

func getCreatorIDFromContext(ctx context.Context) (int, error) {
	creatorID, ok := ctx.Value("userID").(int)
	if !ok {
		return 0, errors.New("could not retrieve user_id from context")
	}
	return creatorID, nil
}

func validateCreateInput(input UserOwnable) error {
	if err := validateCreateInputIsNil(input); err != nil {
		return err
	}
	return nil
}

func validateCreateInputIsNil(input UserOwnable) error {
	if input.getCreatorID() != nil {
		return errors.New("invalid value provided for CreatorID, must be nil")
	}
	return nil
}
