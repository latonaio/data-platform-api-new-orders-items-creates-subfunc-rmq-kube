package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-orders-creates-subfunc-items-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-creates-subfunc-items-rmq-kube/API_Processing_Data_Formatter"
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
	InspectionPlanMap := StructArrayToMap(psdc.InspectionPlan, "Product")
	InspectionOrderMap := StructArrayToMap(psdc.InspectionOrder, "Product")
	productTaxClassificationBillToCountryMap := StructArrayToMap(psdc.ProductTaxClassificationBillToCountry, "Product")
	productTaxClassificationBillFromCountryMap := StructArrayToMap(psdc.ProductTaxClassificationBillFromCountry, "Product")
	definedTaxClassificationMap := StructArrayToMap(psdc.DefinedTaxClassification, "Product")
	taxCodeMap := StructArrayToMap(psdc.TaxCode, "Product")
	taxRateMap := StructArrayToMap(psdc.TaxRate, "TaxCode")
	netAmountMap := StructArrayToMap(psdc.NetAmount, "Product")
	taxAmounteMap := StructArrayToMap(psdc.TaxAmount, "Product")
	grossAmountMap := StructArrayToMap(psdc.GrossAmount, "Product")
	orderQuantityInDeliveryUnitMap := StructArrayToMap(psdc.OrderQuantityInDeliveryUnit, "OrderItem")
	stockConfirmationStatusMap := StructArrayToMap(psdc.StockConfirmationStatus, "OrderItem")

	items := make([]*Item, 0, len(sdc.Header.Item))
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
		orderItem := psdc.OrderItem[i].OrderItemNumber

		item.OrderID = sdc.Header.OrderID
		item.OrderItem = orderItem
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
		item.CreationTime = psdc.CreationTimeItem.CreationTime
		item.LastChangeDate = psdc.LastChangeDateItem.LastChangeDate
		item.LastChangeTime = psdc.LastChangeTimeItem.LastChangeTime

		item.DeliverToPlant = &psdc.SupplyChainRelationshipDeliveryPlantRelation[0].DeliverToPlant
		item.DeliverFromPlant = &psdc.SupplyChainRelationshipDeliveryPlantRelation[0].DeliverFromPlant

		item.OrderQuantityInDeliveryUnit = orderQuantityInDeliveryUnitMap[orderItem].OrderQuantityInDeliveryUnit
		// item.StockConfirmationPolicy =  //TBD
		item.StockConfirmationStatus = stockConfirmationStatusMap[orderItem].StockConfirmationStatus

		item.ItemWeightUnit = productMasterGeneralMap[product].WeightUnit
		item.ProductGrossWeight = productMasterGeneralMap[product].GrossWeight
		item.ItemGrossWeight = itemGrossWeightMap[orderItem].ItemGrossWeight
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
		item.IsCancelled = getBoolPtr(false)
		item.IsMarkedForDeletion = getBoolPtr(false)

		if itemCategoryIsINVPMap[product].ItemCategoryIsINVP {
			item.DeliverToPlantTimeZone = psdc.DeliverToPlantTimeZone[0].TimeZone
			if _, ok := supplyChainRelationshipDeliveryPlantRelationProductMap[product]; ok {
				item.DeliverToPlantStorageLocation = &supplyChainRelationshipDeliveryPlantRelationProductMap[product].DeliverToPlantStorageLocation
				item.DeliverFromPlantStorageLocation = &supplyChainRelationshipDeliveryPlantRelationProductMap[product].DeliverFromPlantStorageLocation
				item.DeliveryUnit = supplyChainRelationshipDeliveryPlantRelationProductMap[product].DeliveryUnit
			}
			if _, ok := supplyChainRelationshipProductMasterBPPlantDeliverToMap[product]; ok {
				item.ProductIsBatchManagedInDeliverToPlant = supplyChainRelationshipProductMasterBPPlantDeliverToMap[product].IsBatchManagementRequired
				item.BatchMgmtPolicyInDeliverToPlant = supplyChainRelationshipProductMasterBPPlantDeliverToMap[product].BatchManagementPolicy
			}

			item.DeliverFromPlantTimeZone = psdc.DeliverFromPlantTimeZone[0].TimeZone
			if _, ok := supplyChainRelationshipProductMasterBPPlantDeliverFromMap[product]; ok {
				item.ProductIsBatchManagedInDeliverFromPlant = supplyChainRelationshipProductMasterBPPlantDeliverFromMap[product].IsBatchManagementRequired
				item.BatchMgmtPolicyInDeliverFromPlant = supplyChainRelationshipProductMasterBPPlantDeliverFromMap[product].BatchManagementPolicy
			}

			if _, ok := stockConfPlantRelationProductMap[product]; ok {
				item.StockConfirmationBusinessPartner = &stockConfPlantRelationProductMap[product].StockConfirmationBusinessPartner
				item.StockConfirmationPlant = &stockConfPlantRelationProductMap[product].StockConfirmationPlant
			}

			for _, v := range psdc.StockConfirmationPlantTimeZone {
				if v.BusinessPartner == *item.StockConfirmationBusinessPartner && v.Plant == *item.StockConfirmationPlant {
					item.StockConfirmationPlantTimeZone = v.TimeZone
					break
				}
			}

			if _, ok := confirmedOrderQuantityInBaseUnitMap[orderItem]; ok {
				item.ConfirmedOrderQuantityInBaseUnit = &confirmedOrderQuantityInBaseUnitMap[orderItem].ConfirmedOrderQuantityInBaseUnit
			}

			if _, ok := stockConfPlantProductMasterBPPlantMap[product]; ok {
				item.ProductIsBatchManagedInStockConfirmationPlant = stockConfPlantProductMasterBPPlantMap[product].IsBatchManagementRequired
				item.BatchMgmtPolicyInStockConfirmationPlant = stockConfPlantProductMasterBPPlantMap[product].BatchManagementPolicy
			}

			if _, ok := productionPlantRelationProductMap[product]; ok {
				item.ProductionPlantBusinessPartner = &productionPlantRelationProductMap[product].ProductionPlantBusinessPartner
				item.ProductionPlant = &productionPlantRelationProductMap[product].ProductionPlant
				item.ProductionPlantStorageLocation = productionPlantRelationProductMap[product].ProductionPlantStorageLocation
			}
			for _, v := range psdc.ProductionPlantTimeZone {
				if v.BusinessPartner == *item.ProductionPlantBusinessPartner && v.Plant == *item.ProductionPlant {
					item.ProductionPlantTimeZone = v.TimeZone
					break
				}
			}

			if _, ok := productionPlantProductMasterBPPlantMap[product]; ok {
				item.ProductIsBatchManagedInProductionPlant = productionPlantProductMasterBPPlantMap[product].IsBatchManagementRequired
				item.BatchMgmtPolicyInProductionPlant = productionPlantProductMasterBPPlantMap[product].BatchManagementPolicy
			}

			if _, ok := InspectionPlanMap[product]; ok {
				item.InspectionPlan = &InspectionPlanMap[product].InspectionPlan
				item.InspectionPlant = &InspectionPlanMap[product].InspectionPlant
			}

			if _, ok := InspectionOrderMap[product]; ok {
				item.InspectionOrder = &InspectionOrderMap[product].InspectionOrder
			}

		}

		items = append(items, item)
	}

	return items, nil
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

	itemPricingElements := make([]*ItemPricingElement, 0, length)
	for i, item := range sdc.Header.Item {
		if item.Product == nil {
			continue
		}
		product := *item.Product

		supplyChainRelationshipID := psdc.SupplyChainRelationshipGeneral[0].SupplyChainRelationshipID

		inputItemPricingElement := item.ItemPricingElement[0]

		pricingProcedureCounterIdx := -1
		for i, v := range psdc.PricingProcedureCounter {
			if v.Product == product && v.SupplyChainRelationshipID == supplyChainRelationshipID {
				pricingProcedureCounterIdx = i
			}
		}
		if pricingProcedureCounterIdx == -1 {
			continue
		}

		conditionTypeIsMWST := false
		for j := range psdc.PricingProcedureCounter[pricingProcedureCounterIdx].PricingProcedureCounter {
			itemPricingElement := &ItemPricingElement{}

			// 入力ファイル
			itemPricingElement, err = jsonTypeConversion(itemPricingElement, inputItemPricingElement)
			if err != nil {
				return nil, err
			}

			itemPricingElement.OrderID = sdc.Header.OrderID
			itemPricingElement.OrderItem = psdc.OrderItem[i].OrderItemNumber
			itemPricingElement.PricingProcedureCounter = psdc.PricingProcedureCounter[pricingProcedureCounterIdx].PricingProcedureCounter[j]
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
					if v.Product == psdc.PricingProcedureCounter[pricingProcedureCounterIdx].Product && v.SupplyChainRelationshipID == psdc.PricingProcedureCounter[pricingProcedureCounterIdx].SupplyChainRelationshipID {
						conditionTypeIsMWST = true
						break
					}
				}
			}

			itemPricingElements = append(itemPricingElements, itemPricingElement)
		}

		// 200-3-2. Orders Item Pricing Elementデータの整列とセット(ConditionTypeが“MWST“の明細)
		if conditionTypeIsMWST {
			conditionRateValueIdx := -1
			for i, v := range psdc.ConditionRateValue {
				if v.Product == product && v.SupplyChainRelationshipID == supplyChainRelationshipID {
					conditionRateValueIdx = i
					break
				}
			}
			if conditionRateValueIdx == -1 {
				continue
			}

			itemPricingElement := &ItemPricingElement{}

			// 入力ファイル
			itemPricingElement, err = jsonTypeConversion(itemPricingElement, inputItemPricingElement)
			if err != nil {
				return nil, err
			}

			itemPricingElement.OrderID = sdc.Header.OrderID
			itemPricingElement.OrderItem = psdc.OrderItem[i].OrderItemNumber
			itemPricingElement.PricingProcedureCounter = len(psdc.PricingProcedureCounter[pricingProcedureCounterIdx].PricingProcedureCounter) + 1
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
				itemPricingElement.ConditionRateValue = psdc.ConditionRateValue[conditionRateValueIdx].ConditionRateValue
				itemPricingElement.ConditionQuantity = conditionAmountMap[product].ConditionQuantity
				itemPricingElement.ConditionAmount = psdc.ConditionRateValue[conditionRateValueIdx].ConditionAmount
				itemPricingElement.ConditionIsManuallyChanged = psdc.ConditionRateValue[conditionRateValueIdx].ConditionIsManuallyChanged
			} else {
				itemPricingElement.ConditionIsManuallyChanged = conditionIsManuallyChangedMap[product].ConditionIsManuallyChanged
			}
			itemPricingElements = append(itemPricingElements, itemPricingElement)
		}

	}

	return itemPricingElements, nil
}

