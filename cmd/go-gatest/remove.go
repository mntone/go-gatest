package main

import (
	"fmt"
	"strconv"

	"github.com/mntone/go-gatest/internal"
	"github.com/spf13/cobra"
)

var removeCommand = &cobra.Command{
	Use:   "remove <id>",
	Short: "Remove a task by task ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return withTaskRepository(cmd, func(repo internal.TaskRepository) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			deleted, err := repo.RemoveByID(cmd.Context(), int64(id))
			if err != nil {
				return err
			}

			if !deleted {
				return fmt.Errorf("failed to remove task: id=%d", id)
			}

			cmd.Printf("remove task: id=%d\n", id)
			return nil
		})
	},
}

func init() {
	rootCommand.AddCommand(removeCommand)
}
