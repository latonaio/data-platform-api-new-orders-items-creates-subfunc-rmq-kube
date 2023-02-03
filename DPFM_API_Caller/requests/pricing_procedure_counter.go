package requests

type PricingProcedureCounter struct {
	Product                   string `json:"Product"`
	SupplyChainRelationshipID int    `json:"SupplyChainRelationshipID"`
	Buyer                     int    `json:"Buyer"`
	Seller                    int    `json:"Seller"`
	PricingProcedureCounter   []int  `json:"PricingProcedureCounter"`
}
