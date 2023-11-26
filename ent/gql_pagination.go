// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"groupinary/ent/definition"
	"groupinary/ent/group"
	"groupinary/ent/user"
	"groupinary/ent/word"
	"io"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// DefinitionEdge is the edge representation of Definition.
type DefinitionEdge struct {
	Node   *Definition `json:"node"`
	Cursor Cursor      `json:"cursor"`
}

// DefinitionConnection is the connection containing edges to Definition.
type DefinitionConnection struct {
	Edges      []*DefinitionEdge `json:"edges"`
	PageInfo   PageInfo          `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

func (c *DefinitionConnection) build(nodes []*Definition, pager *definitionPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Definition
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Definition {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Definition {
			return nodes[i]
		}
	}
	c.Edges = make([]*DefinitionEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &DefinitionEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// DefinitionPaginateOption enables pagination customization.
type DefinitionPaginateOption func(*definitionPager) error

// WithDefinitionOrder configures pagination ordering.
func WithDefinitionOrder(order *DefinitionOrder) DefinitionPaginateOption {
	if order == nil {
		order = DefaultDefinitionOrder
	}
	o := *order
	return func(pager *definitionPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultDefinitionOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithDefinitionFilter configures pagination filter.
func WithDefinitionFilter(filter func(*DefinitionQuery) (*DefinitionQuery, error)) DefinitionPaginateOption {
	return func(pager *definitionPager) error {
		if filter == nil {
			return errors.New("DefinitionQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type definitionPager struct {
	reverse bool
	order   *DefinitionOrder
	filter  func(*DefinitionQuery) (*DefinitionQuery, error)
}

func newDefinitionPager(opts []DefinitionPaginateOption, reverse bool) (*definitionPager, error) {
	pager := &definitionPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultDefinitionOrder
	}
	return pager, nil
}

func (p *definitionPager) applyFilter(query *DefinitionQuery) (*DefinitionQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *definitionPager) toCursor(d *Definition) Cursor {
	return p.order.Field.toCursor(d)
}

func (p *definitionPager) applyCursors(query *DefinitionQuery, after, before *Cursor) (*DefinitionQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultDefinitionOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *definitionPager) applyOrder(query *DefinitionQuery) *DefinitionQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultDefinitionOrder.Field {
		query = query.Order(DefaultDefinitionOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *definitionPager) orderExpr(query *DefinitionQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultDefinitionOrder.Field {
			b.Comma().Ident(DefaultDefinitionOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Definition.
func (d *DefinitionQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...DefinitionPaginateOption,
) (*DefinitionConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newDefinitionPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if d, err = pager.applyFilter(d); err != nil {
		return nil, err
	}
	conn := &DefinitionConnection{Edges: []*DefinitionEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = d.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if d, err = pager.applyCursors(d, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		d.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := d.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	d = pager.applyOrder(d)
	nodes, err := d.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// DefinitionOrderFieldDescription orders Definition by description.
	DefinitionOrderFieldDescription = &DefinitionOrderField{
		Value: func(d *Definition) (ent.Value, error) {
			return d.Description, nil
		},
		column: definition.FieldDescription,
		toTerm: definition.ByDescription,
		toCursor: func(d *Definition) Cursor {
			return Cursor{
				ID:    d.ID,
				Value: d.Description,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f DefinitionOrderField) String() string {
	var str string
	switch f.column {
	case DefinitionOrderFieldDescription.column:
		str = "ALPHA"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f DefinitionOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *DefinitionOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("DefinitionOrderField %T must be a string", v)
	}
	switch str {
	case "ALPHA":
		*f = *DefinitionOrderFieldDescription
	default:
		return fmt.Errorf("%s is not a valid DefinitionOrderField", str)
	}
	return nil
}

// DefinitionOrderField defines the ordering field of Definition.
type DefinitionOrderField struct {
	// Value extracts the ordering value from the given Definition.
	Value    func(*Definition) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) definition.OrderOption
	toCursor func(*Definition) Cursor
}

// DefinitionOrder defines the ordering of Definition.
type DefinitionOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     *DefinitionOrderField `json:"field"`
}

// DefaultDefinitionOrder is the default ordering of Definition.
var DefaultDefinitionOrder = &DefinitionOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &DefinitionOrderField{
		Value: func(d *Definition) (ent.Value, error) {
			return d.ID, nil
		},
		column: definition.FieldID,
		toTerm: definition.ByID,
		toCursor: func(d *Definition) Cursor {
			return Cursor{ID: d.ID}
		},
	},
}

// ToEdge converts Definition into DefinitionEdge.
func (d *Definition) ToEdge(order *DefinitionOrder) *DefinitionEdge {
	if order == nil {
		order = DefaultDefinitionOrder
	}
	return &DefinitionEdge{
		Node:   d,
		Cursor: order.Field.toCursor(d),
	}
}

// GroupEdge is the edge representation of Group.
type GroupEdge struct {
	Node   *Group `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// GroupConnection is the connection containing edges to Group.
type GroupConnection struct {
	Edges      []*GroupEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *GroupConnection) build(nodes []*Group, pager *groupPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Group
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Group {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Group {
			return nodes[i]
		}
	}
	c.Edges = make([]*GroupEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &GroupEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// GroupPaginateOption enables pagination customization.
type GroupPaginateOption func(*groupPager) error

// WithGroupOrder configures pagination ordering.
func WithGroupOrder(order *GroupOrder) GroupPaginateOption {
	if order == nil {
		order = DefaultGroupOrder
	}
	o := *order
	return func(pager *groupPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultGroupOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithGroupFilter configures pagination filter.
func WithGroupFilter(filter func(*GroupQuery) (*GroupQuery, error)) GroupPaginateOption {
	return func(pager *groupPager) error {
		if filter == nil {
			return errors.New("GroupQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type groupPager struct {
	reverse bool
	order   *GroupOrder
	filter  func(*GroupQuery) (*GroupQuery, error)
}

func newGroupPager(opts []GroupPaginateOption, reverse bool) (*groupPager, error) {
	pager := &groupPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultGroupOrder
	}
	return pager, nil
}

func (p *groupPager) applyFilter(query *GroupQuery) (*GroupQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *groupPager) toCursor(gr *Group) Cursor {
	return p.order.Field.toCursor(gr)
}

func (p *groupPager) applyCursors(query *GroupQuery, after, before *Cursor) (*GroupQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultGroupOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *groupPager) applyOrder(query *GroupQuery) *GroupQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultGroupOrder.Field {
		query = query.Order(DefaultGroupOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *groupPager) orderExpr(query *GroupQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultGroupOrder.Field {
			b.Comma().Ident(DefaultGroupOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Group.
func (gr *GroupQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...GroupPaginateOption,
) (*GroupConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newGroupPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if gr, err = pager.applyFilter(gr); err != nil {
		return nil, err
	}
	conn := &GroupConnection{Edges: []*GroupEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = gr.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if gr, err = pager.applyCursors(gr, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		gr.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := gr.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	gr = pager.applyOrder(gr)
	nodes, err := gr.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// GroupOrderField defines the ordering field of Group.
type GroupOrderField struct {
	// Value extracts the ordering value from the given Group.
	Value    func(*Group) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) group.OrderOption
	toCursor func(*Group) Cursor
}

// GroupOrder defines the ordering of Group.
type GroupOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *GroupOrderField `json:"field"`
}

// DefaultGroupOrder is the default ordering of Group.
var DefaultGroupOrder = &GroupOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &GroupOrderField{
		Value: func(gr *Group) (ent.Value, error) {
			return gr.ID, nil
		},
		column: group.FieldID,
		toTerm: group.ByID,
		toCursor: func(gr *Group) Cursor {
			return Cursor{ID: gr.ID}
		},
	},
}

// ToEdge converts Group into GroupEdge.
func (gr *Group) ToEdge(order *GroupOrder) *GroupEdge {
	if order == nil {
		order = DefaultGroupOrder
	}
	return &GroupEdge{
		Node:   gr,
		Cursor: order.Field.toCursor(gr),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	reverse bool
	order   *UserOrder
	filter  func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption, reverse bool) (*userPager, error) {
	pager := &userPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) (*UserQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultUserOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *userPager) applyOrder(query *UserQuery) *UserQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(DefaultUserOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *userPager) orderExpr(query *UserQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = u.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if u, err = pager.applyCursors(u, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	u = pager.applyOrder(u)
	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	// Value extracts the ordering value from the given User.
	Value    func(*User) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) user.OrderOption
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.ID, nil
		},
		column: user.FieldID,
		toTerm: user.ByID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}

// WordEdge is the edge representation of Word.
type WordEdge struct {
	Node   *Word  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// WordConnection is the connection containing edges to Word.
type WordConnection struct {
	Edges      []*WordEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *WordConnection) build(nodes []*Word, pager *wordPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Word
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Word {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Word {
			return nodes[i]
		}
	}
	c.Edges = make([]*WordEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &WordEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// WordPaginateOption enables pagination customization.
type WordPaginateOption func(*wordPager) error

// WithWordOrder configures pagination ordering.
func WithWordOrder(order *WordOrder) WordPaginateOption {
	if order == nil {
		order = DefaultWordOrder
	}
	o := *order
	return func(pager *wordPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultWordOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithWordFilter configures pagination filter.
func WithWordFilter(filter func(*WordQuery) (*WordQuery, error)) WordPaginateOption {
	return func(pager *wordPager) error {
		if filter == nil {
			return errors.New("WordQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type wordPager struct {
	reverse bool
	order   *WordOrder
	filter  func(*WordQuery) (*WordQuery, error)
}

func newWordPager(opts []WordPaginateOption, reverse bool) (*wordPager, error) {
	pager := &wordPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultWordOrder
	}
	return pager, nil
}

func (p *wordPager) applyFilter(query *WordQuery) (*WordQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *wordPager) toCursor(w *Word) Cursor {
	return p.order.Field.toCursor(w)
}

func (p *wordPager) applyCursors(query *WordQuery, after, before *Cursor) (*WordQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultWordOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *wordPager) applyOrder(query *WordQuery) *WordQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultWordOrder.Field {
		query = query.Order(DefaultWordOrder.Field.toTerm(direction.OrderTermOption()))
	}
	switch p.order.Field.column {
	case WordOrderFieldDefinitionsCount.column, WordOrderFieldDescendantsCount.column:
	default:
		if len(query.ctx.Fields) > 0 {
			query.ctx.AppendFieldOnce(p.order.Field.column)
		}
	}
	return query
}

func (p *wordPager) orderExpr(query *WordQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	switch p.order.Field.column {
	case WordOrderFieldDefinitionsCount.column, WordOrderFieldDescendantsCount.column:
		query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	default:
		if len(query.ctx.Fields) > 0 {
			query.ctx.AppendFieldOnce(p.order.Field.column)
		}
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultWordOrder.Field {
			b.Comma().Ident(DefaultWordOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Word.
func (w *WordQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...WordPaginateOption,
) (*WordConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newWordPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if w, err = pager.applyFilter(w); err != nil {
		return nil, err
	}
	conn := &WordConnection{Edges: []*WordEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = w.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if w, err = pager.applyCursors(w, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		w.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := w.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	w = pager.applyOrder(w)
	nodes, err := w.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// WordOrderFieldDescription orders Word by description.
	WordOrderFieldDescription = &WordOrderField{
		Value: func(w *Word) (ent.Value, error) {
			return w.Description, nil
		},
		column: word.FieldDescription,
		toTerm: word.ByDescription,
		toCursor: func(w *Word) Cursor {
			return Cursor{
				ID:    w.ID,
				Value: w.Description,
			}
		},
	}
	// WordOrderFieldDefinitionsCount orders by DEFINITIONS_COUNT.
	WordOrderFieldDefinitionsCount = &WordOrderField{
		Value: func(w *Word) (ent.Value, error) {
			return w.Value("definitions_count")
		},
		column: "definitions_count",
		toTerm: func(opts ...sql.OrderTermOption) word.OrderOption {
			return word.ByDefinitionsCount(
				append(opts, sql.OrderSelectAs("definitions_count"))...,
			)
		},
		toCursor: func(w *Word) Cursor {
			cv, _ := w.Value("definitions_count")
			return Cursor{
				ID:    w.ID,
				Value: cv,
			}
		},
	}
	// WordOrderFieldDescendantsCount orders by DESCENDANTS_COUNT.
	WordOrderFieldDescendantsCount = &WordOrderField{
		Value: func(w *Word) (ent.Value, error) {
			return w.Value("descendants_count")
		},
		column: "descendants_count",
		toTerm: func(opts ...sql.OrderTermOption) word.OrderOption {
			return word.ByDescendantsCount(
				append(opts, sql.OrderSelectAs("descendants_count"))...,
			)
		},
		toCursor: func(w *Word) Cursor {
			cv, _ := w.Value("descendants_count")
			return Cursor{
				ID:    w.ID,
				Value: cv,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f WordOrderField) String() string {
	var str string
	switch f.column {
	case WordOrderFieldDescription.column:
		str = "ALPHA"
	case WordOrderFieldDefinitionsCount.column:
		str = "DEFINITIONS_COUNT"
	case WordOrderFieldDescendantsCount.column:
		str = "DESCENDANTS_COUNT"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f WordOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *WordOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("WordOrderField %T must be a string", v)
	}
	switch str {
	case "ALPHA":
		*f = *WordOrderFieldDescription
	case "DEFINITIONS_COUNT":
		*f = *WordOrderFieldDefinitionsCount
	case "DESCENDANTS_COUNT":
		*f = *WordOrderFieldDescendantsCount
	default:
		return fmt.Errorf("%s is not a valid WordOrderField", str)
	}
	return nil
}

// WordOrderField defines the ordering field of Word.
type WordOrderField struct {
	// Value extracts the ordering value from the given Word.
	Value    func(*Word) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) word.OrderOption
	toCursor func(*Word) Cursor
}

// WordOrder defines the ordering of Word.
type WordOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *WordOrderField `json:"field"`
}

// DefaultWordOrder is the default ordering of Word.
var DefaultWordOrder = &WordOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &WordOrderField{
		Value: func(w *Word) (ent.Value, error) {
			return w.ID, nil
		},
		column: word.FieldID,
		toTerm: word.ByID,
		toCursor: func(w *Word) Cursor {
			return Cursor{ID: w.ID}
		},
	},
}

// ToEdge converts Word into WordEdge.
func (w *Word) ToEdge(order *WordOrder) *WordEdge {
	if order == nil {
		order = DefaultWordOrder
	}
	return &WordEdge{
		Node:   w,
		Cursor: order.Field.toCursor(w),
	}
}
