package requests

type StockConfPlantRelationProduct struct {
	SupplyChainRelationshipID               int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID int    `json:"SupplyChainRelationshipStockConfPlantID"`
	Buyer                                   int    `json:"Buyer"`
	Seller                                  int    `json:"Seller"`
	StockConfirmationBusinessPartner        int    `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                  string `json:"StockConfirmationPlant"`
	Product                                 string `json:"Product"`
}
