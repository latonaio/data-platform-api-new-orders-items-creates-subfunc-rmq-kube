package requests

type Address struct {
	AddressID       int     `json:"AddressID"`
	ValidityEndDate string  `json:"ValidityEndDate"`
	PostalCode      string  `json:"PostalCode"`
	LocalRegion     string  `json:"LocalRegion"`
	Country         string  `json:"Country"`
	District        *string `json:"District"`
	StreetName      string  `json:"StreetName"`
	CityName        string  `json:"CityName"`
	Building        *string `json:"Building"`
	Floor           *int    `json:"Floor"`
	Room            *int    `json:"Room"`
}
