//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
)

func main() {
	ex, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithSchemaGenerator(),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaPath("ent.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ex),
		entc.Extensions(entviz.Extension{}),
	}
	if err := entc.Generate("./ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
