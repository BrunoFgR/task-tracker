package cmd

import (
	"strconv"

	"github.com/BrunoFgR/task-tracker/internal/models"
	"github.com/spf13/cobra"
)

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress [id]",
	Short: "Mark a task as in progress",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if err := appCtx.Storage.UpdateStatusByID(id, models.StatusInProgress); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)
}
