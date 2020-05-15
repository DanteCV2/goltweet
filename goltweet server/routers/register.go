package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DanteCV2/goltweet/db"
	"github.com/DanteCV2/goltweet/models"
)

/*Register creates users in DB*/
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error with recived data"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password must be at least 6 char", 400)
		return
	}

	_, found, _ := db.UserExist(t.Email)

	if found == true {
		http.Error(w, "This email is already registered", 400)
		return
	}

	_, status, err := db.InsertUser(t)

	if err != nil {
		http.Error(w, "Error while register User"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Cant register user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
