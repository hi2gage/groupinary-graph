package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"groupinary/ent"
	"groupinary/graph"
	"groupinary/utils"
)

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, name string, description *string) (*ent.Group, error) {
	group := r.client.Group.Create()
	group.SetName(name).SetNillableDescription(description)

	return group.Save(ctx)
}

// UpdateGroupName is the resolver for the updateGroupName field.
func (r *mutationResolver) UpdateGroupName(ctx context.Context, id int, name string) (*ent.Group, error) {
	groupUpdate := r.client.Group.UpdateOneID(id)

	groupUpdate.SetName(name)

	return groupUpdate.Save(ctx)
}

// DeleteGroup is the resolver for the deleteGroup field.
func (r *mutationResolver) DeleteGroup(ctx context.Context, id int) (bool, error) {
	if err := r.client.Group.DeleteOneID(id).Exec(ctx); err != nil {
		return false, err
	}
	return true, nil
}

// DeleteWord is the resolver for the deleteWord field.
func (r *mutationResolver) DeleteWord(ctx context.Context, id int) (bool, error) {
	if err := r.client.Word.DeleteOneID(id).Exec(ctx); err != nil {
		return false, err
	}
	return true, nil
}

// DeleteDefinition is the resolver for the deleteDefinition field.
func (r *mutationResolver) DeleteDefinition(ctx context.Context, id int) (bool, error) {
	if err := r.client.Definition.DeleteOneID(id).Exec(ctx); err != nil {
		return false, err
	}
	return true, nil
}

// UpdateUserName is the resolver for the updateUserName field.
func (r *mutationResolver) UpdateUserName(ctx context.Context, firstName string, lastName *string, nickName *string) (*ent.User, error) {
	creatorID, err := utils.GetCreatorIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	u := r.client.User.
		UpdateOneID(creatorID).
		SetFirstName(firstName).
		SetNillableLastName(lastName).
		SetNillableNickName(nickName)

	return u.Save(ctx)
}

// JoinGroup is the resolver for the joinGroup field.
func (r *mutationResolver) JoinGroup(ctx context.Context, groupID int) (*ent.Group, error) {
	creatorID, err := utils.GetCreatorIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	groupAppend := r.client.Group.UpdateOneID(groupID)

	groupAppend.AddUserIDs(creatorID)

	return groupAppend.Save(ctx)
}

// AddRootWord is the resolver for the addRootWord field.
func (r *mutationResolver) AddRootWord(ctx context.Context, rootWord string, groupID int, rootDefinition *string) (*ent.Word, error) {
	creatorID, err := utils.GetCreatorIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	wordCreate := r.client.Word.Create().
		SetCreatorID(creatorID).
		SetDescription(rootWord).
		SetGroupID(groupID)

	if rootDefinition != nil {
		definition, err := r.client.Definition.Create().
			SetCreatorID(creatorID).
			SetDescription(*rootDefinition).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		wordCreate.AddDefinitions(definition)
	}

	return wordCreate.Save(ctx)
}

// AddChildWord is the resolver for the addChildWord field.
func (r *mutationResolver) AddChildWord(ctx context.Context, rootIds []int, groupID int, childWord string, childDefinition *string) (*ent.Word, error) {
	creatorID, err := utils.GetCreatorIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	wordCreate := r.client.Word.Create().
		SetCreatorID(creatorID).
		SetGroupID(groupID).
		SetDescription(childWord)

	if childDefinition != nil {
		definition, err := r.client.Definition.Create().
			SetCreatorID(creatorID).
			SetDescription(*childDefinition).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		wordCreate.AddDefinitions(definition)
	}

	wordCreate.AddParentIDs(rootIds...)

	return wordCreate.Save(ctx)
}

// AddDefinition is the resolver for the addDefinition field.
func (r *mutationResolver) AddDefinition(ctx context.Context, wordID int, definition string) (*ent.Definition, error) {
	creatorID, err := utils.GetCreatorIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	definitionCreate := r.client.Definition.Create().
		SetCreatorID(creatorID).
		SetDescription(definition).
		SetWordID(wordID)

	return definitionCreate.Save(ctx)
}

// ConnectWords is the resolver for the connectWords field.
func (r *mutationResolver) ConnectWords(ctx context.Context, parentID int, childID int) (*ent.Word, error) {
	childWord := r.client.Word.UpdateOneID(childID)

	childWord.AddParentIDs(parentID)

	return childWord.Save(ctx)
}

// UpdateWord is the resolver for the updateWord field.
func (r *mutationResolver) UpdateWord(ctx context.Context, id int, wordDescription string) (*ent.Word, error) {
	wordUpdate := r.client.Word.UpdateOneID(id)

	wordUpdate.SetDescription(wordDescription)

	return wordUpdate.Save(ctx)
}

// UpdateDefinition is the resolver for the updateDefinition field.
func (r *mutationResolver) UpdateDefinition(ctx context.Context, id int, definitionDescription string) (*ent.Definition, error) {
	definitionUpdate := r.client.Definition.UpdateOneID(id)

	definitionUpdate.SetDescription(definitionDescription)

	return definitionUpdate.Save(ctx)
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
