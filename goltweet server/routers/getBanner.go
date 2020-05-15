package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/DanteCV2/goltweet/db"
)

/*GetBanner get img from specific profile*/
func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "cannot get id", http.StatusBadRequest)
		return
	}

	profile, err := db.FindProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
}
