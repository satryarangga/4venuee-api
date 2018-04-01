package controllers

import (
	"net/http"
	"encoding/json"

	. "github.com/satryarangga/4venuee-api/models"
)

func CreateVisitEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var visit Visit
	if err := json.NewDecoder(r.Body).Decode(&visit); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	visit.ID = bson.NewObjectId()
    now := time.Now().Unix()
	visit.DateTime = now + 25200 // JAKARTA TIME
	if err := dao.Insert(visit); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, visit)
}

func FindVisitEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	visit, err := dao.FindByVenueId(i)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Venue ID")
		return
	}
	respondWithJson(w, http.StatusOK, visit)
}