package context

import (
	"os"
	"path/filepath"
	"strings"

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
}

// Executable is the path to the currently invoked binary.
// This is useful for logging, but is relatively niche.
func (c *CmdContext) Executable() string {
	if !strings.ContainsRune(c.ExecutableName, os.PathSeparator) {
		c.ExecutableName = executable(c.ExecutableName)
	}
	return c.ExecutableName
}

// Finds the location of the executable for the current process as it's found in PATH, respecting symlinks.
// If the process couldn't determine its location, return fallbackName. If the executable wasn't found in
// PATH, return the absolute location to the program.
//
// The idea is that the result of this function is callable in the future and refers to the same
// installation of gh, even across upgrades. This is needed primarily for Homebrew, which installs software
// under a location such as `/usr/local/Cellar/gh/1.13.1/bin/gh` and symlinks it from `/usr/local/bin/gh`.
// When the version is upgraded, Homebrew will often delete older versions, but keep the symlink. Because of
// this, we want to refer to the `gh` binary as `/usr/local/bin/gh` and not as its internal Homebrew
// location.
//
// None of this would be needed if we could just refer to GitHub CLI as `gh`, i.e. without using an absolute
// path. However, for some reason Homebrew does not include `/usr/local/bin` in PATH when it invokes git
// commands to update its taps. If `gh` (no path) is being used as git credential helper, as set up by `gh
// auth login`, running `brew update` will print out authentication errors as git is unable to locate
// Homebrew-installed `gh`.
func executable(fallbackName string) string {
	exe, err := os.Executable()
	if err != nil {
		return fallbackName
	}

	base := filepath.Base(exe)
	path := os.Getenv("PATH")
	for _, dir := range filepath.SplitList(path) {
		p, err := filepath.Abs(filepath.Join(dir, base))
		if err != nil {
			continue
		}
		f, err := os.Lstat(p)
		if err != nil {
			continue
		}

		if p == exe {
			return p
		} else if f.Mode()&os.ModeSymlink != 0 {
			if t, err := os.Readlink(p); err == nil && t == exe {
				return p
			}
		}
	}

	return exe
}
