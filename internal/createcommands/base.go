package createcommands

import (
	"flag"
	"io"

	"github.com/mitchellh/cli"
)

// Base Command that will be embedded in all commands
//
// It is largely inspired by the `Meta` struct from Terraform:
// https://github.com/hashicorp/terraform/blob/6af6540233404ba8ca5358a4bb4b6d41d54710e0/internal/command/meta.go
type BaseCommand struct {
	Ui cli.Ui // Ui for output
}

// defaultFlagSet creates a default flag set for commands.
// See also command/arguments/default.go
func (m *BaseCommand) defaultFlagSet(n string) *flag.FlagSet {
	f := flag.NewFlagSet(n, flag.ContinueOnError)
	f.SetOutput(io.Discard)

	// Set the default Usage to empty
	f.Usage = func() {}

	return f
}
