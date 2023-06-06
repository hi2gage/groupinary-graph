// Code generated by ent, DO NOT EDIT.

package definition

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the definition type in the database.
	Label = "definition"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeWord holds the string denoting the word edge name in mutations.
	EdgeWord = "word"
	// Table holds the table name of the definition in the database.
	Table = "definitions"
	// WordTable is the table that holds the word relation/edge.
	WordTable = "definitions"
	// WordInverseTable is the table name for the Word entity.
	// It exists in this package in order to avoid circular dependency with the "word" package.
	WordInverseTable = "words"
	// WordColumn is the table column denoting the word relation/edge.
	WordColumn = "word_definitions"
)

// Columns holds all SQL columns for definition fields.
var Columns = []string{
	FieldID,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "definitions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"word_definitions",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)

// OrderOption defines the ordering options for the Definition queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByWordField orders the results by word field.
func ByWordField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWordStep(), sql.OrderByField(field, opts...))
	}
}
func newWordStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WordInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, WordTable, WordColumn),
	)
}
