package api

import (
	"github.com/gofiber/fiber/v2"

	swagger "github.com/gofiber/swagger"
)

func SwaggerRoute(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
