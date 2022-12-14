package augment

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Creates the API template augmentation command.
func NewApiCmd() *cobra.Command {
	apiCmd := &cobra.Command{
		Use:     "api",
		GroupID: "augment",
		Short:   "Augment an existing API template.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("api called")
		},
	}

	return apiCmd
}
