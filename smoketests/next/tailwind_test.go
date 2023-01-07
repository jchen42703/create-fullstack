package ui_test

import (
	"os"
	"strings"
	"testing"

	"github.com/jchen42703/create-fullstack/core/ui"
	"github.com/jchen42703/create-fullstack/internal/directory"
	"github.com/jchen42703/create-fullstack/internal/log"
	"github.com/jchen42703/create-fullstack/internal/testutil"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapio"
)

func TestAddTailwind(t *testing.T) {
	outputDir := "test-next-ts-tailwind"
	logFilePath := "./create-fullstack-tailwind.log"
	logger, err := log.CreateLogger(logFilePath)
	if err != nil {
		t.Fatalf("failed to create logger")
	}

	// Cleanup
	defer testutil.CleanupUiTest(t, outputDir, logFilePath, logger)
	writer := &zapio.Writer{
		Log:   logger,
		Level: zapcore.DebugLevel,
	}

	// 1. Create next.js project
	// 2. Add Tailwind to it
	// 3. Check that files were properly created
	baseTemplate := "test-next-ts"
	err = testutil.TemplateCache.GetTemplateAndCopy(baseTemplate, outputDir)
	// Only create base template if one does not exist already
	if err != nil {
		testutil.CreateTemplate(t, baseTemplate, writer)
	}

	// Copy the base template to outputDir
	err = testutil.TemplateCache.GetTemplateAndCopy(baseTemplate, outputDir)
	if err != nil {
		// Should not ever happen
		t.Fatalf("failed to get template after caching it: %s", err)
	}

	err = os.Chdir(outputDir)
	if err != nil {
		t.Fatalf("failed to change to output directory: %s", err)
	}

	err = ui.AddTailwind(writer, logger.Sugar())
	if err != nil {
		t.Fatalf("failed to augment with tailwind: %s", err.Error())
	}

	// Raw Test that changes to files work
	readBytes, err := os.ReadFile("./styles/globals.css")
	if err != nil {
		t.Fatalf("failed to read styles/globals.css: %s", err.Error())
	}

	tailwindHeader := `@tailwind base;
@tailwind components;
@tailwind utilities;

`

	if !strings.HasPrefix(string(readBytes), tailwindHeader) {
		t.Fatalf("globals.css missing have proper tailwind header")
	}

	if exists, _ := directory.Exists("./tailwind.config.js"); !exists {
		t.Fatalf("missing tailwind.config.js")
	}

	if exists, _ := directory.Exists("./postcss.config.js"); !exists {
		t.Fatalf("missing postcss.config.js")
	}

	// Integration test this by confirming that adding a new component with tailwind works
	// Not easily parallelizable because would need to manage port numbers.
}
