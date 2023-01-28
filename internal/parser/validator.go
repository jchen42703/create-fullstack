package parser

import (
	"github.com/go-playground/validator/v10"
	"github.com/jchen42703/create-fullstack/core/lang"
	"github.com/jchen42703/create-fullstack/core/ui"
)

type TemplateConfig struct {
	OutputDirectoryPath string                    `yaml:"output_dir" validate:"required,file"`
	Base                string                    `yaml:"base" validate:"required"`
	Language            lang.PROGRAMMING_LANGUAGE `yaml:"lang" validate:"required,oneof=go python javascript typescript"`
	AugmentOpts         *ui.UiAugmentOptions      `yaml:"augment" validate:"required"`
}

func (c *TemplateConfig) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	// More validation (i.e. unsupported / invalid shit)
	return err
}
