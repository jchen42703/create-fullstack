package context

import (
	"os"
	"path/filepath"
)

// Following the terraform scheme for specifying a plugins directory:
// Windows: %APPDATA%\.terraform.d\plugins -> %APPDATA%\.create-fullstack\plugins
// Mac OS/Linux: $HOME/.create-fullstack/plugins
func GetGlobalPluginsDir(runtimeOS string) string {
	pluginsDir := filepath.Join(".create-fullstack", "plugins")
	if runtimeOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), pluginsDir)
	}

	// Mac OS/Linux
	return filepath.Join(os.Getenv("HOME"), pluginsDir)
}
