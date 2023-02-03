package requests

type QuantityUnitConversionQueryGets struct {
	Product               string  `json:"Product"`
	QuantityUnitFrom      string  `json:"QuantityUnitFrom"`
	QuantityUnitTo        string  `json:"QuantityUnitTo"`
	ConversionCoefficient float32 `json:"ConversionCoefficient"`
}
