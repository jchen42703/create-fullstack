package parser

import (
	"os"

	"github.com/jchen42703/create-fullstack/core/api"
	"github.com/jchen42703/create-fullstack/core/fullstack"
	"github.com/jchen42703/create-fullstack/core/ui"
	"gopkg.in/yaml.v2"
)

// Generic config for template generation + augmentation.
type GeneralConfig struct {
	FullstackCfg *fullstack.TemplateConfig `yaml:"fullstack"`
	UiCfg        *ui.TemplateConfig        `yaml:"frontend"`
	ApiCfg       *api.TemplateConfig       `yaml:"api"`
}

// convert yaml file to a config
func YamlToTemplateCfg(filename string) (*GeneralConfig, error) {

	file, err := os.ReadFile(filename)

	if err != nil {
		return &GeneralConfig{}, err
	}

	var generalConfig GeneralConfig

	err = yaml.Unmarshal(file, &generalConfig)

	if err != nil {
		return &generalConfig, err
	}

	return &generalConfig, err
}
