package requests

type DeliverToPlantBatch struct {
	Product                   string `json:"Product"`
	BusinessPartner           int    `json:"BusinessPartner"`
	Plant                     string `json:"Plant"`
	Batch                     string `json:"Batch"`
	DeliverToPlantBatchExConf bool   `json:"DeliverToPlantBatchExConf"`
}
