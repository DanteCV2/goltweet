package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DanteCV2/goltweet/db"
)

/*ViewProfile get profile data*/
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	profile, err := db.FindProfile(ID)
	if err != nil {
		http.Error(w, "Errir finding profile"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
