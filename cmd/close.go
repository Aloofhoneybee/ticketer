package cmd

import (
	"fmt"
	"strconv"

	"ticketer/tickets"

	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:   "close [id]",
	Short: "Mark a ticket as closed",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("ticket id must be a number, got %q", args[0])
		}

		list, err := tickets.Load(dataFile)
		if err != nil {
			return err
		}

		if err := tickets.CloseTicket(list, id); err != nil {
			return err
		}

		if err := tickets.Save(dataFile, list); err != nil {
			return err
		}
		fmt.Printf("Closed ticket %d\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
