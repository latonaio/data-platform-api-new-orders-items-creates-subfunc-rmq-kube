package requests

type QuantityUnitConversion struct {
	OrderItem                   int     `json:"OrderItem"`
	Product                     string  `json:"Product"`
	ConversionCoefficient       float32 `json:"ConversionCoefficient"`
	OrderQuantityInDeliveryUnit float32 `json:"OrderQuantityInDeliveryUnit"`
}
