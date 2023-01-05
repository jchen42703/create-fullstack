package configs

type HuskyOptions struct {
	CommitLint *PackageOptions `yaml:"commitlint"`
	// Might need to change these to include specific formatters/linters and their versions
	Format bool `yaml:"format"`
	Lint   bool `yaml:"lint"`
}

type UiAugmentOptions struct {
	AddScss             *PackageOptions `yaml:"scss"`
	AddTailwind         *PackageOptions `yaml:"tailwind"`
	AddStyledComponents *PackageOptions `yaml:"styled_components"`
	HuskyOpts           *HuskyOptions   `yaml:"husky"`
	AddDockerfile       bool            `yaml:"dockerfile"`
	GitOpts             *struct {
		AddIssueTemplates bool `ymal:"issue_templates"`
		AddPRTemplates    bool `yaml:"pr_templates"`
	} `yaml:"git"`

	AddCi string `yaml:"ci"`
}

type UiConfig struct {
	OutputDirectoryPath string               `yaml:"output_dir"`
	Base                string               `yaml:"base"`
	Language            PROGRAMMING_LANGUAGE `yaml:"lang"`
	AugmentOpts         *UiAugmentOptions    `yaml:"augment"`
}
