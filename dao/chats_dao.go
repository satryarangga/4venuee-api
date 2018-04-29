package dao

import (
	"log"
	. "github.com/satryarangga/4venuee-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ChatsDAO struct {
	Server   string
	Database string
}

const (
	COLLECTIONCHAT 	= "chats"
)

// Establish a connection to database
func (m *ChatsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Insert a chat into database
func (m *ChatsDAO) Insert(chat Chat) error {
	err := db.C(COLLECTIONCHAT).Insert(&chat)
	return err
}

// Update an existing chat
func (m *ChatsDAO) Update(chat Chat) error {
	err := db.C(COLLECTIONCHAT).UpdateId(chat.ID, &chat)
	return err
}

func (m *ChatsDAO) FindOwnerChat(ownerid int) ([]Chat, error) {
	var chat []Chat

	pipe := []bson.M{
		{
	        "$match": bson.M{
	            "owner_id": ownerid,
	        },
	    },
		{
	        "$group": bson.M{
	            "_id": "$customer_id",
	            "owner_id": bson.M{
	            	"$last": "$owner_id",
	            },
	            "sender": bson.M{
	            	"$last": "$sender",
	            },
	            "timestamp": bson.M{
	            	"$last": "$timestamp",
	            },
	        },
	    },
	    {
	        "$project": bson.M{
	            "customer_id": "$_id",
	            "owner_id":1,
	            "sender":1,
	            "timestamp":1,
	        },
	    },
	}

	err := db.C(COLLECTIONCHAT).Pipe(pipe).All(&chat)
	return chat, err
}

func (m *ChatsDAO) FindCustomerChat(customerid int) ([]Chat, error) {
	var chat []Chat

	pipe := []bson.M{
		{
	        "$match": bson.M{
	            "customer_id": customerid,
	        },
	    },
		{
	        "$group": bson.M{
	            "_id": "$owner_id",
	            "customer_id": bson.M{
	            	"$last": "$customer_id",
	            },
	            "sender": bson.M{
	            	"$last": "$sender",
	            },
	            "timestamp": bson.M{
	            	"$last": "$timestamp",
	            },
	        },
	    },
	    {
	        "$project": bson.M{
	            "owner_id": "$_id",
	            "customer_id":1,
	            "sender":1,
	            "timestamp":1,
	        },
	    },
	}

	err := db.C(COLLECTIONCHAT).Pipe(pipe).All(&chat)
	return chat, err
}

// Check if customer has favourites venue
func (m *ChatsDAO) FindConversation(customerid int, ownerid int) ([]Chat, error) {
	var chat []Chat
	err := db.C(COLLECTIONCHAT).Find(bson.M{"owner_id":ownerid, "customer_id":customerid}).All(&chat)
	return chat, err
}