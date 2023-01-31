package requests

type SupplyChainRelationshipDeliveryPlantRelationKey struct {
	SupplyChainRelationshipID         []int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID []int `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             []int `json:"Buyer"`
	Seller                            []int `json:"Seller"`
	DeliverToParty                    []int `json:"DeliverToParty"`
	DeliverFromParty                  []int `json:"DeliverFromParty"`
	DefaultRelation                   bool  `json:"DefaultRelation"`
}
