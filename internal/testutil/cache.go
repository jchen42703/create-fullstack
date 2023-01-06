package testutil

import (
	"fmt"
	"os"
	"sync"

	"github.com/jchen42703/create-fullstack/internal/directory"
)

type BaseTemplateCache struct {
	mu            sync.RWMutex
	TemplatePaths map[string]string
}

func NewBaseTemplateCache() *BaseTemplateCache {
	return &BaseTemplateCache{
		TemplatePaths: map[string]string{},
	}
}

func (c *BaseTemplateCache) GetTemplateAndCopy(selectedTemplate, outputDir string) error {
	defer c.mu.Unlock()
	c.mu.Lock()
	cachedTemplatePath, ok := c.TemplatePaths[selectedTemplate]
	if !ok {
		return fmt.Errorf("template %s not found in cache", selectedTemplate)
	}

	// Copy to output dir
	if err := os.MkdirAll(outputDir, directory.READ_WRITE_EXEC_PERM); err != nil {
		return fmt.Errorf("failed to create outputDir: %s", err)
	}

	if err := directory.CopyDirectory(cachedTemplatePath, outputDir); err != nil {
		return fmt.Errorf("GetTemplateAndCopy: failed to copy to outputDir: %s\n", err)
	}

	return nil
}

func (c *BaseTemplateCache) AddTemplate(templateName, templatePath string) {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.TemplatePaths[templateName] = templatePath
}
