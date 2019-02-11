package helpers

import (
	"database/sql"
	"oraclehr.api.com/models"
)

func ConfirmRegion(db *sql.DB, query string) bool {
	var region models.Region
	row := db.QueryRow(query)
	err := row.Scan(&region.RegionId, &region.RegionName)
	if err == sql.ErrNoRows {
		return false
	} else{
		return true
	}
}
