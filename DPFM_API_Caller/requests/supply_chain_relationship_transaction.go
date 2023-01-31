package requests

type SupplyChainRelationshipTransaction struct {
	SupplyChainRelationshipID int     `json:"SupplyChainRelationshipID"`
	Buyer                     int     `json:"Buyer"`
	Seller                    int     `json:"Seller"`
	TransactionCurrency       *string `json:"TransactionCurrency"`
	Incoterms                 *string `json:"Incoterms"`
	PaymentTerms              *string `json:"PaymentTerms"`
	PaymentMethod             *string `json:"PaymentMethod"`
	AccountAssignmentGroup    *string `json:"AccountAssignmentGroup"`
}
