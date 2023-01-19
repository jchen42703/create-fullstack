package context

import (
	"github.com/cli/cli/v2/pkg/iostreams"
	"go.uber.org/zap"
)

// Contains the shared parameters that all commands need access to.
type CmdContext struct {
	Logger *zap.Logger
	// TODO: Include Survey Prompter
	IoStreams *iostreams.IOStreams

	Version        string
	BuildDate      string
	ExecutableName string

	GlobalPluginsDir string // the parent dir storing all plugins directories
}
