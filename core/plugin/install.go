package plugin

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-hclog"
	cp "github.com/otiai10/copy"
	"go.uber.org/zap"

	"github.com/jchen42703/create-fullstack/core/run"
	"github.com/jchen42703/create-fullstack/internal/directory"
	"github.com/jchen42703/create-fullstack/internal/log"
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

func NewAugmentorPluginInstaller(parentPluginDir string, pluginLogger hclog.Logger) (*AugmentorPluginInstaller, error) {
	err := os.MkdirAll(parentPluginDir, directory.READ_WRITE_EXEC_PERM)
	if err != nil {
		return nil, fmt.Errorf("failed to create parent plugin dir %s: %s", parentPluginDir, err)
	}

	// Install plugin prior to running
	logFilepath := filepath.Join(parentPluginDir, "create-fullstack-plugin-installer.log")
	installerLogger, err := log.CreateLogger(logFilepath)
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %s", err)
	}

	installer := &AugmentorPluginInstaller{
		ParentOutputPluginDir: parentPluginDir,
		Writer:                pluginLogger.StandardWriter(&hclog.StandardLoggerOptions{}),
		Logger:                installerLogger,
	}

	return installer, nil
}

// Reads the plugin metadata.
func (i *AugmentorPluginInstaller) ReadMetadata(pluginDir string) (*PluginMeta, error) {
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

	return &metadata, nil
}

func (i *AugmentorPluginInstaller) GetPlugin(pluginUri string) error {
	// Is url?
	isUrl := false
	if isUrl {
		// Download it into i.ParentOutputPluginDir
		return nil
	}

	isDir, err := directory.Exists(pluginUri)
	if err != nil {
		return fmt.Errorf("failed to determine if %s exists: %s", pluginUri, err)
	}

	if isDir {
		metadata, err := i.ReadMetadata(pluginUri)
		if err != nil {
			return fmt.Errorf("failed to read metadata: %s", err)
		}

		// Doesn't copy symlinks (bugged)
		opts := cp.Options{
			OnSymlink: func(src string) cp.SymlinkAction {
				return cp.Skip
			},
		}
		// Is directory?
		// Copy it into i.ParentOutputPluginDir
		err = cp.Copy(pluginUri, filepath.Join(i.ParentOutputPluginDir, metadata.Id), opts)
		if err != nil {
			return fmt.Errorf("failed to copy %s to plugin dir %s: %s", pluginUri, i.ParentOutputPluginDir, err)
		}

		return nil
	}

	return fmt.Errorf("cannot determine fetch method for plugin %s", pluginUri)
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

// Useful if the entrypoint requires you to build the file.
func (i *AugmentorPluginInstaller) getBuildCommand(entrypoint string, outputExecPath string, pluginDir string) *exec.Cmd {
	extension := strings.TrimPrefix(filepath.Ext(entrypoint), ".")

	switch extension {
	case "go":
		return exec.Command("go", "build", "-o", outputExecPath, pluginDir)
	case "python":
		return nil
	default:
		return nil
	}
}

// Creates the command for running the plugin.
// `entrypoint` is not exactly the Entrypoint from PluginMeta. It is the full path to entrypoint.
// `execPath` should only be set if a language requires an executable to run the plugin.
//
// Usage [Go]:
//
//	i.GetRunCmd("main.go", "./build/jchen42703/Example-Augmentor-go")
//
// Usage [Python]: (No Build)
//
//	i.GetRunCmd("./plugin-python/plugin.py", "")
func (i *AugmentorPluginInstaller) GetRunCmd(entrypointPath string, execPath string) (*exec.Cmd, error) {
	extension := strings.TrimPrefix(filepath.Ext(entrypointPath), ".")
	absEntrypoint, err := filepath.Abs(entrypointPath)
	if err != nil {
		return nil, err
	}

	switch extension {
	case "go":
		absExec, err := filepath.Abs(execPath)
		if err != nil {
			return nil, fmt.Errorf("failed convert exec path to abs path: %s", err)
		}

		return exec.Command(absExec), nil
	case "py":
		return exec.Command("python", absEntrypoint), nil
	default:
		return nil, fmt.Errorf("unsupported language extension: %s", extension)
	}
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
func (i *AugmentorPluginInstaller) Install(pluginDir string) (*PluginMeta, error) {
	metadata, err := i.ReadMetadata(pluginDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read plugin meta: %s", err)
	}

	// Build plugin executable to the output directory.
	pluginInstallDir, err := filepath.Abs(pluginDir)
	if err != nil {
		return nil, fmt.Errorf("filepath.Abs: %s", pluginInstallDir)
	}

	outputExecPath := filepath.Join(pluginInstallDir, metadata.Id)
	cmd := i.getBuildCommand(metadata.Entrypoint, outputExecPath, pluginInstallDir)
	if cmd != nil {
		err = run.Cmd(cmd, i.Writer)

		if err != nil {
			return nil, fmt.Errorf("plugin build failed: %s", err)
		}
	}

	return metadata, nil
}

// Use this when you want to install multiple plugins in a directory.
// parentPluginDir
//
//	plugin1
//	plugin2
func (i *AugmentorPluginInstaller) InstallAll() ([]*PluginMeta, error) {
	pluginMetas, err := i.GetAllPlugins()
	if err != nil {
		return nil, fmt.Errorf("failed get all plugins: %s", err)
	}

	for _, meta := range pluginMetas {
		pluginDir := filepath.Join(i.ParentOutputPluginDir, meta.Id)
		_, err := i.Install(pluginDir)
		if err != nil {
			return nil, fmt.Errorf("failed to install plugin %s: %s", meta.Id, err)
		}
	}

	return pluginMetas, nil
}
