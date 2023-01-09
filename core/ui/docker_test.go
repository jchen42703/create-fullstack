package ui_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jchen42703/create-fullstack/core/ui"
	"github.com/jchen42703/create-fullstack/internal/directory"
)

func initDockerTest(t *testing.T, testDir string) string {
	// Create test files and test in a separate directory.
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to Getwd: %s", err)
	}

	// Create test working dir + logger
	testWd := filepath.Join(wd, testDir)
	err = os.Mkdir(testWd, directory.READ_WRITE_EXEC_PERM)
	if err != nil {
		t.Fatalf("failed to mk test dir: %s", err)
	}

	return testWd
}

func TestInitializeNextDocker(t *testing.T) {
	t.Run("FileCreation", func(t *testing.T) {
		t.Parallel()
		// 1. Add files
		// 2. Test that files are being written correctly
		// 3. Cleanup
		testDir := "test-next-ts-docker"
		testWd := initDockerTest(t, testDir)
		defer func() {
			err := os.RemoveAll(testWd)
			if err != nil {
				t.Errorf("failed to cleanup wd: %s", testWd)
			}
		}()

		err := ui.InitializeNextDocker(testWd, 3000, true)
		if err != nil {
			t.Fatalf("InitializeNextDocker: %s", err)
		}

		// TODO: check that the files are exactly how you want them to be created.
	})
}
