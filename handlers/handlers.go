package handlers

import (
	"example-ci/models"

	"github.com/gofiber/fiber/v2"
)

var samplePosts = &[]models.Post{
	{
		Id:          1,
		Content:     "This is a sample post",
		PublishedAt: "2021-06-09",
	},
	{
		Id:          2,
		Content:     "This is another post",
		PublishedAt: "2021-08-09",
	},
}

func HomeHandler(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Welcome Home!",
	})
}

func PostsHandler(c *fiber.Ctx) error {
	return c.JSON(samplePosts)
}
