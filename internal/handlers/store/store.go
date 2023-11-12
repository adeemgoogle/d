package store

import (
	"github.com/jmoiron/sqlx"
	"tz/internal/models"
)

type AuthorDB struct {
	DB *sqlx.DB
}

func NewStore(db *sqlx.DB) *AuthorDB {
	return &AuthorDB{DB: db}

}

func (a *AuthorDB) AuthorDelete(c int) error {

	_, err := a.DB.Exec("DELETE from author where id=$1", c)
	if err != nil {
		return err
	}
	return err
}
func (a *AuthorDB) BookDelete(c int) error {
	_, err := a.DB.Exec("delete from book where id=$1", c)
	if err != nil {
		return err
	}
	return err
}
func (a *AuthorDB) MemberDelete(c int) error {
	_, err := a.DB.Exec("delete from members where id=$1", c)
	if err != nil {
		return err
	}
	return err
}

func (a *AuthorDB) GetAuthors() ([]models.Author, error) {
	var author []models.Author
	err := a.DB.Select(&author, "Select id, fullname, specialization, pseudonym from author")
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (a *AuthorDB) GetBooks() ([]models.Book, error) {
	var books []models.Book
	err := a.DB.Select(&books, "Select id, title, genre, isbn, authorid, memberid from book")
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (a *AuthorDB) GetMember() ([]models.Member, error) {
	var members []models.Member
	err := a.DB.Select(&members, "Select id, fullname from members")
	if err != nil {
		return nil, err
	}
	return members, nil
}
func (a *AuthorDB) PostAuthor(p models.Author) error {

	_, err := a.DB.Exec("INSERT INTO author VALUES ($1, $2, $3, $4)", p.ID, p.FullName, p.Pseudonym, p.Specialization)
	if err != nil {
		return err
	}
	return err
}

func (a *AuthorDB) CreateBook(p models.Book) error {

	_, err := a.DB.Exec("INSERT INTO book VALUES ($1, $2, $3, $4, $5, $6)", p.ID, p.Title, p.Genre, p.ISBN, p.AuthorID, p.MemberID)
	if err != nil {
		return err
	}
	return err
}
func (a *AuthorDB) CreateMember(p models.Member) error {
	_, err := a.DB.Exec("insert into members values ($1, $2)", p.MemberID, p.FullName)
	if err != nil {
		return err
	}
	return err
}
func (a *AuthorDB) GetAuthorsById(authorId int) (models.Author, error) {
	var author models.Author
	err := a.DB.Get(&author, "SELECT id, fullname, pseudonym, specialization FROM author where id=$1", authorId)
	if err != nil {
		panic(err)
	}
	return author, err

}
func (a *AuthorDB) GetBookById(bookId int) (models.Book, error) {
	var book models.Book
	err := a.DB.Get(&book, "SELECT id, title, genre, isbn, authorid, memberid FROM book where id=$1", bookId)
	return book, err

}
func (a *AuthorDB) GetMembersById(memberId int) (models.Member, error) {
	var member models.Member
	err := a.DB.Get(&member, "SELECT id, fullname from members where id=$1", memberId)
	return member, err

}
func (a *AuthorDB) UpdAuthors(updauthor int, upd models.Author) error {
	_, err := a.DB.Exec("UPDATE author SET fullname=$1, pseudonym=$2, specialization=$3 WHERE id=$4", upd.FullName, upd.Pseudonym, upd.Specialization, updauthor)
	return err

}
func (a *AuthorDB) UpdBook(updbook int, upd models.Book) error {
	_, err := a.DB.Exec("UPDATE book SET id=$1 ,title=$2, genre=$3, isbn=$4, authorid= $5, memberid=$6 where id=$7", upd.ID, upd.Title, upd.Genre, upd.ISBN, upd.AuthorID, upd.MemberID, updbook)
	if err != nil {
		return err
	}
	return err

}
func (a *AuthorDB) UpdMember(updbook int, upd models.Member) error {
	_, err := a.DB.Exec("UPDATE members SET id=$1 ,fullname=$2 where id=$3", upd.MemberID, upd.FullName, updbook)
	if err != nil {
		return err
	}
	return err

}
func (a *AuthorDB) AuthorsBook(bookId int) ([]models.Book, error) {
	var book []models.Book
	err := a.DB.Select(&book, "Select id, title, authorid, memberid from book where authorid=$1", bookId)
	if err != nil {
		panic(err)
	}
	return book, err
}
func (a *AuthorDB) MembersBook(bookId int) ([]models.Book, error) {
	var book []models.Book
	err := a.DB.Select(&book, "Select id, title, genre, isbn, authorid, memberid from book where authorid=$1", bookId)
	if err != nil {
		panic(err)
	}
	return book, err
}
