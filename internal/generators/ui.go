// Creates the UI based on the frontend and UI options
package generators

import (
	"fmt"

	"github.com/jchen42703/create-fullstack/internal/configs"
)

type BaseUiGenerator struct {
}

func NewBaseUiGenerator() *BaseUiGenerator {
	return &BaseUiGenerator{}
}

func (g *BaseUiGenerator) GenerateTemplate(cfg *configs.UiConfig) {
	fmt.Println(cfg)
}
