package main

import (
	"time"

	"github.com/mntone/go-gatest/internal"
	"github.com/spf13/cobra"
)

var searchCommand = &cobra.Command{
	Use:     "search K",
	Aliases: []string{"find"},
	Short:   "Search tasks by keyword",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return withTaskRepository(cmd, func(repo internal.TaskRepository) error {
			tasks, err := repo.Find(cmd.Context(), args[0])
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
	rootCommand.AddCommand(searchCommand)
}
