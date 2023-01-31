package requests

type OrderItemText struct {
	Product       string  `json:"Product"`
	Language      string  `json:"Language"`
	OrderItemText *string `json:"OrderItemText"`
}
