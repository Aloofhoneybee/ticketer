package cmd

import (
	"fmt"

	"ticketer/tickets"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Create a new open ticket",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		list, err := tickets.Load(dataFile)
		if err != nil {
			return err
		}

		list, id := tickets.Add(list, args[0])

		if err := tickets.Save(dataFile, list); err != nil {
			return err
		}
		fmt.Println("Created ticket with ID:", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
