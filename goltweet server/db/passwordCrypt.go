package db

import "golang.org/x/crypto/bcrypt"

/*PasswordCrypt Crypt pass*/
func PasswordCrypt(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
