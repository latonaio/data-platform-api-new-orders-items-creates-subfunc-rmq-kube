package requests

type ProductionPlantBatchKey struct {
	Product                        string `json:"Product"`
	ProductionPlantBatch           string `json:"ProductionPlantBatch"`
	ProductionPlantBusinessPartner int    `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                string `json:"ProductionPlant"`
	ValidityStartDate              string `json:"ValidityStartDate"`
	ValidityEndDate                string `json:"ValidityEndDate"`
}
