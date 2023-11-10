package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"

	"tz/pkg/config"
)

var Db *sqlx.DB

func init() {
	config := config.Get()
	var err error
	Db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable"), config.UserName, config.Password, config.DataBaseName, config.Host, config.Port)

	if err != nil {
		log.Fatal(err)
	}
	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
