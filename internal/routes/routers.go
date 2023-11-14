package server

import (
	"github.com/gofiber/fiber/v2"
	"tz/internal/handlers"
)

func Requests(h *handlers.Handlers) {

	Author = fiber.New()
	Author.Get("/authors", h.GetAuthors)
	Author.Get("/author/:id", h.GetAuthorsById())
	Author.Get("/authors/:id/book", h.AuthorsBook())
	Author.Delete("/author/:id", h.AuthorDelete())
	Author.Post("/author", h.CreateAuthor())
	Author.Patch("/author/:id", h.UpdAuthors())

	//Books := api.Group("/book")
	//Books.Get("/books", handlers.GetBooks(store.NewStore(db)))
	//Books.Get("/book/:id", handlers.GetBooksById(store.NewStore(db)))
	//Books.Post("/bookp", handlers.CreateBook(store.NewStore(db)))
	//Books.Delete("/book/:id", handlers.BookDelete(store.NewStore(db)))
	//Books.Patch("/book/:id", handlers.UpdBook(store.NewStore(db)))
	//
	//Member := api.Group("/member")
	//Member.Get("/members", handlers.GetMember(store.NewStore(db)))
	//Member.Get("/member/:id/book", handlers.MembersBook(store.NewStore(db)))
	//Member.Get("/members/:id", handlers.GetMembersByid(store.NewStore(db)))
	//Member.Post("/member", handlers.CreateMember(store.NewStore(db)))
	//Member.Delete("/members/:id", handlers.MemberDelete(store.NewStore(db)))
	//Member.Patch("/members/:id", handlers.UpdMember(store.NewStore(db)))

}
