package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"encoding/json"
	"reflect"

	"golang.org/x/xerrors"
)

func ConvertToItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Item, error) {
	var err error

	itemCategoryIsINVPMap := StructArrayToMap(psdc.ItemCategoryIsINVP, "Product")
	productMasterGeneralMap := StructArrayToMap(psdc.ProductMasterGeneral, "Product")
	orderItemTextMap := StructArrayToMap(psdc.OrderItemText, "Product")
	orderItemTextByBuyerMap := StructArrayToMap(psdc.OrderItemTextByBuyer, "Product")
	orderItemTextBySellerMap := StructArrayToMap(psdc.OrderItemTextBySeller, "Product")
	supplyChainRelationshipDeliveryPlantRelationProductMap := StructArrayToMap(psdc.SupplyChainRelationshipDeliveryPlantRelationProduct, "Product")
	supplyChainRelationshipProductMasterBPPlantDeliverToMap := StructArrayToMap(psdc.SupplyChainRelationshipProductMasterBPPlantDeliverTo, "Product")
	supplyChainRelationshipProductMasterBPPlantDeliverFromMap := StructArrayToMap(psdc.SupplyChainRelationshipProductMasterBPPlantDeliverFrom, "Product")
	stockConfPlantRelationProductMap := StructArrayToMap(psdc.StockConfPlantRelationProduct, "Product")
	stockConfPlantProductMasterBPPlantMap := StructArrayToMap(psdc.StockConfPlantProductMasterBPPlant, "Product")
	confirmedOrderQuantityInBaseUnitMap := StructArrayToMap(psdc.ConfirmedOrderQuantityInBaseUnit, "OrderItem")
	itemGrossWeightMap := StructArrayToMap(psdc.ItemGrossWeight, "OrderItem")
	itemNetWeightMap := StructArrayToMap(psdc.ItemNetWeight, "Product")
	productionPlantRelationProductMap := StructArrayToMap(psdc.ProductionPlantRelationProduct, "Product")
	productionPlantProductMasterBPPlantMap := StructArrayToMap(psdc.ProductionPlantProductMasterBPPlant, "Product")
	productTaxClassificationBillToCountryMap := StructArrayToMap(psdc.ProductTaxClassificationBillToCountry, "Product")
	productTaxClassificationBillFromCountryMap := StructArrayToMap(psdc.ProductTaxClassificationBillFromCountry, "Product")
	definedTaxClassificationMap := StructArrayToMap(psdc.DefinedTaxClassification, "Product")
	taxCodeMap := StructArrayToMap(psdc.TaxCode, "Product")
	taxRateMap := StructArrayToMap(psdc.TaxRate, "TaxCode")
	netAmountMap := StructArrayToMap(psdc.NetAmount, "Product")
	taxAmounteMap := StructArrayToMap(psdc.TaxAmount, "Product")
	grossAmountMap := StructArrayToMap(psdc.GrossAmount, "Product")
	orderQuantityInDeliveryUnitMap := StructArrayToMap(psdc.OrderQuantityInDeliveryUnit, "OrderItem")

	res := make([]*Item, 0, len(sdc.Header.Item))
	for i, v := range sdc.Header.Item {
		item := &Item{}
		inputItem := sdc.Header.Item[i]

		// 入力ファイル
		item, err = jsonTypeConversion(item, inputItem)
		if err != nil {
			return nil, err
		}

		if v.Product == nil {
			continue
		}
		product := *v.Product

		item.OrderID = psdc.CalculateOrderID.OrderID
		item.OrderItem = psdc.OrderItem[i].OrderItemNumber
		item.OrderItemCategory = *productMasterGeneralMap[product].ItemCategory
		item.SupplyChainRelationshipID = psdc.SupplyChainRelationshipGeneral[0].SupplyChainRelationshipID
		item.SupplyChainRelationshipDeliveryPlantID = &psdc.SupplyChainRelationshipDeliveryPlantRelation[0].SupplyChainRelationshipDeliveryPlantID
		item.SupplyChainRelationshipDeliveryID = &psdc.SupplyChainRelationshipDeliveryRelation[0].SupplyChainRelationshipDeliveryID
		item.SupplyChainRelationshipStockConfPlantID = &psdc.StockConfPlantRelationProduct[0].SupplyChainRelationshipStockConfPlantID
		item.SupplyChainRelationshipProductionPlantID = &psdc.ProductionPlantRelationProduct[0].SupplyChainRelationshipProductionPlantID
		item.OrderItemText = *orderItemTextMap[product].OrderItemText
		item.OrderItemTextByBuyer = *orderItemTextByBuyerMap[product].ProductDescription
		item.OrderItemTextBySeller = *orderItemTextBySellerMap[product].ProductDescription
		item.Product = productMasterGeneralMap[product].Product
		item.ProductStandardID = *productMasterGeneralMap[product].ProductStandardID
		item.ProductGroup = productMasterGeneralMap[product].ProductGroup
		item.BaseUnit = *productMasterGeneralMap[product].BaseUnit
		item.PricingDate = psdc.ItemPricingDate[i].PricingDate
		item.PriceDetnExchangeRate = psdc.ItemPriceDetnExchangeRate[i].PriceDetnExchangeRate
		item.DeliverToParty = &psdc.SupplyChainRelationshipDeliveryRelation[0].DeliverToParty
		item.DeliverFromParty = &psdc.SupplyChainRelationshipDeliveryRelation[0].DeliverFromParty
		item.DeliverToPlant = &psdc.SupplyChainRelationshipDeliveryPlantRelation[0].DeliverToPlant
		item.CreationDate = psdc.CreationDateItem.CreationDate
		item.LastChangeDate = psdc.LastChangeDateItem.LastChangeDate

		item.DeliverToPlant = &psdc.SupplyChainRelationshipDeliveryPlantRelation[0].DeliverToPlant
		item.DeliverFromPlant = &psdc.SupplyChainRelationshipDeliveryPlantRelation[0].DeliverFromPlant

		item.OrderQuantityInDeliveryUnit = orderQuantityInDeliveryUnitMap[item.OrderItem].OrderQuantityInDeliveryUnit
		// item.StockConfirmationPolicy =  //TBD
		// item.StockConfirmationStatus =  //TBD

		item.ItemWeightUnit = productMasterGeneralMap[product].WeightUnit
		item.ProductGrossWeight = productMasterGeneralMap[product].GrossWeight
		item.ItemGrossWeight = itemGrossWeightMap[item.OrderItem].ItemGrossWeight
		item.ProductNetWeight = productMasterGeneralMap[product].NetWeight
		item.ItemNetWeight = itemNetWeightMap[product].ItemNetWeight
		item.InternalCapacityQuantity = productMasterGeneralMap[product].InternalCapacityQuantity
		item.InternalCapacityQuantityUnit = productMasterGeneralMap[product].InternalCapacityQuantityUnit
		item.NetAmount = netAmountMap[product].NetAmount
		item.TaxAmount = taxAmounteMap[product].TaxAmount
		item.GrossAmount = grossAmountMap[product].GrossAmount
		item.InvoiceDocumentDate = &psdc.ItemInvoiceDocumentDate[i].InvoiceDocumentDate

		item.Incoterms = psdc.Incoterms[i].Incoterms
		item.TransactionTaxClassification = *psdc.SupplyChainRelationshipBillingRelation[0].TransactionTaxClassification
		item.ProductTaxClassificationBillToCountry = *productTaxClassificationBillToCountryMap[product].ProductTaxClassifiication
		item.ProductTaxClassificationBillFromCountry = *productTaxClassificationBillFromCountryMap[product].ProductTaxClassifiication
		item.DefinedTaxClassification = definedTaxClassificationMap[product].DefinedTaxClassification
		item.AccountAssignmentGroup = *psdc.SupplyChainRelationshipTransaction[0].AccountAssignmentGroup
		item.ProductAccountAssignmentGroup = *productMasterGeneralMap[product].ProductAccountAssignmentGroup
		item.PaymentTerms = *psdc.ItemPaymentTerms[i].PaymentTerms
		// item.DueCalculationBaseDate =  //TBD
		// item.PaymentDueDate =  //TBD
		// item.NetPaymentDays =  //TBD
		item.PaymentMethod = *psdc.PaymentMethod[i].PaymentMethod

		item.AccountingExchangeRate = psdc.ItemAccountingExchangeRate[i].AccountingExchangeRate

		item.ItemCompleteDeliveryIsDefined = getBoolPtr(false)
		item.ItemDeliveryStatus = getStringPtr("NP")
		item.IssuingStatus = getStringPtr("NP")
		item.ReceivingStatus = getStringPtr("NP")
		item.ItemBillingStatus = getStringPtr("NP")
		item.TaxCode = taxCodeMap[product].TaxCode
		item.TaxRate = taxRateMap[*item.TaxCode].TaxRate
		item.CountryOfOrigin = productMasterGeneralMap[product].CountryOfOrigin
		item.CountryOfOriginLanguage = productMasterGeneralMap[product].CountryOfOriginLanguage
		item.ItemBlockStatus = getBoolPtr(false)
		item.ItemDeliveryBlockStatus = getBoolPtr(false)
		item.ItemBillingBlockStatus = getBoolPtr(false)
		item.ItemIsCancelled = getBoolPtr(false)
		item.ItemIsDeleted = getBoolPtr(false)

		if itemCategoryIsINVPMap[product].ItemCategoryIsINVP {
			item.DeliverToPlantTimeZone = psdc.DeliverToPlantTimeZone[0].TimeZone
			item.DeliverToPlantStorageLocation = &supplyChainRelationshipDeliveryPlantRelationProductMap[product].DeliverToPlantStorageLocation
			item.ProductIsBatchManagedInDeliverToPlant = supplyChainRelationshipProductMasterBPPlantDeliverToMap[product].IsBatchManagementRequired
			item.BatchMgmtPolicyInDeliverToPlant = supplyChainRelationshipProductMasterBPPlantDeliverToMap[product].BatchManagementPolicy

			item.DeliverFromPlantTimeZone = psdc.DeliverFromPlantTimeZone[0].TimeZone
			item.DeliverFromPlantStorageLocation = &supplyChainRelationshipDeliveryPlantRelationProductMap[product].DeliverFromPlantStorageLocation
			item.ProductIsBatchManagedInDeliverFromPlant = supplyChainRelationshipProductMasterBPPlantDeliverFromMap[product].IsBatchManagementRequired
			item.BatchMgmtPolicyInDeliverFromPlant = supplyChainRelationshipProductMasterBPPlantDeliverFromMap[product].BatchManagementPolicy

			item.DeliveryUnit = supplyChainRelationshipDeliveryPlantRelationProductMap[product].DeliveryUnit
			item.StockConfirmationBusinessPartner = &stockConfPlantRelationProductMap[product].StockConfirmationBusinessPartner
			item.StockConfirmationPlant = &stockConfPlantRelationProductMap[product].StockConfirmationPlant
			for _, v := range psdc.StockConfirmationPlantTimeZone {
				if v.BusinessPartner == *item.StockConfirmationBusinessPartner && v.Plant == *item.StockConfirmationPlant {
					item.StockConfirmationPlantTimeZone = v.TimeZone
					break
				}
			}

			item.ConfirmedOrderQuantityInBaseUnit = &confirmedOrderQuantityInBaseUnitMap[item.OrderItem].ConfirmedOrderQuantityInBaseUnit

			item.ProductIsBatchManagedInStockConfirmationPlant = stockConfPlantProductMasterBPPlantMap[product].IsBatchManagementRequired
			item.BatchMgmtPolicyInStockConfirmationPlant = stockConfPlantProductMasterBPPlantMap[product].BatchManagementPolicy

			item.ProductionPlantBusinessPartner = &productionPlantRelationProductMap[product].ProductionPlantBusinessPartner
			item.ProductionPlant = &productionPlantRelationProductMap[product].ProductionPlant
			for _, v := range psdc.ProductionPlantTimeZone {
				if v.BusinessPartner == *item.ProductionPlantBusinessPartner && v.Plant == *item.ProductionPlant {
					item.ProductionPlantTimeZone = v.TimeZone
					break
				}
			}
			item.ProductionPlantStorageLocation = productionPlantRelationProductMap[product].ProductionPlantStorageLocation
			item.ProductIsBatchManagedInProductionPlant = productionPlantProductMasterBPPlantMap[product].IsBatchManagementRequired
			item.BatchMgmtPolicyInProductionPlant = productionPlantProductMasterBPPlantMap[product].BatchManagementPolicy
		}

		res = append(res, item)
	}

	return res, nil
}

