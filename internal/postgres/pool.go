package postgres

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// background context
var bgCtx = context.Background()

// ParseConfig handler
var ParseConfig = pgxpool.ParseConfig

// Logger for pgx
type Logger = pgx.Logger

// LogLevel for pgx
type LogLevel = pgx.LogLevel

// Config struct
type Config struct {
	*pgxpool.Config
	ctx context.Context
}

// Dial fn
func Dial(url string, withConfig ...func(*Config)) (*Pool, error) {
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}
	config := &Config{cfg, bgCtx}
	for _, fn := range withConfig {
		fn(config)
	}
	pool, err := pgxpool.ConnectConfig(config.ctx, config.Config)
	if err != nil {
		return nil, err
	}
	return &Pool{
		pool,
		bgCtx,
	}, nil
}

// WithLogger sets the log level
func WithLogger(logger Logger, level LogLevel) func(*Config) {
	return func(c *Config) {
		c.ConnConfig.LogLevel = level
		c.ConnConfig.Logger = logger
	}
}

// WithContext adds a context to the pool
func WithContext(ctx context.Context) func(*Config) {
	return func(c *Config) {
		c.ctx = ctx
	}
}

// Pool struct
type Pool struct {
	p   *pgxpool.Pool
	ctx context.Context
}

var _ DB = (*Pool)(nil)

// DB interface
type DB interface {
	Query(sql string, args ...interface{}) (pgx.Rows, error)
}

// WithContext fn create a new pool with a custom context
func (p *Pool) WithContext(ctx context.Context) *Pool {
	return &Pool{p.p, ctx}
}

// Acquire a co
// Acquire fnnnection in the pool
func (p *Pool) Acquire() (*Conn, error) {
	conn, err := p.p.Acquire(p.ctx)
	if err != nil {
		return nil, err
	}
	return &Conn{conn, p.ctx}, nil
}

// AcquireAllIdle atomically acquires all currently idle connections. Its intended use is for health check and
// AcquireAllIdle fnunctionality. It does not update pool statistics.
func (p *Pool) AcquireAllIdle() (conns []*Conn) {
	idles := p.p.AcquireAllIdle(p.ctx)
	for _, idle := range idles {
		conns = append(conns, &Conn{idle, p.ctx})
	}
	return conns
}

// Begin fn
func (p *Pool) Begin() (*Tx, error) {
	tx, err := p.p.Begin(p.ctx)
	if err != nil {
		return nil, err
	}
	return &Tx{tx, p.ctx}, nil
}

// BeginTx fn
func (p *Pool) BeginTx(txOptions pgx.TxOptions) (*Tx, error) {
	tx, err := p.p.BeginTx(p.ctx, txOptions)
	if err != nil {
		return nil, err
	}
	return &Tx{tx, p.ctx}, nil
}

// Close fn
func (p *Pool) Close() {
	p.p.Close()
}

// CopyFrom fn
func (p *Pool) CopyFrom(tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return p.p.CopyFrom(p.ctx, tableName, columnNames, rowSrc)
}

// Exec fn
func (p *Pool) Exec(sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return p.p.Exec(p.ctx, sql, arguments...)
}

// Query fn
func (p *Pool) Query(query string, v ...interface{}) (pgx.Rows, error) {
	return p.p.Query(p.ctx, query, v...)
}

// QueryRow fn
func (p *Pool) QueryRow(sql string, args ...interface{}) pgx.Row {
	return p.p.QueryRow(p.ctx, sql, args...)
}

// SendBatch fn
func (p *Pool) SendBatch(b *pgx.Batch) pgx.BatchResults {
	return p.p.SendBatch(p.ctx, b)
}

// Stat fn
func (p *Pool) Stat() *pgxpool.Stat {
	return p.p.Stat()
}

// Conn struct
type Conn struct {
	conn *pgxpool.Conn
	ctx  context.Context
}

// WithContext fn
func (c *Conn) WithContext(ctx context.Context) *Conn {
	return &Conn{c.conn, ctx}
}

