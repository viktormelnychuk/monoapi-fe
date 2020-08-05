package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"github.com/viktormelnychuk/monoapi/monoapi/pkg/cfg"
	"github.com/viktormelnychuk/monoapi/monoapi/pkg/db"
	"os"
	"time"
)

var (
	cfgFile          string
	migrationsFolder string
	dbConfig         db.Config
)

var migrateCommand = &cobra.Command{
	Use:          "migrate",
	Short:        "migrate database",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := log.NewLogfmtLogger(os.Stderr)

		err := cfg.Init("config", cfgFile, &dbConfig)
		if err != nil {
			logger.Log("err:", err)
			return errors.Wrap(err, "load db config")
		}
		pg, err := db.New(&dbConfig)
		if err != nil {
			return errors.Wrap(err, "failed to init postgres client")
		}

		logger.Log("Starting migrations")

		errs := make(chan error, 1)

		go func() {
			defer close(errs)

			migrations := &migrate.FileMigrationSource{Dir: migrationsFolder}
			n, err := migrate.Exec(pg, "postgres", migrations, migrate.Up)
			if err != nil {
				errs <- err
				return
			}
			logger.Log(fmt.Sprintf("Applied migrations %d", n))
		}()

		select {
		case err := <-errs:
			if err != nil {
				return errors.Wrap(err, "failed to run migrations")
			}
		case <-time.After(10 * time.Minute):
			return errors.New("failed to run migrations, timeout after 10 minutes")
		}
		return nil
	},
}

func init() {
	migrateCommand.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")
	migrateCommand.PersistentFlags().StringVar(&migrationsFolder, "migrations-dir", "migrations", "directory with migrations files")
}

func main() {
	if err := migrateCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
