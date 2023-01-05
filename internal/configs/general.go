package configs

// Generic config for template generation + augmentation.
type GeneralConfig struct {
	FullstackCfg *FullstackConfig `yaml:"fullstack"`
	UiCfg        *UiConfig        `yaml:"frontend"`
	ApiCfg       *ApiConfig       `yaml:"api"`
}
