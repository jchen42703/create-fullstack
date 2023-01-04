// Creates the UI based on the frontend and UI options
package generators

import (
	"fmt"

	"github.com/jchen42703/create-fullstack/internal/configs"
)

type BaseAPIGenerator struct {
}

func NewBaseAPIGenerator() *BaseFrontendGenerator {
	return &BaseFrontendGenerator{}
}

func (g *BaseAPIGenerator) GenerateTemplate(cfg configs.BackendConfig) {
	fmt.Println(cfg)
}
