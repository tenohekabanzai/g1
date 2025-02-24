package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Netflix struct {
	ID      bson.ObjectID `json:"_id" bson:`
	Movie   string        `json:"movie"`
	Watched bool          `json:"watched"`
}
