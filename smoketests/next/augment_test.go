package next_test

import (
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
	"go.uber.org/zap/zapio"
)

func TestNextAugmentations(t *testing.T) {
	t.Cleanup(func() {
		// Remove all cached templates
		for _, cachedTemplatePath := range testutil.TemplateCache.TemplatePaths {
			os.RemoveAll(cachedTemplatePath)
		}
	})

	t.Run("InitializeDocker", func(t *testing.T) {
		testDir := "test-next-ts-docker"
		logger, testWd, err := testutil.CreateTestDirAndLogger(testDir)
		if err != nil {
			t.Fatalf("CreateTestDirAndLogger: %s", err)
		}

		// Cleanup
		defer func() {
			err := testutil.CleanupBaseTest(testWd, logger)
			if err != nil {
				t.Errorf("CleanupBaseTest: %s", err)
			}

			// Cleanup docker image
			cleanupDocker := exec.Command("docker", "image", "rm", "jchen42703/nextjs-test-docker")
			stdout, err := cleanupDocker.Output()
			if err != nil {
				t.Log(stdout)
				t.Errorf("failed to cleanup docker image")
			}
		}()

		writer := &zapio.Writer{
			Log:   logger,
			Level: zapcore.DebugLevel,
		}

		baseTemplate := "test-next-ts"
		err = testutil.TemplateCache.GetTemplateAndCopy(baseTemplate, testWd)
		// Only create base template if one does not exist already
		if err != nil {
			testutil.CreateTemplate(t, baseTemplate, writer)
		}

		// Copy the base template to outputDir
		err = testutil.TemplateCache.GetTemplateAndCopy(baseTemplate, testWd)
		if err != nil {
			// Should not ever happen
			t.Fatalf("failed to get template after caching it: %s", err)
		}

		err = ui.InitializeNextDocker(testWd, 3000, true)
		if err != nil {
			t.Fatalf("failed to initialize docker configs: %s", err.Error())
		}

		dockerBuildCmd := exec.Command("docker", "build", "-t", "jchen42703/nextjs-test-docker", testWd)
		err = run.Cmd(dockerBuildCmd, writer)
		if err != nil {
			t.Fatalf("failed to build nextjs docker container: %s", err.Error())
		}

		// docker run -p 3000:3000 jchen42703/nextjs-test-docker
	})

	t.Run("AddTailwind", func(t *testing.T) {
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

		// 1. Create next.js project
		// 2. Add Tailwind to it
		// 3. Check that files were properly created
		baseTemplate := "test-next-ts"
		err = testutil.TemplateCache.GetTemplateAndCopy(baseTemplate, testWd)
		// Only create base template if one does not exist already
		if err != nil {
			testutil.CreateTemplate(t, baseTemplate, augmenter.LogWriter)
		}

		// Copy the base template to outputDir
		err = testutil.TemplateCache.GetTemplateAndCopy(baseTemplate, testWd)
		if err != nil {
			// Should not ever happen
			t.Fatalf("failed to get template after caching it: %s", err)
		}

		err = augmenter.Augment(testWd)
		if err != nil {
			t.Fatalf("failed to augment with tailwind: %s", err.Error())
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

		// Integration test this by confirming that adding a new component with tailwind works
		// Not easily parallelizable because would need to manage port numbers.
	},
	)
}
