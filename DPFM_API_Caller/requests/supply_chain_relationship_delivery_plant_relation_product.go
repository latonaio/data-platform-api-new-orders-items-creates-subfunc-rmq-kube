package requests

type SupplyChainRelationshipDeliveryPlantRelationProduct struct {
	SupplyChainRelationshipID              int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int    `json:"SupplyChainRelationshipDeliveryPlantID"`
	DeliverToParty                         int    `json:"DeliverToParty"`
	DeliverFromParty                       int    `json:"DeliverFromParty"`
	DeliverToPlant                         string `json:"DeliverToPlant"`
	DeliverFromPlant                       string `json:"DeliverFromPlant"`
	Product                                string `json:"Product"`
	DeliverToPlantStorageLocation          string `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation        string `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                           string `json:"DeliveryUnit"`
	StandardDeliveryDurationInDays         *int   `json:"StandardDeliveryDurationInDays"`
	IsMarkedForDeletion                    *bool  `json:"IsMarkedForDeletion"`
}
