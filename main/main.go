package main

import (
	"github.com/micro/go-config"
	"github.com/gorilla/mux"
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

	host := config.Get("host").String("localhost")
	port := config.Get("port").Int(3000)

	url := fmt.Sprintf("%s:%d", host, port)
	dbqueries := getQueries()
	router := mux.NewRouter()

	regionController := controllers.RegionController{}
	countryController := controllers.CountryController{}
	departmentController := controllers.DepartmentController{}
	employeeController := controllers.EmployeeController{}
	jobController := controllers.JobController{}
	locationController := controllers.LocationController{}
	

	router.HandleFunc("/api/regions/getAllRegions", regionController.GetAllRegions(dbStr, dbqueries["GETALLREGIONS"])).Methods("GET")
	router.HandleFunc("/api/regions/getSingleRegion/{id}", regionController.GetSingleRegion(dbStr, dbqueries["GETSINGLEREGION"])).Methods("GET")

	router.HandleFunc("/api/countries/getAllCountries", countryController.GetAllCountries(dbStr, dbqueries[""])).Methods("GET")
	router.HandleFunc("/api/countries/getSingleCountry/{id}", countryController.GetSingleCountry(dbStr, dbqueries[""])).Methods("GET")

	router.HandleFunc("/api/departments/getAllDepartments", departmentController.GetAllDepartments(dbStr, dbqueries[""])).Methods("GET")
	router.HandleFunc("/api/departments/getSingleDepartment/{id}", departmentController.GetSingleDepartment(dbStr, dbqueries[""])).Methods("GET")

	router.HandleFunc("/api/employees/getAllEmployees", employeeController.GetAllEmployees(dbStr, dbqueries[""])).Methods("GET")
	router.HandleFunc("/api/employees/getSingleEmployee/{id}", employeeController.GetSingleEmployee(dbStr, dbqueries[""])).Methods("GET")

	router.HandleFunc("/api/jobs/getAllJobs", jobController.GetAllJobs(dbStr, dbqueries[""])).Methods("GET")
	router.HandleFunc("/api/jobs/getSingleJob/{id}", jobController.GetSingleJob(dbStr, dbqueries[""])).Methods("GET")

	router.HandleFunc("/api/locations/getAllLocations", locationController.GetAllLocations(dbStr, dbqueries[""])).Methods("GET")
	router.HandleFunc("/api/locations/getSingleLocation/{id}", locationController.GetSingleLocation(dbStr, dbqueries[""])).Methods("GET")

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
