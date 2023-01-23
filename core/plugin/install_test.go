package plugin_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/jchen42703/create-fullstack/core/plugin"
	"github.com/jchen42703/create-fullstack/internal/directory"
	"github.com/jchen42703/create-fullstack/internal/log"
)

func createValidPlugin(parentPluginDir string, pluginId string) error {
	pluginDir := filepath.Join(parentPluginDir, pluginId)
	err := os.Mkdir(pluginDir, directory.READ_WRITE_EXEC_PERM)
	if err != nil {
		return fmt.Errorf("failed to create plugin dir %s: %s", pluginDir, err)
	}

	// Create valid metadata
	meta := &plugin.PluginMeta{
		Id:          pluginId,
		Version:     "dev",
		InstallLink: "N/A",
	}

	// Write to directory
	f, err := json.Marshal(meta)
	if err != nil {
		return fmt.Errorf("failed to marshal augmentor_metadata.json for %s: %s", pluginId, err)
	}

	metaPath := filepath.Join(pluginDir, "augmentor_metadata.json")
	err = os.WriteFile(metaPath, f, directory.READ_WRITE_EXEC_PERM)
	if err != nil {
		return fmt.Errorf("failed to write metadata: %s", err)
	}

	return nil
}

func createInvalidPlugin(parentPluginDir string, pluginId string, badMeta bool) error {
	pluginDir := filepath.Join(parentPluginDir, pluginId)
	err := os.Mkdir(pluginDir, directory.READ_WRITE_EXEC_PERM)
	if err != nil {
		return fmt.Errorf("failed to create plugin dir %s: %s", pluginDir, err)
	}

	if badMeta {
		// Create valid metadata
		meta := &plugin.PluginMeta{
			InstallLink: "N/A",
		}

		// Write to directory
		f, err := json.Marshal(meta)
		if err != nil {
			return fmt.Errorf("failed to marshal augmentor_metadata.json for %s: %s", pluginId, err)
		}

		metaPath := filepath.Join(pluginDir, "augmentor_metadata.json")
		err = os.WriteFile(metaPath, f, directory.READ_WRITE_EXEC_PERM)
		if err != nil {
			return fmt.Errorf("failed to write metadata: %s", err)
		}
	}

	return nil
}

func createPluginsDir(pluginDir string, numValidPlugins int, numInvalidPlugins int) error {
	// Create pluginDir
	err := os.MkdirAll(pluginDir, directory.READ_WRITE_EXEC_PERM)
	if err != nil {
		return fmt.Errorf("failed to create plugin dir: %s", err)
	}

	pluginId := "valid-plugin"
	// Create valid plugin dirs
	// - make sure to include a valid augmentor_metadata.json
	for i := 0; i < numValidPlugins; i++ {
		childPluginId := pluginId + strconv.Itoa(i)
		err := createValidPlugin(pluginDir, childPluginId)
		if err != nil {
			return fmt.Errorf("failed to create plugin %s", childPluginId)
		}
	}

	// Create invalid plugin dirs
	pluginId = "invalid-plugin"
	for i := 0; i < numInvalidPlugins; i++ {
		childPluginId := pluginId + strconv.Itoa(i)
		err := createInvalidPlugin(pluginDir, childPluginId, false)
		if err != nil {
			return fmt.Errorf("failed to create plugin %s", childPluginId)
		}
	}

	return nil
}
func TestGetAllPlugins(t *testing.T) {
	logger, err := log.CreateLogger("./create-fullstack.log")
	if err != nil {
		t.Fatalf("failed to create logger: %s", err)
	}

	// Plugin dir DNE
	t.Run("DirDNE", func(t *testing.T) {
		parentPluginDir := "./plugins"
		installer := &plugin.AugmentorPluginInstaller{
			ParentOutputPluginDir: parentPluginDir,
			Logger:                logger,
		}
		_, err := installer.GetAllPlugins()
		if err == nil || !strings.Contains(err.Error(), "no such file or directory") {
			t.Fatalf("shouldn't be able to get plugins with plugin dir not created")
		}
	})

	tests := map[string]struct {
		PluginDir  string
		NumValid   int
		NumInvalid int
	}{
		"AllValid": {
			PluginDir:  "all-valid",
			NumValid:   6,
			NumInvalid: 0,
		},
		"SomeInvalid": {
			PluginDir:  "some-invalid",
			NumValid:   3,
			NumInvalid: 3,
		},
		"AllInvalid": {
			PluginDir:  "all-invalid",
			NumValid:   0,
			NumInvalid: 6,
		},
	}

	for name, testEx := range tests {
		t.Run(name, func(t *testing.T) {
			installer := &plugin.AugmentorPluginInstaller{
				ParentOutputPluginDir: testEx.PluginDir,
				Logger:                logger,
			}
			err := createPluginsDir(testEx.PluginDir, testEx.NumValid, testEx.NumInvalid)
			defer func() {
				err := os.RemoveAll(testEx.PluginDir)
				if err != nil {
					t.Errorf("failed to cleanup %s: %s", testEx.PluginDir, err)
				}
			}()

			if err != nil {
				t.Fatalf("failed to create plugin dirs: %s", err)
			}

			pluginsMeta, err := installer.GetAllPlugins()
			if err != nil {
				t.Fatalf("failed to GetAllPlugins: %s", err)
			}

			if len(pluginsMeta) != testEx.NumValid {
				t.Fatalf("found %d plugin metadata json configs, but should be %d", len(pluginsMeta), testEx.NumValid)
			}
		})
	}
}
