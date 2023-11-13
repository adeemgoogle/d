package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"tz/internal/handlers"
	"tz/internal/store"
)

func Requests(app *fiber.App, db *sqlx.DB) {
	api := app.Group("/api")
	Author := api.Group("/Auth")
	Author.Get("/authors", handlers.GetAuthors(store.NewStore(db)))
	Author.Get("/author/:id", handlers.GetAuthorsById(store.NewStore(db)))
	Author.Get("/authors/:id/book", handlers.AuthorsBook(store.NewStore(db)))
	Author.Delete("/author/:id", handlers.AuthorDelete(store.NewStore(db)))
	Author.Post("/author", handlers.CreateAuthor(store.NewStore(db)))
	Author.Patch("/author/:id", handlers.UpdAuthors(store.NewStore(db)))

	Books := api.Group("/book")
	Books.Get("/books", handlers.GetBooks(store.NewStore(db)))
	Books.Get("/book/:id", handlers.GetBooksById(store.NewStore(db)))
	Books.Post("/bookp", handlers.CreateBook(store.NewStore(db)))
	Books.Delete("/book/:id", handlers.BookDelete(store.NewStore(db)))
	Books.Patch("/book/:id", handlers.UpdBook(store.NewStore(db)))

	Member := api.Group("/member")
	Member.Get("/members", handlers.GetMember(store.NewStore(db)))
	Member.Get("/member/:id/book", handlers.MembersBook(store.NewStore(db)))
	Member.Get("/members/:id", handlers.GetMembersByid(store.NewStore(db)))
	Member.Post("/member", handlers.CreateMember(store.NewStore(db)))
	Member.Delete("/members/:id", handlers.MemberDelete(store.NewStore(db)))
	Member.Patch("/members/:id", handlers.UpdMember(store.NewStore(db)))
}
