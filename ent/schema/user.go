package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("authID").
			Immutable().
			NotEmpty(),
		field.String("firstName").Optional(),
		field.String("lastName").Optional(),
		field.String("nickName").Optional(),
		field.String("name").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", Group.Type),
		edge.To("definitions", Definition.Type).
			Immutable().
			Annotations(entgql.RelayConnection()),
		edge.To("words", Word.Type).
			Immutable().
			Annotations(entgql.RelayConnection()),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationUpdate(),
		),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
