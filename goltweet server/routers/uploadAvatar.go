package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/DanteCV2/goltweet/db"
	"github.com/DanteCV2/goltweet/models"
)

/*UploadAvatar set profile img*/
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileDir string = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(fileDir, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error uploading image! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error when copy image! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUsuario + "." + extension
	status, err = db.UpdateRegister(user, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error saving avatar in db! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
