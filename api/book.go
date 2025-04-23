package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vnFuhung2903/vcs-openapi-grpc/model"
)

var (
	books = map[string]model.Book{}
)

func BookRoute(app *fiber.App) {
	app.Post("/api/v1/books", createBook)
	app.Get("api/v1/books", readBooks)
	app.Get("/api/v1/books/:id", readBook)
	app.Put("/api/v1/books/:id", updateBook)
	app.Delete("/api/v1/books/:id", deleteBook)
}

// Create book godoc
// @Summary create a new book
// @ID createBook
// @Tags Create a new book
// @Accept	json
// @Produce json
// @Body
// @Param data body string true "New books"
// @Success 200 {object} map[string]model.Book{}
// @Router /api/v1/books [post]
func createBook(c *fiber.Ctx) error {
	newBook := new(model.Book)

	err := c.BodyParser(&newBook)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})

	}

	newBook.ID = uuid.New().String()
	books[newBook.ID] = *newBook
	c.Status(200).JSON(&fiber.Map{
		"book": newBook,
	})
	return nil
}

// Get all books godoc
// @Summary get all books in the book list
// @ID readBooks
// @Tags Get all books
// @Produce json
// @Success 200 {object} map[string]model.Book{}
// @Router /api/v1/books [get]
func readBooks(c *fiber.Ctx) error {
	c.Status(200).JSON(&fiber.Map{
		"books": books,
	})
	return nil
}

// Read an book godoc
// @Summary read an book by ID
// @ID readBook
// @Tags Get book by ID
// @Produce json
// @Param id path string true "books ID"
// @Success 200 {object} map[string]model.Book{}
// @Router /api/v1/books/{id} [get]
func readBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if book, ok := books[id]; ok {
		c.Status(200).JSON(&fiber.Map{
			"book": book,
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "book not found",
		})
	}
	return nil
}

// Update book godoc
// @Summary update an book by ID
// @ID updateBook
// @Tags Update an book
// @Accept json
// @Produce json
// @Body
// @Param id path string true "books ID"
// @Param data body string true "books ID"
// @Success 200 {object} map[string]model.Book{}
// @Router /api/v1/books/{id} [put]
func updateBook(c *fiber.Ctx) error {
	updateBook := new(model.Book)

	err := c.BodyParser(updateBook)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	id := c.Params("id")
	if book, ok := books[id]; ok {
		books[id] = book
		book.Title = updateBook.Title
		book.Author = updateBook.Author
		book.Publisher = updateBook.Publisher
		book.Page = updateBook.Page
		book.Year = updateBook.Year
		c.Status(200).JSON(&fiber.Map{
			"book": book,
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "book not found",
		})

	}
	return nil
}

// Delete an book godoc
// @Summary delete an book by ID
// @ID deleteBook
// @Tags Delete book
// @Produce json
// @Param id path string true "books ID"
// @Success 200 {object} map[string]model.Book{}
// @Router /api/v1/books/{id} [delete]
func deleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, ok := books[id]; ok {
		delete(books, id)
		c.Status(200).JSON(&fiber.Map{
			"message": "book deleted successfully",
		})
	} else {
		c.Status(404).JSON(&fiber.Map{
			"error": "book not found",
		})
	}
	return nil
}
