package requests

type QuantityUnitConversionKey struct {
	Product      string `json:"Product"`
	BaseUnit     string `json:"BaseUnit"`
	DeliveryUnit string `json:"DeliveryUnit"`
}
