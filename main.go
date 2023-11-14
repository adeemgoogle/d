package main

import (
	"log"
	"tz/internal/server"
)

func main() {

	err := server.Start()
	if err != nil {
		log.Fatal("App couldnt start")
	}
}
