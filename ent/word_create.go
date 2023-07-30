// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"shrektionary_api/ent/definition"
	"shrektionary_api/ent/user"
	"shrektionary_api/ent/word"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WordCreate is the builder for creating a Word entity.
type WordCreate struct {
	config
	mutation *WordMutation
	hooks    []Hook
}

// SetDescription sets the "description" field.
func (wc *WordCreate) SetDescription(s string) *WordCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetRoot sets the "root" field.
func (wc *WordCreate) SetRoot(b bool) *WordCreate {
	wc.mutation.SetRoot(b)
	return wc
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (wc *WordCreate) SetCreatorID(id int) *WordCreate {
	wc.mutation.SetCreatorID(id)
	return wc
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (wc *WordCreate) SetNillableCreatorID(id *int) *WordCreate {
	if id != nil {
		wc = wc.SetCreatorID(*id)
	}
	return wc
}

// SetCreator sets the "creator" edge to the User entity.
func (wc *WordCreate) SetCreator(u *User) *WordCreate {
	return wc.SetCreatorID(u.ID)
}

// AddDefinitionIDs adds the "definitions" edge to the Definition entity by IDs.
func (wc *WordCreate) AddDefinitionIDs(ids ...int) *WordCreate {
	wc.mutation.AddDefinitionIDs(ids...)
	return wc
}

// AddDefinitions adds the "definitions" edges to the Definition entity.
func (wc *WordCreate) AddDefinitions(d ...*Definition) *WordCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return wc.AddDefinitionIDs(ids...)
}

// AddDescendantIDs adds the "descendants" edge to the Word entity by IDs.
func (wc *WordCreate) AddDescendantIDs(ids ...int) *WordCreate {
	wc.mutation.AddDescendantIDs(ids...)
	return wc
}

// AddDescendants adds the "descendants" edges to the Word entity.
func (wc *WordCreate) AddDescendants(w ...*Word) *WordCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wc.AddDescendantIDs(ids...)
}

// SetParentID sets the "parent" edge to the Word entity by ID.
func (wc *WordCreate) SetParentID(id int) *WordCreate {
	wc.mutation.SetParentID(id)
	return wc
}

// SetNillableParentID sets the "parent" edge to the Word entity by ID if the given value is not nil.
func (wc *WordCreate) SetNillableParentID(id *int) *WordCreate {
	if id != nil {
		wc = wc.SetParentID(*id)
	}
	return wc
}

// SetParent sets the "parent" edge to the Word entity.
func (wc *WordCreate) SetParent(w *Word) *WordCreate {
	return wc.SetParentID(w.ID)
}

// Mutation returns the WordMutation object of the builder.
func (wc *WordCreate) Mutation() *WordMutation {
	return wc.mutation
}

// Save creates the Word in the database.
func (wc *WordCreate) Save(ctx context.Context) (*Word, error) {
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WordCreate) SaveX(ctx context.Context) *Word {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WordCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WordCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WordCreate) check() error {
	if _, ok := wc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Word.description"`)}
	}
	if v, ok := wc.mutation.Description(); ok {
		if err := word.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Word.description": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Root(); !ok {
		return &ValidationError{Name: "root", err: errors.New(`ent: missing required field "Word.root"`)}
	}
	return nil
}

func (wc *WordCreate) sqlSave(ctx context.Context) (*Word, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WordCreate) createSpec() (*Word, *sqlgraph.CreateSpec) {
	var (
		_node = &Word{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(word.Table, sqlgraph.NewFieldSpec(word.FieldID, field.TypeInt))
	)
	if value, ok := wc.mutation.Description(); ok {
		_spec.SetField(word.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := wc.mutation.Root(); ok {
		_spec.SetField(word.FieldRoot, field.TypeBool, value)
		_node.Root = value
	}
	if nodes := wc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   word.CreatorTable,
			Columns: []string{word.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_words = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.DefinitionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   word.DefinitionsTable,
			Columns: []string{word.DefinitionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(definition.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.DescendantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   word.DescendantsTable,
			Columns: []string{word.DescendantsColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(word.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   word.ParentTable,
			Columns: []string{word.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(word.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.word_descendants = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WordCreateBulk is the builder for creating many Word entities in bulk.
type WordCreateBulk struct {
	config
	builders []*WordCreate
}

// Save creates the Word entities in the database.
func (wcb *WordCreateBulk) Save(ctx context.Context) ([]*Word, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Word, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WordMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WordCreateBulk) SaveX(ctx context.Context) []*Word {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WordCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WordCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
