package requests

type ProductionPlantRelationProduct struct {
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  string  `json:"Product"`
	ProductionPlantBusinessPartner           int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation           *string `json:"ProductionPlantStorageLocation"`
}
