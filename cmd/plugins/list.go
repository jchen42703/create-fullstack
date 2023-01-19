package plugins

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Creates the command to list all installed plugins.
func NewListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "list",
		GroupID: "plugins",
		Short:   "List all installed plugins.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("List all installed plugins")
		},
	}

	return listCmd
}
