// Generates the API template.
package api

import (
	"fmt"
)

type ApiTemplateGenerator interface {
	GenerateTemplate(config *TemplateConfig) error
}

type BaseApiGenerator struct {
}

func NewBaseGenerator() *BaseApiGenerator {
	return &BaseApiGenerator{}
}

func (g *BaseApiGenerator) GenerateTemplate(cfg *TemplateConfig) error {
	fmt.Println(cfg)
	return nil
}
