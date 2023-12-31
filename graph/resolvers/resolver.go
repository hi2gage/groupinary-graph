package resolvers

import (
	"groupinary/ent"

	"groupinary/graph"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return graph.NewExecutableSchema(graph.Config{
		Resolvers: &Resolver{client},
	})
}
