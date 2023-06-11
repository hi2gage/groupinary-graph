package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// WordConnections holds the schema definition for the WordConnections entity.
type WordConnections struct {
	ent.Schema
}

// Fields of the WordConnections.
func (WordConnections) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").NotEmpty(),
	}
}

// Edges of the WordConnections.
func (WordConnections) Edges() []ent.Edge {
	return nil
}

func (WordConnections) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
