package cmd

import (
	"fmt"
	"os"

	"github.com/BrunoFgR/task-tracker/internal/context"
	"github.com/BrunoFgR/task-tracker/internal/storage/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	appCtx      *context.AppContext
	storagePath string
	cfgFile     string
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A simple CLI task tracker",
	Long: `Task Tracker is a command line tool for managing your tasks.
You can add, list, and manage the status of your tasks through a simple interface.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Initialize storage once for all commands
		storagePath, _ := cmd.Flags().GetString("storage")
		storage, err := file.New(storagePath)
		if err != nil {
			return err
		}

		appCtx = context.New(storage)
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task-tracker.yaml)")
	rootCmd.PersistentFlags().StringP("storage", "s", "tasks.json", "Storage file for tasks")
	viper.BindPFlag("storage", rootCmd.PersistentFlags().Lookup("storage"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".task-tracker")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
