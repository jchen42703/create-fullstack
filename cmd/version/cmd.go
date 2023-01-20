package version

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jchen42703/create-fullstack/cmd/context"
	"github.com/jchen42703/create-fullstack/internal/getproviders"
	"github.com/spf13/cobra"
)

// Creates the version command.
// Based off: https://github.com/cli/cli/blob/trunk/pkg/cmd/version/version.go
func NewCmd(cmdCtx *context.CmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			cmdCtx.CliUi.Log(cmd.Root().Annotations["versionInfo"])
		},
	}

	return cmd
}

// Builds the version string to display.
func Format(version, buildDate string) string {
	version = strings.TrimPrefix(version, "v")

	var dateStr string
	if buildDate != "" {
		dateStr = fmt.Sprintf(" (%s)", buildDate)
	}

	platform := getproviders.CurrentPlatform
	return fmt.Sprintf("create-fullstack version %s%s\n%s\n%s\n", version, dateStr, changelogURL(version), platform)
}

// Creates the release url.
func changelogURL(version string) string {
	path := "https://github.com/jchen42703/create-fullstack"
	r := regexp.MustCompile(`^v?\d+\.\d+\.\d+(-[\w.]+)?$`)
	if !r.MatchString(version) {
		return fmt.Sprintf("%s/releases/latest", path)
	}

	url := fmt.Sprintf("%s/releases/tag/v%s", path, strings.TrimPrefix(version, "v"))
	return url
}
