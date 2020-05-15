package models

/*Tweet Tweet content*/
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
