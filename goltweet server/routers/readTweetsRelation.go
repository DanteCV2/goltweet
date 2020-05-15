package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DanteCV2/goltweet/db"
)

/*ReadTweetsRelation ..*/
func ReadTweetsRelation(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Cannot get Page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Cannot get Page > 0", http.StatusBadRequest)
		return
	}

	response, correct := db.GetFolllowTweets(IDUsuario, page)
	if correct == false {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
