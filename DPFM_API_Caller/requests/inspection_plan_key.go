package requests

type InspectionPlanKey struct {
	Product         []string `json:"Product"`
	BusinessPartner []int    `json:"BusinessPartner"`
	Plant           []string `json:"Plant"`
}
