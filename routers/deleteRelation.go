package routers

import (
	"net/http"

	"github.com/DanteCV2/goltweet/db"
	"github.com/DanteCV2/goltweet/models"
)

/*DeleteRelation ..*/
func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relation
	t.UserID = IDUsuario
	t.UserRelationID = ID

	status, err := db.DeleteRelation(t)

	if err != nil {
		http.Error(w, "Error deleting relation"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Error deleting relation"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
