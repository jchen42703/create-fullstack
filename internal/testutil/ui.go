package testutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jchen42703/create-fullstack/internal/directory"
	"github.com/jchen42703/create-fullstack/internal/log"
	"go.uber.org/zap"
)

// Creates test directory and logger.
// Returns the logger, test working directory path, or error
func CreateTestDirAndLogger(testDir string) (*zap.Logger, string, error) {
	// Create test files and test in a separate directory.
	wd, err := os.Getwd()
	if err != nil {
		return nil, "", fmt.Errorf("failed to Getwd: %s", err)
	}

	// Create test working dir + logger
	testWd := filepath.Join(wd, testDir)
	err = os.Mkdir(testWd, directory.READ_WRITE_EXEC_PERM)
	if err != nil {
		return nil, "", fmt.Errorf("failed to mk test dir: %s", err)
	}

	logFilePath := filepath.Join(testWd, "create-fullstack.log")
	logger, err := log.CreateLogger(logFilePath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create logger")
	}

	return logger, testWd, nil
}

// Generic test to cleanup the working directory and logger.
func CleanupBaseTest(testWd string, logger *zap.Logger) error {
	err := os.RemoveAll(testWd)
	if err != nil {
		return fmt.Errorf("failed to cleanup log: %s", err)
	}

	if err := logger.Sync(); err != nil {
		// this sync error is safe to ignore, since stdout doesn't support syncing in Linux/OS X
		if !strings.HasSuffix(err.Error(), "sync /dev/stdout: invalid argument") {
			return fmt.Errorf("Error cleaning up logger: %s", err.Error())
		}
	}

	return nil
}
