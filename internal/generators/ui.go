// Creates the UI based on the frontend and UI options
package generators

import (
	"fmt"

	"github.com/jchen42703/create-fullstack/internal/configs"
)

type BaseFrontendGenerator struct {
}

func NewFrontendGenerator() *BaseFrontendGenerator {
	return &BaseFrontendGenerator{}
}

func (g *BaseFrontendGenerator) GenerateTemplate(cfg configs.FrontendConfig) {
	fmt.Println(cfg)
}
