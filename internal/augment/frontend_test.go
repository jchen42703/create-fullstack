package augment_test

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/jchen42703/create-fullstack/internal/augment"
	"github.com/jchen42703/create-fullstack/internal/directory"
)

func TestAddTailwind(t *testing.T) {
	// 1. Create next.js project
	// 2. Add Tailwind to it
	// 3. Check that files were properly created
	outputDir := "test-next-ts"
	// createNextJsCmd := exec.Command("yarn", "create", "next-app", "--example", "with-typescript", outputDir)
	createNextJsCmd := exec.Command("yarn", "create", "next-app", "--typescript", "--eslint", outputDir)

	// Cleanup
	defer func() {
		os.Chdir("../")
		err := os.RemoveAll(outputDir)
		if err != nil {
			t.Fatalf("failed to clean up test directory: %s", err)
		}

	}()

	writer := log.Writer()
	createNextJsCmd.Stderr = writer
	createNextJsCmd.Stdout = writer

	err := createNextJsCmd.Run() //blocks until sub process is complete
	if err != nil {
		t.Fatalf("failed to create next js app: %s", err.Error())
	}

	os.Chdir(outputDir)

	err = augment.AddTailwind()
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
