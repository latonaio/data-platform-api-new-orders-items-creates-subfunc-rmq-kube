package requests

type ProductMasterQuality struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	Plant               string `json:"Plant"`
	IsMarkedForDeletion bool   `json:"IsMarkedForDeletion"`
}
