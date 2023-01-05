package configs

type GeneralConfig struct {
	FullstackCfg *FullstackConfig `yaml:"fullstack"`
	UICfg        *FrontendConfig  `yaml:"frontend"`
	APICfg       *BackendConfig   `yaml:"api"`
}
