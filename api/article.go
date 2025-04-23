package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vnFuhung2903/vcs-openapi-grpc/model"
)

var (
	articles = map[string]model.Article{}
)

func SetupRoute() *fiber.App {
	app := *fiber.New()

	app.Post("/api/v1/articles", createArticle)
	app.Get("api/v1/articles", readArticles)
	app.Get("/api/v1/articles/:id", readArticle)
	app.Put("/api/v1/articles/:id", updateArticle)
	app.Delete("/api/v1/articles/:id", deleteArticle)

	return &app
}

// Create article godoc
// @Summary create a new article
// @ID createArticle
// @Tags Create a new article
// @Accept	json
// @Produce json
// @Body
// @Param data body string true "New articles"
// @Success 200 {object} map[string]model.Article{}
// @Router /api/v1/articles [post]
func createArticle(c *fiber.Ctx) error {
	article := new(model.Article)

	err := c.BodyParser(&article)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})

	}

	article.ID = uuid.New().String()

	articles[article.ID] = *article

	c.Status(200).JSON(&fiber.Map{
		"article": article,
	})

	return nil
}

// Get all articles godoc
// @Summary get all articles in the article list
// @ID readArticles
// @Tags Get all articles
// @Produce json
// @Success 200 {object} map[string]model.Article{}
// @Router /api/v1/articles [get]
func readArticles(c *fiber.Ctx) error {
	c.Status(200).JSON(&fiber.Map{
		"articles": articles,
	})
	return nil
}

// Read an article godoc
// @Summary read an article by ID
// @ID readArticle
// @Tags Get article by ID
// @Produce json
// @Param id path string true "articles ID"
// @Success 200 {object} map[string]model.Article{}
// @Router /api/v1/articles/{id} [get]
func readArticle(c *fiber.Ctx) error {
	id := c.Params("id")

	if article, ok := articles[id]; ok {
		c.Status(200).JSON(&fiber.Map{
			"article": article,
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "article not found",
		})
	}
	return nil
}

// Update article godoc
// @Summary update an article by ID
// @ID updateArticle
// @Tags Update an article
// @Accept json
// @Produce json
// @Body
// @Param id path string true "articles ID"
// @Param data body string true "articles ID"
// @Success 200 {object} map[string]model.Article{}
// @Router /api/v1/articles/{id} [put]
func updateArticle(c *fiber.Ctx) error {
	updateArticle := new(model.Article)

	err := c.BodyParser(updateArticle)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	id := c.Params("id")
	if article, ok := articles[id]; ok {
		article.Description = updateArticle.Description
		articles[id] = article
		article.Title = updateArticle.Title
		article.Rate = updateArticle.Rate
		c.Status(200).JSON(&fiber.Map{
			"article": article,
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "article not found",
		})

	}
	return nil
}

// Delete an article godoc
// @Summary delete an article by ID
// @ID deleteArticle
// @Tags Delete article
// @Produce json
// @Param id path string true "articles ID"
// @Success 200 {object} map[string]model.Article{}
// @Router /api/v1/articles/{id} [delete]
func deleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")

	if _, ok := articles[id]; ok {
		delete(articles, id)
		c.Status(200).JSON(&fiber.Map{
			"message": "article deleted successfully",
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "article not found",
		})
	}

	return nil
}
