package controllers


import (
	"net/http"
	"oraclehr.api.com/models"
	"encoding/json"
	dbDriver "oraclehr.api.com/db"
	"github.com/gorilla/mux"
	"fmt"
	"database/sql"
)

type CountryController struct{}


func (rc CountryController) GetAllCountries(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		var country models.Country
		countries := []models.Country{}
		rows := dbDriver.ExecuteRowsQuery(db, query)
		for rows.Next() {
			err := rows.Scan(&country.CountryId, &country.CountryName,
				&country.RegionId)
			logFatal(err, db,  rows)
			countries = append(countries, country)
		}
		defer rows.Close()
		defer db.Close()
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(countries)
	}
}

func (rc CountryController) GetSingleCountry(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		params := mux.Vars(request)
		var country models.Country

		fqry := fmt.Sprintf("%s %s", query, params["id"])
		row := dbDriver.ExecuteRowQuery(db, fqry)
		err := row.Scan(&country.CountryId, &country.CountryName, &country.RegionId)
		if err == sql.ErrNoRows {
			logFatal(err, db, nil)
			var response = models.Response { false, "No results" }
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(response)
		} else {
			defer db.Close()
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(country)
		}

	}
}
