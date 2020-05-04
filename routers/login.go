package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DanteCV2/goltweet/db"
	"github.com/DanteCV2/goltweet/jwt"
	"github.com/DanteCV2/goltweet/models"
)

/*Login log user in app*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "User or Password Invalid"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	document, exist := db.LoginIntent(t.Email, t.Password)
	if exist == false {
		http.Error(w, "User or Password Invalid"+err.Error(), 400)
		return
	}

	jwtKey, err := jwt.CreateJWT(document)

	if err != nil {
		http.Error(w, "Error creating token"+err.Error(), 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//Set cookie form backend with expiraton time
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
