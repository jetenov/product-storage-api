package db

import (
	"context"
	"errors"
	"time"

	"gitlab.dg.ru/platform/database-go/entrypoint"
	"gitlab.dg.ru/platform/database-go/sql"
	"gitlab.dg.ru/platform/database-go/sql/balancer/role"
	"gitlab.dg.ru/platform/scratch/closer"
)

const (
	defaultMaxIdleCon = 10
	defaultMaxOpenCon = 20
)

// PgOption an additional parameter for `connectOptions` constructor
type (
	PgOption func(*connectOptions)

	connectOptions struct {
		maxOpenCon int
		maxIdleCon int
	}
)

// NewDB creates database connection with tracing and metrics
func NewDB(ctx context.Context, dsn string, cl *closer.Closer, hc healthcheck.Handler, opts ...PgOption) (sql.Balancer, error) {
	if len(dsn) == 0 {
		return nil, errors.New("empty dsn")
	}

	c := &connectOptions{
		maxIdleCon: defaultMaxIdleCon,
		maxOpenCon: defaultMaxOpenCon,
	}
	for _, opt := range opts {
		opt(c)
	}

	b, err := entrypoint.New(
		dsn,
		entrypoint.WithMaxOpenConnections(c.maxOpenCon),
		entrypoint.WithMaxIdleConnections(c.maxIdleCon),
		entrypoint.WithOIDCache(),
		entrypoint.WithMasterMappingRoles(role.Read, role.Write),
		entrypoint.WithInitTimeout(time.Second),
	)
	if err != nil {
		return nil, err
	}

	if err := b.Ping(ctx); err != nil {
		return nil, err
	}

	if hc != nil {
		hc.AddReadinessCheck(b.Healthcheck())
	}

	if cl == nil {
		closer.Add(b.Close)
	} else {
		cl.Add(b.Close)
	}

	return b, nil
}

// DB ...
type DB struct {
	Instance sql.Balancer
	Name     string
}

// WithMaxIdleConns sets the maximum number of connections in the idle connection pool.
func WithMaxIdleConns(n int) PgOption {
	return func(b *connectOptions) {
		maxIdleCon := n
		if n <= 0 {
			maxIdleCon = defaultMaxIdleCon
		}
		b.maxIdleCon = maxIdleCon
	}
}

// WithMaxOpenConns sets the maximum number of open connections to the database.
func WithMaxOpenConns(n int) PgOption {
	return func(b *connectOptions) {
		maxOpenCon := n
		if n <= 0 {
			maxOpenCon = defaultMaxOpenCon
		}
		b.maxOpenCon = maxOpenCon
	}
}
