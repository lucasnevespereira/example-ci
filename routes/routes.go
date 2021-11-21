package routes

import (
	"example-ci/handlers"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Get("/", handlers.HomeHandler)
	app.Get("/posts", handlers.PostsHandler)
}
