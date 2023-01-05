package generators

import "github.com/jchen42703/create-fullstack/internal/configs"

type UiTemplateGenerator interface {
	GenerateTemplate(config *configs.UiConfig)
}

type ApiTemplateGenerator interface {
	GenerateTemplate(config *configs.ApiConfig)
}

type FullstackTemplateGenerator interface {
	GenerateTemplate(config *configs.FullstackConfig)
}
