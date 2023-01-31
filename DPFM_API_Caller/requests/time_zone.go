package requests

type TimeZone struct {
	BusinessPartner int     `json:"BusinessPartner"`
	Plant           string  `json:"Plant"`
	TimeZone        *string `json:"TimeZone"`
}
