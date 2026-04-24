package main

import (
	"fmt"
	"time"

	"github.com/mntone/go-gatest/internal"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add D",
	Short: "Add a task with Description",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return withTaskRepository(cmd, func(repo internal.TaskRepository) error {
			desc := args[0]
			created, err := repo.Add(cmd.Context(), desc, time.Now())
			if err != nil {
				return err
			}

			if !created {
				return fmt.Errorf("failed to create task: %s", desc)
			}

			cmd.Printf("added a task: %s\n", desc)
			return nil
		})
	},
}

func init() {
	rootCommand.AddCommand(addCommand)
}
