package main

import (
	"github.com/micro/go-config"
	"github.com/gorilla/mux"
	dbDriver "oraclehr.api.com/db"
	"oraclehr.api.com/controllers"
	"net/http"
	"fmt"
	"log"
)

type Connection struct {
	ConnectionString string `json:"connectionString"`
	Queries map[string]string `json:"queries"`
}


var conn Connection

func main(){
	config.LoadFile("../config/config.json")
	dbStr := getDBConnStr()
	db := dbDriver.DBConnection(dbStr)
	host := config.Get("host").String("localhost")
	port := config.Get("port").Int(3000)

	url := fmt.Sprintf("%s:%d", host, port)
	dbqueries := getQueries()
	router := mux.NewRouter()

	regionController := controllers.RegionController{}
	router.HandleFunc("/api/getAllRegions", regionController.GetAllRegions(db, dbqueries["GETALLREGIONS"])).Methods("GET")
	//router.HandleFunc("/api/SingleRegion/{id}", regionController.GetAllRegions(db, dbqueries["GETSINGLEREGION"])).Methods("GET")
	log.Fatal(http.ListenAndServe(url, router))

}

func getQueries() map[string]string {
	config.Get("db").Scan(&conn)
	queries := conn.Queries
	return queries
}


func getDBConnStr() string {
	config.Get("db").Scan(&conn)
	connStr := conn.ConnectionString
	return connStr
}
