package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"errors"
	"shrektionary_api/ent"
)

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input ent.CreateGroupInput) (*ent.Group, error) {
	return r.client.Group.Create().SetInput(input).Save(ctx)
}

// CreateDefinition is the resolver for the createDefinition field.
func (r *mutationResolver) CreateDefinition(ctx context.Context, input ent.CreateDefinitionInput) (*ent.Definition, error) {
	err := input.ValidateCreateInput()
	if err != nil {
		return nil, err
	}

	input.SetCreatorID(ctx)

	return r.client.Definition.Create().SetInput(input).Save(ctx)
}

// CreateWord is the resolver for the createWord field.
func (r *mutationResolver) CreateWord(ctx context.Context, input ent.CreateWordInput) (*ent.Word, error) {
	err := input.ValidateCreateInput()
	if err != nil {
		return nil, err
	}

	input.SetCreatorID(ctx)

	return r.client.Word.Create().SetInput(input).Save(ctx)
}

// CreateRootWord is the resolver for the createRootWord field.
func (r *mutationResolver) CreateRootWord(ctx context.Context, input ent.CreateWordInput) (*ent.Word, error) {
	err := input.ValidateCreateInput()
	if err != nil {
		return nil, err
	}

	input.SetCreatorID(ctx)

	return r.client.Word.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input ent.UpdateUserInput) (*ent.User, error) {
	creatorID, err := getCreatorIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.User.UpdateOneID(creatorID).SetInput(input).Save(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func getCreatorIDFromContext(ctx context.Context) (int, error) {
	creatorID, ok := ctx.Value("userID").(int)
	if !ok {
		return 0, errors.New("could not retrieve user_id from context")
	}
	return creatorID, nil
}
