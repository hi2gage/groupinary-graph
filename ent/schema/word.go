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
	DescendantCount int `json:"descendantCount"` // Field for descendant count.
}

// Fields of the Word.
func (Word) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").NotEmpty(),
		field.Int("descendantCount").
			StorageKey("descendant_count").
			Default(0).
			Immutable(), // Make the field immutable.
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

// Annotations for the Word.
func (Word) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}

// // Hooks for the Word.
// func (Word) Hooks() []ent.Hook {
// 	return []ent.Hook{
// 		CountDescendantsHook(),
// 	}
// }

// // CountDescendantsHook is a custom hook to update the descendant count field.
// func CountDescendantsHook() ent.Hook {
// 	return func(next ent.Mutator) ent.Mutator {
// 		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
// 			// Call the original mutation operation.
// 			v, err := next.Mutate(ctx, m)
// 			if err != nil {
// 				return nil, err
// 			}

// 			// Update the descendant count after the mutation is applied.
// 			word, ok := v.(*Word)
// 			if ok {
// 				count, err := ctx.Ent().Word.
// 					Query().
// 					Where(word.HasParentsWith(word.Field("id"))).
// 					Count(ctx)
// 				if err != nil {
// 					return nil, err
// 				}
// 				word.DescendantCount = count
// 			}

// 			return v, nil
// 		})
// 	}
// }
