package db

import (
	"context"
	"time"

	"github.com/DanteCV2/goltweet/models"
)

/*InsertRelation beetwen 2 users*/
func InsertRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("goTweet")
	col := db.Collection("relation")
	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