// Begin fn
func (c *Conn) Begin() (*Tx, error) {
	tx, err := c.conn.Begin(c.ctx)
	if err != nil {
		return nil, err
	}
	return &Tx{tx, c.ctx}, nil
}

// BeginTx fn
func (c *Conn) BeginTx(txOptions pgx.TxOptions) (pgx.Tx, error) {
	return c.conn.BeginTx(c.ctx, txOptions)
}

// Conn is a PostgreSQL connection handle. It is not safe for concurrent usage.
// Use a connection pool to manage access to multiple database connections
// from multipl
// Conn fne goroutines.
func (c *Conn) Conn() *pgx.Conn {
	return c.conn.Conn()
}

// CopyFrom fn
func (c *Conn) CopyFrom(tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return c.conn.CopyFrom(c.ctx, tableName, columnNames, rowSrc)
}

// Exec fn
func (c *Conn) Exec(sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return c.conn.Exec(c.ctx, sql, arguments...)
}

// Query fn
func (c *Conn) Query(sql string, args ...interface{}) (pgx.Rows, error) {
	return c.conn.Query(c.ctx, sql, args...)
}

// QueryRow fn
func (c *Conn) QueryRow(sql string, args ...interface{}) pgx.Row {
	return c.conn.QueryRow(c.ctx, sql, args...)
}

// Release fn
func (c *Conn) Release() {
	c.conn.Release()
}

// SendBatch fn
func (c *Conn) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return c.conn.SendBatch(ctx, b)
}

// Tx struct
type Tx struct {
	tx  pgx.Tx
	ctx context.Context
}

var _ DB = (*Tx)(nil)

// Begin star
// Begin fnts a pseudo nested transaction.
func (t *Tx) Begin() (*Tx, error) {
	tx, err := t.tx.Begin(t.ctx)
	if err != nil {
		return nil, err
	}
	return &Tx{tx, t.ctx}, nil
}

// Commit commits the transaction if this is a real transaction or releases the savepoint if this is a pseudo nested
// transaction. Commit will return ErrTxClosed if the Tx is already closed, but is otherwise safe to call multiple
// times. If the commit fails with a rollback status (e.g. a deferred constraint was violated) then
// ErrTxCommitRollback will be returned. Any other failure of a real transaction will result in the connection being
// closed.
// Commit fn
func (t *Tx) Commit() error {
	return t.tx.Commit(t.ctx)
}

// Rollback rolls back the transaction if this is a real transaction or rolls back to the savepoint if this is a
// pseudo nested transaction. Rollback will return ErrTxClosed if the Tx is already closed, but is otherwise safe to
// call multiple times. Hence, a defer tx.Rollback() is safe even if tx.Commit() will be called first in a non-error
// condition.
// Rollback fn Any other failure of a real transaction will result in the connection being closed.
func (t *Tx) Rollback() error {
	return t.tx.Rollback(t.ctx)
}

// CopyFrom fn
func (t *Tx) CopyFrom(tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return t.tx.CopyFrom(t.ctx, tableName, columnNames, rowSrc)
}

// SendBatch fn
func (t *Tx) SendBatch(b *pgx.Batch) pgx.BatchResults {
	return t.tx.SendBatch(t.ctx, b)
}

// LargeObjects fn
func (t *Tx) LargeObjects() pgx.LargeObjects {
	return t.tx.LargeObjects()
}

// Prepare fn
func (t *Tx) Prepare(name, sql string) (*pgconn.StatementDescription, error) {
	return t.tx.Prepare(t.ctx, name, sql)
}

// Exec fn
func (t *Tx) Exec(sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error) {
	return t.tx.Exec(t.ctx, sql, arguments...)
}

// Query fn
func (t *Tx) Query(sql string, args ...interface{}) (pgx.Rows, error) {
	return t.tx.Query(t.ctx, sql, args...)
}

// QueryRow fn
func (t *Tx) QueryRow(sql string, args ...interface{}) pgx.Row {
	return t.tx.QueryRow(t.ctx, sql, args...)
}

// // Conn retur
// // Conn fnns the underlying *Conn that on which this transaction is executing.
// func (t *Tx) Conn() *pgx.Conn {
// 	return t.tx.Conn()
// }
