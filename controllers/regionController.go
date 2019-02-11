package controllers

import (
	"database/sql"
	"net/http"
	"oraclehr.api.com/models"
	"fmt"
	"encoding/json"
	dbDriver "oraclehr.api.com/db"
	"github.com/gorilla/mux"
	"strconv"
	"oraclehr.api.com/helpers"
)

type RegionController struct{}

func logFatal(err error, db *sql.DB, rows *sql.Rows){
	if err != nil {
		fmt.Println(err)
		if rows != nil {
			rows.Close()
		}
		db.Close()
		//panic(err)
	}
	return
}

func (rc RegionController) GetAllRegions(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		var region models.Region
		regions := []models.Region{}
		rows := dbDriver.ExecuteRowsQuery(db, query)
		for rows.Next() {
			err := rows.Scan(&region.RegionId, &region.RegionName)
			logFatal(err, db,  rows)
			regions = append(regions, region)
		}
		defer rows.Close()
		defer db.Close()
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(regions)
	}
}

func (rc RegionController) GetSingleRegion(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		params := mux.Vars(request)
		var region models.Region
		intParams, _ := strconv.ParseInt(params["id"], 10, 0)

		fqry := fmt.Sprintf("%s %d", query, intParams)
		row := dbDriver.ExecuteRowQuery(db, fqry)
		err := row.Scan(&region.RegionId, &region.RegionName)
		if err == sql.ErrNoRows {
			logFatal(err, db, nil)
			var response = models.Response { false, "No results" }
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(response)
		} else {
			defer db.Close()
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(region)
		}
	}
}


func (rc RegionController) AddNewRegion(dbStr string, insQuery string, lastIdQuery string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		var region models.Region
		var lastId int64;
		newId := lastId + 1
		json.NewDecoder(request.Body).Decode(&region)
		dbDriver.ExecuteRowQuery(db, lastIdQuery).Scan(&lastId)
		if region.RegionName == "" {
			var response = models.Response { false, "Region Name is required!" }
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(response)
		} else {
			fquery := fmt.Sprintf(insQuery, newId, region.RegionName)
			_ := dbDriver.ExecuteDataManipulationQuery(db, fquery)
		}

		validate := helpers.ConfirmRegionInsert(db, fmt.Sprintf("SELECT * FROM REGION WHERE REGION_ID = ", newId))
		if validate {
			var response = models.Response { true, "New Region Created!" }
			writer.WriteHeader(http.StatusCreated)
			json.NewEncoder(writer).Encode(response)
		}else{
			var response = models.Response { false, "Error Creating Region!" }
			writer.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(writer).Encode(response)
		}
	}
}

func (rc RegionController) RemoveRegion(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {


	}
}

func (rc RegionController) UpdateRegion(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {


	}
}


