package controllers

import (
	"net/http"
	"oraclehr.api.com/models"
	"github.com/gorilla/mux"
	"encoding/json"
	dbDriver "oraclehr.api.com/db"
	"strconv"
	"fmt"
	"database/sql"
)

type LocationController struct{}


func (rc LocationController) GetAllLocations(dbStr string, query string) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			db := dbDriver.DBConnection(dbStr)
			var location models.Location
			locations := []models.Location{}
			rows := dbDriver.ExecuteRowsQuery(db, query)
			for rows.Next() {
				err := rows.Scan(&location.LocationId, &location.StreetAddress,
					&location.PostalCode.String, &location.City, &location.State.String, &location.CountryId)
				logFatal(err, db,  rows)
				locations = append(locations, location)
			}
			defer rows.Close()
			defer db.Close()
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(locations)
		}
}

func (rc LocationController) GetSingleLocation(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		params := mux.Vars(request)
		var location models.Location
		intParams, _ := strconv.ParseInt(params["id"], 10, 0)

		fqry := fmt.Sprintf("%s %d", query, intParams)
		row := dbDriver.ExecuteRowQuery(db, fqry)
		err := row.Scan(&location.LocationId, &location.StreetAddress,
			&location.PostalCode.String, &location.City, &location.State.String, &location.CountryId)
		if err == sql.ErrNoRows {
			logFatal(err, db, nil)
			var response = models.Response { false, "No results" }
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(response)
		} else {
			defer db.Close()
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(location)
		}
	}
}



