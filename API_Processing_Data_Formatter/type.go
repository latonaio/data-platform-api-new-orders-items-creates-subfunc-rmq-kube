package api_processing_data_formatter

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	MetaData                                            *MetaData                                              `json:"MetaData"`
	SupplyChainRelationshipGeneral                      []*SupplyChainRelationshipGeneral                      `json:"SupplyChainRelationshipGeneral"`
	SupplyChainRelationshipDeliveryRelation             []*SupplyChainRelationshipDeliveryRelation             `json:"SupplyChainRelationshipDeliveryRelation"`
	SupplyChainRelationshipDeliveryPlantRelation        []*SupplyChainRelationshipDeliveryPlantRelation        `json:"SupplyChainRelationshipDeliveryPlantRelation"`
	SupplyChainRelationshipTransaction                  []*SupplyChainRelationshipTransaction                  `json:"SupplyChainRelationshipTransaction"`
	SupplyChainRelationshipBillingRelation              []*SupplyChainRelationshipBillingRelation              `json:"SupplyChainRelationshipBillingRelation"`
	CalculateOrderID                                    *CalculateOrderID                                      `json:"CalculateOrderID"`
	OrderItem                                           []*OrderItem                                           `json:"OrderItem"`
	ProductTaxClassificationBillToCountry               []*ProductTaxClassificationBillToCountry               `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry             []*ProductTaxClassificationBillFromCountry             `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                            []*DefinedTaxClassification                            `json:"DefinedTaxClassification"`
	ProductMasterGeneral                                []*ProductMasterGeneral                                `json:"ProductMasterGeneral"`
	OrderItemText                                       []*OrderItemText                                       `json:"OrderItemText"`
	ItemCategoryIsINVP                                  []*ItemCategoryIsINVP                                  `json:"ItemCategoryIsINVP"`
	StockConfPlantRelationProduct                       []*StockConfPlantRelationProduct                       `json:"StockConfPlantRelationProduct"`
	StockConfPlantProductMasterBPPlant                  []*ProductMasterBPPlant                                `json:"StockConfPlantProductMasterBPPlant"`
	StockConfPlantBPGeneral                             []*BusinessPartnerGeneral                              `json:"StockConfPlantBPGeneral"`
	ProductionPlantRelationProduct                      []*ProductionPlantRelationProduct                      `json:"ProductionPlantRelationProduct"`
	ProductionPlantProductMasterBPPlant                 []*ProductMasterBPPlant                                `json:"ProductionPlantProductMasterBPPlant"`
	ProductionPlantBPGeneral                            []*BusinessPartnerGeneral                              `json:"ProductionPlantBPGeneral"`
	SupplyChainRelationshipDeliveryPlantRelationProduct []*SupplyChainRelationshipDeliveryPlantRelationProduct `json:"SupplyChainRelationshipDeliveryPlantRelationProduct"`
	SupplyChainRelationshipProductMasterBPPlant         []*ProductMasterBPPlant                                `json:"SupplyChainRelationshipProductMasterBPPlant"`
	ProductionPlantTimeZone                             []*TimeZone                                            `json:"ProductionPlantTimeZone"`
	DeliverToPlantTimeZone                              []*TimeZone                                            `json:"DeliverToPlantTimeZone"`
	DeliverFromPlantTimeZone                            []*TimeZone                                            `json:"DeliverFromPlantTimeZone"`
	StockConfirmationPlantTimeZone                      []*TimeZone                                            `json:"StockConfirmationPlantTimeZone"`
	Incoterms                                           []*Incoterms                                           `json:"Incoterms"`
	PaymentTerms                                        []*PaymentTerms                                        `json:"PaymentTerms"`
	PaymentMethod                                       []*PaymentMethod                                       `json:"PaymentMethod"`
	ItemGrossWeight                                     []*ItemGrossWeight                                     `json:"ItemGrossWeight"`
	ItemNetWeight                                       []*ItemNetWeight                                       `json:"ItemNetWeight"`
	CreationDateItem                                    *CreationDateItem                                      `json:"CreationDateItem"`
	LastChangeDateItem                                  *LastChangeDateItem                                    `json:"LastChangeDateItem"`
}

// Initializer
type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

// Header
type SupplyChainRelationshipGeneral struct {
	SupplyChainRelationshipID int `json:"SupplyChainRelationshipID"`
	Buyer                     int `json:"Buyer"`
	Seller                    int `json:"Seller"`
}

type SupplyChainRelationshipDeliveryRelationKey struct {
	SupplyChainRelationshipID []int `json:"SupplyChainRelationshipID"`
	Buyer                     []int `json:"Buyer"`
	Seller                    []int `json:"Seller"`
	DeliverToParty            []int `json:"DeliverToParty"`
	DeliverFromParty          []int `json:"DeliverFromParty"`
}

