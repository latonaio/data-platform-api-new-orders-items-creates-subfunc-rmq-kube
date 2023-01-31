package requests

type ProductTaxClassificationBillFromCountry struct {
	Product                   string  `json:"Product"`
	Country                   string  `json:"Country"`
	ProductTaxCategory        string  `json:"ProductTaxCategory"`
	ProductTaxClassifiication *string `json:"ProductTaxClassification"`
}
