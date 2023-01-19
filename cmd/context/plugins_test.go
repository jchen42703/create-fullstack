package context_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jchen42703/create-fullstack/cmd/context"
)

func TestGetGlobalPluginsDir(t *testing.T) {

	runtimeOs := runtime.GOOS
	if runtimeOs != "windows" {
		expected := filepath.Join(os.Getenv("HOME"), ".create-fullstack", "plugins")
		res := context.GetGlobalPluginsDir(runtimeOs)
		if res != expected {
			t.Fatalf("expected %s but go %s", expected, res)
		}
	}

	// Not testing windows since unit tests run on Linux, but could add it down the line
}
