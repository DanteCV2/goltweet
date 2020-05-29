package main

import (
	"log"

	"github.com/DanteCV2/goltweet/db"
	"github.com/DanteCV2/goltweet/handlers"
)

func main() {
	if db.CheckConn() == 0 {
		log.Fatal("Cant connect with db")
		return
	}
	handlers.Manage()
}