func ConvertToItemPricingElement(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*ItemPricingElement, error) {
	var err error

	priceMasterMap := StructArrayToMap(psdc.PriceMaster, "Product")
	conditionAmountMap := StructArrayToMap(psdc.ConditionAmount, "Product")
	conditionIsManuallyChangedMap := StructArrayToMap(psdc.ConditionIsManuallyChanged, "Product")
	taxCodeMap := StructArrayToMap(psdc.TaxCode, "Product")

	length := 0
	for _, v := range sdc.Header.Item {
		length += len(v.ItemPricingElement)
	}

	res := make([]*ItemPricingElement, 0, length)
	for i, item := range sdc.Header.Item {
		if item.Product == nil {
			continue
		}
		product := *item.Product

		supplyChainRelationshipID := psdc.SupplyChainRelationshipGeneral[0].SupplyChainRelationshipID

		inputItemPricingElement := item.ItemPricingElement[0]

		idx := -1
		for i, v := range psdc.PricingProcedureCounter {
			if v.Product == product && v.SupplyChainRelationshipID == supplyChainRelationshipID {
				idx = i
			}
		}
		if idx == -1 {
			continue
		}

		conditionTypeIsMWST := false
		for j := range psdc.PricingProcedureCounter[idx].PricingProcedureCounter {
			itemPricingElement := &ItemPricingElement{}

			// 入力ファイル
			itemPricingElement, err = jsonTypeConversion(itemPricingElement, inputItemPricingElement)
			if err != nil {
				return nil, err
			}

			itemPricingElement.OrderID = psdc.CalculateOrderID.OrderID
			itemPricingElement.OrderItem = psdc.OrderItem[i].OrderItemNumber
			itemPricingElement.PricingProcedureCounter = psdc.PricingProcedureCounter[idx].PricingProcedureCounter[j]
			itemPricingElement.PricingDate = &psdc.ItemPricingDate[i].PricingDate
			itemPricingElement.ConditionCurrency = psdc.SupplyChainRelationshipTransaction[0].TransactionCurrency
			itemPricingElement.ConditionQuantityUnit = item.BaseUnit
			itemPricingElement.TaxCode = taxCodeMap[product].TaxCode
			itemPricingElement.TransactionCurrency = psdc.SupplyChainRelationshipTransaction[0].TransactionCurrency

			if inputItemPricingElement.ConditionAmount == nil {
				itemPricingElement.SupplyChainRelationshipID = priceMasterMap[product].SupplyChainRelationshipID
				itemPricingElement.Buyer = priceMasterMap[product].Buyer
				itemPricingElement.Seller = priceMasterMap[product].Seller
				itemPricingElement.ConditionRecord = &priceMasterMap[product].ConditionRecord
				itemPricingElement.ConditionSequentialNumber = &priceMasterMap[product].ConditionSequentialNumber
				itemPricingElement.ConditionType = &priceMasterMap[product].ConditionType
				itemPricingElement.ConditionRateValue = priceMasterMap[product].ConditionRateValue
				itemPricingElement.ConditionQuantity = conditionAmountMap[product].ConditionQuantity
				itemPricingElement.ConditionAmount = conditionAmountMap[product].ConditionAmount
				itemPricingElement.ConditionIsManuallyChanged = conditionAmountMap[product].ConditionIsManuallyChanged
			} else {
				itemPricingElement.ConditionIsManuallyChanged = conditionIsManuallyChangedMap[product].ConditionIsManuallyChanged
			}

			if !conditionTypeIsMWST {
				for _, v := range psdc.ConditionRateValue {
					if v.Product == psdc.PricingProcedureCounter[idx].Product && v.SupplyChainRelationshipID == psdc.PricingProcedureCounter[idx].SupplyChainRelationshipID {
						conditionTypeIsMWST = true
						break
					}
				}
			}

			res = append(res, itemPricingElement)
		}

		// 200-3-2. Orders Item Pricing Elementデータの整列とセット(ConditionTypeが“MWST“の明細)
		if conditionTypeIsMWST {
			idx := -1
			for i, v := range psdc.ConditionRateValue {
				if v.Product == product && v.SupplyChainRelationshipID == supplyChainRelationshipID {
					idx = i
				}
			}
			if idx == -1 {
				continue
			}

			itemPricingElement := &ItemPricingElement{}

			// 入力ファイル
			itemPricingElement, err = jsonTypeConversion(itemPricingElement, inputItemPricingElement)
			if err != nil {
				return nil, err
			}

			itemPricingElement.OrderID = psdc.CalculateOrderID.OrderID
			itemPricingElement.OrderItem = psdc.OrderItem[i].OrderItemNumber
			itemPricingElement.PricingProcedureCounter = len(psdc.PricingProcedureCounter[idx].PricingProcedureCounter) + 1
			itemPricingElement.PricingDate = &psdc.ItemPricingDate[i].PricingDate
			itemPricingElement.ConditionCurrency = psdc.SupplyChainRelationshipTransaction[0].TransactionCurrency
			itemPricingElement.TaxCode = taxCodeMap[product].TaxCode
			itemPricingElement.TransactionCurrency = psdc.SupplyChainRelationshipTransaction[0].TransactionCurrency

			if inputItemPricingElement.ConditionAmount == nil {
				itemPricingElement.SupplyChainRelationshipID = priceMasterMap[product].SupplyChainRelationshipID
				itemPricingElement.Buyer = priceMasterMap[product].Buyer
				itemPricingElement.Seller = priceMasterMap[product].Seller
				itemPricingElement.ConditionRecord = &priceMasterMap[product].ConditionRecord
				itemPricingElement.ConditionSequentialNumber = getIntPtr(maxConditionSequentialNumber(psdc) + 1)
				itemPricingElement.ConditionType = getStringPtr("MWST")
				itemPricingElement.ConditionRateValue = psdc.ConditionRateValue[idx].ConditionRateValue
				itemPricingElement.ConditionQuantity = conditionAmountMap[product].ConditionQuantity
				itemPricingElement.ConditionAmount = psdc.ConditionRateValue[idx].ConditionAmount
				itemPricingElement.ConditionIsManuallyChanged = psdc.ConditionRateValue[idx].ConditionIsManuallyChanged
			} else {
				itemPricingElement.ConditionIsManuallyChanged = conditionIsManuallyChangedMap[product].ConditionIsManuallyChanged
			}
			res = append(res, itemPricingElement)
		}

	}

	return res, nil
}

