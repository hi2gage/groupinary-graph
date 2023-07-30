// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"shrektionary_api/ent/user"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AuthID holds the value of the "authID" field.
	AuthID string `json:"authID,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// Definitions holds the value of the definitions edge.
	Definitions []*Definition `json:"definitions,omitempty"`
	// Words holds the value of the words edge.
	Words []*Word `json:"words,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedGroups      map[string][]*Group
	namedDefinitions map[string][]*Definition
	namedWords       map[string][]*Word
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[0] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// DefinitionsOrErr returns the Definitions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DefinitionsOrErr() ([]*Definition, error) {
	if e.loadedTypes[1] {
		return e.Definitions, nil
	}
	return nil, &NotLoadedError{edge: "definitions"}
}

// WordsOrErr returns the Words value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WordsOrErr() ([]*Word, error) {
	if e.loadedTypes[2] {
		return e.Words, nil
	}
	return nil, &NotLoadedError{edge: "words"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldAuthID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldAuthID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field authID", values[i])
			} else if value.Valid {
				u.AuthID = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryGroups queries the "groups" edge of the User entity.
func (u *User) QueryGroups() *GroupQuery {
	return NewUserClient(u.config).QueryGroups(u)
}

// QueryDefinitions queries the "definitions" edge of the User entity.
func (u *User) QueryDefinitions() *DefinitionQuery {
	return NewUserClient(u.config).QueryDefinitions(u)
}

// QueryWords queries the "words" edge of the User entity.
func (u *User) QueryWords() *WordQuery {
	return NewUserClient(u.config).QueryWords(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("authID=")
	builder.WriteString(u.AuthID)
	builder.WriteByte(')')
	return builder.String()
}

// NamedGroups returns the Groups named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedGroups(name string) ([]*Group, error) {
	if u.Edges.namedGroups == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedGroups[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedGroups(name string, edges ...*Group) {
	if u.Edges.namedGroups == nil {
		u.Edges.namedGroups = make(map[string][]*Group)
	}
	if len(edges) == 0 {
		u.Edges.namedGroups[name] = []*Group{}
	} else {
		u.Edges.namedGroups[name] = append(u.Edges.namedGroups[name], edges...)
	}
}

// NamedDefinitions returns the Definitions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedDefinitions(name string) ([]*Definition, error) {
	if u.Edges.namedDefinitions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedDefinitions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedDefinitions(name string, edges ...*Definition) {
	if u.Edges.namedDefinitions == nil {
		u.Edges.namedDefinitions = make(map[string][]*Definition)
	}
	if len(edges) == 0 {
		u.Edges.namedDefinitions[name] = []*Definition{}
	} else {
		u.Edges.namedDefinitions[name] = append(u.Edges.namedDefinitions[name], edges...)
	}
}

// NamedWords returns the Words named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedWords(name string) ([]*Word, error) {
	if u.Edges.namedWords == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedWords[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedWords(name string, edges ...*Word) {
	if u.Edges.namedWords == nil {
		u.Edges.namedWords = make(map[string][]*Word)
	}
	if len(edges) == 0 {
		u.Edges.namedWords[name] = []*Word{}
	} else {
		u.Edges.namedWords[name] = append(u.Edges.namedWords[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User