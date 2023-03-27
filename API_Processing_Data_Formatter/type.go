package api_processing_data_formatter

type SDC struct {
	MetaData                                               *MetaData                                              `json:"MetaData"`
	SupplyChainRelationshipGeneral                         []*SupplyChainRelationshipGeneral                      `json:"SupplyChainRelationshipGeneral"`
	SupplyChainRelationshipDeliveryRelation                []*SupplyChainRelationshipDeliveryRelation             `json:"SupplyChainRelationshipDeliveryRelation"`
	SupplyChainRelationshipDeliveryPlantRelation           []*SupplyChainRelationshipDeliveryPlantRelation        `json:"SupplyChainRelationshipDeliveryPlantRelation"`
	SupplyChainRelationshipTransaction                     []*SupplyChainRelationshipTransaction                  `json:"SupplyChainRelationshipTransaction"`
	SupplyChainRelationshipBillingRelation                 []*SupplyChainRelationshipBillingRelation              `json:"SupplyChainRelationshipBillingRelation"`
	SupplyChainRelationshipPaymentRelation                 []*SupplyChainRelationshipPaymentRelation              `json:"SupplyChainRelationshipPaymentRelation"`
	PaymentTerms                                           []*PaymentTerms                                        `json:"PaymentTerms"`
	HeaderInvoiceDocumentDate                              *HeaderInvoiceDocumentDate                             `json:"HeaderInvoiceDocumentDate"`
	HeaderPricingDate                                      *PricingDate                                           `json:"HeaderPricingDate"`
	HeaderPriceDetnExchangeRate                            *PriceDetnExchangeRate                                 `json:"HeaderPriceDetnExchangeRate"`
	HeaderAccountingExchangeRate                           *AccountingExchangeRate                                `json:"HeaderAccountingExchangeRate"`
	BusinessPartnerGeneralBuyer                            []*BusinessPartnerGeneral                              `json:"BusinessPartnerGeneralBuyer"`
	BusinessPartnerGeneralSeller                           []*BusinessPartnerGeneral                              `json:"BusinessPartnerGeneralSeller"`
	BusinessPartnerGeneralDeliverToParty                   []*BusinessPartnerGeneral                              `json:"BusinessPartnerGeneralDeliverToParty"`
	BusinessPartnerGeneralDeliverFromParty                 []*BusinessPartnerGeneral                              `json:"BusinessPartnerGeneralDeliverFromParty"`
	BusinessPartnerGeneralBillToParty                      []*BusinessPartnerGeneral                              `json:"BusinessPartnerGeneralBillToParty"`
	BusinessPartnerGeneralBillFromParty                    []*BusinessPartnerGeneral                              `json:"BusinessPartnerGeneralBillFromParty"`
	BusinessPartnerGeneralPayer                            []*BusinessPartnerGeneral                              `json:"BusinessPartnerGeneralPayer"`
	BusinessPartnerGeneralPayee                            []*BusinessPartnerGeneral                              `json:"BusinessPartnerGeneralPayee"`
	OrderItem                                              []*OrderItem                                           `json:"OrderItem"`
	ProductTaxClassificationBillToCountry                  []*ProductTaxClassificationBillToCountry               `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry                []*ProductTaxClassificationBillFromCountry             `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                               []*DefinedTaxClassification                            `json:"DefinedTaxClassification"`
	ProductMasterGeneral                                   []*ProductMasterGeneral                                `json:"ProductMasterGeneral"`
	OrderItemText                                          []*OrderItemText                                       `json:"OrderItemText"`
	ItemCategoryIsINVP                                     []*ItemCategoryIsINVP                                  `json:"ItemCategoryIsINVP"`
	StockConfPlantRelationProduct                          []*StockConfPlantRelationProduct                       `json:"StockConfPlantRelationProduct"`
	StockConfPlantProductMasterBPPlant                     []*ProductMasterBPPlant                                `json:"StockConfPlantProductMasterBPPlant"`
	StockConfPlantBPGeneral                                []*BusinessPartnerGeneral                              `json:"StockConfPlantBPGeneral"`
	ProductionPlantRelationProduct                         []*ProductionPlantRelationProduct                      `json:"ProductionPlantRelationProduct"`
	ProductionPlantProductMasterBPPlant                    []*ProductMasterBPPlant                                `json:"ProductionPlantProductMasterBPPlant"`
	ProductionPlantBPGeneral                               []*BusinessPartnerGeneral                              `json:"ProductionPlantBPGeneral"`
	SupplyChainRelationshipDeliveryPlantRelationProduct    []*SupplyChainRelationshipDeliveryPlantRelationProduct `json:"SupplyChainRelationshipDeliveryPlantRelationProduct"`
	SupplyChainRelationshipProductMasterBPPlantDeliverTo   []*ProductMasterBPPlant                                `json:"SupplyChainRelationshipProductMasterBPPlantDeliverTo"`
	SupplyChainRelationshipProductMasterBPPlantDeliverFrom []*ProductMasterBPPlant                                `json:"SupplyChainRelationshipProductMasterBPPlantDeliverFrom"`
	ProductionPlantTimeZone                                []*TimeZone                                            `json:"ProductionPlantTimeZone"`
	DeliverToPlantTimeZone                                 []*TimeZone                                            `json:"DeliverToPlantTimeZone"`
	DeliverFromPlantTimeZone                               []*TimeZone                                            `json:"DeliverFromPlantTimeZone"`
	StockConfirmationPlantTimeZone                         []*TimeZone                                            `json:"StockConfirmationPlantTimeZone"`
	Incoterms                                              []*Incoterms                                           `json:"Incoterms"`
	ItemPaymentTerms                                       []*ItemPaymentTerms                                    `json:"ItemPaymentTerms"`
	PaymentMethod                                          []*PaymentMethod                                       `json:"PaymentMethod"`
	ItemGrossWeight                                        []*ItemGrossWeight                                     `json:"ItemGrossWeight"`
	ItemNetWeight                                          []*ItemNetWeight                                       `json:"ItemNetWeight"`
	TaxCode                                                []*TaxCode                                             `json:"TaxCode"`
	TaxRate                                                []*TaxRate                                             `json:"TaxRate"`
	StockConfirmation                                      []*StockConfirmation                                   `json:"StockConfirmation"`
	StockConfirmationOrdersItemScheduleLine                []*OrdersItemScheduleLine                              `json:"StockConfirmationOrdersItemScheduleLine"`
	StockConfirmationStatus                                []*StockConfirmationStatus                             `json:"StockConfirmationStatus"`
	ConfirmedOrderQuantityInBaseUnit                       []*ConfirmedOrderQuantityInBaseUnit                    `json:"ConfirmedOrderQuantityInBaseUnit"`
	ItemPricingDate                                        []*PricingDate                                         `json:"ItemPricingDate"`
	ItemInvoiceDocumentDate                                []*ItemInvoiceDocumentDate                             `json:"ItemInvoiceDocumentDate"`
	ItemPriceDetnExchangeRate                              []*PriceDetnExchangeRate                               `json:"ItemPriceDetnExchangeRate"`
	ItemAccountingExchangeRate                             []*AccountingExchangeRate                              `json:"ItemAccountingExchangeRate"`
	ItemReferenceDocument                                  []*ItemReferenceDocument                               `json:"ItemReferenceDocument"`
	OrderItemTextByBuyer                                   []*OrderItemTextByBuyerSeller                          `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                                  []*OrderItemTextByBuyerSeller                          `json:"OrderItemTextBySeller"`
	ProductMasterQuality                                   []*ProductMasterQuality                                `json:"ProductMasterQuality"`
	InspectionPlan                                         []*InspectionPlan                                      `json:"InspectionPlan"`
	InspectionOrder                                        []*InspectionOrder                                     `json:"InspectionOrder"`
	PriceMaster                                            []*PriceMaster                                         `json:"PriceMaster"`
	ConditionAmount                                        []*ConditionAmount                                     `json:"ConditionAmount"`
	ConditionRateValue                                     []*ConditionRateValue                                  `json:"ConditionRateValue"`
	ConditionIsManuallyChanged                             []*ConditionIsManuallyChanged                          `json:"ConditionIsManuallyChanged"`
	PricingProcedureCounter                                []*PricingProcedureCounter                             `json:"PricingProcedureCounter"`
	NetAmount                                              []*NetAmount                                           `json:"NetAmount"`
	TaxAmount                                              []*TaxAmount                                           `json:"TaxAmount"`
	GrossAmount                                            []*GrossAmount                                         `json:"GrossAmount"`
	Address                                                []*Address                                             `json:"Address"`
	AddressMaster                                          []*AddressMaster                                       `json:"AddressMaster"`
	QuantityUnitConversion                                 []*QuantityUnitConversion                              `json:"QuantityUnitConversion"`
	OrderQuantityInDeliveryUnit                            []*OrderQuantityInDeliveryUnit                         `json:"OrderQuantityInDeliveryUnit"`
	Partner                                                []*Partner                                             `json:"Partner"`
	CreationDateItem                                       *CreationDateItem                                      `json:"CreationDateItem"`
	CreationTimeItem                                       *CreationTimeItem                                      `json:"CreationTimeItem"`
	LastChangeDateItem                                     *LastChangeDateItem                                    `json:"LastChangeDateItem"`
	LastChangeTimeItem                                     *LastChangeTimeItem                                    `json:"LastChangeTimeItem"`
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

type SupplyChainRelationshipPaymentRelationKey struct {
	SupplyChainRelationshipID []int `json:"SupplyChainRelationshipID"`
	Buyer                     []int `json:"Buyer"`
	Seller                    []int `json:"Seller"`
	BillToParty               []int `json:"BillToParty"`
	BillFromParty             []int `json:"BillFromParty"`
	DefaultRelation           bool  `json:"DefaultRelation"`
}

type SupplyChainRelationshipPaymentRelation struct {
	SupplyChainRelationshipID        int   `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID int   `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID int   `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int   `json:"Buyer"`
	Seller                           int   `json:"Seller"`
	BillToParty                      int   `json:"BillToParty"`
	BillFromParty                    int   `json:"BillFromParty"`
	Payer                            int   `json:"Payer"`
	Payee                            int   `json:"Payee"`
	DefaultRelation                  *bool `json:"DefaultRelation"`
}

