package api_input_reader

type SDC struct {
	ConnectionKey         string                `json:"connection_key"`
	Result                bool                  `json:"result"`
	RedisKey              string                `json:"redis_key"`
	Filepath              string                `json:"filepath"`
	APIStatusCode         int                   `json:"api_status_code"`
	RuntimeSessionID      string                `json:"runtime_session_id"`
	BusinessPartnerID     *int                  `json:"business_partner"`
	ServiceLabel          string                `json:"service_label"`
	OrdersInputParameters OrdersInputParameters `json:"OrdersInputParameters"`
	Header                Header                `json:"Orders"`
	APISchema             string                `json:"api_schema"`
	Accepter              []string              `json:"accepter"`
	Deleted               bool                  `json:"deleted"`
}

type OrdersInputParameters struct {
	ReferenceDocument     *int `json:"ReferenceDocument"`
	ReferenceDocumentItem *int `json:"ReferenceDocumentItem"`
}

type Header struct {
	OrderID                         int       `json:"OrderID"`
	OrderDate                       *string   `json:"OrderDate"`
	OrderType                       *string   `json:"OrderType"`
	Buyer                           *int      `json:"Buyer"`
	Seller                          *int      `json:"Seller"`
	CreationDate                    *string   `json:"CreationDate"`
	LastChangeDate                  *string   `json:"LastChangeDate"`
	ContractType                    *string   `json:"ContractType"`
	ValidityStartDate               *string   `json:"ValidityStartDate"`
	ValidityEndDate                 *string   `json:"ValidityEndDate"`
	InvoicePeriodStartDate          *string   `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate            *string   `json:"InvoicePeriodEndDate"`
	TotalNetAmount                  *float32  `json:"TotalNetAmount"`
	TotalTaxAmount                  *float32  `json:"TotalTaxAmount"`
	TotalGrossAmount                *float32  `json:"TotalGrossAmount"`
	HeaderDeliveryStatus            *string   `json:"HeaderDeliveryStatus"`
	HeaderBlockStatus               *bool     `json:"HeaderBlockStatus"`
	HeaderBillingStatus             *string   `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus        *string   `json:"HeaderDocReferenceStatus"`
	TransactionCurrency             *string   `json:"TransactionCurrency"`
	PricingDate                     *string   `json:"PricingDate"`
	PriceDetnExchangeRate           *float32  `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate           *string   `json:"RequestedDeliveryDate"`
	HeaderCompleteDeliveryIsDefined *bool     `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderBillingBlockStatus        *bool     `json:"HeaderBillingBlockStatus"`
	HeaderDeliveryBlockStatus       *bool     `json:"HeaderDeliveryBlockStatus"`
	Incoterms                       *string   `json:"Incoterms"`
	BillFromParty                   *int      `json:"BillFromParty"`
	BillToParty                     *int      `json:"BillToParty"`
	BillFromCountry                 *string   `json:"BillFromCountry"`
	BillToCountry                   *string   `json:"BillToCountry"`
	Payer                           *int      `json:"Payer"`
	Payee                           *int      `json:"Payee"`
	PaymentTerms                    *string   `json:"PaymentTerms"`
	PaymentMethod                   *string   `json:"PaymentMethod"`
	ReferenceDocument               *int      `json:"ReferenceDocument"`
	ReferenceDocumentItem           *int      `json:"ReferenceDocumentItem"`
	BPAccountAssignmentGroup        *string   `json:"BPAccountAssignmentGroup"`
	AccountingExchangeRate          *float32  `json:"AccountingExchangeRate"`
	InvoiceDocumentDate             *string   `json:"InvoiceDocumentDate"`
	IsExportImportDelivery          *bool     `json:"IsExportImportDelivery"`
	HeaderText                      *string   `json:"HeaderText"`
	Item                            []Item    `json:"Item"`
	Partner                         []Partner `json:"Partner"`
	Address                         []Address `json:"Address"`
}

