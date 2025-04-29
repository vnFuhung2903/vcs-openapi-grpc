package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vnFuhung2903/vcs-openapi-grpc/api"
)

func createSwaggerDocs() {
	app := fiber.New()
	api.SwaggerRoute(app)
	api.BookRoute(app)
	app.Listen(":8080")
	app.Shutdown()
}

func main() {
	createSwaggerDocs()
}
