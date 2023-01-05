package frameworks

// Not sure if this is needed.
type ApiFrameworks string

const (
	// JS
	Express ApiFrameworks = "express"
	Fastify ApiFrameworks = "fastify"
	Nestjs  ApiFrameworks = "nestjs"

	// Go
	Echo ApiFrameworks = "echo"
	Gin  ApiFrameworks = "gin"
	Chi  ApiFrameworks = "chi"

	// Python
	FastApi ApiFrameworks = "fastapi"

	// Java
	Spring ApiFrameworks = "springboot"

	// Other
	Other ApiFrameworks = "other"
)
