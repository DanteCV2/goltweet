package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/DanteCV2/goltweet/middlew"
	"github.com/DanteCV2/goltweet/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manage set PORT, handler and listen PORT*/
func Manage() {
	router := mux.NewRouter()
	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
