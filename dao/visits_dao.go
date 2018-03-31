package dao

import (
	"log"

	. "github.com/satryarangga/4venuee-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type VisitsDAO struct {
	Server   string
	Database string
}

const (
	COLLECTIONS 	= "visits"
)

// Establish a connection to database
func (m *VisitsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of visits
func (m *VisitsDAO) FindAll() ([]Visit, error) {
	var visits []Visit
	err := db.C(COLLECTIONS).Find(bson.M{}).All(&visits)
	return visits, err
}

// Find a visit by its id
func (m *VisitsDAO) FindById(id string) (Visit, error) {
	var visit Visit
	err := db.C(COLLECTIONS).FindId(bson.ObjectIdHex(id)).One(&visit)
	return visit, err
}

// Find all visit for venue id
func (m *VisitsDAO) FindByVenueId(id int) ([]Visit, error) {
	var visits []Visit
	err := db.C(COLLECTIONS).Find(bson.M{"venue_id":id}).All(&visits)
	// err := db.C(COLLECTIONS).Find(bson.M{}).All(&visits)
	return visits, err
}

// Insert a visit into database
func (m *VisitsDAO) Insert(visit Visit) error {
	err := db.C(COLLECTIONS).Insert(&visit)
	return err
}

// Delete an existing visit
func (m *VisitsDAO) Delete(visit Visit) error {
	err := db.C(COLLECTIONS).Remove(&visit)
	return err
}

// Update an existing visit
func (m *VisitsDAO) Update(visit Visit) error {
	err := db.C(COLLECTIONS).UpdateId(visit.ID, &visit)
	return err
}