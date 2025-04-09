package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete an existing task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if err := appCtx.Storage.DeleteByID(id); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
