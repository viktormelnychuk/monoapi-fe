package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"
)

type options struct {
	maxOpenConns int
}
type Option interface{ apply(*options) }
type optionFunc func(*options)

func (f optionFunc) apply(o *options) { f(o) }

func New(config *Config, opts ...Option) (*sql.DB, error) {
	// default options
	o := options{maxOpenConns: 10}
	driverName := "pgx"

	for _, op := range opts {
		op.apply(&o)
	}

	db, err := sql.Open(driverName, config.DB.Postgres.URL())
	if err != nil {
		return nil, errors.Wrap(err, "open connection")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "ping")
	}

	db.SetMaxOpenConns(o.maxOpenConns)

	return db, nil
}
