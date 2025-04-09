package cmd

import (
	"fmt"

	"github.com/BrunoFgR/task-tracker/internal/models"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks or list by status",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			tasks, err := appCtx.Storage.List()
			if err != nil {
				return err
			}
			for _, task := range tasks {
				fmt.Printf("%d. %s | %s\n", task.ID, task.Description, task.Status)
			}
		} else {
			status := args[0]
			switch status {
			case "todo":
				tasks, err := appCtx.Storage.ListByStatus(models.StatusTodo)
				if err != nil {
					return err
				}
				for _, task := range tasks {
					fmt.Printf("%d. %s | %s\n", task.ID, task.Description, task.Status)
				}
			case "in-progress":
				tasks, err := appCtx.Storage.ListByStatus(models.StatusInProgress)
				if err != nil {
					return err
				}
				for _, task := range tasks {
					fmt.Printf("%d. %s | %s\n", task.ID, task.Description, task.Status)
				}
			case "done":
				tasks, err := appCtx.Storage.ListByStatus(models.StatusDone)
				if err != nil {
					return err
				}
				for _, task := range tasks {
					fmt.Printf("%d. %s | %s\n", task.ID, task.Description, task.Status)
				}
			default:
				return fmt.Errorf("invalid status: %s", status)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
