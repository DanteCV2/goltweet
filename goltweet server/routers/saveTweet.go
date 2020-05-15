package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DanteCV2/goltweet/db"
	"github.com/DanteCV2/goltweet/models"
)

/*SaveTweet save tweets in db*/
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)
	register := models.SaveTweet{
		UserID:  IDUsuario,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(register)
	if err != nil {
		http.Error(w, "Error inserting tweet"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Error saving tweet"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
