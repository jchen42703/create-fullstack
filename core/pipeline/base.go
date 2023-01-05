package pipeline

import (
	"github.com/jchen42703/create-fullstack/internal/parser"
)

type TemplatePipeline interface {
	Read(filePath string) (*parser.GeneralConfig, error)
	Generate(*parser.GeneralConfig) error
}
