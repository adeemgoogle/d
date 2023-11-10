package models

type Book struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Genre    string `json:"genre" db:"genre"`
	ISBN     int    `json:"isbn" db:"isbn"`
	AuthorID int    `json:"authorid" db:"authorid"`
	MemberID int    `json:"memberID" db:"memberid"`
}
