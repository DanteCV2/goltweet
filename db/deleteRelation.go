package db

import (
	"context"
	"time"

	"github.com/DanteCV2/goltweet/models"
)

/*DeleteRelation ..*/
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("goTweet")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
