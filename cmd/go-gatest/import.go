package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mntone/go-gatest/internal"
	"github.com/spf13/cobra"
)

type importTask struct {
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

var importCommand = &cobra.Command{
	Use:   "import <path>",
	Short: "Import tasks from JSON file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return withTaskRepository(cmd, func(repo internal.TaskRepository) error {
			data, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			var inputTasks []importTask
			if err := json.Unmarshal(data, &inputTasks); err != nil {
				return err
			}

			now := time.Now()
			importedCount := 0
			for i, task := range inputTasks {
				if task.Description == "" {
					return fmt.Errorf("description is empty at index %d", i)
				}

				createdAt := task.CreatedAt
				if createdAt.IsZero() {
					createdAt = now
				}

				created, err := repo.Add(cmd.Context(), task.Description, createdAt)
				if err != nil {
					return err
				}
				if !created {
					return fmt.Errorf("failed to import task at index %d", i)
				}

				importedCount++
			}

			cmd.Printf("imported %d task(s) from %s\n", importedCount, args[0])
			return nil
		})
	},
}

func init() {
	rootCommand.AddCommand(importCommand)
}
