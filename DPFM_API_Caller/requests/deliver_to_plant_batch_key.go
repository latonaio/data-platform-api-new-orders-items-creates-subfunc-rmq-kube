package requests

type DeliverToPlantBatchKey struct {
	Product             string `json:"Product"`
	DeliverToPlantBatch string `json:"DeliverToPlantBatch"`
	DeliverToParty      int    `json:"DeliverToParty"`
	DeliverToPlant      string `json:"DeliverToPlant"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	OrderItem           int    `json:"OrderItem"`
}
