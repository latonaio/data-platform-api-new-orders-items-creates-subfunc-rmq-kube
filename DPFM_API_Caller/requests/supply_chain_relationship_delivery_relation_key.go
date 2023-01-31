package requests

type SupplyChainRelationshipDeliveryRelationKey struct {
	SupplyChainRelationshipID []int `json:"SupplyChainRelationshipID"`
	Buyer                     []int `json:"Buyer"`
	Seller                    []int `json:"Seller"`
	DeliverToParty            []int `json:"DeliverToParty"`
	DeliverFromParty          []int `json:"DeliverFromParty"`
}
