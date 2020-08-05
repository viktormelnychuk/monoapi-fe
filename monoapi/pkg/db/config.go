package db

import (
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

type Postgres struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func (p *Postgres) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		p.User,
		url.QueryEscape(p.Password),
		p.Host,
		p.Port,
		p.Database,
	)
}

// Validate config.
func (p *Postgres) Validate() error {
	var errs []string

	if p.Database == "" {
		errs = append(errs, "Database shouldn't be empty")
	}

	if p.Host == "" {
		errs = append(errs, "Host shouldn't be empty")
	}

	if p.Port <= 0 {
		errs = append(errs, "Port shouldn't be empty")
	}

	if p.User == "" {
		errs = append(errs, "User shouldn't be empty")
	}

	if p.Password == "" {
		errs = append(errs, "Password shouldn't be empty")
	}

	if len(errs) > 0 {
		return errors.Errorf(strings.Join(errs, ","))
	}

	return nil
}

type PGConfig struct {
	Postgres Postgres
}

func (c *PGConfig) Validate() error {
	return c.Postgres.Validate()
}

type Config struct {
	DB PGConfig
}

func (c *Config) Validate() error {
	return c.DB.Validate()
}
