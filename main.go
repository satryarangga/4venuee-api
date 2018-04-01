package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	. "github.com/satryarangga/4venuee-api/config"
	. "github.com/satryarangga/4venuee-api/dao"
	. "github.com/satryarangga/4venuee-api/controllers"
)

var config = Config{}
var dao = VisitsDAO{}

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
	r.HandleFunc("/customer-favourites/{id}/{customerid}", FindFavEndpoint).Methods("GET")
	r.HandleFunc("/customer-favourites-delete/{id}/{customerid}", DeleteCustomerFavEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}