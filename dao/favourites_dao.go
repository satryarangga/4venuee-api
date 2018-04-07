package dao

import (
	"log"
	. "github.com/satryarangga/4venuee-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FavouritesDAO struct {
	Server   string
	Database string
}

const (
	COLLECTIONFAV 	= "favourites"
)

// Establish a connection to database
func (m *FavouritesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of favourites
func (m *FavouritesDAO) FindAll() ([]Favourite, error) {
	var favourites []Favourite
	err := db.C(COLLECTIONFAV).Find(bson.M{}).All(&favourites)
	return favourites, err
}

// Find a favourite by its id
func (m *FavouritesDAO) FindById(id string) (Favourite, error) {
	var favourite Favourite
	err := db.C(COLLECTIONFAV).FindId(bson.ObjectIdHex(id)).One(&favourite)
	return favourite, err
}

// Find all favourite for venue id
func (m *FavouritesDAO) FindByVenueId(id int) ([]Favourite, error) {
	var favourites []Favourite
	err := db.C(COLLECTIONFAV).Find(bson.M{"venue_id":id}).All(&favourites)
	return favourites, err
}

// Count all favourite for venue id
func (m *FavouritesDAO) CountByVenueId(id int) (total int) {
	totals, err := db.C(COLLECTIONFAV).Find(bson.M{"venue_id":id}).Count()
	if err != nil {
		return 0
	}
	return totals
}

// Check if customer has favourites venue
func (m *FavouritesDAO) FindByVenueAndCustomer(id int, customerid int) (Favourite, error) {
	var favourite Favourite
	err := db.C(COLLECTIONFAV).Find(bson.M{"venue_id":id, "customer_id":customerid}).One(&favourite)
	return favourite, err
}

// Insert a favourite into database
func (m *FavouritesDAO) Insert(favourite Favourite) error {
	err := db.C(COLLECTIONFAV).Insert(&favourite)
	return err
}

// Delete an existing favourite
func (m *FavouritesDAO) Delete(id int, customerid int) error {
	err := db.C(COLLECTIONFAV).Remove(bson.M{"venue_id":id, "customer_id":customerid})
	return err
}

// Update an existing favourite
func (m *FavouritesDAO) Update(favourite Favourite) error {
	err := db.C(COLLECTIONFAV).UpdateId(favourite.ID, &favourite)
	return err
}