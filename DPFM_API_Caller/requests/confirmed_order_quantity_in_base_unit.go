package requests

type ConfirmedOrderQuantityInBaseUnit struct {
	OrderItem                        int     `json:"OrderItem"`
	ConfirmedOrderQuantityInBaseUnit float32 `json:"ConfirmedOrderQuantityInBaseUnit"`
}
