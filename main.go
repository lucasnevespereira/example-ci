package main

import (
	"example-ci/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

const (
	PORT = ":3000"
)

func main() {

	engine := html.New("./templates", ".html")

	config := &fiber.Config{
		AppName: "Example CI",
		Views:   engine,
	}

	app := fiber.New(*config)

	app.Static("/", "./public")

	routes.Init(app)

	if err := app.Listen(PORT); err != nil {
		panic(err)
	}
}
