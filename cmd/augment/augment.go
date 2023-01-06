package augment

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Creates the `augment` command with the appropriate subcommands.
func NewCmd() *cobra.Command {
	augCmd := &cobra.Command{
		Use:     "augment",
		GroupID: "core",
		Short:   "Augment existing project templates",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Error: must also specify a resource like api, ui, or fullstack.")
		},
	}
	augCmd.AddGroup(&cobra.Group{
		ID:    "augment",
		Title: "Augmentation Commands",
	})

	augCmd.AddCommand(NewApiCmd(), NewUiCmd(), NewFullstackCmd())

	return augCmd
}
