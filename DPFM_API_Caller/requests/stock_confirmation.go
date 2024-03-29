package requests

type StockConfirmation struct {
	BusinessPartner                 int     `json:"BusinessPartner"`
	Product                         string  `json:"Product"`
	Plant                           string  `json:"Plant"`
	StorageLocation                 string  `json:"StorageLocation"`
	StorageBin                      string  `json:"StorageBin"`
	Batch                           string  `json:"Batch"`
	RequestedQuantity               float32 `json:"RequestedQuantity"`
	ProductStockAvailabilityDate    string  `json:"ProductStockAvailabilityDate"`
	OrderID                         int     `json:"OrderID"`
	OrderItem                       int     `json:"OrderItem"`
	InventoryStockType              string  `json:"InventoryStockType"`
	InventorySpecialStockType       string  `json:"InventorySpecialStockType"`
	AvailableProductStock           float32 `json:"AvailableProductStock"`
	CheckedQuantity                 float32 `json:"CheckedQuantity"`
	CheckedDate                     string  `json:"CheckedDate"`
	OpenConfirmedQuantityInBaseUnit float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyChecked             bool    `json:"StockIsFullyChecked"`
	Suffix                          string  `json:"Suffix"`
	StockConfirmationIsLotUnit      bool    `json:"StockConfirmationIsLotUnit"`
	StockConfirmationIsOrdinary     bool    `json:"StockConfirmationIsOrdinary"`
}
