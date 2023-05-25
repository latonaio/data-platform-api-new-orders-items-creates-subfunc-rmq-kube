package requests

type DeliverFromPlantBatch struct {
	Product                     string `json:"Product"`
	BusinessPartner             int    `json:"BusinessPartner"`
	Plant                       string `json:"Plant"`
	Batch                       string `json:"Batch"`
	DeliverFromPlantBatchExConf bool   `json:"DeliverFromPlantBatchExConf"`
}
