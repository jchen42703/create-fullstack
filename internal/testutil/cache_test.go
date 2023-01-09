package testutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jchen42703/create-fullstack/internal/directory"
	"github.com/jchen42703/create-fullstack/internal/testutil"
)

func TestBaseTemplateCache(t *testing.T) {
	cache := testutil.NewBaseTemplateCache()
	// Create test dir
	testDirPath := "./test-dir"
	err := os.MkdirAll(testDirPath, directory.READ_WRITE_EXEC_PERM)
	if err != nil {
		t.Fatalf("failed to create test dir: %s", err)
	}

	defer func() {
		os.RemoveAll(testDirPath)
	}()

	// Populate test dir with files
	testFilePath := ".gitignore"
	_, err = os.Create(filepath.Join(testDirPath, testFilePath))
	if err != nil {
		t.Fatalf("failed to create .gitignore: %s", err)
	}

	cache.AddTemplate("test", testDirPath)

	outputDir := "./output-dir"
	err = cache.GetTemplateAndCopy("test", outputDir)
	if err != nil {
		t.Fatalf("GetTemplateAndCopy failed: %s", err)
	}
	defer func() {
		os.RemoveAll(outputDir)
	}()

	exists, err := directory.Exists(filepath.Join(outputDir, testFilePath))
	if !exists {
		t.Fatalf("did not properly copy hidden file: %s", err)
	}
}
