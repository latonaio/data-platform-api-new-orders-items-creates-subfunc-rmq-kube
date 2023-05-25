package requests

type DeliverFromPlantBatchKey struct {
	Product               string `json:"Product"`
	DeliverFromPlantBatch string `json:"DeliverFromPlantBatch"`
	DeliverFromParty      int    `json:"DeliverFromParty"`
	DeliverFromPlant      string `json:"DeliverFromPlant"`
	ValidityStartDate     string `json:"ValidityStartDate"`
	ValidityEndDate       string `json:"ValidityEndDate"`
}
