package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"tz/internal/handlers/store"
	"tz/internal/models"
)

type Handlers struct {
	Store *store.NewStoreDB
}

func NewHandlers(s *store.NewStoreDB) *Handlers {
	return &Handlers{
		Store: s,
	}
}

func (h *Handlers) GetAuthors() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authors, err := h.Store.GetAuthors()
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": "select not corrct",
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(authors)
	}
}

func (h *Handlers) GetAuthorsById() fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve authors",
			})
		}

		authors, err := h.Store.GetAuthorsById(idInt)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "author not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(authors)
	}

}

func (h *Handlers) AuthorDelete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorId := ctx.Params("id")
		authorIdInt, err := strconv.Atoi(authorId)

		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": "delete id err",
			})
		}
		err = h.Store.BookDelete(authorIdInt)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"err": "erre del",
			})
		}
		err = h.Store.AuthorDelete(authorIdInt)
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

func (h *Handlers) UpdAuthors() fiber.Handler {
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
		err = h.Store.UpdAuthors(authorIDint, updAuthor)
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

func (h *Handlers) CreateAuthor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var author models.Author

		if err := c.BodyParser(&author); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Not right format",
			})
		}
		err := h.Store.PostAuthor(author)
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

func (h *Handlers) AuthorsBook() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve authors",
			})
		}

		authors, err := h.Store.AuthorsBook(idInt)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "author not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(authors)
	}
}
