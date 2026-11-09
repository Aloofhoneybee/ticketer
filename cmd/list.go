package cmd

import (
	"fmt"

	"ticketer/tickets"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show open tickets (use --all for every ticket)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			return err
		}

		list, err := tickets.Load(dataFile)
		if err != nil {
			return err
		}

		shown := tickets.Filter(list, all)
		if len(shown) == 0 {
			fmt.Println("No tickets.")
			return nil
		}
		for _, t := range shown {
			fmt.Println(t)
		}
		return nil
	},
}

func init() {
	listCmd.Flags().BoolP("all", "a", false, "show closed tickets too")
	rootCmd.AddCommand(listCmd)
}
