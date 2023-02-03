package requests

type BusinessPartnerGeneralBillingRelationKey struct {
	BillToParty   int `json:"BillToParty"`
	BillFromParty int `json:"BillFromParty"`
}
