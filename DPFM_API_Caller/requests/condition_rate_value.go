package requests

type ConditionRateValue struct {
	Product                       string   `json:"Product"`
	SupplyChainRelationshipID     int      `json:"SupplyChainRelationshipID"`
	TaxCode                       string   `json:"TaxCode"`
	PriceMasterConditionRateValue *float32 `json:"PriceMasterConditionRateValue"`
	TaxRate                       *float32 `json:"TaxRate"`
	ConditionRateValue            *float32 `json:"ConditionRateValue"`
	ConditionQuantity             *float32 `json:"ConditionQuantity"`
	ConditionAmount               *float32 `json:"ConditionAmount"`
	ConditionIsManuallyChanged    *bool    `json:"ConditionIsManuallyChanged"`
}
