package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"tz/pkg/config"
)

var DBUrl *sqlx.DB

func InitDB() (*sqlx.DB, error) {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	DBUrl, err = sqlx.Open("postgres", conf.DB)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return DBUrl, nil

}
