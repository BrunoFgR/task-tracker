package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] [description]",
	Short: "Update an existing task",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		description := args[1]
		if err := appCtx.Storage.UpdateByID(id, description); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