func ConvertToItemScheduleLine(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*ItemScheduleLine, error) {
	var err error

	ordersItemScheduleLineMap := StructArrayToMap(psdc.OrdinaryStockConfirmationOrdersItemScheduleLine, "Product")

	length := 0
	for _, v := range sdc.Header.Item {
		length += len(v.ItemScheduleLine)
	}

	res := make([]*ItemScheduleLine, 0, length)
	for _, item := range sdc.Header.Item {
		for j := range item.ItemScheduleLine {
			if _, ok := ordersItemScheduleLineMap[*item.Product]; !ok {
				continue
			}

			itemScheduleLine := &ItemScheduleLine{}
			inputItemScheduleLine := item.ItemScheduleLine[j]

			// 入力ファイル
			itemScheduleLine, err = jsonTypeConversion(itemScheduleLine, inputItemScheduleLine)
			if err != nil {
				return nil, err
			}

			if item.Product == nil {
				continue
			}
			product := *item.Product

			itemScheduleLine.OrderID = ordersItemScheduleLineMap[product].OrderID
			itemScheduleLine.OrderItem = ordersItemScheduleLineMap[product].OrderItem
			itemScheduleLine.ScheduleLine = ordersItemScheduleLineMap[product].ScheduleLine
			itemScheduleLine.SupplyChainRelationshipID = ordersItemScheduleLineMap[product].SupplyChainRelationshipID
			itemScheduleLine.SupplyChainRelationshipStockConfPlantID = ordersItemScheduleLineMap[product].SupplyChainRelationshipStockConfPlantID
			itemScheduleLine.Product = ordersItemScheduleLineMap[product].Product
			itemScheduleLine.StockConfirmationBussinessPartner = ordersItemScheduleLineMap[product].StockConfirmationBussinessPartner
			itemScheduleLine.StockConfirmationPlant = ordersItemScheduleLineMap[product].StockConfirmationPlant
			itemScheduleLine.StockConfirmationPlantTimeZone = ordersItemScheduleLineMap[product].StockConfirmationPlantTimeZone
			itemScheduleLine.StockConfirmationPlantBatch = ordersItemScheduleLineMap[product].StockConfirmationPlantBatch
			itemScheduleLine.StockConfirmationPlantBatchValidityStartDate = ordersItemScheduleLineMap[product].StockConfirmationPlantBatchValidityStartDate
			itemScheduleLine.StockConfirmationPlantBatchValidityEndDate = ordersItemScheduleLineMap[product].StockConfirmationPlantBatchValidityEndDate
			itemScheduleLine.RequestedDeliveryDate = ordersItemScheduleLineMap[product].RequestedDeliveryDate
			itemScheduleLine.ConfirmedDeliveryDate = ordersItemScheduleLineMap[product].ConfirmedDeliveryDate
			itemScheduleLine.OrderQuantityInBaseUnit = ordersItemScheduleLineMap[product].OrderQuantityInBaseUnit
			itemScheduleLine.ConfirmedOrderQuantityByPDTAvailCheck = ordersItemScheduleLineMap[product].ConfirmedOrderQuantityByPDTAvailCheck
			itemScheduleLine.DeliveredQuantityInBaseUnit = ordersItemScheduleLineMap[product].DeliveredQuantityInBaseUnit
			itemScheduleLine.OpenConfirmedQuantityInBaseUnit = ordersItemScheduleLineMap[product].OpenConfirmedQuantityInBaseUnit
			itemScheduleLine.StockIsFullyConfirmed = ordersItemScheduleLineMap[product].StockIsFullyConfirmed
			itemScheduleLine.PlusMinusFlag = ordersItemScheduleLineMap[product].PlusMinusFlag
			itemScheduleLine.ItemScheduleLineDeliveryBlockStatus = ordersItemScheduleLineMap[product].ItemScheduleLineDeliveryBlockStatus

			res = append(res, itemScheduleLine)
		}
	}

	return res, nil
}

func ConvertToPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Partner, error) {
	var err error

	res := make([]*Partner, 0, len(psdc.Partner))
	for _, v := range psdc.Partner {
		partner := &Partner{}
		inputPartner := sdc.Header.Partner[0]

		// 入力ファイル
		partner, err = jsonTypeConversion(partner, inputPartner)
		if err != nil {
			return nil, err
		}

		partner.OrderID = psdc.CalculateOrderID.OrderID
		partner.PartnerFunction = v.PartnerFunction
		partner.BusinessPartner = v.BusinessPartner
		partner.BusinessPartnerFullName = v.BusinessPartnerFullName
		partner.BusinessPartnerName = v.BusinessPartnerName
		partner.Country = v.Country
		partner.Language = v.Language
		partner.Currency = v.Currency
		partner.AddressID = v.AddressID

		res = append(res, partner)
	}

	return res, nil
}

func ConvertToAddress(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Address, error) {
	var err error

	res := make([]*Address, 0, len(psdc.Address))
	for _, v := range psdc.Address {
		address := &Address{}
		inputAddress := sdc.Header.Address[0]

		// 入力ファイル
		address, err = jsonTypeConversion(address, inputAddress)
		if err != nil {
			return nil, err
		}

		address.OrderID = psdc.CalculateOrderID.OrderID
		address.AddressID = v.AddressID
		address.PostalCode = &v.PostalCode
		address.LocalRegion = &v.LocalRegion
		address.Country = &v.Country
		address.District = v.District
		address.StreetName = &v.StreetName
		address.CityName = &v.CityName
		address.Building = v.Building
		address.Floor = v.Floor
		address.Room = v.Room

		res = append(res, address)
	}

	return res, nil
}

