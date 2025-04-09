package cmd

import (
	"time"

	"github.com/BrunoFgR/task-tracker/internal/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		description := args[0]

		task := models.Task{
			Description: description,
			Status:      models.StatusTodo,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := appCtx.Storage.Create(task); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
