// Creates the UI based on the UiConfig
package ui

import (
	"fmt"
)

type UiTemplateGenerator interface {
	GenerateTemplate(config *TemplateConfig) error
}

type BaseUiGenerator struct {
}

func NewBaseGenerator() *BaseUiGenerator {
	return &BaseUiGenerator{}
}

func (g *BaseUiGenerator) GenerateTemplate(cfg *TemplateConfig) error {
	fmt.Println(cfg)
	return nil
}
