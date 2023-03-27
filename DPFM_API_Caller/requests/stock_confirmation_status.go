package requests

type StockConfirmationStatus struct {
	OrderItem                                       int     `json:"OrderItem"`
	StockIsFullyConfirmed                           *bool   `json:"StockIsFullyConfirmed"`
	ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit float32 `json:"ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit"`
	StockConfirmationStatus                         *string `json:"StockConfirmationStatus"`
}
