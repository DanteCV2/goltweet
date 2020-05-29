package routers

import (
	"encoding/json"
	"net/http"

	"github.com/DanteCV2/goltweet/db"
	"github.com/DanteCV2/goltweet/models"
)

/*GetRelation .*/
func GetRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = IDUsuario
	t.UserRelationID = ID

	var resp models.ResponseGetRelation

	status, err := db.GetRalation(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
