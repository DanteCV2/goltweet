package routers

import (
	"net/http"

	"github.com/DanteCV2/goltweet/db"
)

/*DeleteTweet delete specific tweeet from db*/
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Cannot get ID", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Cannot delete tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
