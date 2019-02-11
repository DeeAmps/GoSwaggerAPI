package models

import "database/sql"

type Location struct {
	LocationId int `json:"LOCATION_ID"`
	StreetAddress string `json:"STREET_ADDRESS"`
	PostalCode sql.NullString `json:"POSTAL_CODE"`
	City string `json:"CITY"`
	State sql.NullString `json:"STATE_PROVINCE"`
	CountryId string `json:"COUNTRY_ID"`
}