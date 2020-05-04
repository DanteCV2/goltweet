package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Claim process the jwt*/
type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:":id json: _id,imitempty"`
	jwt.StandardClaims
}
