package author

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"tz/internal/handlers/store"
	"tz/internal/models"
)

func GetAuthors(s *store.AuthorDB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authors, err := s.GetAuthors()
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": "select not corrct",
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(authors)
	}
}

func GetAuthorsById(s *store.AuthorDB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve authors",
			})
		}

		authors, err := s.GetAuthorsById(idInt)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "author not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(authors)
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
				"err": "erre del",
			})
		}
		err = s.AuthorDelete(authorIdInt)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"err": "erer del",
			})
		}
		return ctx.Status(200).JSON(fiber.Map{
			"message": "Deleted",
		})
	}

}
func UpdAuthors(s *store.AuthorDB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var updAuthor models.Author

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
		err = s.UpdAuthors(authorIDint, updAuthor)
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
func CreateAuthor(s *store.AuthorDB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var author models.Author

		if err := c.BodyParser(&author); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Not right format",
			})
		}
		err := s.PostAuthor(author)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to create member",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "created",
			"author":  author,
		})
	}
}
func AuthorsBook(s *store.AuthorDB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve authors",
			})
		}

		authors, err := s.AuthorsBook(idInt)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "author not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(authors)
	}
}
