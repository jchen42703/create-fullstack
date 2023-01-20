package context

import (
	"github.com/jchen42703/create-fullstack/cmd/cliui"
	"go.uber.org/zap"
)

// Contains the shared parameters that all commands need access to.
type CmdContext struct {
	Logger *zap.Logger
	CliUi  *cliui.ColorUi

	Version        string
	BuildDate      string
	ExecutableName string

	GlobalPluginsDir string // the parent dir storing all plugins directories
}