type PaymentTerms struct {
	PaymentTerms                string `json:"PaymentTerms"`
	BaseDate                    int    `json:"BaseDate"`
	BaseDateCalcAddMonth        *int   `json:"BaseDateCalcAddMonth"`
	BaseDateCalcFixedDate       *int   `json:"BaseDateCalcFixedDate"`
	PaymentDueDateCalcAddMonth  *int   `json:"PaymentDueDateCalcAddMonth"`
	PaymentDueDateCalcFixedDate *int   `json:"PaymentDueDateCalcFixedDate"`
}

type HeaderInvoiceDocumentDate struct {
	RequestedDeliveryDate string `json:"RequestedDeliveryDate"`
	InvoiceDocumentDate   string `json:"InvoiceDocumentDate"`
}

type PricingDate struct {
	PricingDate string `json:"PricingDate"`
}

type PriceDetnExchangeRate struct {
	PriceDetnExchangeRate *float32 `json:"PriceDetnExchangeRate"`
}

type AccountingExchangeRate struct {
	AccountingExchangeRate *float32 `json:"AccountingExchangeRate"`
}

type BusinessPartnerGeneralDeliveryRelationKey struct {
	Buyer            int `json:"Buyer"`
	Seller           int `json:"Seller"`
	DeliverToParty   int `json:"DeliverToParty"`
	DeliverFromParty int `json:"DeliverFromParty"`
}

