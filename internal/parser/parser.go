package parser

import (
	"os"

	"github.com/jchen42703/create-fullstack/internal/configs"
	"gopkg.in/yaml.v2"
)

// convert yaml file to a config
func YamlToAugmentCfg(filename string) (*configs.GeneralConfig, error) {

	file, err := os.ReadFile(filename)

	if err != nil {
		return &configs.GeneralConfig{}, err
	}

	var generalConfig configs.GeneralConfig

	err = yaml.Unmarshal(file, &generalConfig)

	if err != nil {
		return &generalConfig, err
	}

	return &generalConfig, err
}
