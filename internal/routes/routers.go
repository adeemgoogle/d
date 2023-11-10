package routes

import (
	"github.com/gofiber/fiber/v2"
	"tz/internal/handlers/author"
	"tz/internal/handlers/book"
	"tz/internal/handlers/members"
	"tz/internal/handlers/store"
	"tz/pkg/database"
)

func Requests(app *fiber.App) {
	app.Get("/authors", author.GetAuthors(store.NewStore(database.Db)))
	app.Get("/books", book.GetBooks(store.NewStore(database.Db)))
	app.Get("/members", members.GetMember(store.NewStore(database.Db)))
	app.Get("/authors/:id/book", author.AuthorsBook(store.NewStore(database.Db)))
	app.Get("/member/:id/book", members.MembersBook(store.NewStore(database.Db)))
	app.Get("/author/:id", author.GetAuthorsById(store.NewStore(database.Db)))
	app.Get("/book/:id", book.GetBooksById(store.NewStore(database.Db)))
	app.Get("/members/:id", members.GetMembersByid(store.NewStore(database.Db)))

	app.Post("/member", members.CreateMember(store.NewStore(database.Db)))
	app.Post("/author", author.CreateAuthor(store.NewStore(database.Db)))
	app.Post("/bookp", book.CreateBook(store.NewStore(database.Db)))

	app.Delete("/members/:id", members.MemberDelete(store.NewStore(database.Db)))
	app.Delete("/book/:id", book.AuthorDelete(store.NewStore(database.Db)))
	app.Delete("/auhtor/:id", author.AuthorDelete(store.NewStore(database.Db)))

	app.Patch("/members/:id", members.UpdMember(store.NewStore(database.Db)))
	app.Patch("/author/:id", author.UpdAuthors(store.NewStore(database.Db)))
	app.Patch("/book/:id", book.UpdBook(store.NewStore(database.Db)))
}
