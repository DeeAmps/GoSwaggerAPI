package models


type Country struct {
	CountryId int `json:"COUNTRY_ID"`
	CountryName string `json:"COUNTRY_NAME"`
	RegionId int `json:"REGION_ID"`
}
