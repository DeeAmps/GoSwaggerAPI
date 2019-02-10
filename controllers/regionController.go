package controllers

import (
	"database/sql"
	"net/http"
	"oraclehr.api.com/models"
	"fmt"
	"encoding/json"
	dbDriver "oraclehr.api.com/db"

)

type RegionController struct{}

func logFatal(err error){
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return
}

func (rc RegionController) GetAllRegions(db *sql.DB, query string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var region models.Region
		regions := []models.Region{}
		rows := dbDriver.ExecuteQuery(db, query)
		for rows.Next() {
			err := rows.Scan(&region.RegionId, &region.RegionName)
			logFatal(err)
			regions = append(regions, region)
		}
		defer rows.Close()
		defer db.Close()
		json.NewEncoder(writer).Encode(regions)
	}
}

//func (rc RegionController) GetSingleRegion(db *sql.DB, query string) http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		params := mux.Vars(r)
//
//	}
//}


