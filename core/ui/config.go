package ui

import (
	"github.com/jchen42703/create-fullstack/core/lang"
	"github.com/jchen42703/create-fullstack/core/opts"
)

type HuskyOptions struct {
	CommitLint *opts.PackageOptions `yaml:"commitlint"`
	// Might need to change these to include specific formatters/linters and their versions
	Format bool `yaml:"format" validate:"required,boolean"`
	Lint   bool `yaml:"lint" validate:"required,boolean"`
}

type UiAugmentOptions struct {
	AddScss             *opts.PackageOptions `yaml:"scss" validate:"required"`
	AddTailwind         *opts.PackageOptions `yaml:"tailwind" validate:"required"`
	AddStyledComponents *opts.PackageOptions `yaml:"styled_components" validate:"required"`
	HuskyOpts           *HuskyOptions        `yaml:"husky" validate:"required"`
	AddDockerfile       bool                 `yaml:"dockerfile" validate:"required,boolean"`
	GitOpts             *struct {
		AddIssueTemplates bool `yaml:"issue_templates" validate:"required,boolean"`
		AddPRTemplates    bool `yaml:"pr_templates" validate:"required,boolean"`
	} `yaml:"git" validate:"required"`

	AddCi string `yaml:"ci" validate:"required,oneof=circleci travisci jenkins git_workflows"`
}

type TemplateConfig struct {
	OutputDirectoryPath string                    `yaml:"output_dir"`
	Base                string                    `yaml:"base"`
	Language            lang.PROGRAMMING_LANGUAGE `yaml:"lang"`
	AugmentOpts         *UiAugmentOptions         `yaml:"augment"`
}
