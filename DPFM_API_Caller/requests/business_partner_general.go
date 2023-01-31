package requests

type BusinessPartnerGeneral struct {
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string  `json:"BusinessPartnerName"`
	OrganizationBPName1     *string `json:"OrganizationBPName1"`
	Language                string  `json:"Language"`
	AddressID               *int    `json:"AddressID"`
}
