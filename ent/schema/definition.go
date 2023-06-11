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
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("word", Word.Type).
			Ref("definitions").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}
}

func (Definition) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
