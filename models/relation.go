package models

/*Relation set relation beetwen 2 users*/
type Relation struct {
	UserID         string `bson:"userid" json:"userId"`
	UserRelationID string `bson:"userrelationid" json:"userRelationId"`
}
