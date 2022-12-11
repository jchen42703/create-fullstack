package augment

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jchen42703/create-fullstack/internal/directory"
)

// Adds Tailwind to the project following the official docs:
// https://tailwindcss.com/docs/guides/nextjs
// This is better than the regular with-tailwind Next.js template because
// this approach works with Typescript too.
func AddTailwind(isSCSS bool) error {
	commands := []*exec.Cmd{
		exec.Command("yarn", "-D", "tailwindcss", "postcss", "autoprefixer"),
		exec.Command("npx", "tailwindcss", "init", "-p"),
	}

	for _, cmd := range commands {
		stdout, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("AddTailwind: '%s' failed with err: %s", cmd.String(), err.Error())
		}

		fmt.Println(stdout)
	}

	tailwindConfig := `
	/** @type {import('tailwindcss').Config} */
	module.exports = {
	content: [
		"./pages/**/*.{js,ts,jsx,tsx}",
		"./components/**/*.{js,ts,jsx,tsx}",
	],
	theme: {
		extend: {},
	},
	plugins: [],
	}
	`

	err := os.WriteFile("./tailwind.config.js", []byte(tailwindConfig), 0644)
	if err != nil {
		return fmt.Errorf("AddTailwind: writing config failed: %s", err.Error())
	}

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
	if globalStylesPath != "" && err == nil {
		return fmt.Errorf("AddTailwind: template must have a globals css or scss file for path '%s'", globalStylesPath)
	} else if err != nil {
		return fmt.Errorf("AddTailwind: validating styles path failed: %s", err.Error())
	}

	// Exists, so add tailwind styles to global styles
	tailwindHeader := `
	@tailwind base;
	@tailwind components;
	@tailwind utilities;

	`

	readBytes, err := os.ReadFile(globalStylesPath)
	if err != nil {
		return fmt.Errorf("AddTailwind: failed to read global styles: %s", err.Error())
	}

	newGlobalStyles := append([]byte(tailwindHeader), readBytes...)
	err = os.WriteFile(globalStylesPath, []byte(newGlobalStyles), 0644)
	if err != nil {
		return fmt.Errorf("AddTailwind: writing global styles failed: %s", err.Error())
	}

	return nil
}
