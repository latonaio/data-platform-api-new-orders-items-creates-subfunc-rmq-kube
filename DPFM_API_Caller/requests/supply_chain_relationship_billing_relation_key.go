package requests

type SupplyChainRelationshipBillingRelationKey struct {
	SupplyChainRelationshipID []int `json:"SupplyChainRelationshipID"`
	Buyer                     []int `json:"Buyer"`
	Seller                    []int `json:"Seller"`
	DefaultRelation           bool  `json:"DefaultRelation"`
}
