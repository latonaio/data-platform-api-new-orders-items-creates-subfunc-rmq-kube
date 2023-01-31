package requests

type CalculateOrderIDQueryGets struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	LatestNumber             *int   `json:"LatestNumber"`
}
