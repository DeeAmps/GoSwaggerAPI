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
)

type RegionController struct{}

func logFatal(err error, db *sql.DB, rows *sql.Rows){
	if err != nil {
		fmt.Println(err)
		if rows != nil {
			rows.Close()
		}
		db.Close()
		panic(err)
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
		json.NewEncoder(writer).Encode(regions)
	}
}

func (rc RegionController) GetSingleRegion(dbStr string, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		db := dbDriver.DBConnection(dbStr)
		params := mux.Vars(request)
		var region models.Region
		intParams, _ := strconv.ParseInt(params["id"], 10, 0)
		row := dbDriver.ExecuteRowQuery(db, query, intParams)
		fmt.Println(row)
		err := row.Scan(&region.RegionId, &region.RegionName)
		logFatal(err, db,nil)
		defer db.Close()
		json.NewEncoder(writer).Encode(region)

	}
}


