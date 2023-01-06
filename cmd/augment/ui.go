package augment

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Creates the UI template augmentation command.
func NewUiCmd() *cobra.Command {
	uiCmd := &cobra.Command{
		Use:     "ui",
		GroupID: "augment",
		Short:   "Augment an existing UI template.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ui called")
		},
	}
	return uiCmd
}
