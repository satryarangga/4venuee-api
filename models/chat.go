package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Chat struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	OwnerId        int        `bson:"owner_id" json:"owner_id"`
	CustomerId  int        `bson:"customer_id" json:"customer_id"`
	Sender  string        `bson:"sender" json:"sender"`
	Message  string        `bson:"message" json:"message"`
	Timestamp  int64        `bson:"timestamp" json:"timestamp"`
}