package db

import (
	"context"
	"time"

	"github.com/DanteCV2/goltweet/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*UpdateRegister modify profile info*/
func UpdateRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("goTweet")
	col := db.Collection("user")

	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}

	if len(u.Surname) > 0 {
		register["surname"] = u.Surname
	}

	register["surname"] = u.BirthDate

	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}

	if len(u.Biography) > 0 {
		register["biography"] = u.Surname
	}

	if len(u.Location) > 0 {
		register["location"] = u.Location
	}

	if len(u.Website) > 0 {
		register["website"] = u.Website
	}

	updtString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}