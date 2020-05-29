package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DanteCV2/goltweet/db"
)

/*Readtweets get tweeets from db*/
func Readtweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Cannot get ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Cannot get page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Cannot get page with value 1+", http.StatusBadRequest)
		return
	}

	pag := int64(page)
	response, correct := db.ReadTweets(ID, pag)
	if correct == false {
		http.Error(w, "Error while reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
