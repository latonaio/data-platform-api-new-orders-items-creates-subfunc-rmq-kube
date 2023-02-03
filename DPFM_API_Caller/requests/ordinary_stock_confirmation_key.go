package requests

type OrdinaryStockConfirmationKey struct {
	Product                          string `json:"Product"`
	StockConfirmationBusinessPartner int    `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           string `json:"StockConfirmationPlant"`
	RequestedDeliveryDate            string `json:"RequestedDeliveryDate"`
}
