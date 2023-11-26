// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldAuthID holds the string denoting the authid field in the database.
	FieldAuthID = "auth_id"
	// FieldFirstName holds the string denoting the firstname field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the lastname field in the database.
	FieldLastName = "last_name"
	// FieldNickName holds the string denoting the nickname field in the database.
	FieldNickName = "nick_name"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeDefinitions holds the string denoting the definitions edge name in mutations.
	EdgeDefinitions = "definitions"
	// EdgeWords holds the string denoting the words edge name in mutations.
	EdgeWords = "words"
	// Table holds the table name of the user in the database.
	Table = "users"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "user_groups"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// DefinitionsTable is the table that holds the definitions relation/edge.
	DefinitionsTable = "definitions"
	// DefinitionsInverseTable is the table name for the Definition entity.
	// It exists in this package in order to avoid circular dependency with the "definition" package.
	DefinitionsInverseTable = "definitions"
	// DefinitionsColumn is the table column denoting the definitions relation/edge.
	DefinitionsColumn = "user_definitions"
	// WordsTable is the table that holds the words relation/edge.
	WordsTable = "words"
	// WordsInverseTable is the table name for the Word entity.
	// It exists in this package in order to avoid circular dependency with the "word" package.
	WordsInverseTable = "words"
	// WordsColumn is the table column denoting the words relation/edge.
	WordsColumn = "user_words"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldAuthID,
	FieldFirstName,
	FieldLastName,
	FieldNickName,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"user_id", "group_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// AuthIDValidator is a validator for the "authID" field. It is called by the builders before save.
	AuthIDValidator func(string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByAuthID orders the results by the authID field.
func ByAuthID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthID, opts...).ToFunc()
}

// ByFirstName orders the results by the firstName field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the lastName field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByNickName orders the results by the nickName field.
func ByNickName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickName, opts...).ToFunc()
}

// ByGroupsCount orders the results by groups count.
func ByGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupsStep(), opts...)
	}
}

// ByGroups orders the results by groups terms.
func ByGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDefinitionsCount orders the results by definitions count.
func ByDefinitionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDefinitionsStep(), opts...)
	}
}

// ByDefinitions orders the results by definitions terms.
func ByDefinitions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDefinitionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWordsCount orders the results by words count.
func ByWordsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWordsStep(), opts...)
	}
}

// ByWords orders the results by words terms.
func ByWords(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWordsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, GroupsTable, GroupsPrimaryKey...),
	)
}
func newDefinitionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DefinitionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DefinitionsTable, DefinitionsColumn),
	)
}
func newWordsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WordsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WordsTable, WordsColumn),
	)
}