type SupplyChainRelationshipDeliveryRelation struct {
	SupplyChainRelationshipID         int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID int `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             int `json:"Buyer"`
	Seller                            int `json:"Seller"`
	DeliverToParty                    int `json:"DeliverToParty"`
	DeliverFromParty                  int `json:"DeliverFromParty"`
}

type SupplyChainRelationshipDeliveryPlantRelationKey struct {
	SupplyChainRelationshipID         []int `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID []int `json:"SupplyChainRelationshipDeliveryID"`
	Buyer                             []int `json:"Buyer"`
	Seller                            []int `json:"Seller"`
	DeliverToParty                    []int `json:"DeliverToParty"`
	DeliverFromParty                  []int `json:"DeliverFromParty"`
	DefaultRelation                   bool  `json:"DefaultRelation"`
}

type SupplyChainRelationshipDeliveryPlantRelation struct {
	SupplyChainRelationshipID              int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int    `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int    `json:"Buyer"`
	Seller                                 int    `json:"Seller"`
	DeliverToParty                         int    `json:"DeliverToParty"`
	DeliverFromParty                       int    `json:"DeliverFromParty"`
	DeliverToPlant                         string `json:"DeliverToPlant"`
	DeliverFromPlant                       string `json:"DeliverFromPlant"`
	DefaultRelation                        *bool  `json:"DefaultRelation"`
}

type SupplyChainRelationshipTransaction struct {
	SupplyChainRelationshipID int     `json:"SupplyChainRelationshipID"`
	Buyer                     int     `json:"Buyer"`
	Seller                    int     `json:"Seller"`
	TransactionCurrency       *string `json:"TransactionCurrency"`
	Incoterms                 *string `json:"Incoterms"`
	PaymentTerms              *string `json:"PaymentTerms"`
	PaymentMethod             *string `json:"PaymentMethod"`
	AccountAssignmentGroup    *string `json:"AccountAssignmentGroup"`
}

type SupplyChainRelationshipBillingRelationKey struct {
	SupplyChainRelationshipID []int `json:"SupplyChainRelationshipID"`
	Buyer                     []int `json:"Buyer"`
	Seller                    []int `json:"Seller"`
	DefaultRelation           bool  `json:"DefaultRelation"`
}

type SupplyChainRelationshipBillingRelation struct {
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int     `json:"SupplyChainRelationshipBillingID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	BillToParty                      int     `json:"BillToParty"`
	BillFromParty                    int     `json:"BillFromParty"`
	DefaultRelation                  *bool   `json:"DefaultRelation"`
	BillToCountry                    string  `json:"BillToCountry"`
	BillFromCountry                  string  `json:"BillFromCountry"`
	IsExportImport                   *bool   `json:"IsExportImport"`
	TransactionTaxCategory           *string `json:"TransactionTaxCategory"`
	TransactionTaxClassification     *string `json:"TransactionTaxClassification"`
}

type CalculateOrderIDKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
}

type CalculateOrderIDQueryGets struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	LatestNumber             *int   `json:"LatestNumber"`
}

type CalculateOrderID struct {
	OrderIDLatestNumber *int `json:"OrderIDLatestNumber"`
	OrderID             int  `json:"OrderID"`
}

// Item
type OrderItem struct {
	OrderItemNumber int `json:"OrderItemNumber"`
}

type ProductTaxClassificationKey struct {
	Product            []*string `json:"Product"`
	Country            string    `json:"Country"`
	ProductTaxCategory string    `json:"ProductTaxCategory"`
}

type ProductTaxClassificationBillToCountry struct {
	Product                   string  `json:"Product"`
	Country                   string  `json:"Country"`
	ProductTaxCategory        string  `json:"ProductTaxCategory"`
	ProductTaxClassifiication *string `json:"ProductTaxClassification"`
}

type ProductTaxClassificationBillFromCountry struct {
	Product                   string  `json:"Product"`
	Country                   string  `json:"Country"`
	ProductTaxCategory        string  `json:"ProductTaxCategory"`
	ProductTaxClassifiication *string `json:"ProductTaxClassification"`
}

type DefinedTaxClassification struct {
	Product                                 string  `json:"Product"`
	TransactionTaxClassification            *string `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   *string `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry *string `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                string  `json:"DefinedTaxClassification"`
}

type ProductMasterGeneralKey struct {
	Product             []*string `json:"Product"`
	ValidityStartDate   string    `json:"ValidityStartDate"`
	IsMarkedForDeletion bool      `json:"IsMarkedForDeletion"`
}

type ProductMasterGeneral struct {
	Product                       string   `json:"Product"`
	BaseUnit                      *string  `json:"BaseUnit"`
	ProductGroup                  *string  `json:"ProductGroup"`
	ProductStandardID             *string  `json:"ProductStandardID"`
	GrossWeight                   *float32 `json:"GrossWeight"`
	NetWeight                     *float32 `json:"NetWeight"`
	WeightUnit                    *string  `json:"WeightUnit"`
	ItemCategory                  *string  `json:"ItemCategory"`
	ProductAccountAssignmentGroup *string  `json:"ProductAccountAssignmentGroup"`
	CountryOfOrigin               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string  `json:"CountryOfOriginLanguage"`
}

type OrderItemTextKey struct {
	Product  string `json:"Product"`
	Language string `json:"Language"`
}

type OrderItemText struct {
	Product       string  `json:"Product"`
	Language      string  `json:"Language"`
	OrderItemText *string `json:"OrderItemText"`
}

type ItemCategoryIsINVP struct {
	Product            string `json:"Product"`
	ItemCategoryIsINVP bool   `json:"ItemCategoryIsINVP"`
}

type StockConfPlantRelationProductKey struct {
	SupplyChainRelationshipID int      `json:"SupplyChainRelationshipID"`
	Buyer                     int      `json:"Buyer"`
	Seller                    int      `json:"Seller"`
	Product                   []string `json:"Product"`
}

type StockConfPlantRelationProduct struct {
	SupplyChainRelationshipID               int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID int    `json:"SupplyChainRelationshipStockConfPlantID"`
	Buyer                                   int    `json:"Buyer"`
	Seller                                  int    `json:"Seller"`
	StockConfirmationBusinessPartner        int    `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                  string `json:"StockConfirmationPlant"`
	Product                                 string `json:"Product"`
}

