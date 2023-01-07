package testutil

import (
	"io"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/jchen42703/create-fullstack/core/run"
)

// Creates the next.js typescript template
func CreateTemplate(t *testing.T, baseTemplate string, writer io.Writer) {
	t.Logf("Creating template %s\n", baseTemplate)
	createNextJsCmd := exec.Command("yarn", "create", "next-app", "--typescript", "--eslint", baseTemplate)
	err := run.Cmd(createNextJsCmd, writer)
	if err != nil {
		t.Fatalf("failed to create next js app: %s", err.Error())
	}

	absPath, err := filepath.Abs(baseTemplate)
	if err != nil {
		t.Fatalf("failed to get absolute path for %s", baseTemplate)
	}

	TemplateCache.AddTemplate("test-next-ts", absPath)
}
