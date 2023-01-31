package requests

type SupplyChainRelationshipDeliveryPlantRelationProductKey struct {
	SupplyChainRelationshipID              []int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      []int    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID []int    `json:"SupplyChainRelationshipDeliveryPlantID"`
	DeliverToParty                         []int    `json:"DeliverToParty"`
	DeliverFromParty                       []int    `json:"DeliverFromParty"`
	DeliverToPlant                         []string `json:"DeliverToPlant"`
	DeliverFromPlant                       []string `json:"DeliverFromPlant"`
	Product                                []string `json:"Product"`
	IsMarkedForDeletion                    bool     `json:"IsMarkedForDeletion"`
}
