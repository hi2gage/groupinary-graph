package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Definition holds the schema definition for the Definition entity.
type Definition struct {
	ent.Schema
}

// Fields of the Definition.
func (Definition) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").NotEmpty().Annotations(
			entgql.OrderField("ALPHA"), // Specify the ordering tag
		),
	}
}

// Edges of the Definition.
func (Definition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("word", Word.Type).
			Ref("definitions").
			Unique(),
		edge.From("creator", User.Type).
			Ref("definitions").
			Unique(),
	}
}

func (Definition) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
