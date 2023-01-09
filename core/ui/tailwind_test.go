package ui_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jchen42703/create-fullstack/core/lang"
	"github.com/jchen42703/create-fullstack/core/run"
	"github.com/jchen42703/create-fullstack/core/ui"
	"github.com/jchen42703/create-fullstack/internal/directory"
	"github.com/jchen42703/create-fullstack/internal/testutil"
	"go.uber.org/zap/zapcore"
)

// Basic npm project with only a package.json
func initYarnProj(workingDirectory string, writer io.Writer) error {
	// Yarn init empty directory
	yarnInit := exec.Command("yarn", "--cwd", workingDirectory, "init", "-y")
	err := run.Cmd(yarnInit, writer)
	if err != nil {
		return fmt.Errorf("initYarnProj: failed to run init: %s", err)
	}

	return nil
}

// Basic npm project with only a package.json
func initProjWithGlobalCss(workingDirectory string, writer io.Writer) error {
	// Yarn init empty directory
	err := initYarnProj(workingDirectory, writer)
	if err != nil {
		return fmt.Errorf("createBaseYarnProj: %s", err)
	}

	err = os.Mkdir(filepath.Join(workingDirectory, "styles"), directory.READ_WRITE_EXEC_PERM)
	if err != nil {
		return fmt.Errorf("failed to make test dir: %s", err)
	}

	_, err = os.Create(filepath.Join(workingDirectory, "styles", "globals.css"))
	if err != nil {
		return fmt.Errorf("failed to create styles/globals.css: %s", err)
	}

	return nil
}

func TestTailwindAugmenter(t *testing.T) {
	t.Run("NoGlobalsStyles", func(t *testing.T) {
		t.Parallel()
		testDir := "test-next-ts-tailwind"
		logger, testWd, err := testutil.CreateTestDirAndLogger(testDir)
		if err != nil {
			t.Fatalf("CreateTestDirAndLogger: %s", err)
		}

		defer func() {
			err := testutil.CleanupBaseTest(testWd, logger)
			if err != nil {
				t.Errorf("CleanupBaseTest: %s", err)
			}
		}()

		augmenter := ui.NewTailwindAugmenter(lang.Typescript, logger, zapcore.DebugLevel)

		// Initialize test dir
		err = initYarnProj(testWd, augmenter.LogWriter)
		if err != nil {
			t.Fatalf("failed to init test dir: %s", err)
		}

		// Run tests
		err = augmenter.Augment(testWd)
		if err == nil || !strings.HasSuffix(err.Error(), "template must have a globals css or scss file for path ''") {
			// Prints logs on error
			t.Log(testutil.GetLogs(filepath.Join(testWd, "create-fullstack.log")))
			t.Fatalf("should raise err looking for globals css/scss file, err: %s", err)
		}
	})

	t.Run("WithGlobalsStylesCss", func(t *testing.T) {
		t.Parallel()
		testDir := "test-next-ts-tailwind-css"
		logger, testWd, err := testutil.CreateTestDirAndLogger(testDir)
		if err != nil {
			t.Fatalf("CreateTestDirAndLogger: %s", err)
		}

		defer func() {
			err := testutil.CleanupBaseTest(testWd, logger)
			if err != nil {
				t.Errorf("CleanupBaseTest: %s", err)
			}
		}()

		augmenter := ui.NewTailwindAugmenter(lang.Typescript, logger, zapcore.DebugLevel)

		// Initialize test dir
		err = initProjWithGlobalCss(testWd, augmenter.LogWriter)
		if err != nil {
			t.Fatalf("failed to init test dir: %s", err)
		}

		// Run tests
		err = augmenter.Augment(testWd)
		if err != nil {
			t.Fatalf("err augmenting dir: %s", err)
		}

		// Raw Test that changes to files work
		stylesPath := filepath.Join(testWd, "styles", "globals.css")
		readBytes, err := os.ReadFile(stylesPath)
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

		tailwindCfgPath := filepath.Join(testWd, "tailwind.config.js")
		if exists, _ := directory.Exists(tailwindCfgPath); !exists {
			t.Fatalf("missing tailwind.config.js")
		}

		postCssCfgPath := filepath.Join(testWd, "postcss.config.js")
		if exists, _ := directory.Exists(postCssCfgPath); !exists {
			t.Fatalf("missing postcss.config.js")
		}
	})
}
