package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DanteCV2/goltweet/db"
	"github.com/DanteCV2/goltweet/models"
)

/*UpdateProfile update profile info*/
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect Data"+err.Error(), 400)
		return
	}

	var status bool
	status, err = db.UpdateRegister(t, IDUsuario)
	if err != nil {
		http.Error(w, "Error updating info, please try again"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Cant update profile"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
