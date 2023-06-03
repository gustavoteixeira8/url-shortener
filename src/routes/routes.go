package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustavoteixeira8/url-shortener/src/controllers"
)

func SetupRoutes(app *fiber.App) {
	route := app.Group("/api")

	route.Post("/", controllers.CreateUrlShortController)
}
