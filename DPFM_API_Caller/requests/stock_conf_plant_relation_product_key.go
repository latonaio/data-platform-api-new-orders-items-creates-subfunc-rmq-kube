package requests

type StockConfPlantRelationProductKey struct {
	SupplyChainRelationshipID int      `json:"SupplyChainRelationshipID"`
	Buyer                     int      `json:"Buyer"`
	Seller                    int      `json:"Seller"`
	Product                   []string `json:"Product"`
}
