package requests

type SupplyChainRelationshipPaymentRelation struct {
	SupplyChainRelationshipID        int   `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int   `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID int   `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int   `json:"Buyer"`
	Seller                           int   `json:"Seller"`
	BillToParty                      int   `json:"BillToParty"`
	BillFromParty                    int   `json:"BillFromParty"`
	Payer                            int   `json:"Payer"`
	Payee                            int   `json:"Payee"`
	DefaultRelation                  *bool `json:"DefaultRelation"`
}
