package requests

type ConditionIsManuallyChanged struct {
	Product                    string `json:"Product"`
	ConditionIsManuallyChanged *bool  `json:"ConditionIsManuallyChanged"`
}
