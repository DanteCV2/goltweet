package routers

import (
	"errors"
	"strings"

	"github.com/DanteCV2/goltweet/db"
	"github.com/DanteCV2/goltweet/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email get email value in endpoint*/
var Email string

/*IDUsuario get IDUsuario value in endpoint*/
var IDUsuario string

/*ProcessToken Get tk values*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myPass := []byte("MasterDevs")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myPass, nil
	})
	if err == nil {
		_, found, _ := db.UserExist(claims.Email)
		if found == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, found, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}
