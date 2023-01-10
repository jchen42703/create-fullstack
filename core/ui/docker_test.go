package ui_test

import (
	"testing"

	"github.com/jchen42703/create-fullstack/core/ui"
	"github.com/jchen42703/create-fullstack/internal/testutil"
)

func TestInitializeNextDocker(t *testing.T) {
	t.Run("FileCreation", func(t *testing.T) {
		t.Parallel()
		// 1. Add files
		// 2. Test that files are being written correctly
		// 3. Cleanup
		testDir := "test-next-ts-docker"
		logger, testWd, err := testutil.CreateTestDirAndLogger(testDir)
		if err != nil {
			t.Fatalf("failed to create test dir/logger: %s", err)
		}

		defer func() {
			err := testutil.CleanupBaseTest(testWd, logger)
			if err != nil {
				t.Errorf("CleanupBaseTest: %s", err)
			}
		}()

		err = ui.InitializeNextDocker(testWd, 3000, true)
		if err != nil {
			t.Fatalf("InitializeNextDocker: %s", err)
		}

		// TODO: check that the files are exactly how you want them to be created.
	})
}
