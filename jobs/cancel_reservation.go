package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	. "github.com/satryarangga/4venuee-api/config"
	_ "github.com/go-sql-driver/mysql"
)

var config = Config{}

func init() {
	config.Read()
}

func main() {
	maximumUnpaidDays := 1
	reservationStatusCancelled := 3
	currentTime := time.Now().AddDate(0, 0, -maximumUnpaidDays).Format("2006-01-02 12:00:00")
	connection := fmt.Sprintf("%s:%s@/%s", config.MysqlUser, config.MysqlPassword, config.MysqlDatabase)
	db, err := sql.Open("mysql", connection)
	var query = "SELECT reservation_id from reservation_payment where status = 0 and created_at < ?"
	rows, err := db.Query(query, currentTime)
	if err != nil {
	        log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
	        var reservation_id string
	        if err := rows.Scan(&reservation_id); err != nil {
	                log.Fatal(err)
	        }
	        db.Query("update reservation_payment set status = ? where reservation_id = ?", reservationStatusCancelled, reservation_id)
	        fmt.Printf("Reservation id %s is cancelled \n", reservation_id)
	}
	if err := rows.Err(); err != nil {
	        log.Fatal(err)
	}
}