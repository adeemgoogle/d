package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"tz/internal/handlers/author"
	"tz/internal/handlers/book"
	"tz/internal/handlers/members"
	"tz/internal/handlers/store"
)

func Requests(app *fiber.App, db *sqlx.DB) {
	app.Get("/authors", author.GetAuthors(store.NewStore(db)))
	app.Get("/books", book.GetBooks(store.NewStore(db)))
	app.Get("/members", members.GetMember(store.NewStore(db)))
	app.Get("/authors/:id/book", author.AuthorsBook(store.NewStore(db)))
	app.Get("/member/:id/book", members.MembersBook(store.NewStore(db)))
	app.Get("/author/:id", author.GetAuthorsById(store.NewStore(db)))
	app.Get("/book/:id", book.GetBooksById(store.NewStore(db)))
	app.Get("/members/:id", members.GetMembersByid(store.NewStore(db)))

	app.Post("/member", members.CreateMember(store.NewStore(db)))
	app.Post("/author", author.CreateAuthor(store.NewStore(db)))
	app.Post("/bookp", book.CreateBook(store.NewStore(db)))

	app.Delete("/members/:id", members.MemberDelete(store.NewStore(db)))
	app.Delete("/book/:id", book.AuthorDelete(store.NewStore(db)))
	app.Delete("/auhtor/:id", author.AuthorDelete(store.NewStore(db)))

	app.Patch("/members/:id", members.UpdMember(store.NewStore(db)))
	app.Patch("/author/:id", author.UpdAuthors(store.NewStore(db)))
	app.Patch("/book/:id", book.UpdBook(store.NewStore(db)))
}
