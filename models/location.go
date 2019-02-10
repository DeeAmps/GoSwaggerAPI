package models


type Location struct {
	LocationId int `json:"LOCATION_ID"`
	StreetAddress string `json:"STREET_ADDRESS"`
	PostalCode string `json:"POSTAL_CODE"`
	City string `json:"CITY"`
	State string `json:"STATE_PROVINCE"`
	CountryId string `json:"COUNTRY_ID"`
}