func maxConditionSequentialNumber(psdc *api_processing_data_formatter.SDC) int {
	priceMaster := psdc.PriceMaster

	maxValue := priceMaster[0].ConditionSequentialNumber
	for _, v := range priceMaster {
		value := v.ConditionSequentialNumber
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func getBoolPtr(b bool) *bool {
	return &b
}

func getStringPtr(s string) *string {
	return &s
}

func getIntPtr(i int) *int {
	return &i
}

func StructArrayToMap[T any](data []T, key string) map[any]T {
	res := make(map[any]T, len(data))

	for _, value := range data {
		m := StructToMap[T](&value, key)
		for k, v := range m {
			res[k] = v
		}
	}

	return res
}

func StructToMap[T any](data interface{}, key string) map[any]T {
	res := make(map[any]T)
	elem := reflect.Indirect(reflect.ValueOf(data).Elem())
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		if field == key {
			rv := reflect.ValueOf(elem.Field(i).Interface())
			if rv.Kind() == reflect.Ptr {
				if rv.IsNil() {
					return nil
				}
			}
			value := reflect.Indirect(elem.Field(i)).Interface()
			var dist T
			res[value], _ = jsonTypeConversion(dist, elem.Interface())
			break
		}
	}

	return res
}

func jsonTypeConversion[T any](dist T, data interface{}) (T, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
