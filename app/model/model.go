package model

import "github.com/globalsign/mgo/bson"

type Task struct {
	ID           bson.ObjectId       `bson:"_id" json:"id"`
	Name         string              `bson:"name" json:"name"`
	CreationDate bson.MongoTimestamp `bson:"creation_date" json:"creation_date"`
	Description  string              `bson:"description" json:"description"`
}