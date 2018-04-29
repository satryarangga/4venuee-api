package controllers

import (
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"time"
	"strconv"
	. "github.com/satryarangga/4venuee-api/models"
	. "github.com/satryarangga/4venuee-api/dao"
	. "github.com/satryarangga/4venuee-api/helpers"
)


var chatDao = ChatsDAO{}

func CreateChatEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var chat Chat
	if err := json.NewDecoder(r.Body).Decode(&chat); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	chat.ID = bson.NewObjectId()
    now := time.Now().Unix()
	chat.Timestamp = now + 25200 // JAKARTA TIME
	if err := chatDao.Insert(chat); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, chat, "Success to create chat")
}

func GetOwnerChatsEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ownerid, err := strconv.Atoi(params["ownerid"])
	chat, err := chatDao.FindOwnerChat(ownerid)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "No chat for this owner")
		return
	}
	RespondWithJson(w, http.StatusOK, chat, "Success")
}

func GetCustomerChatsEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerid, err := strconv.Atoi(params["customerid"])
	chat, err := chatDao.FindCustomerChat(customerid)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "No chat for this customer")
		return
	}
	RespondWithJson(w, http.StatusOK, chat, "Success")
}

func GetConversations(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerid, err := strconv.Atoi(params["customerid"])
	ownerid, err := strconv.Atoi(params["ownerid"])
	chat, err := chatDao.FindConversation(customerid, ownerid)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "No conversation")
		return
	}
	RespondWithJson(w, http.StatusOK, chat, "Success")
}