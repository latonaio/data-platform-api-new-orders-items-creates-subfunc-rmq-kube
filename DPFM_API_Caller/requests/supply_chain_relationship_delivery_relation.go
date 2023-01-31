package requests

type SupplyChainRelationshipDeliveryRelation struct {
	SupplyChainRelationshipID         int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID int `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             int `json:"Buyer"`
	Seller                            int `json:"Seller"`
	DeliverToParty                    int `json:"DeliverToParty"`
	DeliverFromParty                  int `json:"DeliverFromParty"`
}
