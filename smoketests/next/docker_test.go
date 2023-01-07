package ui_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/jchen42703/create-fullstack/core/run"
	"github.com/jchen42703/create-fullstack/core/ui"
	"github.com/jchen42703/create-fullstack/internal/log"
	"github.com/jchen42703/create-fullstack/internal/testutil"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapio"
)

func TestInitializeNextDocker(t *testing.T) {
	outputDir := "test-next-ts-docker"
	logFilePath := "./create-fullstack-docker.log"
	logger, err := log.CreateLogger(logFilePath)
	if err != nil {
		t.Fatalf("failed to create logger")
	}

	// Cleanup
	defer testutil.CleanupUiTest(t, outputDir, logFilePath, logger)

	// Cleanup docker image
	defer func() {
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

	err = ui.InitializeNextDocker(3000)
	if err != nil {
		t.Fatalf("failed to initialize docker configs: %s", err.Error())
	}

	// docker build -t jchen42703/nextjs-test-docker:latest .
	dockerBuildCmd := exec.Command("docker", "build", "-t", "jchen42703/nextjs-test-docker", ".")
	err = run.Cmd(dockerBuildCmd, writer)
	if err != nil {
		t.Fatalf("failed to build nextjs docker container: %s", err.Error())
	}

	// docker run -p 3000:3000 jchen42703/nextjs-test-docker
}