type ProductMasterBPPlantKey struct {
	Product         string `json:"Product"`
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}

type ProductMasterBPPlant struct {
	Product                   string  `json:"Product"`
	BusinessPartner           int     `json:"BusinessPartner"`
	Plant                     string  `json:"Plant"`
	IsBatchManagementRequired *bool   `json:"IsBatchManagementRequired"`
	BatchManagementPolicy     *string `json:"BatchManagementPolicy"`
}

type BusinessPartnerGeneral struct {
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string  `json:"BusinessPartnerName"`
	OrganizationBPName1     *string `json:"OrganizationBPName1"`
	Language                string  `json:"Language"`
	AddressID               *int    `json:"AddressID"`
}

type ProductionPlantRelationProductKey struct {
	SupplyChainRelationshipID int      `json:"SupplyChainRelationshipID"`
	Buyer                     int      `json:"Buyer"`
	Seller                    int      `json:"Seller"`
	Product                   []string `json:"Product"`
}

type ProductionPlantRelationProduct struct {
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  string  `json:"Product"`
	ProductionPlantBusinessPartner           int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                          string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation           *string `json:"ProductionPlantStorageLocation"`
}

type SupplyChainRelationshipDeliveryPlantRelationProductKey struct {
	SupplyChainRelationshipID              []int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      []int    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID []int    `json:"SupplyChainRelationshipDeliveryPlantID"`
	DeliverToParty                         []int    `json:"DeliverToParty"`
	DeliverFromParty                       []int    `json:"DeliverFromParty"`
	DeliverToPlant                         []string `json:"DeliverToPlant"`
	DeliverFromPlant                       []string `json:"DeliverFromPlant"`
	Product                                []string `json:"Product"`
	IsMarkedForDeletion                    bool     `json:"IsMarkedForDeletion"`
}

type SupplyChainRelationshipDeliveryPlantRelationProduct struct {
	SupplyChainRelationshipID              int    `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int    `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int    `json:"SupplyChainRelationshipDeliveryPlantID"`
	DeliverToParty                         int    `json:"DeliverToParty"`
	DeliverFromParty                       int    `json:"DeliverFromParty"`
	DeliverToPlant                         string `json:"DeliverToPlant"`
	DeliverFromPlant                       string `json:"DeliverFromPlant"`
	Product                                string `json:"Product"`
	DeliverToPlantStorageLocation          string `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation        string `json:"DeliverFromPlantStorageLocation"`
	DeliveryUnit                           string `json:"DeliveryUnit"`
	StandardDeliveryDurationInDays         *int   `json:"StandardDeliveryDurationInDays"`
	IsMarkedForDeletion                    *bool  `json:"IsMarkedForDeletion"`
}

type TimeZoneKey struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}

type TimeZone struct {
	BusinessPartner int     `json:"BusinessPartner"`
	Plant           string  `json:"Plant"`
	TimeZone        *string `json:"TimeZone"`
}

type Incoterms struct {
	Incoterms *string `json:"Incoterms"`
}

type PaymentTerms struct {
	PaymentTerms *string `json:"PaymentTerms"`
}

type PaymentMethod struct {
	PaymentMethod *string `json:"PaymentMethod"`
}

type ItemGrossWeight struct {
	Product                 string   `json:"Product"`
	ProductGrossWeight      *float32 `json:"ProductGrossWeight"`
	OrderQuantityInBaseUnit *float32 `json:"OrderQuantityInBaseUnit"`
	ItemGrossWeight         *float32 `json:"ItemGrossWeight"`
}

type ItemNetWeight struct {
	Product                 string   `json:"Product"`
	ProductNetWeight        *float32 `json:"ProductNetWeight"`
	OrderQuantityInBaseUnit *float32 `json:"OrderQuantityInBaseUnit"`
	ItemNetWeight           *float32 `json:"ItemNetWeight"`
}

type CreationDateItem struct {
	CreationDate string `json:"CreationDate"`
}

type LastChangeDateItem struct {
	LastChangeDate string `json:"LastChangeDate"`
}
