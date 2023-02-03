package requests

type SupplyChainRelationshipPaymentRelationKey struct {
	SupplyChainRelationshipID []int `json:"SupplyChainRelationshipID"`
	Buyer                     []int `json:"Buyer"`
	Seller                    []int `json:"Seller"`
	BillToParty               []int `json:"BillToParty"`
	BillFromParty             []int `json:"BillFromParty"`
	DefaultRelation           bool  `json:"DefaultRelation"`
}
