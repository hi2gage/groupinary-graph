package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Word holds the schema definition for the Word entity.
type Word struct {
	ent.Schema
	DescendantCount int `json:"descendantCount"` // Field for descendant count.
}

// Fields of the Word.
func (Word) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").
			NotEmpty().
			Annotations(entgql.OrderField("ALPHA")),
	}
}

// Edges of the Word.
func (Word) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).
			Ref("words").
			Immutable().
			Unique(),
		edge.From("group", Group.Type).
			Ref("words").
			Required().
			Unique(),
		edge.To("definitions", Definition.Type).
			Immutable().
			Annotations(
				entgql.RelayConnection(),
				entgql.OrderField("DEFINITIONS_COUNT"),
			),
		edge.To("descendants", Word.Type).
			Immutable().
			Annotations(
				entgql.RelayConnection(),
				entgql.OrderField("DESCENDANTS_COUNT"),
				// entsql.Annotation{OnDelete: entsql.Cascade},
			).
			From("parents"),
		// edge.From("parents", Word.Type).
		// 	Ref("descendants"),
	}
}

// Annotations for the Word.
func (Word) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Word) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
