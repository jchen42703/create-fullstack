package fullstack

import "fmt"

type FullstackTemplateGenerator interface {
	GenerateTemplate(config *TemplateConfig)
}

type BaseFullstackGenerator struct {
}

func NewBaseFullstackGenerator() *BaseFullstackGenerator {
	return &BaseFullstackGenerator{}
}

func (g *BaseFullstackGenerator) GenerateTemplate(cfg TemplateConfig) {
	fmt.Println(cfg)
}
