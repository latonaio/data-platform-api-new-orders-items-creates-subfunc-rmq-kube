package requests

type ItemGrossWeight struct {
	OrderItem               int      `json:"OrderItem"`
	Product                 string   `json:"Product"`
	ProductGrossWeight      *float32 `json:"ProductGrossWeight"`
	OrderQuantityInBaseUnit *float32 `json:"OrderQuantityInBaseUnit"`
	ItemGrossWeight         *float32 `json:"ItemGrossWeight"`
}
