package requests

type ProductionPlantBatch struct {
	Product                    string `json:"Product"`
	BusinessPartner            int    `json:"BusinessPartner"`
	Plant                      string `json:"Plant"`
	Batch                      string `json:"Batch"`
	ProductionPlantBatchExConf bool   `json:"ProductionPlantBatchExConf"`
}
