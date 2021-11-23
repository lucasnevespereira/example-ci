package handlers

import (
	"example-ci/helpers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HomeHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "My Web App",
	})
}

func PostsHandler(c *fiber.Ctx) error {
	posts, err := helpers.DecodePosts("data/posts.json")
	if err != nil {
		fmt.Printf("getting posts: %v", err)
	}

	return c.Render("posts", fiber.Map{
		"Title": "My Posts",
		"Posts": posts,
	})
}
