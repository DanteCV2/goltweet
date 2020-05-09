package routers

import (
	"net/http"

	"github.com/DanteCV2/goltweet/db"
	"github.com/DanteCV2/goltweet/models"
)

/*SetRelation ..*/
func SetRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Connot get id", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUsuario
	t.UserRelationID = ID

	status, err := db.InsertRelation(t)

	if err != nil {
		http.Error(w, "Error with relation"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Error with relation"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