type Partner struct {
	OrderID                 int     `json:"OrderID"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

type Address struct {
	OrderID     int     `json:"OrderID"`
	AddressID   int     `json:"AddressID"`
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}

type Item struct {
	OrderID                                       int                  `json:"OrderID"`
	OrderItem                                     int                  `json:"OrderItem"`
	OrderItemCategory                             *string              `json:"OrderItemCategory"`
	SupplyChainRelationshipID                     *int                 `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int                 `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int                 `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int                 `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int                 `json:"SupplyChainRelationshipProductionPlantID"`
	OrderItemText                                 *string              `json:"OrderItemText"`
	OrderItemTextByBuyer                          *string              `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                         *string              `json:"OrderItemTextBySeller"`
	Product                                       *string              `json:"Product"`
	ProductStandardID                             *string              `json:"ProductStandardID"`
	ProductGroup                                  *string              `json:"ProductGroup"`
	BaseUnit                                      *string              `json:"BaseUnit"`
	PricingDate                                   *string              `json:"PricingDate"`
	PriceDetnExchangeRate                         *float32             `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                         *string              `json:"RequestedDeliveryDate"`
	DeliverToParty                                *int                 `json:"DeliverToParty"`
	DeliverFromParty                              *int                 `json:"DeliverFromParty"`
	CreationDate                                  *string              `json:"CreationDate"`
	LastChangeDate                                *string              `json:"LastChangeDate"`
	DeliverToPlant                                *string              `json:"DeliverToPlant"`
	DeliverToPlantTimeZone                        *string              `json:"DeliverToPlantTimeZone"`
	DeliverToPlantStorageLocation                 *string              `json:"DeliverToPlantStorageLocation"`
	ProductIsBatchManagedInDeliverToPlant         *bool                `json:"ProductIsBatchManagedInDeliverToPlant"`
	BatchMgmtPolicyInDeliverToPlant               *string              `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                           *string              `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate          *string              `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityEndDate            *string              `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverFromPlant                              *string              `json:"DeliverFromPlant"`
	DeliverFromPlantTimeZone                      *string              `json:"DeliverFromPlantTimeZone"`
	DeliverFromPlantStorageLocation               *string              `json:"DeliverFromPlantStorageLocation"`
	ProductIsBatchManagedInDeliverFromPlant       *bool                `json:"ProductIsBatchManagedInDeliverFromPlant"`
	BatchMgmtPolicyInDeliverFromPlant             *string              `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                         *string              `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate        *string              `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityEndDate          *string              `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliveryUnit                                  *string              `json:"DeliveryUnit"`
	StockConfirmationBusinessPartner              *int                 `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string              `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone                *string              `json:"StockConfirmationPlantTimeZone"`
	ProductIsBatchManagedInStockConfirmationPlant *bool                `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	BatchMgmtPolicyStockConfirmationInPlant       *string              `json:"BatchMgmtPolicyStockConfirmationInPlant"`
	StockConfirmationPlantBatch                   *string              `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string              `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate    *string              `json:"StockConfirmationPlantBatchValidityEndDate"`
	ServicesRenderingDate                         *string              `json:"ServicesRenderingDate"`
	OrderQuantityInBaseUnit                       *float32             `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit                   *float32             `json:"OrderQuantityInDeliveryUnit"`
	StockConfirmationPolicy                       *string              `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       *string              `json:"StockConfirmationStatus"`
	ConfirmedOrderQuantityInBaseUnit              *float32             `json:"ConfirmedOrderQuantityInBaseUnit"`
	ItemWeightUnit                                *string              `json:"ItemWeightUnit"`
	ProductGrossWeight                            *float32             `json:"ProductGrossWeight"`
	ItemGrossWeight                               *float32             `json:"ItemGrossWeight"`
	ProductNetWeight                              *float32             `json:"ProductNetWeight"`
	ItemNetWeight                                 *float32             `json:"ItemNetWeight"`
	NetAmount                                     *float32             `json:"NetAmount"`
	TaxAmount                                     *float32             `json:"TaxAmount"`
	GrossAmount                                   *float32             `json:"GrossAmount"`
	InvoiceDocumentDate                           *string              `json:"InvoiceDocumentDate"`
	ProductionPlantBusinessPartner                *int                 `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string              `json:"ProductionPlant"`
	ProductionPlantTimeZone                       *string              `json:"ProductionPlantTimeZone"`
	ProductionPlantStorageLocation                *string              `json:"ProductionPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *string              `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionPlant              *bool                `json:"BatchMgmtPolicyInProductionPlant"`
	ProductionPlantBatch                          *string              `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate         *string              `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityEndDate           *string              `json:"ProductionPlantBatchValidityEndDate"`
	Incoterms                                     *string              `json:"Incoterms"`
	TransactionTaxClassification                  *string              `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         *string              `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       *string              `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      *string              `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                        *string              `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 *string              `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                                  *string              `json:"PaymentTerms"`
	DueCalculationBaseDate                        *string              `json:"DueCalculationBaseDate"`
	PaymentDueDate                                *string              `json:"PaymentDueDate"`
	NetPaymentDays                                *int                 `json:"NetPaymentDays"`
	PaymentMethod                                 *string              `json:"PaymentMethod"`
	Project                                       *string              `json:"Project"`
	AccountingExchangeRate                        *float32             `json:"AccountingExchangeRate"`
	ReferenceDocument                             *int                 `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *int                 `json:"ReferenceDocumentItem"`
	ItemCompleteDeliveryIsDefined                 *bool                `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus                            *string              `json:"ItemDeliveryStatus"`
	IssuingStatus                                 *string              `json:"IssuingStatus"`
	ReceivingStatus                               *string              `json:"ReceivingStatus"`
	ItemBillingStatus                             *string              `json:"ItemBillingStatus"`
	TaxCode                                       *string              `json:"TaxCode"`
	TaxRate                                       *float32             `json:"TaxRate"`
	CountryOfOrigin                               *string              `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                       *string              `json:"CountryOfOriginLanguage"`
	ItemBlockStatus                               *bool                `json:"ItemBlockStatus"`
	ItemDeliveryBlockStatus                       *bool                `json:"ItemDeliveryBlockStatus"`
	ItemBillingBlockStatus                        *bool                `json:"ItemBillingBlockStatus"`
	ItemPricingElement                            []ItemPricingElement `json:"ItemPricingElement"`
	ItemScheduleLine                              []ItemScheduleLine   `json:"ItemScheduleLine"`
}

type ItemPricingElement struct {
	OrderID                    int      `json:"OrderID"`
	OrderItem                  int      `json:"OrderItem"`
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}

type ItemScheduleLine struct {
	OrderID                                      int      `json:"OrderID"`
	OrderItem                                    int      `json:"OrderItem"`
	ScheduleLine                                 int      `json:"ScheduleLine"`
	SupplyChainRelationshipID                    *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID      *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                      *string  `json:"Product"`
	StockConfirmationBussinessPartner            *string  `json:"StockConfirmationBussinessPartner"`
	StockConfirmationPlant                       *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone               *string  `json:"StockConfirmationPlantTimeZone"`
	StockConfirmationPlantBatch                  *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate   *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	RequestedDeliveryDate                        *string  `json:"RequestedDeliveryDate"`
	ConfirmedDeliveryDate                        *string  `json:"ConfirmedDeliveryDate"`
	OrderQuantityInBaseUnit                      *float32 `json:"OrderQuantityInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheck        *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheck"`
	DeliveredQuantityInBaseUnit                  *float32 `json:"DeliveredQuantityInBaseUnit"`
	OpenConfirmedQuantityInBaseUnit              *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyConfirmed                        *bool    `json:"StockIsFullyConfirmed"`
	PlusMinusFlag                                *string  `json:"PlusMinusFlag"`
	ItemScheduleLineDeliveryBlockStatus          *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
}
