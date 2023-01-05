package ui

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/jchen42703/create-fullstack/core/lang"
	"github.com/jchen42703/create-fullstack/core/run"
	"github.com/jchen42703/create-fullstack/internal/directory"
	"go.uber.org/zap"
)

// Adds Tailwind to the repository
type TailwindAugmenter struct {
	Lang lang.PROGRAMMING_LANGUAGE
}

func (a *TailwindAugmenter) Augment() error {
	// var err error
	// switch a.Lang {
	// case lang.Go:
	// 	err = addGoDockerfile()
	// case lang.Javascript:
	// 	err = addJavascriptDockerfile()
	// case lang.Python:
	// 	err = addPythonDockerfile()
	// case lang.Typescript:
	// 	err = addTypescriptDockerfile()
	// default:
	// 	return fmt.Errorf("unsupported programming language '%s' for generating dockerfile", a.Lang)
	// }

	// if err != nil {
	// 	return fmt.Errorf("failed to gen api dockerfile: %s", err)
	// }

	return nil
}

// Adds Tailwind to the project following the official docs:
// https://tailwindcss.com/docs/guides/nextjs
// This is better than the regular with-tailwind Next.js template because
// this approach works with Typescript too.
func AddTailwind(logWriter io.Writer, logger *zap.SugaredLogger) error {
	logger.Debug("Adding Tailwind and peer dependencies...\n")
	commands := []*exec.Cmd{
		exec.Command("yarn", "add", "-D", "tailwindcss", "postcss", "autoprefixer"),
		exec.Command("npx", "tailwindcss", "init", "-p"),
	}

	for _, cmd := range commands {
		err := run.Cmd(cmd, logWriter) //blocks until sub process is complete
		if err != nil {
			return fmt.Errorf("AddTailwind: %s", err.Error())
		}
	}

	logger.Debug("Creating tailwind config...\n")

	tailwindConfig := `/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		"./pages/**/*.{js,ts,jsx,tsx}",
		"./components/**/*.{js,ts,jsx,tsx}",
	],
	theme: {
		extend: {},
	},
	plugins: [],
};
`

	err := os.WriteFile("./tailwind.config.js", []byte(tailwindConfig), directory.READ_WRITE_PERM)
	if err != nil {
		return fmt.Errorf("AddTailwind: writing config failed: %s", err.Error())
	}

	logger.Debug("Attempting to add the tailwind styles to the global styles...\n")

	// Assume globals.scss/css is in styles/
	possibleStylesPaths := []string{
		"./styles/globals.css",
		"./styles/globals.scss",
	}

	// Checks if any of the path exists and tries to match
	globalStylesPath := ""
	for _, stylesPath := range possibleStylesPaths {
		var exists bool
		if exists, err = directory.Exists(stylesPath); exists {
			globalStylesPath = stylesPath
			break
		}
	}

	// DNE
	if globalStylesPath == "" && err == nil {
		return fmt.Errorf("AddTailwind: template must have a globals css or scss file for path '%s'", globalStylesPath)
	} else if err != nil {
		return fmt.Errorf("AddTailwind: validating styles path failed: %s", err.Error())
	}

	// Exists, so add tailwind styles to global styles
	tailwindHeader := `@tailwind base;
@tailwind components;
@tailwind utilities;

`

	readBytes, err := os.ReadFile(globalStylesPath)
	if err != nil {
		return fmt.Errorf("AddTailwind: failed to read global styles: %s", err.Error())
	}

	newGlobalStyles := append([]byte(tailwindHeader), readBytes...)
	err = os.WriteFile(globalStylesPath, []byte(newGlobalStyles), directory.READ_WRITE_PERM)
	if err != nil {
		return fmt.Errorf("AddTailwind: writing global styles failed: %s", err.Error())
	}

	logger.Debug("Successfully added the tailwind styles to the global styles...\n")

	return nil
}
