package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

var Db *sqlx.DB

func InitDB() *sqlx.DB {
	if Db != nil {
		return Db
	}

	dbHost := os.Getenv("HOST")
	dbPort := os.Getenv("PORT")
	dbUser := os.Getenv("USER")
	dbPassword := os.Getenv("PASSWORD")
	dbName := os.Getenv("DATABASENAME")

	dbURL := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)

	Db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
	}

	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	return Db

}
