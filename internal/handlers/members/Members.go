package members

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"tz/internal/handlers/store"
	"tz/internal/models"
)

func GetMember(s *store.AuthorDB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		members, err := s.GetMember()
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": "select not corrct",
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(members)
	}
}
func GetMembersByid(s *store.AuthorDB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "error",
			})
		}
		books, err := s.GetMembersById(idInt)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "book not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(books)
	}
}
func MemberDelete(s *store.AuthorDB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorId := ctx.Params("id")
		authorIdInt, err := strconv.Atoi(authorId)

		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": "delete id err",
			})
		}
		err = s.MemberDelete(authorIdInt)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"err": "err del",
			})
		}
		return ctx.Status(200).JSON(fiber.Map{
			"message": "Deleted",
		})
	}

}
func CreateMember(s *store.AuthorDB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var member models.Member
		if err := c.BodyParser(&member); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Not right format",
			})
		}

		err := s.CreateMember(member)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to create member",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "created",
		})
	}
}
func UpdMember(s *store.AuthorDB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var updAuthor models.Member

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
		err = s.UpdMember(authorIDint, updAuthor)
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
func MembersBook(s *store.AuthorDB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve authors",
			})
		}

		authors, err := s.MembersBook(idInt)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "author not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(authors)
	}
}
