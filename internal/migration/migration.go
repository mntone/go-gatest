package migration

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	driver "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed *.sql
var fs embed.FS

func MigrateUp(db *sql.DB) (err error) {
	source, err := iofs.New(fs, ".")
	if err != nil {
		return fmt.Errorf("create migration source: %w", err)
	}
	defer func() {
		sourceCloseErr := source.Close()
		if sourceCloseErr != nil {
			err = errors.Join(
				err,
				fmt.Errorf("close migration source: %w", sourceCloseErr),
			)
		}
	}()

	database, err := driver.WithInstance(db, &driver.Config{})
	if err != nil {
		err = fmt.Errorf("create migration driver: %w", err)
		return
	}

	migrateInst, err := migrate.NewWithInstance("iofs", source, "sqlite3", database)
	if err != nil {
		err = fmt.Errorf("create migration instance: %w", err)
	}

	err = migrateInst.Up()
	if err != nil && err != migrate.ErrNoChange {
		err = fmt.Errorf("migration up: %w", err)
	}

	err = nil
	return
}
