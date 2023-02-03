package requests

type BusinessPartnerGeneralDeliveryRelationKey struct {
	Buyer            int `json:"Buyer"`
	Seller           int `json:"Seller"`
	DeliverToParty   int `json:"DeliverToParty"`
	DeliverFromParty int `json:"DeliverFromParty"`
}
