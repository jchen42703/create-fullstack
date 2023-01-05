// Creates the UI based on the UiConfig
package ui

import (
	"fmt"
)

type UiTemplateGenerator interface {
	GenerateTemplate(config *TemplateConfig)
}

type BaseUiGenerator struct {
}

func NewBaseUiGenerator() *BaseUiGenerator {
	return &BaseUiGenerator{}
}

func (g *BaseUiGenerator) GenerateTemplate(cfg TemplateConfig) {
	fmt.Println(cfg)
}
