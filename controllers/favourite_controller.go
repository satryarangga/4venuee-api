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


var favDao = FavouritesDAO{}

func CreateFavEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var favourite Favourite
	if err := json.NewDecoder(r.Body).Decode(&favourite); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	favourite.ID = bson.NewObjectId()
    now := time.Now().Unix()
	favourite.DateTime = now + 25200 // JAKARTA TIME
	if err := favDao.Insert(favourite); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, favourite, "Success to create favourite")
}

func FindFavEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	favourite, err := favDao.FindByVenueId(i)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Venue ID")
		return
	}
	RespondWithJson(w, http.StatusOK, favourite, "Success")
}

func CountFavEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Venue ID")
		return
	}
	favourite := favDao.CountByVenueId(i)
	RespondWithJson(w, http.StatusOK, favourite, "Success")
}

func CheckCustomerFavEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	c, err := strconv.Atoi(params["customerid"])
	favourite, err := favDao.FindByVenueAndCustomer(i, c)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "No Favourite Venue for Customer")
		return
	}
	RespondWithJson(w, http.StatusOK, favourite, "Success")
}

func DeleteCustomerFavEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	c, err := strconv.Atoi(params["customerid"])
	err = favDao.Delete(i, c)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "No data found")
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{}, "Success to remove favourite")
}