package models

type Member struct {
	MemberID int    `json:"id" db:"id"`
	FullName string `json:"fullname" db:"fullname"`
}
