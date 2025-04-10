package cmd

import (
	"fmt"

	"github.com/BrunoFgR/task-tracker/internal/models"
	"github.com/spf13/cobra"
)

var listInProgressCmd = &cobra.Command{
	Use:   "in-progress",
	Short: "List tasks that status is in progress",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := appCtx.Storage.ListByStatus(models.StatusInProgress)
		if err != nil {
			return err
		}
		for _, task := range tasks {
			fmt.Printf("%d. %s | %s\n", task.ID, task.Description, task.Status)
		}

		return nil
	},
}

func init() {
	listCmd.AddCommand(listInProgressCmd)
}
