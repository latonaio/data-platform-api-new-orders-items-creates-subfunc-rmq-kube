package requests

type AddressKey struct {
	AddressID       []*int `json:"AddressID"`
	ValidityEndDate string `json:"ValidityEndDate"`
}
