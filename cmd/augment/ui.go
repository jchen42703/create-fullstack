package augment

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewUiCmd() *cobra.Command {
	// uiCmd represents the ui command
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
