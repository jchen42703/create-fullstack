package plugin

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	cp "github.com/otiai10/copy"
	"go.uber.org/zap"

	"github.com/jchen42703/create-fullstack/core/run"
)

// Abstracting the plugin installation process.
// Plugins are supplied with the following file structure:
//
//	 pluginDir
//		main.go
//		augmentor_metadata.json
//
// pluginDirs is a collection of paths to directories like the one above.
//
// The installer will install all of those plugins in `parentOutputPluginDir`, which has a similar
// directory structure:
//
// parentOutputPluginDir
//
//		 plugin1
//			plugin1.exe
//			augmentor_metadata.json
//	  	 plugin2
//			...
type AugmentorPluginInstaller struct {
	ParentOutputPluginDir string

	// Separating because the logger != writer
	// i.e. when running plugins, the logger will not be a zap logger.
	// Uses hclog
	Writer io.Writer
	Logger *zap.Logger
}

// Gets all installed plugins. It assumes that all plugins with a valid PluginMeta json file.
func (i *AugmentorPluginInstaller) GetAllPlugins() ([]*PluginMeta, error) {
	dirs, err := os.ReadDir(i.ParentOutputPluginDir)
	if err != nil {
		return nil, err
	}

	allMeta := []*PluginMeta{}
	for _, pluginDir := range dirs {
		// Ignore not directories
		if !pluginDir.IsDir() {
			continue
		}

		// Check if it has a metadata JSON
		f, err := os.ReadFile(filepath.Join(i.ParentOutputPluginDir, pluginDir.Name(), "augmentor_metadata.json"))
		if err != nil {
			// Log this
			i.Logger.Debug(fmt.Sprintf("SKIPPING plugin %s: is missing augmentor_metadata.json", pluginDir.Name()))
			continue
		}

		var metadata PluginMeta
		err = json.Unmarshal(f, &metadata)
		if err != nil {
			i.Logger.Debug(fmt.Sprintf("failed to parse augmentor_metadata.json for plugin %s", pluginDir.Name()))
			continue
		}

		// Is valid, so add to slice
		allMeta = append(allMeta, &metadata)
	}

	return allMeta, nil
}

// This function installs a single Go RPC Plugin.
// CFS expects the pluginDir to lead to a directory that looks like:
//
//		...
//	 main.go
//	 augmentor_metadata.json
//
// It also expects `outputPluginDir` to be a child directory of `i.parentOutputPluginDir`.
//
// To install the plugin, this function:
// 1. Builds the plugin.
// 2. Validates the metadata.
// 3. Moves the built executable and metadata into the output plugin directory.
func (i *AugmentorPluginInstaller) Install(pluginDir string, outputPluginDir string) (*PluginMeta, error) {
	// Reads the plugin metadata
	augMetadataName := "augmentor_metadata.json"
	metadataPath := filepath.Join(pluginDir, augMetadataName)
	f, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata: %s", err)
	}

	var metadata PluginMeta
	err = json.Unmarshal(f, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to parse metadata: %s", err)
	}

	// Build plugin executable to the output directory.
	outputExecPath := filepath.Join(outputPluginDir, metadata.Id)
	cmd := exec.Command("go", "build", "-o", outputExecPath, pluginDir)
	err = run.Cmd(cmd, i.Writer)

	if err != nil {
		return nil, fmt.Errorf("plugin build failed: %s", err)
	}

	err = cp.Copy(metadataPath, filepath.Join(outputPluginDir, augMetadataName))
	if err != nil {
		return nil, fmt.Errorf("failed to copy metadata: %s", err)
	}

	return &metadata, nil
}
