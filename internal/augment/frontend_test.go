package augment_test

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/jchen42703/create-fullstack/internal/augment"
	"github.com/jchen42703/create-fullstack/internal/directory"
	"github.com/jchen42703/create-fullstack/internal/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapio"
)

// General cleanup function that removes an output directory, and removes the logger file.
func cleanupFunc(t *testing.T, outputDir string, logFilePath string, logger *zap.Logger) {
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

func TestAddTailwind(t *testing.T) {
	// 1. Create next.js project
	// 2. Add Tailwind to it
	// 3. Check that files were properly created
	outputDir := "test-next-ts"
	// createNextJsCmd := exec.Command("yarn", "create", "next-app", "--example", "with-typescript", outputDir)
	createNextJsCmd := exec.Command("yarn", "create", "next-app", "--typescript", "--eslint", outputDir)

	// Cleanup
	logFilePath := "./create-fullstack.log"
	logger, err := log.CreateLogger(logFilePath)
	if err != nil {
		t.Fatalf("failed to create logger")
	}

	defer cleanupFunc(t, outputDir, logFilePath, logger)

	writer := &zapio.Writer{
		Log:   logger,
		Level: zapcore.DebugLevel,
	}

	err = augment.RunCommand(createNextJsCmd, writer)
	if err != nil {
		t.Fatalf("failed to create next js app: %s", err.Error())
	}

	err = os.Chdir(outputDir)
	if err != nil {
		t.Fatalf("failed to change to output directory: %s", err)
	}

	err = augment.AddTailwind(writer, logger.Sugar())
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

func TestInitializeNextDocker(t *testing.T) {
	outputDir := "test-next-ts-docker"
	// createNextJsCmd := exec.Command("yarn", "create", "next-app", "--example", "with-typescript", outputDir)
	createNextJsCmd := exec.Command("yarn", "create", "next-app", "--typescript", "--eslint", outputDir)

	logFilePath := "./create-fullstack.log"
	logger, err := log.CreateLogger(logFilePath)
	if err != nil {
		t.Fatalf("failed to create logger")
	}

	// Cleanup
	defer cleanupFunc(t, outputDir, logFilePath, logger)
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

	err = augment.RunCommand(createNextJsCmd, writer)
	if err != nil {
		t.Fatalf("failed to create next js app: %s", err.Error())
	}

	err = os.Chdir(outputDir)
	if err != nil {
		t.Fatalf("failed to change to output directory: %s", err)
	}

	err = augment.InitializeNextDocker(3000)
	if err != nil {
		t.Fatalf("failed to initialize docker configs: %s", err.Error())
	}

	// docker build -t jchen42703/nextjs-test-docker:latest .
	dockerBuildCmd := exec.Command("docker", "build", "-t", "jchen42703/nextjs-test-docker", ".")
	err = augment.RunCommand(dockerBuildCmd, writer)
	if err != nil {
		t.Fatalf("failed to build nextjs docker container: %s", err.Error())
	}

	// docker run -p 3000:3000 jchen42703/nextjs-test-docker
}
