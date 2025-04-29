package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vnFuhung2903/vcs-openapi-grpc/api"
	"github.com/vnFuhung2903/vcs-openapi-grpc/client"
)

func createSwaggerDocs() {
	app := fiber.New()
	api.SwaggerRoute(app)
	api.BookRoute(app)
	app.Listen(":8080")
	app.Shutdown()
}

func checkRPC() {
	client := client.NewClient()
	log.Printf("%v", client)
}

func main() {
	createSwaggerDocs()
	checkRPC()
}
