package requests

type DeliverFromPlantBatchMasterdata struct {
	Product                     string  `json:"Product"`
	BusinessPartner             *int    `json:"BusinessPartner"`
	Plant                       string  `json:"Plant"`
	Batch                       string  `json:"Batch"`
	CountryOfOrigin             string  `json:"CountryOfOrigin"`
	ValidityStartDate           *string `json:"ValidityStartDate"`
	ManufactureDate             *string `json:"ManufactureDate"`
	CreationDateTime            string  `json:"CreationDateTime"`
	LastChangeDateTime          string  `json:"LastChangeDateTime"`
	IsMarkedForDeletion         bool    `json:"IsMarkedForDeletion"`
	DeliverFromPlantBatchExConf bool    `json:"DeliverFromPlantBatchExConf"`
}
