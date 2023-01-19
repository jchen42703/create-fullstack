package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	cfsplugin "github.com/jchen42703/create-fullstack/core/plugin"
)

type ExampleAugmentor struct {
	logger hclog.Logger
}

func (a *ExampleAugmentor) Id() string {
	return "jchen42703-ExampleAugmentor-go"
}

func (a *ExampleAugmentor) Augment() error {
	a.logger.Debug("message from jchen42703-ExampleAugmentor-go.Augment")
	return nil
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	augmentor := &ExampleAugmentor{
		logger: logger,
	}

	err := cfsplugin.ServeAugmentor(augmentor, logger)
	if err != nil {
		panic(err)
	}
}
