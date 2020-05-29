package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DanteCV2/goltweet/db"
)

/*ListUsers ..*/
func ListUsers(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Page needs > 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)
	result, status := db.GetAllUsers(IDUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error getting users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
