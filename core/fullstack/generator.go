package fullstack

import "fmt"

type FullstackTemplateGenerator interface {
	GenerateTemplate(config *TemplateConfig) error
}

type BaseFullstackGenerator struct {
}

func NewBaseGenerator() *BaseFullstackGenerator {
	return &BaseFullstackGenerator{}
}

func (g *BaseFullstackGenerator) GenerateTemplate(cfg *TemplateConfig) error {
	fmt.Println(cfg)
	return nil
}
