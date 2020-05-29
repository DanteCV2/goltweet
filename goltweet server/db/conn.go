package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN is the conection object to DB*/
var MongoCN = ConnDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://dantecv:268Dc487*@gotweet-ox8ek.mongodb.net/test?retryWrites=true&w=majority")

/*ConnDB allows conection to DB*/
func ConnDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conection successful with DB")
	return client
}

/*CheckConn PING to DB*/
func CheckConn() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
