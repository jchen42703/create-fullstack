package ui

import (
	"github.com/jchen42703/create-fullstack/cmd/lang"
	"github.com/jchen42703/create-fullstack/cmd/opts"
)

type HuskyOptions struct {
	CommitLint *opts.PackageOptions `yaml:"commitlint"`
	// Might need to change these to include specific formatters/linters and their versions
	Format bool `yaml:"format"`
	Lint   bool `yaml:"lint"`
}

type UiAugmentOptions struct {
	AddScss             *opts.PackageOptions `yaml:"scss"`
	AddTailwind         *opts.PackageOptions `yaml:"tailwind"`
	AddStyledComponents *opts.PackageOptions `yaml:"styled_components"`
	HuskyOpts           *HuskyOptions        `yaml:"husky"`
	AddDockerfile       bool                 `yaml:"dockerfile"`
	GitOpts             *struct {
		AddIssueTemplates bool `ymal:"issue_templates"`
		AddPRTemplates    bool `yaml:"pr_templates"`
	} `yaml:"git"`

	AddCi string `yaml:"ci"`
}

type TemplateConfig struct {
	OutputDirectoryPath string                    `yaml:"output_dir"`
	Base                string                    `yaml:"base"`
	Language            lang.PROGRAMMING_LANGUAGE `yaml:"lang"`
	AugmentOpts         *UiAugmentOptions         `yaml:"augment"`
}
