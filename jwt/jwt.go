package jwt

import (
	"time"

	"github.com/DanteCV2/goltweet/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*CreateJWT creates a jwt*/
func CreateJWT(t models.User) (string, error) {
	myPass := []byte("MasterDevs")
	payload := jwt.MapClaims{
		"email":      t.Email,
		"name":       t.Name,
		"surname":    t.Surname,
		"birth_date": t.BirthDate,
		"biography":  t.Biography,
		"location":   t.Location,
		"website":    t.Website,
		"_id":        t.ID.Hex(),
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myPass)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
