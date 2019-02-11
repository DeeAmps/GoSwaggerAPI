package main

// author Daniel Bennin @danyelamps.db@gmail.com

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
	fmt.Printf("API running on %s/api\n", url)
	fmt.Printf("Swagger Documentation running on %s/api-docs\n", url)
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
	router.HandleFunc("/api/regions/addNewRegion", regionController.AddNewRegion(dbStr, dbqueries["INSERT_NEW_REGION"], dbqueries["GET_LAST_REGION_ID"])).Methods("POST")
	router.HandleFunc("/api/regions/removeRegion", regionController.RemoveRegion(dbStr, dbqueries["REMOVE_REGION"])).Methods("DELETE")
	router.HandleFunc("/api/regions/updateRegion/{id}", regionController.UpdateRegion(dbStr, dbqueries["UPDATE_REGION"])).Methods("PUT")



	router.HandleFunc("/api/countries/getAllCountries", countryController.GetAllCountries(dbStr, dbqueries["GET_COUNTRIES"])).Methods("GET")
	router.HandleFunc("/api/countries/getSingleCountry/{id}", countryController.GetSingleCountry(dbStr, dbqueries["GET_COUNTRY"])).Methods("GET")

	router.HandleFunc("/api/departments/getAllDepartments", departmentController.GetAllDepartments(dbStr, dbqueries["GET_DEPARTMENTS"])).Methods("GET")
	router.HandleFunc("/api/departments/getSingleDepartment/{id}", departmentController.GetSingleDepartment(dbStr, dbqueries["GET_DEPARTMENT"])).Methods("GET")

	router.HandleFunc("/api/employees/getAllEmployees", employeeController.GetAllEmployees(dbStr, dbqueries["GET_EMPLOYEES"])).Methods("GET")
	router.HandleFunc("/api/employees/getSingleEmployee/{id}", employeeController.GetSingleEmployee(dbStr, dbqueries["GET_EMPLOYEE"])).Methods("GET")

	router.HandleFunc("/api/jobs/getAllJobs", jobController.GetAllJobs(dbStr, dbqueries["GET_JOBS"])).Methods("GET")
	router.HandleFunc("/api/jobs/getSingleJob/{id}", jobController.GetSingleJob(dbStr, dbqueries["GET_JOB"])).Methods("GET")

	router.HandleFunc("/api/locations/getAllLocations", locationController.GetAllLocations(dbStr, dbqueries["GET_LOCATIONS"])).Methods("GET")
	router.HandleFunc("/api/locations/getSingleLocation/{id}", locationController.GetSingleLocation(dbStr, dbqueries[""])).Methods("GET")



	sh := http.StripPrefix("/api-docs/", http.FileServer(http.Dir("../swagger/")))
	router.PathPrefix("/api-docs/").Handler(sh)
	router.PathPrefix("/api-docs").Handler(sh)
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
