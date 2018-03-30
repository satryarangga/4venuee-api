package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Visit struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	VenueId        int        `bson:"venue_id" json:"venue_id"`
	CustomerId  int        `bson:"customer_id" json:"customer_id"`
	DateTime bson.MongoTimestamp        `bson:"_timestamp" json:"timestamps,omitempty"`
}