// Creates the UI based on the frontend and UI options
package generators

import (
	"fmt"

	"github.com/jchen42703/create-fullstack/internal/configs"
)

type BaseApiGenerator struct {
}

func NewBaseAPIGenerator() *BaseApiGenerator {
	return &BaseApiGenerator{}
}

func (g *BaseApiGenerator) GenerateTemplate(cfg *configs.ApiConfig) {
	fmt.Println(cfg)
}
