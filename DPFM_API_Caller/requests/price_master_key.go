package requests

type PriceMasterKey struct {
	Product                    []*string `json:"Product"`
	SupplyChainRelationshipID  []int     `json:"SupplyChainRelationshipID"`
	Buyer                      []int     `json:"Buyer"`
	Seller                     []int     `json:"Seller"`
	ConditionValidityEndDate   string    `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate string    `json:"ConditionValidityStartDate"`
}
