package requests

type TaxCode struct {
	Product                  string  `json:"Product"`
	DefinedTaxClassification string  `json:"DefinedTaxClassification"`
	IsExportImport           *bool   `json:"IsExportImport"`
	TaxCode                  *string `json:"TaxCode"`
}
