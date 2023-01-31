package requests

type ProductMasterGeneralKey struct {
	Product             []*string `json:"Product"`
	ValidityStartDate   string    `json:"ValidityStartDate"`
	IsMarkedForDeletion bool      `json:"IsMarkedForDeletion"`
}
