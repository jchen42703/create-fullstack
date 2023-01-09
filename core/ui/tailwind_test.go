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
	"github.com/jchen42703/create-fullstack/internal/log"
	"go.uber.org/zap/zapcore"
)

// Basic npm project with only a package.json
func createBaseYarnProj(workingDirectory string, writer io.Writer) error {
	// Yarn init empty directory
	yarnInit := exec.Command("yarn", "--cwd", workingDirectory, "init", "-y")
	err := run.Cmd(yarnInit, writer)
	if err != nil {
		return fmt.Errorf("createBaseNpmProj: failed to run init: %s", err)
	}

	return nil
}

// Basic npm project with only a package.json
func createProjWithGlobalCss(workingDirectory string, writer io.Writer) error {
	// Yarn init empty directory
	err := createBaseYarnProj(workingDirectory, writer)
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

func initTailwindTest(t *testing.T, testDir string) (string, *ui.TailwindAugmenter) {
	// Create test files and test in a separate directory.
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to Getwd: %s", err)
	}

	// Create test working dir + logger
	testWd := filepath.Join(wd, testDir)
	err = os.Mkdir(testWd, directory.READ_WRITE_EXEC_PERM)
	if err != nil {
		t.Fatalf("failed to make test dir: %s", err)
	}

	logFilePath := filepath.Join(testWd, "create-fullstack-tailwind.log")
	logger, err := log.CreateLogger(logFilePath)
	if err != nil {
		t.Fatalf("failed to create logger")
	}

	augmenter := ui.NewTailwindAugmenter(lang.Typescript, logger, zapcore.DebugLevel)
	return testWd, augmenter
}

func TestTailwindAugmenter(t *testing.T) {
	t.Run("NoGlobalsStyles", func(t *testing.T) {
		t.Parallel()
		testDir := "test-next-ts-tailwind"
		testWd, augmenter := initTailwindTest(t, testDir)
		defer func() {
			err := os.RemoveAll(testWd)
			if err != nil {
				t.Errorf("failed to cleanup wd: %s", testWd)
			}
		}()

		// Initialize test dir
		err := createBaseYarnProj(testWd, augmenter.LogWriter)
		if err != nil {
			t.Fatalf("failed to init test dir: %s", err)
		}

		// Run tests
		err = augmenter.Augment(testWd)
		if err == nil || !strings.HasSuffix(err.Error(), "template must have a globals css or scss file for path ''") {
			t.Fatalf("should raise err looking for globals css/scss file")
		}
	})

	t.Run("WithGlobalsStylesCss", func(t *testing.T) {
		t.Parallel()
		testDir := "test-next-ts-tailwind-css"
		testWd, augmenter := initTailwindTest(t, testDir)
		defer func() {
			err := os.RemoveAll(testWd)
			if err != nil {
				t.Errorf("failed to cleanup wd: %s", testWd)
			}
		}()

		// Initialize test dir
		err := createProjWithGlobalCss(testWd, augmenter.LogWriter)
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
