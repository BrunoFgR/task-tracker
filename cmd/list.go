package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := appCtx.Storage.List()
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
	rootCmd.AddCommand(listCmd)
}
