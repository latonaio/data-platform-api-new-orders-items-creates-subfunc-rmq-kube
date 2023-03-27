package requests

type InspectionPlan struct {
	InspectionPlantBusinessPartner int    `json:"InspectionPlantBusinessPartner"`
	InspectionPlan                 int    `json:"InspectionPlan"`
	InspectionPlant                string `json:"InspectionPlant"`
	Product                        string `json:"Product"`
}
