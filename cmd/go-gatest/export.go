package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/mntone/go-gatest/internal"
	"github.com/spf13/cobra"
)

type exportTask struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var exportCommand = &cobra.Command{
	Use:   "export <path>",
	Short: "Export tasks to JSON file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return withTaskRepository(cmd, func(repo internal.TaskRepository) error {
			tasks, err := repo.All(cmd.Context())
			if err != nil {
				return err
			}

			outputTasks := make([]exportTask, 0, len(tasks))
			for _, task := range tasks {
				outputTasks = append(outputTasks, exportTask{
					ID:          task.ID,
					Description: task.Description,
					CreatedAt:   task.CreatedAt,
					UpdatedAt:   task.UpdatedAt,
				})
			}

			file, err := os.Create(args[0])
			if err != nil {
				return err
			}
			defer func() {
				_ = file.Close()
			}()

			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "  ")
			if err := encoder.Encode(outputTasks); err != nil {
				return err
			}

			cmd.Printf("exported %d task(s) to %s\n", len(outputTasks), args[0])
			return nil
		})
	},
}

func init() {
	rootCommand.AddCommand(exportCommand)
}
