package testutil

import (
	"os"
	"strings"
	"testing"

	"go.uber.org/zap"
)

// General cleanup function that removes an output directory, and removes the logger file.
func CleanupUiTest(t *testing.T, outputDir string, logFilePath string, logger *zap.Logger) {
	err := os.Chdir("../")
	if err != nil {
		t.Fatalf("failed to change back to reg directory during cleanup: %s", err)
	}

	err = os.RemoveAll(outputDir)
	if err != nil {
		t.Fatalf("failed to clean up test directory: %s", err)
	}

	err = os.Remove(logFilePath)
	if err != nil {
		t.Fatalf("failed to cleanup log: %s", err)
	}

	if err := logger.Sync(); err != nil {
		// this sync error is safe to ignore, since stdout doesn't support syncing in Linux/OS X
		if !strings.HasSuffix(err.Error(), "sync /dev/stdout: invalid argument") {
			t.Fatalf("Error cleaning up logger: %s", err.Error())
		}
	}
}
