package db

import (
	"github.com/DanteCV2/goltweet/models"
	"golang.org/x/crypto/bcrypt"
)

/*LoginIntent login*/
func LoginIntent(email string, password string) (models.User, bool) {
	usu, found, _ := UserExist(email)
	if found == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
