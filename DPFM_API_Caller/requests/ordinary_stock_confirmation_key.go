package requests

type OrdinaryStockConfirmationKey struct {
	OrderID                          int     `json:"OrderID"`
	OrderItem                        int     `json:"OrderItem"`
	Product                          string  `json:"Product"`
	StockConfirmationBusinessPartner int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           string  `json:"StockConfirmationPlant"`
	ScheduleLineOrderQuantity        float32 `json:"ScheduleLineOrderQuantity"`
	RequestedDeliveryDate            string  `json:"RequestedDeliveryDate"`
}
