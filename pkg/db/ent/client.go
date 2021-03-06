// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/login-door/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/login-door/pkg/db/ent/loginrecord"
	"github.com/NpoolPlatform/login-door/pkg/db/ent/provider"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// LoginRecord is the client for interacting with the LoginRecord builders.
	LoginRecord *LoginRecordClient
	// Provider is the client for interacting with the Provider builders.
	Provider *ProviderClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.LoginRecord = NewLoginRecordClient(c.config)
	c.Provider = NewProviderClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		LoginRecord: NewLoginRecordClient(cfg),
		Provider:    NewProviderClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		config:      cfg,
		LoginRecord: NewLoginRecordClient(cfg),
		Provider:    NewProviderClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		LoginRecord.
//		Query().
//		Count(ctx)
//
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
	c.LoginRecord.Use(hooks...)
	c.Provider.Use(hooks...)
}

// LoginRecordClient is a client for the LoginRecord schema.
type LoginRecordClient struct {
	config
}

// NewLoginRecordClient returns a client for the LoginRecord from the given config.
func NewLoginRecordClient(c config) *LoginRecordClient {
	return &LoginRecordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `loginrecord.Hooks(f(g(h())))`.
func (c *LoginRecordClient) Use(hooks ...Hook) {
	c.hooks.LoginRecord = append(c.hooks.LoginRecord, hooks...)
}

// Create returns a create builder for LoginRecord.
func (c *LoginRecordClient) Create() *LoginRecordCreate {
	mutation := newLoginRecordMutation(c.config, OpCreate)
	return &LoginRecordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LoginRecord entities.
func (c *LoginRecordClient) CreateBulk(builders ...*LoginRecordCreate) *LoginRecordCreateBulk {
	return &LoginRecordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LoginRecord.
func (c *LoginRecordClient) Update() *LoginRecordUpdate {
	mutation := newLoginRecordMutation(c.config, OpUpdate)
	return &LoginRecordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LoginRecordClient) UpdateOne(lr *LoginRecord) *LoginRecordUpdateOne {
	mutation := newLoginRecordMutation(c.config, OpUpdateOne, withLoginRecord(lr))
	return &LoginRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LoginRecordClient) UpdateOneID(id uuid.UUID) *LoginRecordUpdateOne {
	mutation := newLoginRecordMutation(c.config, OpUpdateOne, withLoginRecordID(id))
	return &LoginRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LoginRecord.
func (c *LoginRecordClient) Delete() *LoginRecordDelete {
	mutation := newLoginRecordMutation(c.config, OpDelete)
	return &LoginRecordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *LoginRecordClient) DeleteOne(lr *LoginRecord) *LoginRecordDeleteOne {
	return c.DeleteOneID(lr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *LoginRecordClient) DeleteOneID(id uuid.UUID) *LoginRecordDeleteOne {
	builder := c.Delete().Where(loginrecord.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LoginRecordDeleteOne{builder}
}

// Query returns a query builder for LoginRecord.
func (c *LoginRecordClient) Query() *LoginRecordQuery {
	return &LoginRecordQuery{
		config: c.config,
	}
}

// Get returns a LoginRecord entity by its id.
func (c *LoginRecordClient) Get(ctx context.Context, id uuid.UUID) (*LoginRecord, error) {
	return c.Query().Where(loginrecord.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LoginRecordClient) GetX(ctx context.Context, id uuid.UUID) *LoginRecord {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LoginRecordClient) Hooks() []Hook {
	return c.hooks.LoginRecord
}

// ProviderClient is a client for the Provider schema.
type ProviderClient struct {
	config
}

// NewProviderClient returns a client for the Provider from the given config.
func NewProviderClient(c config) *ProviderClient {
	return &ProviderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `provider.Hooks(f(g(h())))`.
func (c *ProviderClient) Use(hooks ...Hook) {
	c.hooks.Provider = append(c.hooks.Provider, hooks...)
}

// Create returns a create builder for Provider.
func (c *ProviderClient) Create() *ProviderCreate {
	mutation := newProviderMutation(c.config, OpCreate)
	return &ProviderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Provider entities.
func (c *ProviderClient) CreateBulk(builders ...*ProviderCreate) *ProviderCreateBulk {
	return &ProviderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Provider.
func (c *ProviderClient) Update() *ProviderUpdate {
	mutation := newProviderMutation(c.config, OpUpdate)
	return &ProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProviderClient) UpdateOne(pr *Provider) *ProviderUpdateOne {
	mutation := newProviderMutation(c.config, OpUpdateOne, withProvider(pr))
	return &ProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProviderClient) UpdateOneID(id uuid.UUID) *ProviderUpdateOne {
	mutation := newProviderMutation(c.config, OpUpdateOne, withProviderID(id))
	return &ProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Provider.
func (c *ProviderClient) Delete() *ProviderDelete {
	mutation := newProviderMutation(c.config, OpDelete)
	return &ProviderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProviderClient) DeleteOne(pr *Provider) *ProviderDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProviderClient) DeleteOneID(id uuid.UUID) *ProviderDeleteOne {
	builder := c.Delete().Where(provider.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProviderDeleteOne{builder}
}

// Query returns a query builder for Provider.
func (c *ProviderClient) Query() *ProviderQuery {
	return &ProviderQuery{
		config: c.config,
	}
}

// Get returns a Provider entity by its id.
func (c *ProviderClient) Get(ctx context.Context, id uuid.UUID) (*Provider, error) {
	return c.Query().Where(provider.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProviderClient) GetX(ctx context.Context, id uuid.UUID) *Provider {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ProviderClient) Hooks() []Hook {
	return c.hooks.Provider
}
