package main

import (
	"fmt"
	"os"

	"github.com/mntone/go-gatest/internal"
	"github.com/mntone/go-gatest/internal/migration"
	"github.com/mntone/go-gatest/internal/sqlite"
	"github.com/spf13/cobra"
)

var version = "0.0.0+dev"

func withTaskRepository(
	cmd *cobra.Command,
	callback func(repo internal.TaskRepository) error,
) error {
	db, err := internal.OpenDatabase(cmd.Context(), "file:go-gatest.sqlite")
	if err != nil {
		return err
	}

	err = migration.MigrateUp(db)
	if err != nil {
		return err
	}

	return callback(sqlite.NewTaskRepository(db))
}

func main() {
	fmt.Printf("Go GitHub Actions Test version: %s\n", version)

	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
