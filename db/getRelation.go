package db

import (
	"context"
	"time"

	"github.com/DanteCV2/goltweet/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*GetRalation .*/
func GetRalation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("goTweet")
	col := db.Collection("relation")

	condition := bson.M{
		"userid":         t.UserID,
		"Userrelationid": t.UserRelationID,
	}

	var result models.Relation
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil
}
