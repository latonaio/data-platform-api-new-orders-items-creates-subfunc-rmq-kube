package requests

type TaxAmount struct {
	OrderItem int      `json:"OrderItem"`
	Product   string   `json:"Product"`
	TaxCode   string   `json:"TaxCode"`
	TaxRate   *float32 `json:"TaxRate"`
	NetAmount *float32 `json:"NetAmount"`
	TaxAmount *float32 `json:"TaxAmount"`
}
