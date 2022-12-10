package generators

import "github.com/jchen42703/create-fullstack/internal/configs"

type FrontendTemplateGenerator interface {
	GenerateTemplate(config configs.FrontendConfig)
}

type BackendTemplateGenerator interface {
	GenerateTemplate(config configs.BackendConfig)
}

type FullstackTemplateGenerator interface {
	GenerateTemplate(config configs.FullstackConfig)
}
