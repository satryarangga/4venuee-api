package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Favourite struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	VenueId        int        `bson:"venue_id" json:"venue_id"`
	CustomerId  int        `bson:"customer_id" json:"customer_id"`
	DateTime int64        `bson:"unix_timestamp" json:"timestamps"`
}