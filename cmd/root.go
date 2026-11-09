package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var dataFile string

func dataPath() (string, error) {
	base := os.Getenv("XDG_DATA_HOME")
	if base == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		base = filepath.Join(home, ".local", "share")
	}
	return filepath.Join(base, "ticketer", "tickets.json"), nil
}

var rootCmd = &cobra.Command{
	Use:          "ticketer",
	Short:        "A tiny helpdesk CLI",
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		path, err := dataPath()
		if err != nil {
			return err
		}
		dataFile = path
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
