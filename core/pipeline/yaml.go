package pipeline

import (
	"fmt"

	"github.com/jchen42703/create-fullstack/core/api"
	"github.com/jchen42703/create-fullstack/core/fullstack"
	"github.com/jchen42703/create-fullstack/core/ui"
	"github.com/jchen42703/create-fullstack/internal/parser"
)

// Pipeline to read and run the template generation based on a YAML config file.
type YamlPipeline struct {
	ConfigPath         string
	FullstackGenerator fullstack.TemplateGenerator
	UiGenerator        ui.TemplateGenerator
	ApiGenerator       api.TemplateGenerator
}

// Should read and validate the template config.
func (p *YamlPipeline) Read(configPath string) (*parser.GeneralConfig, error) {
	return parser.YamlToTemplateCfg(configPath)
}

// 1. Creates and parses the base template.
// 2. Augments the base template according to the config.
func (p *YamlPipeline) Generate(cfg *parser.GeneralConfig) error {
	if cfg.FullstackCfg != nil {
		// Don't augment ui/api if fullstack exists
		return p.FullstackGenerator.GenerateTemplate(cfg.FullstackCfg)
	}

	// Allow users to generate and augment templates separately.
	if cfg.UiCfg != nil {
		err := p.UiGenerator.GenerateTemplate(cfg.UiCfg)
		if err != nil {
			return fmt.Errorf("failed to generate ui: %s", err)
		}
	}

	if cfg.ApiCfg != nil {
		err := p.ApiGenerator.GenerateTemplate(cfg.ApiCfg)
		if err != nil {
			return fmt.Errorf("failed to generate ui: %s", err)
		}
	}

	return nil
}
