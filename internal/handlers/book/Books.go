package book

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"tz/internal/handlers/store"
	"tz/internal/models"
)

func GetBooks(s *store.AuthorDB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		books, err := s.GetBooks()
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": "select not work",
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(books)
	}

}
func GetBooksById(s *store.AuthorDB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "error",
			})
		}
		books, err := s.GetBookById(idInt)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "book not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(books)
	}
}
func AuthorDelete(s *store.AuthorDB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorId := ctx.Params("id")
		authorIdInt, err := strconv.Atoi(authorId)

		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": "delete id err",
			})
		}
		err = s.BookDelete(authorIdInt)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"err": "errd del",
			})
		}
		err = s.MemberDelete(authorIdInt)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"err": "erro del",
			})
		}
		err = s.AuthorDelete(authorIdInt)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"err": "errg del",
			})
		}

		return ctx.Status(200).JSON(fiber.Map{
			"message": "Deleted",
		})
	}

}
func UpdBook(s *store.AuthorDB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var updAuthor models.Book

		authorID := c.Params("id")
		authorIDint, err := strconv.Atoi(authorID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error":   true,
				"message": "failed to upd author",
			})
		}

		if err := c.BodyParser(&updAuthor); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error":   true,
				"message": "invalid json req",
			})
		}
		err = s.UpdBook(authorIDint, updAuthor)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"err": "auth not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"err": "gdeto err",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "good",
		})
	}
}
func CreateBook(s *store.AuthorDB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var book models.Book

		if err := c.BodyParser(&book); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Not right format",
			})
		}

		err := s.CreateBook(book)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to create book",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "created",
		})
	}
}
