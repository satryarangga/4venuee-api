package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/gorilla/mux"
	. "github.com/satryarangga/4venuee-api/config"
	. "github.com/satryarangga/4venuee-api/dao"
	. "github.com/satryarangga/4venuee-api/models"
)

var config = Config{}
var dao = VisitsDAO{}
var favDao = FavouritesDAO{}

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

func CreateFavEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var favourite Favourite
	if err := json.NewDecoder(r.Body).Decode(&favourite); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	favourite.ID = bson.NewObjectId()
    now := time.Now().Unix()
	favourite.DateTime = now + 25200 // JAKARTA TIME
	if err := favDao.Insert(favourite); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, favourite)
}

func FindFavEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])
	favourite, err := favDao.FindByVenueId(i)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Venue ID")
		return
	}
	respondWithJson(w, http.StatusOK, favourite)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/visits", CreateVisitEndpoint).Methods("POST")
	r.HandleFunc("/visits/{id}", FindVisitEndpoint).Methods("GET")
	r.HandleFunc("/favourites", CreateFavEndpoint).Methods("POST")
	r.HandleFunc("/favourites/{id}", FindFavEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}