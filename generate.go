//go:build ignore

package shrektionary_api

//go:generate go run -mod=mod ./ent/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration ./ent/schema
