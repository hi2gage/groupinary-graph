package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Word holds the schema definition for the Word entity.
type Word struct {
	ent.Schema
}

// Fields of the Word.
func (Word) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").NotEmpty(),
	}

}

// Edges of the Word.
func (Word) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).
			Ref("words").
			Unique(),
		edge.From("group", Group.Type).
			Ref("rootWords").
			Unique(),
		edge.To("definitions", Definition.Type).
			Annotations(entgql.RelayConnection()),
		edge.To("descendants", Word.Type).
			Annotations(entgql.RelayConnection()),
		edge.From("parents", Word.Type).
			Ref("descendants"),
	}
}

func (Word) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
