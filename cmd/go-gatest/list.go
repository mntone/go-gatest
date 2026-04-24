package main

import (
	"time"

	"github.com/mntone/go-gatest/internal"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Show all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return withTaskRepository(cmd, func(repo internal.TaskRepository) error {
			tasks, err := repo.All(cmd.Context())
			if err != nil {
				return err
			}

			for _, task := range tasks {
				cmd.Printf(
					"%d: %s (%s)\n",
					task.ID,
					task.Description,
					task.UpdatedAt.Format(time.RFC3339Nano),
				)
			}

			return nil
		})
	},
}

func init() {
	rootCommand.AddCommand(listCommand)
}
