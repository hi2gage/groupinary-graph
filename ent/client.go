// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"shrektionary_api/ent/migrate"

	"shrektionary_api/ent/definition"
	"shrektionary_api/ent/group"
	"shrektionary_api/ent/word"
	"shrektionary_api/ent/wordconnections"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Definition is the client for interacting with the Definition builders.
	Definition *DefinitionClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// Word is the client for interacting with the Word builders.
	Word *WordClient
	// WordConnections is the client for interacting with the WordConnections builders.
	WordConnections *WordConnectionsClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Definition = NewDefinitionClient(c.config)
	c.Group = NewGroupClient(c.config)
	c.Word = NewWordClient(c.config)
	c.WordConnections = NewWordConnectionsClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Definition:      NewDefinitionClient(cfg),
		Group:           NewGroupClient(cfg),
		Word:            NewWordClient(cfg),
		WordConnections: NewWordConnectionsClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Definition:      NewDefinitionClient(cfg),
		Group:           NewGroupClient(cfg),
		Word:            NewWordClient(cfg),
		WordConnections: NewWordConnectionsClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Definition.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Definition.Use(hooks...)
	c.Group.Use(hooks...)
	c.Word.Use(hooks...)
	c.WordConnections.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Definition.Intercept(interceptors...)
	c.Group.Intercept(interceptors...)
	c.Word.Intercept(interceptors...)
	c.WordConnections.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *DefinitionMutation:
		return c.Definition.mutate(ctx, m)
	case *GroupMutation:
		return c.Group.mutate(ctx, m)
	case *WordMutation:
		return c.Word.mutate(ctx, m)
	case *WordConnectionsMutation:
		return c.WordConnections.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// DefinitionClient is a client for the Definition schema.
type DefinitionClient struct {
	config
}

