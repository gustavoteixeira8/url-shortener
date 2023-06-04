package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustavoteixeira8/url-shortener/src/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/", controllers.CreateUrlShortController)
	app.Get("/:name", controllers.RedirectURLShortController)
	app.Get("/:name/details", controllers.GetURLShortDetailsController)
}
