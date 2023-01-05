package api

import (
	"fmt"

	"github.com/jchen42703/create-fullstack/cmd/lang"
)

// Add a dockerfile and .dockerignore.
type DockerAugmenter struct {
	Lang lang.PROGRAMMING_LANGUAGE
}

func (a *DockerAugmenter) Augment() error {
	var err error
	switch a.Lang {
	case lang.Go:
		err = addGoDockerfile()
	case lang.Javascript:
		err = addJavascriptDockerfile()
	case lang.Python:
		err = addPythonDockerfile()
	case lang.Typescript:
		err = addTypescriptDockerfile()
	default:
		return fmt.Errorf("unsupported programming language '%s' for generating dockerfile", a.Lang)
	}

	if err != nil {
		return fmt.Errorf("failed to gen api dockerfile: %s", err)
	}

	return nil
}

func addGoDockerfile() error {
	return nil
}

func addJavascriptDockerfile() error {
	return nil
}

func addPythonDockerfile() error {
	return nil
}

func addTypescriptDockerfile() error {
	return nil
}