// NewDefinitionClient returns a client for the Definition from the given config.
func NewDefinitionClient(c config) *DefinitionClient {
	return &DefinitionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `definition.Hooks(f(g(h())))`.
func (c *DefinitionClient) Use(hooks ...Hook) {
	c.hooks.Definition = append(c.hooks.Definition, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `definition.Intercept(f(g(h())))`.
func (c *DefinitionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Definition = append(c.inters.Definition, interceptors...)
}

// Create returns a builder for creating a Definition entity.
func (c *DefinitionClient) Create() *DefinitionCreate {
	mutation := newDefinitionMutation(c.config, OpCreate)
	return &DefinitionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Definition entities.
func (c *DefinitionClient) CreateBulk(builders ...*DefinitionCreate) *DefinitionCreateBulk {
	return &DefinitionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Definition.
func (c *DefinitionClient) Update() *DefinitionUpdate {
	mutation := newDefinitionMutation(c.config, OpUpdate)
	return &DefinitionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DefinitionClient) UpdateOne(d *Definition) *DefinitionUpdateOne {
	mutation := newDefinitionMutation(c.config, OpUpdateOne, withDefinition(d))
	return &DefinitionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DefinitionClient) UpdateOneID(id int) *DefinitionUpdateOne {
	mutation := newDefinitionMutation(c.config, OpUpdateOne, withDefinitionID(id))
	return &DefinitionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Definition.
func (c *DefinitionClient) Delete() *DefinitionDelete {
	mutation := newDefinitionMutation(c.config, OpDelete)
	return &DefinitionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DefinitionClient) DeleteOne(d *Definition) *DefinitionDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DefinitionClient) DeleteOneID(id int) *DefinitionDeleteOne {
	builder := c.Delete().Where(definition.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DefinitionDeleteOne{builder}
}

// Query returns a query builder for Definition.
func (c *DefinitionClient) Query() *DefinitionQuery {
	return &DefinitionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDefinition},
		inters: c.Interceptors(),
	}
}

// Get returns a Definition entity by its id.
func (c *DefinitionClient) Get(ctx context.Context, id int) (*Definition, error) {
	return c.Query().Where(definition.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DefinitionClient) GetX(ctx context.Context, id int) *Definition {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWord queries the word edge of a Definition.
func (c *DefinitionClient) QueryWord(d *Definition) *WordQuery {
	query := (&WordClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(definition.Table, definition.FieldID, id),
			sqlgraph.To(word.Table, word.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, definition.WordTable, definition.WordColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DefinitionClient) Hooks() []Hook {
	return c.hooks.Definition
}

// Interceptors returns the client interceptors.
func (c *DefinitionClient) Interceptors() []Interceptor {
	return c.inters.Definition
}

func (c *DefinitionClient) mutate(ctx context.Context, m *DefinitionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DefinitionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DefinitionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DefinitionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DefinitionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Definition mutation op: %q", m.Op())
	}
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `group.Intercept(f(g(h())))`.
func (c *GroupClient) Intercept(interceptors ...Interceptor) {
	c.inters.Group = append(c.inters.Group, interceptors...)
}

// Create returns a builder for creating a Group entity.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGroup},
		inters: c.Interceptors(),
	}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// Interceptors returns the client interceptors.
func (c *GroupClient) Interceptors() []Interceptor {
	return c.inters.Group
}

func (c *GroupClient) mutate(ctx context.Context, m *GroupMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GroupCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GroupDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Group mutation op: %q", m.Op())
	}
}

// WordClient is a client for the Word schema.
type WordClient struct {
	config
}

// NewWordClient returns a client for the Word from the given config.
func NewWordClient(c config) *WordClient {
	return &WordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `word.Hooks(f(g(h())))`.
func (c *WordClient) Use(hooks ...Hook) {
	c.hooks.Word = append(c.hooks.Word, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `word.Intercept(f(g(h())))`.
func (c *WordClient) Intercept(interceptors ...Interceptor) {
	c.inters.Word = append(c.inters.Word, interceptors...)
}

// Create returns a builder for creating a Word entity.
func (c *WordClient) Create() *WordCreate {
	mutation := newWordMutation(c.config, OpCreate)
	return &WordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Word entities.
func (c *WordClient) CreateBulk(builders ...*WordCreate) *WordCreateBulk {
	return &WordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Word.
func (c *WordClient) Update() *WordUpdate {
	mutation := newWordMutation(c.config, OpUpdate)
	return &WordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WordClient) UpdateOne(w *Word) *WordUpdateOne {
	mutation := newWordMutation(c.config, OpUpdateOne, withWord(w))
	return &WordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WordClient) UpdateOneID(id int) *WordUpdateOne {
	mutation := newWordMutation(c.config, OpUpdateOne, withWordID(id))
	return &WordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Word.
func (c *WordClient) Delete() *WordDelete {
	mutation := newWordMutation(c.config, OpDelete)
	return &WordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *WordClient) DeleteOne(w *Word) *WordDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *WordClient) DeleteOneID(id int) *WordDeleteOne {
	builder := c.Delete().Where(word.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WordDeleteOne{builder}
}

// Query returns a query builder for Word.
func (c *WordClient) Query() *WordQuery {
	return &WordQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeWord},
		inters: c.Interceptors(),
	}
}

// Get returns a Word entity by its id.
func (c *WordClient) Get(ctx context.Context, id int) (*Word, error) {
	return c.Query().Where(word.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WordClient) GetX(ctx context.Context, id int) *Word {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDefinitions queries the definitions edge of a Word.
func (c *WordClient) QueryDefinitions(w *Word) *DefinitionQuery {
	query := (&DefinitionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, id),
			sqlgraph.To(definition.Table, definition.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, word.DefinitionsTable, word.DefinitionsColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WordClient) Hooks() []Hook {
	return c.hooks.Word
}

// Interceptors returns the client interceptors.
func (c *WordClient) Interceptors() []Interceptor {
	return c.inters.Word
}

func (c *WordClient) mutate(ctx context.Context, m *WordMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&WordCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&WordUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&WordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&WordDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Word mutation op: %q", m.Op())
	}
}

// WordConnectionsClient is a client for the WordConnections schema.
type WordConnectionsClient struct {
	config
}

// NewWordConnectionsClient returns a client for the WordConnections from the given config.
func NewWordConnectionsClient(c config) *WordConnectionsClient {
	return &WordConnectionsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `wordconnections.Hooks(f(g(h())))`.
func (c *WordConnectionsClient) Use(hooks ...Hook) {
	c.hooks.WordConnections = append(c.hooks.WordConnections, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `wordconnections.Intercept(f(g(h())))`.
func (c *WordConnectionsClient) Intercept(interceptors ...Interceptor) {
	c.inters.WordConnections = append(c.inters.WordConnections, interceptors...)
}

// Create returns a builder for creating a WordConnections entity.
func (c *WordConnectionsClient) Create() *WordConnectionsCreate {
	mutation := newWordConnectionsMutation(c.config, OpCreate)
	return &WordConnectionsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of WordConnections entities.
func (c *WordConnectionsClient) CreateBulk(builders ...*WordConnectionsCreate) *WordConnectionsCreateBulk {
	return &WordConnectionsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for WordConnections.
func (c *WordConnectionsClient) Update() *WordConnectionsUpdate {
	mutation := newWordConnectionsMutation(c.config, OpUpdate)
	return &WordConnectionsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WordConnectionsClient) UpdateOne(wc *WordConnections) *WordConnectionsUpdateOne {
	mutation := newWordConnectionsMutation(c.config, OpUpdateOne, withWordConnections(wc))
	return &WordConnectionsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WordConnectionsClient) UpdateOneID(id int) *WordConnectionsUpdateOne {
	mutation := newWordConnectionsMutation(c.config, OpUpdateOne, withWordConnectionsID(id))
	return &WordConnectionsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for WordConnections.
func (c *WordConnectionsClient) Delete() *WordConnectionsDelete {
	mutation := newWordConnectionsMutation(c.config, OpDelete)
	return &WordConnectionsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *WordConnectionsClient) DeleteOne(wc *WordConnections) *WordConnectionsDeleteOne {
	return c.DeleteOneID(wc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *WordConnectionsClient) DeleteOneID(id int) *WordConnectionsDeleteOne {
	builder := c.Delete().Where(wordconnections.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WordConnectionsDeleteOne{builder}
}

// Query returns a query builder for WordConnections.
func (c *WordConnectionsClient) Query() *WordConnectionsQuery {
	return &WordConnectionsQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeWordConnections},
		inters: c.Interceptors(),
	}
}

// Get returns a WordConnections entity by its id.
func (c *WordConnectionsClient) Get(ctx context.Context, id int) (*WordConnections, error) {
	return c.Query().Where(wordconnections.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WordConnectionsClient) GetX(ctx context.Context, id int) *WordConnections {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *WordConnectionsClient) Hooks() []Hook {
	return c.hooks.WordConnections
}

// Interceptors returns the client interceptors.
func (c *WordConnectionsClient) Interceptors() []Interceptor {
	return c.inters.WordConnections
}

func (c *WordConnectionsClient) mutate(ctx context.Context, m *WordConnectionsMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&WordConnectionsCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&WordConnectionsUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&WordConnectionsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&WordConnectionsDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown WordConnections mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Definition, Group, Word, WordConnections []ent.Hook
	}
	inters struct {
		Definition, Group, Word, WordConnections []ent.Interceptor
	}
)
