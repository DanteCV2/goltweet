package db

import (
	"context"
	"fmt"
	"time"

	"github.com/DanteCV2/goltweet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*GetAllUsers ..*/
func GetAllUsers(ID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("goTweet")
	col := db.Collection("user")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, include bool

	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		found, err = GetRalation(r)
		if tipo == "new" && found == false {
			include = true
		}

		if tipo == "follow" && found == true {
			include = true
		}

		if r.UserRelationID == ID {
			include = false
		}

		if include == true {
			s.Password = ""
			s.Biography = ""
			s.Website = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)
	return results, true
}
