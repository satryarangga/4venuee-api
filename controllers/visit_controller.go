package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"time"
	"strconv"
	. "github.com/satryarangga/4venuee-api/models"
	. "github.com/satryarangga/4venuee-api/dao"
	. "github.com/satryarangga/4venuee-api/helpers"
)

var dao = VisitsDAO{}

func CreateVisitEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var visit Visit
	if err := json.NewDecoder(r.Body).Decode(&visit); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	visit.ID = bson.NewObjectId()
    now := time.Now().Unix()
	visit.DateTime = now + 25200 // JAKARTA TIME
	if err := dao.Insert(visit); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, visit, "Success to create visit")
}

func FindVisitEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	visit, err := dao.FindByVenueId(i)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Venue ID")
		return
	}
	RespondWithJson(w, http.StatusOK, visit, "Success")
}

func CountVisitEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Venue ID")
		return
	}
	visit:= dao.CountByVenueId(i)
	RespondWithJson(w, http.StatusOK, visit, "Success")
}