type BusinessPartnerGeneralBillingRelationKey struct {
	BillToParty   int `json:"BillToParty"`
	BillFromParty int `json:"BillFromParty"`
}

type BusinessPartnerGeneralPaymentRelationKey struct {
	Payer int `json:"Payer"`
	Payee int `json:"Payee"`
}

type BusinessPartnerGeneral struct {
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string  `json:"BusinessPartnerName"`
	Country                 string  `json:"Country"`
	Language                string  `json:"Language"`
	Currency                string  `json:"Currency"`
	AddressID               *int    `json:"AddressID"`
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
	InternalCapacityQuantity      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit  *string  `json:"InternalCapacityQuantityUnit"`
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
	SupplyChainRelationshipID []int    `json:"SupplyChainRelationshipID"`
	Buyer                     []int    `json:"Buyer"`
	Seller                    []int    `json:"Seller"`
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

type ProductionPlantRelationProductKey struct {
	SupplyChainRelationshipID []int    `json:"SupplyChainRelationshipID"`
	Buyer                     []int    `json:"Buyer"`
	Seller                    []int    `json:"Seller"`
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

type ItemPaymentTerms struct {
	PaymentTerms *string `json:"PaymentTerms"`
}

type PaymentMethod struct {
	PaymentMethod *string `json:"PaymentMethod"`
}

type ItemGrossWeight struct {
	OrderItem               int      `json:"OrderItem"`
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

type TaxCode struct {
	Product                  string  `json:"Product"`
	DefinedTaxClassification string  `json:"DefinedTaxClassification"`
	IsExportImport           *bool   `json:"IsExportImport"`
	TaxCode                  *string `json:"TaxCode"`
}

type TaxRateKey struct {
	Country           string    `json:"Country"`
	TaxCode           []*string `json:"TaxCode"`
	ValidityEndDate   string    `json:"ValidityEndDate"`
	ValidityStartDate string    `json:"ValidityStartDate"`
}

type TaxRate struct {
	Country           string   `json:"Country"`
	TaxCode           string   `json:"TaxCode"`
	ValidityEndDate   string   `json:"ValidityEndDate"`
	ValidityStartDate string   `json:"ValidityStartDate"`
	TaxRate           *float32 `json:"TaxRate"`
}

type LotUnitStockConfirmationKey struct {
	OrderID                                      int     `json:"OrderID"`
	OrderItem                                    int     `json:"OrderItem"`
	Product                                      string  `json:"Product"`
	StockConfirmationBusinessPartner             int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                       string  `json:"StockConfirmationPlant"`
	ScheduleLineOrderQuantity                    float32 `json:"ScheduleLineOrderQuantity"`
	RequestedDeliveryDate                        string  `json:"RequestedDeliveryDate"`
	StockConfirmationPlantBatch                  string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   string  `json:"StockConfirmationPlantBatchValidityEndDate"`
}

type OrdinaryStockConfirmationKey struct {
	OrderID                          int     `json:"OrderID"`
	OrderItem                        int     `json:"OrderItem"`
	Product                          string  `json:"Product"`
	StockConfirmationBusinessPartner int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant           string  `json:"StockConfirmationPlant"`
	ScheduleLineOrderQuantity        float32 `json:"ScheduleLineOrderQuantity"`
	RequestedDeliveryDate            string  `json:"RequestedDeliveryDate"`
}

type StockConfirmationKey struct {
	OrderID                                      int     `json:"OrderID"`
	OrderItem                                    int     `json:"OrderItem"`
	Product                                      string  `json:"Product"`
	StockConfirmationBusinessPartner             int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                       string  `json:"StockConfirmationPlant"`
	ScheduleLineOrderQuantity                    float32 `json:"ScheduleLineOrderQuantity"`
	RequestedDeliveryDate                        string  `json:"RequestedDeliveryDate"`
	StockConfirmationPlantBatch                  string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationIsLotUnit                   bool    `json:"StockConfirmationIsLotUnit"`
	StockConfirmationIsOrdinary                  bool    `json:"StockConfirmationIsOrdinary"`
}

type StockConfirmation struct {
	BusinessPartner                 int     `json:"BusinessPartner"`
	Product                         string  `json:"Product"`
	Plant                           string  `json:"Plant"`
	Batch                           string  `json:"Batch"`
	RequestedQuantity               float32 `json:"RequestedQuantity"`
	ProductStockAvailabilityDate    string  `json:"ProductStockAvailabilityDate"`
	OrderID                         int     `json:"OrderID"`
	OrderItem                       int     `json:"OrderItem"`
	Project                         string  `json:"Project"`
	InventoryStockType              string  `json:"InventoryStockType"`
	InventorySpecialStockType       string  `json:"InventorySpecialStockType"`
	AvailableProductStock           float32 `json:"AvailableProductStock"`
	CheckedQuantity                 float32 `json:"CheckedQuantity"`
	CheckedDate                     string  `json:"CheckedDate"`
	OpenConfirmedQuantityInBaseUnit float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyChecked             bool    `json:"StockIsFullyChecked"`
	Suffix                          string  `json:"Suffix"`
	StockConfirmationIsLotUnit      bool    `json:"StockConfirmationIsLotUnit"`
	StockConfirmationIsOrdinary     bool    `json:"StockConfirmationIsOrdinary"`
}

type ProductAvailabilityCheck struct {
	ConnectionKey                 string `json:"connection_key"`
	Result                        bool   `json:"result"`
	RedisKey                      string `json:"redis_key"`
	Filepath                      string `json:"filepath"`
	APIStatusCode                 int    `json:"api_status_code"`
	RuntimeSessionID              string `json:"runtime_session_id"`
	BusinessPartnerID             *int   `json:"business_partner"`
	ServiceLabel                  string `json:"service_label"`
	ProductStockAvailabilityCheck struct {
		BusinessPartner                 *int     `json:"BusinessPartner"`
		Product                         *string  `json:"Product"`
		Plant                           *string  `json:"Plant"`
		Batch                           *string  `json:"Batch"`
		RequestedQuantity               *float32 `json:"RequestedQuantity"`
		ProductStockAvailabilityDate    *string  `json:"ProductStockAvailabilityDate"`
		OrderID                         *int     `json:"OrderID"`
		OrderItem                       *int     `json:"OrderItem"`
		Project                         *string  `json:"Project"`
		InventoryStockType              *string  `json:"InventoryStockType"`
		InventorySpecialStockType       *string  `json:"InventorySpecialStockType"`
		AvailableProductStock           *float32 `json:"AvailableProductStock"`
		CheckedQuantity                 *float32 `json:"CheckedQuantity"`
		CheckedDate                     *string  `json:"CheckedDate"`
		OpenConfirmedQuantityInBaseUnit *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
		StockIsFullyChecked             *bool    `json:"StockIsFullyChecked"`
		Suffix                          *string  `json:"Suffix"`
	} `json:"ProductStockAvailabilityCheck"`
	APISchema        string   `json:"api_schema"`
	Accepter         []string `json:"accepter"`
	ProductStockCode string   `json:"product_stock_code"`
	Deleted          bool     `json:"deleted"`
}

type OrdersItemScheduleLine struct {
	OrderID                                         int      `json:"OrderID"`
	OrderItem                                       int      `json:"OrderItem"`
	ScheduleLine                                    int      `json:"ScheduleLine"`
	SupplyChainRelationshipID                       int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID         int      `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                         string   `json:"Product"`
	StockConfirmationBussinessPartner               int      `json:"StockConfirmationBussinessPartner"`
	StockConfirmationPlant                          string   `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone                  *string  `json:"StockConfirmationPlantTimeZone"`
	StockConfirmationPlantBatch                     *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate    *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate      *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	RequestedDeliveryDate                           string   `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                           string   `json:"RequestedDeliveryTime"`
	ConfirmedDeliveryDate                           string   `json:"ConfirmedDeliveryDate"`
	ScheduleLineOrderQuantity                       float32  `json:"ScheduleLineOrderQuantity"`
	OriginalOrderQuantityInBaseUnit                 float32  `json:"OriginalOrderQuantityInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit float32  `json:"ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheck           float32  `json:"ConfirmedOrderQuantityByPDTAvailCheck"`
	DeliveredQuantityInBaseUnit                     *float32 `json:"DeliveredQuantityInBaseUnit"`
	OpenConfirmedQuantityInBaseUnit                 *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyConfirmed                           *bool    `json:"StockIsFullyConfirmed"`
	PlusMinusFlag                                   string   `json:"PlusMinusFlag"`
	ItemScheduleLineDeliveryBlockStatus             *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
	IsCancelled                                     *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion"`
}

type StockConfirmationStatus struct {
	OrderItem                                       int     `json:"OrderItem"`
	StockIsFullyConfirmed                           *bool   `json:"StockIsFullyConfirmed"`
	ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit float32 `json:"ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit"`
	StockConfirmationStatus                         *string `json:"StockConfirmationStatus"`
}

type ConfirmedOrderQuantityInBaseUnit struct {
	OrderItem                        int     `json:"OrderItem"`
	ConfirmedOrderQuantityInBaseUnit float32 `json:"ConfirmedOrderQuantityInBaseUnit"`
}

type ItemInvoiceDocumentDate struct {
	InvoiceDocumentDate string `json:"InvoiceDocumentDate"`
}

type ItemReferenceDocument struct {
	ReferenceDocument     *int `json:"ReferenceDocument"`
	ReferenceDocumentItem *int `json:"ReferenceDocumentItem"`
}

type OrderItemTextByBuyerSellerKey struct {
	Product         []*string `json:"Product"`
	BusinessPartner []int     `json:"BusinessPartner"`
	Language        []string  `json:"Language"`
}

type OrderItemTextByBuyerSeller struct {
	Product            string  `json:"Product"`
	BusinessPartner    int     `json:"BusinessPartner"`
	Language           string  `json:"Language"`
	ProductDescription *string `json:"ProductDescription"`
}

type ProductMasterQualityKey struct {
	Product             []string `json:"Product"`
	BusinessPartner     []int    `json:"BusinessPartner"`
	Plant               []string `json:"Plant"`
	IsMarkedForDeletion bool     `json:"IsMarkedForDeletion"`
}

type ProductMasterQuality struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	Plant               string `json:"Plant"`
	IsMarkedForDeletion bool   `json:"IsMarkedForDeletion"`
}

type InspectionPlanKey struct {
	Product         []string `json:"Product"`
	BusinessPartner []int    `json:"BusinessPartner"`
	Plant           []string `json:"Plant"`
}

type InspectionPlan struct {
	InspectionPlantBusinessPartner int    `json:"InspectionPlantBusinessPartner"`
	InspectionPlan                 int    `json:"InspectionPlan"`
	InspectionPlant                string `json:"InspectionPlant"`
	Product                        string `json:"Product"`
}

type InspectionOrderKey struct {
	Product         []string `json:"Product"`
	BusinessPartner []int    `json:"BusinessPartner"`
	Plant           []string `json:"Plant"`
}

type InspectionOrder struct {
	InspectionOrder                int    `json:"InspectionOrder"`
	Product                        string `json:"Product"`
	InspectionPlantBusinessPartner int    `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                string `json:"InspectionPlant"`
}

// Item Pricing Element
type PriceMasterKey struct {
	Product                    []*string `json:"Product"`
	SupplyChainRelationshipID  []int     `json:"SupplyChainRelationshipID"`
	Buyer                      []int     `json:"Buyer"`
	Seller                     []int     `json:"Seller"`
	ConditionValidityEndDate   string    `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate string    `json:"ConditionValidityStartDate"`
}

type PriceMaster struct {
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	ConditionRecord            int      `json:"ConditionRecord"`
	ConditionSequentialNumber  int      `json:"ConditionSequentialNumber"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	Product                    string   `json:"Product"`
	ConditionType              string   `json:"ConditionType"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
}

type ConditionAmount struct {
	OrderItem                  int      `json:"OrderItem"`
	Product                    string   `json:"Product"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}

type ConditionRateValue struct {
	Product                       string   `json:"Product"`
	SupplyChainRelationshipID     int      `json:"SupplyChainRelationshipID"`
	TaxCode                       string   `json:"TaxCode"`
	PriceMasterConditionRateValue *float32 `json:"PriceMasterConditionRateValue"`
	TaxRate                       *float32 `json:"TaxRate"`
	ConditionRateValue            *float32 `json:"ConditionRateValue"`
	ConditionQuantity             *float32 `json:"ConditionQuantity"`
	ConditionAmount               *float32 `json:"ConditionAmount"`
	ConditionIsManuallyChanged    *bool    `json:"ConditionIsManuallyChanged"`
}

type ConditionIsManuallyChanged struct {
	Product                    string `json:"Product"`
	ConditionIsManuallyChanged *bool  `json:"ConditionIsManuallyChanged"`
}

type PricingProcedureCounter struct {
	Product                   string `json:"Product"`
	SupplyChainRelationshipID int    `json:"SupplyChainRelationshipID"`
	Buyer                     int    `json:"Buyer"`
	Seller                    int    `json:"Seller"`
	PricingProcedureCounter   []int  `json:"PricingProcedureCounter"`
}

// Amount関連の計算
type NetAmount struct {
	OrderItem int      `json:"OrderItem"`
	Product   string   `json:"Product"`
	NetAmount *float32 `json:"NetAmount"`
}

type TaxAmount struct {
	OrderItem int      `json:"OrderItem"`
	Product   string   `json:"Product"`
	TaxCode   string   `json:"TaxCode"`
	TaxRate   *float32 `json:"TaxRate"`
	NetAmount *float32 `json:"NetAmount"`
	TaxAmount *float32 `json:"TaxAmount"`
}

type GrossAmount struct {
	OrderItem   int      `json:"OrderItem"`
	Product     string   `json:"Product"`
	NetAmount   *float32 `json:"NetAmount"`
	TaxAmount   *float32 `json:"TaxAmount"`
	GrossAmount *float32 `json:"GrossAmount"`
}

// Address
type AddressKey struct {
	AddressID       []*int `json:"AddressID"`
	ValidityEndDate string `json:"ValidityEndDate"`
}

type Address struct {
	AddressID       int     `json:"AddressID"`
	ValidityEndDate string  `json:"ValidityEndDate"`
	PostalCode      string  `json:"PostalCode"`
	LocalRegion     string  `json:"LocalRegion"`
	Country         string  `json:"Country"`
	District        *string `json:"District"`
	StreetName      string  `json:"StreetName"`
	CityName        string  `json:"CityName"`
	Building        *string `json:"Building"`
	Floor           *int    `json:"Floor"`
	Room            *int    `json:"Room"`
}

type AddressMaster struct {
	AddressID         int     `json:"AddressID"`
	ValidityEndDate   string  `json:"ValidityEndDate"`
	ValidityStartDate string  `json:"ValidityStartDate"`
	PostalCode        string  `json:"PostalCode"`
	LocalRegion       string  `json:"LocalRegion"`
	Country           string  `json:"Country"`
	GlobalRegion      string  `json:"GlobalRegion"`
	TimeZone          string  `json:"TimeZone"`
	District          *string `json:"District"`
	StreetName        string  `json:"StreetName"`
	CityName          string  `json:"CityName"`
	Building          *string `json:"Building"`
	Floor             *int    `json:"Floor"`
	Room              *int    `json:"Room"`
}

// 数量単位変換実行の是非の判定
type QuantityUnitConversionKey struct {
	Product      string `json:"Product"`
	BaseUnit     string `json:"BaseUnit"`
	DeliveryUnit string `json:"DeliveryUnit"`
}

type QuantityUnitConversionQueryGets struct {
	Product               string  `json:"Product"`
	QuantityUnitFrom      string  `json:"QuantityUnitFrom"`
	QuantityUnitTo        string  `json:"QuantityUnitTo"`
	ConversionCoefficient float32 `json:"ConversionCoefficient"`
}

type QuantityUnitConversion struct {
	OrderItem                   int     `json:"OrderItem"`
	Product                     string  `json:"Product"`
	ConversionCoefficient       float32 `json:"ConversionCoefficient"`
	OrderQuantityInDeliveryUnit float32 `json:"OrderQuantityInDeliveryUnit"`
}

type OrderQuantityInDeliveryUnit struct {
	OrderItem                   int     `json:"OrderItem"`
	OrderQuantityInDeliveryUnit float32 `json:"OrderQuantityInDeliveryUnit"`
}

// Partner
type Partner struct {
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	AddressID               *int    `json:"AddressID"`
}

// 日付等の処理
type CreationDateItem struct {
	CreationDate string `json:"CreationDate"`
}

type CreationTimeItem struct {
	CreationTime string `json:"CreationTime"`
}

type LastChangeDateItem struct {
	LastChangeDate string `json:"LastChangeDate"`
}

type LastChangeTimeItem struct {
	LastChangeTime string `json:"LastChangeTime"`
}
