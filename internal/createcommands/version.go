package createcommands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jchen42703/create-fullstack/internal/getproviders"
)

// VersionCommand is a Command implementation prints the version.
type VersionCommand struct {
	BaseCommand

	Version           string
	VersionPrerelease string
	CheckFunc         VersionCheckFunc
	Platform          getproviders.Platform
}

// VersionCheckInfo is the return value for the VersionCheckFunc callback
// and tells the Version command information about the latest version
// of create-fullstack.
type VersionCheckInfo struct {
	Outdated bool
	Latest   string
	Alerts   []string
}

type VersionOutput struct {
	Version  string `json:"create_fullstack_version"`
	Platform string `json:"platform"`
	// ProviderSelections map[string]string `json:"provider_selections"`
	Outdated bool `json:"create_fullstack_outdated"`
}

// VersionCheckFunc is the callback called by the Version command to
// check if there is a new version of create-fullstack.
type VersionCheckFunc func() (VersionCheckInfo, error)

func (c *VersionCommand) Help() string {
	helpText := `
Usage: create-fullstack [global options] version [options]
  Displays the version of create-fullstack and all installed plugins.
Options:
  -json       Output the version information as a JSON object.
`
	return strings.TrimSpace(helpText)
}

// TODO: Print out plugin versions once the plugin architecture is configured.
func (c *VersionCommand) Run(args []string) int {
	var outdated bool
	var latest string
	var versionString bytes.Buffer
	var jsonOutput bool
	cmdFlags := c.BaseCommand.defaultFlagSet("version")
	cmdFlags.BoolVar(&jsonOutput, "json", false, "json")
	// Enable but ignore the global version flags. In main.go, if any of the
	// arguments are -v, -version, or --version, this command will be called
	// with the rest of the arguments, so we need to be able to cope with
	// those.
	cmdFlags.Bool("v", true, "version")
	cmdFlags.Bool("version", true, "version")
	cmdFlags.Usage = func() { c.Ui.Error(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		c.Ui.Error(fmt.Sprintf("Error parsing command-line flags: %s\n", err.Error()))
		return 1
	}

	fmt.Fprintf(&versionString, "create-fullstack v%s", c.Version)
	if c.VersionPrerelease != "" {
		fmt.Fprintf(&versionString, "-%s", c.VersionPrerelease)
	}

	// If we have a version check function, then let's check for
	// the latest version as well.
	if c.CheckFunc != nil {
		// Check the latest version
		info, err := c.CheckFunc()
		if err != nil && !jsonOutput {
			c.Ui.Error(fmt.Sprintf(
				"\nError checking latest version: %s", err))
		}
		if info.Outdated {
			outdated = true
			latest = info.Latest
		}
	}

	if jsonOutput {
		// selectionsOutput := make(map[string]string)
		// for providerAddr, lock := range providerLocks {
		// 	version := lock.Version().String()
		// 	selectionsOutput[providerAddr.String()] = version
		// }

		var versionOutput string
		if c.VersionPrerelease != "" {
			versionOutput = c.Version + "-" + c.VersionPrerelease
		} else {
			versionOutput = c.Version
		}

		output := VersionOutput{
			Version:  versionOutput,
			Platform: c.Platform.String(),
			// ProviderSelections: selectionsOutput,
			Outdated: outdated,
		}

		jsonOutput, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			c.Ui.Error(fmt.Sprintf("\nError marshalling JSON: %s", err))
			return 1
		}
		c.Ui.Output(string(jsonOutput))
		return 0
	} else {
		c.Ui.Output(versionString.String())
		c.Ui.Output(fmt.Sprintf("on %s", c.Platform))

		// if len(providerVersions) != 0 {
		// 	sort.Strings(providerVersions)
		// 	for _, str := range providerVersions {
		// 		c.Ui.Output(str)
		// 	}
		// }
		if outdated {
			c.Ui.Output(fmt.Sprintf(
				"\nYour version of Terraform is out of date! The latest version\n"+
					"is %s. You can update by downloading from https://www.terraform.io/downloads.html",
				latest))
		}

	}

	return 0
}

func (c *VersionCommand) Synopsis() string {
	return "Show the current create-fullstack version"
}
