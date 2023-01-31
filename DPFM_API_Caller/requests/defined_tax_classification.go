package requests

type DefinedTaxClassification struct {
	Product                                 string  `json:"Product"`
	TransactionTaxClassification            *string `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   *string `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry *string `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                string  `json:"DefinedTaxClassification"`
}
