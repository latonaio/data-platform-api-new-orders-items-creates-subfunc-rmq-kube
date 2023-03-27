package requests

type OrdersItemScheduleLine struct {
	OrderID                                         int      `json:"OrderID"`
	OrderItem                                       int      `json:"OrderItem"`
	ScheduleLine                                    int      `json:"ScheduleLine"`
	SupplyChainRelationshipID                       int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID         int      `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                         string   `json:"Product"`
	StockConfirmationBussinessPartner               int      `json:"StockConfirmationBussinessPartner"`
	StockConfirmationPlant                          string   `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone                  *string  `json:"StockConfirmationPlantTimeZone"`
	StockConfirmationPlantBatch                     *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate    *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate      *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	RequestedDeliveryDate                           string   `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                           string   `json:"RequestedDeliveryTime"`
	ConfirmedDeliveryDate                           string   `json:"ConfirmedDeliveryDate"`
	ScheduleLineOrderQuantity                       float32  `json:"ScheduleLineOrderQuantity"`
	OriginalOrderQuantityInBaseUnit                 float32  `json:"OriginalOrderQuantityInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit float32  `json:"ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheck           float32  `json:"ConfirmedOrderQuantityByPDTAvailCheck"`
	DeliveredQuantityInBaseUnit                     *float32 `json:"DeliveredQuantityInBaseUnit"`
	OpenConfirmedQuantityInBaseUnit                 *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyConfirmed                           *bool    `json:"StockIsFullyConfirmed"`
	PlusMinusFlag                                   string   `json:"PlusMinusFlag"`
	ItemScheduleLineDeliveryBlockStatus             *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
	IsCancelled                                     *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion"`
}
