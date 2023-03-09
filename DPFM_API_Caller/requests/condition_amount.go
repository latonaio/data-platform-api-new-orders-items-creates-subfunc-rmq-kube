package requests

type ConditionAmount struct {
	OrderItem                  int      `json:"OrderItem"`
	Product                    string   `json:"Product"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}