func ConvertToItemScheduleLine(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*ItemScheduleLine, error) {
	var err error

	ordersItemScheduleLineMap := StructArrayToMap(psdc.StockConfirmationOrdersItemScheduleLine, "OrderItem")

	length := 0
	for _, v := range sdc.Header.Item {
		length += len(v.ItemScheduleLine)
	}

	itemScheduleLines := make([]*ItemScheduleLine, 0, length)
	for i, orderItem := range psdc.OrderItem {
		item := sdc.Header.Item[i]
		orderItemNumber := orderItem.OrderItemNumber
		for j := range item.ItemScheduleLine {
			if _, ok := ordersItemScheduleLineMap[orderItemNumber]; !ok {
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

			itemScheduleLine.OrderID = ordersItemScheduleLineMap[orderItemNumber].OrderID
			itemScheduleLine.OrderItem = ordersItemScheduleLineMap[orderItemNumber].OrderItem
			itemScheduleLine.ScheduleLine = ordersItemScheduleLineMap[orderItemNumber].ScheduleLine
			itemScheduleLine.SupplyChainRelationshipID = ordersItemScheduleLineMap[orderItemNumber].SupplyChainRelationshipID
			itemScheduleLine.SupplyChainRelationshipStockConfPlantID = ordersItemScheduleLineMap[orderItemNumber].SupplyChainRelationshipStockConfPlantID
			itemScheduleLine.Product = ordersItemScheduleLineMap[orderItemNumber].Product
			itemScheduleLine.StockConfirmationBussinessPartner = ordersItemScheduleLineMap[orderItemNumber].StockConfirmationBussinessPartner
			itemScheduleLine.StockConfirmationPlant = ordersItemScheduleLineMap[orderItemNumber].StockConfirmationPlant
			itemScheduleLine.StockConfirmationPlantTimeZone = ordersItemScheduleLineMap[orderItemNumber].StockConfirmationPlantTimeZone
			itemScheduleLine.StockConfirmationPlantBatch = ordersItemScheduleLineMap[orderItemNumber].StockConfirmationPlantBatch
			itemScheduleLine.StockConfirmationPlantBatchValidityStartDate = ordersItemScheduleLineMap[orderItemNumber].StockConfirmationPlantBatchValidityStartDate
			itemScheduleLine.StockConfirmationPlantBatchValidityEndDate = ordersItemScheduleLineMap[orderItemNumber].StockConfirmationPlantBatchValidityEndDate
			itemScheduleLine.RequestedDeliveryDate = ordersItemScheduleLineMap[orderItemNumber].RequestedDeliveryDate
			itemScheduleLine.RequestedDeliveryTime = ordersItemScheduleLineMap[orderItemNumber].RequestedDeliveryTime
			itemScheduleLine.ConfirmedDeliveryDate = ordersItemScheduleLineMap[orderItemNumber].ConfirmedDeliveryDate
			itemScheduleLine.ScheduleLineOrderQuantity = ordersItemScheduleLineMap[orderItemNumber].ScheduleLineOrderQuantity
			itemScheduleLine.OriginalOrderQuantityInBaseUnit = ordersItemScheduleLineMap[orderItemNumber].OriginalOrderQuantityInBaseUnit
			itemScheduleLine.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit = ordersItemScheduleLineMap[orderItemNumber].ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit
			itemScheduleLine.ConfirmedOrderQuantityByPDTAvailCheck = ordersItemScheduleLineMap[orderItemNumber].ConfirmedOrderQuantityByPDTAvailCheck
			itemScheduleLine.DeliveredQuantityInBaseUnit = ordersItemScheduleLineMap[orderItemNumber].DeliveredQuantityInBaseUnit
			itemScheduleLine.OpenConfirmedQuantityInBaseUnit = ordersItemScheduleLineMap[orderItemNumber].OpenConfirmedQuantityInBaseUnit
			itemScheduleLine.StockIsFullyConfirmed = ordersItemScheduleLineMap[orderItemNumber].StockIsFullyConfirmed
			itemScheduleLine.PlusMinusFlag = ordersItemScheduleLineMap[orderItemNumber].PlusMinusFlag
			itemScheduleLine.ItemScheduleLineDeliveryBlockStatus = ordersItemScheduleLineMap[orderItemNumber].ItemScheduleLineDeliveryBlockStatus
			itemScheduleLine.IsCancelled = ordersItemScheduleLineMap[orderItemNumber].IsCancelled
			itemScheduleLine.IsMarkedForDeletion = ordersItemScheduleLineMap[orderItemNumber].IsMarkedForDeletion

			itemScheduleLines = append(itemScheduleLines, itemScheduleLine)
		}
	}

	return itemScheduleLines, nil
}

func ConvertToPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Partner, error) {
	var err error

	partners := make([]*Partner, 0, len(psdc.Partner))
	for _, v := range psdc.Partner {
		partner := &Partner{}
		inputPartner := sdc.Header.Partner[0]

		// 入力ファイル
		partner, err = jsonTypeConversion(partner, inputPartner)
		if err != nil {
			return nil, err
		}

		partner.OrderID = sdc.Header.OrderID
		partner.PartnerFunction = v.PartnerFunction
		partner.BusinessPartner = v.BusinessPartner
		partner.BusinessPartnerFullName = v.BusinessPartnerFullName
		partner.BusinessPartnerName = v.BusinessPartnerName
		partner.Country = v.Country
		partner.Language = v.Language
		partner.Currency = v.Currency
		partner.AddressID = v.AddressID

		partners = append(partners, partner)
	}

	return partners, nil
}

func ConvertToAddress(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Address, error) {
	var err error

	addresses := make([]*Address, 0, len(psdc.Address))
	for _, v := range psdc.Address {
		address := &Address{}
		inputAddress := sdc.Header.Address[0]

		// 入力ファイル
		address, err = jsonTypeConversion(address, inputAddress)
		if err != nil {
			return nil, err
		}

		address.OrderID = sdc.Header.OrderID
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

		addresses = append(addresses, address)
	}

	return addresses, nil
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
