package requests

type InspectionOrder struct {
	InspectionOrder                int    `json:"InspectionOrder"`
	Product                        string `json:"Product"`
	InspectionPlantBusinessPartner int    `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                string `json:"InspectionPlant"`
}
