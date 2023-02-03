package requests

type OrderQuantityInDeliveryUnit struct {
	OrderItem                   int     `json:"OrderItem"`
	OrderQuantityInDeliveryUnit float32 `json:"OrderQuantityInDeliveryUnit"`
}
