package requests

type NetAmount struct {
	OrderItem int      `json:"OrderItem"`
	Product   string   `json:"Product"`
	NetAmount *float32 `json:"NetAmount"`
}
