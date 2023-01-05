// Generates the API template.
package api

import (
	"fmt"
)

type ApiTemplateGenerator interface {
	GenerateTemplate(config *TemplateConfig)
}

type BaseApiGenerator struct {
}

func NewBaseAPIGenerator() *BaseApiGenerator {
	return &BaseApiGenerator{}
}

func (g *BaseApiGenerator) GenerateTemplate(cfg *TemplateConfig) {
	fmt.Println(cfg)
}
