package main

import (
	"example-ci/routes"

	"github.com/gofiber/fiber/v2"
)

const (
	PORT = ":3000"
)

func main() {

	config := &fiber.Config{
		AppName: "Example CI",
	}

	app := fiber.New(*config)

	routes.Init(app)

	if err := app.Listen(PORT); err != nil {
		panic(err)
	}
}
