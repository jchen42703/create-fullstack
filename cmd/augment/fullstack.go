package augment

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewFullstackCmd() *cobra.Command {
	// apiCmd represents the api command
	fullstackCmd := &cobra.Command{
		Use:     "fullstack",
		GroupID: "augment",
		Short:   "Augment an existing fullstack template.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("fullstack called")
		},
	}

	return fullstackCmd
}
