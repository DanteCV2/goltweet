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
	router.HandleFunc("/viewprofile", middlew.CheckDB(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/updateProfile", middlew.CheckDB(middlew.ValidateJWT(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/readTweets", middlew.CheckDB(middlew.ValidateJWT(routers.Readtweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlew.CheckDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/setRelation", middlew.CheckDB(middlew.ValidateJWT(routers.SetRelation))).Methods("POST")
	router.HandleFunc("/deleteRelation", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/getRelation", middlew.CheckDB(middlew.ValidateJWT(routers.GetRelation))).Methods("GET")

	router.HandleFunc("/usersList", middlew.CheckDB(middlew.ValidateJWT(routers.ListUsers))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
