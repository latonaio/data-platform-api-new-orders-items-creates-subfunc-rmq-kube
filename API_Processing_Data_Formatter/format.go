package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	"data-platform-api-orders-items-creates-subfunc-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
	// "data-platform-api-orders-items-creates-subfunc-rmq-kube/DPFM_API_Caller/requests"
)

// Initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) *MetaData {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}

	data := pm
	res := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &res
}

// Header
func (psdc *SDC) ConvertToSupplyChainRelationshipGeneral(rows *sql.Rows) ([]*SupplyChainRelationshipGeneral, error) {
	var res []*SupplyChainRelationshipGeneral

	for i := 0; true; i++ {
		pm := &requests.SupplyChainRelationshipGeneral{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_supply_chain_relationship_general_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &SupplyChainRelationshipGeneral{
			SupplyChainRelationshipID: data.SupplyChainRelationshipID,
			Buyer:                     data.Buyer,
			Seller:                    data.Seller,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToSupplyChainRelationshipDeliveryRelationKey() *SupplyChainRelationshipDeliveryRelationKey {
	pm := &requests.SupplyChainRelationshipDeliveryRelationKey{}

	data := pm
	res := SupplyChainRelationshipDeliveryRelationKey{
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
		DeliverToParty:            data.DeliverToParty,
		DeliverFromParty:          data.DeliverFromParty,
	}

	return &res
}

func (psdc *SDC) ConvertToSupplyChainRelationshipDeliveryRelation(rows *sql.Rows) ([]*SupplyChainRelationshipDeliveryRelation, error) {
	var res []*SupplyChainRelationshipDeliveryRelation

	for i := 0; true; i++ {
		pm := &requests.SupplyChainRelationshipDeliveryRelation{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_supply_chain_relationship_delivery_relation_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &SupplyChainRelationshipDeliveryRelation{
			SupplyChainRelationshipID:         data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID: data.SupplyChainRelationshipDeliveryID,
			Buyer:                             data.Buyer,
			Seller:                            data.Seller,
			DeliverToParty:                    data.DeliverToParty,
			DeliverFromParty:                  data.DeliverFromParty,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToSupplyChainRelationshipDeliveryPlantRelationKey() *SupplyChainRelationshipDeliveryPlantRelationKey {
	pm := &requests.SupplyChainRelationshipDeliveryPlantRelationKey{
		DefaultRelation: true,
	}

	data := pm
	res := SupplyChainRelationshipDeliveryPlantRelationKey{
		SupplyChainRelationshipID:         data.SupplyChainRelationshipID,
		SupplyChainRelationshipDeliveryID: data.SupplyChainRelationshipDeliveryID,
		Buyer:                             data.Buyer,
		Seller:                            data.Seller,
		DeliverToParty:                    data.DeliverToParty,
		DeliverFromParty:                  data.DeliverFromParty,
		DefaultRelation:                   data.DefaultRelation,
	}

	return &res
}

func (psdc *SDC) ConvertToSupplyChainRelationshipDeliveryPlantRelation(rows *sql.Rows) ([]*SupplyChainRelationshipDeliveryPlantRelation, error) {
	var res []*SupplyChainRelationshipDeliveryPlantRelation

	for i := 0; true; i++ {
		pm := &requests.SupplyChainRelationshipDeliveryPlantRelation{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_supply_chain_relationship_delivery_plant_relation_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.DefaultRelation,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &SupplyChainRelationshipDeliveryPlantRelation{
			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
			Buyer:                                  data.Buyer,
			Seller:                                 data.Seller,
			DeliverToParty:                         data.DeliverToParty,
			DeliverFromParty:                       data.DeliverFromParty,
			DeliverToPlant:                         data.DeliverToPlant,
			DeliverFromPlant:                       data.DeliverFromPlant,
			DefaultRelation:                        data.DefaultRelation,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToSupplyChainRelationshipTransaction(rows *sql.Rows) ([]*SupplyChainRelationshipTransaction, error) {
	var res []*SupplyChainRelationshipTransaction

	for i := 0; true; i++ {
		pm := &requests.SupplyChainRelationshipTransaction{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_supply_chain_relationship_transaction_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.TransactionCurrency,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.AccountAssignmentGroup,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &SupplyChainRelationshipTransaction{
			SupplyChainRelationshipID: data.SupplyChainRelationshipID,
			Buyer:                     data.Buyer,
			Seller:                    data.Seller,
			TransactionCurrency:       data.TransactionCurrency,
			Incoterms:                 data.Incoterms,
			PaymentTerms:              data.PaymentTerms,
			PaymentMethod:             data.PaymentMethod,
			AccountAssignmentGroup:    data.AccountAssignmentGroup,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToSupplyChainRelationshipBillingRelationKey() *SupplyChainRelationshipBillingRelationKey {
	pm := &requests.SupplyChainRelationshipBillingRelationKey{
		DefaultRelation: true,
	}

	data := pm
	res := SupplyChainRelationshipBillingRelationKey{
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
		DefaultRelation:           data.DefaultRelation,
	}

	return &res
}

func (psdc *SDC) ConvertToSupplyChainRelationshipBillingRelation(rows *sql.Rows) ([]*SupplyChainRelationshipBillingRelation, error) {
	var res []*SupplyChainRelationshipBillingRelation

	for i := 0; true; i++ {
		pm := &requests.SupplyChainRelationshipBillingRelation{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_supply_chain_relationship_delivery_plant_relation_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.DefaultRelation,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.IsExportImport,
			&pm.TransactionTaxCategory,
			&pm.TransactionTaxClassification,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &SupplyChainRelationshipBillingRelation{
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			BillToParty:                      data.BillToParty,
			BillFromParty:                    data.BillFromParty,
			DefaultRelation:                  data.DefaultRelation,
			BillToCountry:                    data.BillToCountry,
			BillFromCountry:                  data.BillFromCountry,
			IsExportImport:                   data.IsExportImport,
			TransactionTaxCategory:           data.TransactionTaxCategory,
			TransactionTaxClassification:     data.TransactionTaxClassification,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToCalculateOrderIDKey() *CalculateOrderIDKey {
	pm := &requests.CalculateOrderIDKey{
		FieldNameWithNumberRange: "OrderID",
	}

	data := pm
	res := CalculateOrderIDKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &res
}

func (psdc *SDC) ConvertToCalculateOrderIDQueryGets(rows *sql.Rows) (*CalculateOrderIDQueryGets, error) {
	pm := &requests.CalculateOrderIDQueryGets{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_number_range_latest_number_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.LatestNumber,
		)
		if err != nil {
			return nil, err
		}
	}

	data := pm
	res := CalculateOrderIDQueryGets{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
		LatestNumber:             data.LatestNumber,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCalculateOrderID(orderIDLatestNumber *int, orderID int) *CalculateOrderID {
	pm := &requests.CalculateOrderID{}

	pm.OrderIDLatestNumber = orderIDLatestNumber
	pm.OrderID = orderID

	data := pm
	res := CalculateOrderID{
		OrderIDLatestNumber: data.OrderIDLatestNumber,
		OrderID:             data.OrderID,
	}

	return &res
}

// Item
func (psdc *SDC) ConvertToOrderItem(sdc *api_input_reader.SDC) []*OrderItem {
	var res []*OrderItem

	for i := range sdc.Header.Item {
		pm := &requests.OrderItem{}

		pm.OrderItemNumber = i + 1

		data := pm
		res = append(res, &OrderItem{
			OrderItemNumber: data.OrderItemNumber,
		})
	}

	return res
}

func (psdc *SDC) ConvertToProductTaxClassificationKey() *ProductTaxClassificationKey {
	pm := &requests.ProductTaxClassificationKey{
		ProductTaxCategory: "MWST",
	}

	data := pm
	res := ProductTaxClassificationKey{
		Product:            data.Product,
		Country:            data.Country,
		ProductTaxCategory: data.ProductTaxCategory,
	}

	return &res
}

func (psdc *SDC) ConvertToProductTaxClassificationBillToCountry(rows *sql.Rows) ([]*ProductTaxClassificationBillToCountry, error) {
	var res []*ProductTaxClassificationBillToCountry

	for i := 0; true; i++ {
		pm := &requests.ProductTaxClassificationBillToCountry{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_product_master_tax_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Product,
			&pm.Country,
			&pm.ProductTaxCategory,
			&pm.ProductTaxClassifiication,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &ProductTaxClassificationBillToCountry{
			Product:                   data.Product,
			Country:                   data.Country,
			ProductTaxCategory:        data.ProductTaxCategory,
			ProductTaxClassifiication: data.ProductTaxClassifiication,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToProductTaxClassificationBillFromCountry(rows *sql.Rows) ([]*ProductTaxClassificationBillFromCountry, error) {
	var res []*ProductTaxClassificationBillFromCountry

	for i := 0; true; i++ {
		pm := &requests.ProductTaxClassificationBillFromCountry{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_product_master_tax_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Product,
			&pm.Country,
			&pm.ProductTaxCategory,
			&pm.ProductTaxClassifiication,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &ProductTaxClassificationBillFromCountry{
			Product:                   data.Product,
			Country:                   data.Country,
			ProductTaxCategory:        data.ProductTaxCategory,
			ProductTaxClassifiication: data.ProductTaxClassifiication,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToDefinedTaxClassification(product string, transactionTaxClassification, productTaxClassificationBillToCountry, productTaxClassificationBillFromCountry *string, definedTaxClassification string) *DefinedTaxClassification {
	pm := &requests.DefinedTaxClassification{}

	pm.Product = product
	pm.TransactionTaxClassification = transactionTaxClassification
	pm.ProductTaxClassificationBillToCountry = productTaxClassificationBillToCountry
	pm.ProductTaxClassificationBillFromCountry = productTaxClassificationBillFromCountry
	pm.DefinedTaxClassification = definedTaxClassification

	data := pm
	res := DefinedTaxClassification{
		Product:                                 data.Product,
		TransactionTaxClassification:            data.TransactionTaxClassification,
		ProductTaxClassificationBillToCountry:   data.ProductTaxClassificationBillToCountry,
		ProductTaxClassificationBillFromCountry: data.ProductTaxClassificationBillFromCountry,
		DefinedTaxClassification:                data.DefinedTaxClassification,
	}

	return &res
}

func (psdc *SDC) ConvertToProductMasterGeneralKey() *ProductMasterGeneralKey {
	pm := &requests.ProductMasterGeneralKey{
		IsMarkedForDeletion: false,
	}

	data := pm
	res := ProductMasterGeneralKey{
		Product:             data.Product,
		ValidityStartDate:   data.ValidityStartDate,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToProductMasterGeneral(rows *sql.Rows) ([]*ProductMasterGeneral, error) {
	var res []*ProductMasterGeneral

	for i := 0; true; i++ {
		pm := &requests.ProductMasterGeneral{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_product_master_general_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Product,
			&pm.BaseUnit,
			&pm.ProductGroup,
			&pm.ProductStandardID,
			&pm.GrossWeight,
			&pm.NetWeight,
			&pm.WeightUnit,
			&pm.ItemCategory,
			&pm.ProductAccountAssignmentGroup,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &ProductMasterGeneral{
			Product:                       data.Product,
			BaseUnit:                      data.BaseUnit,
			ProductGroup:                  data.ProductGroup,
			ProductStandardID:             data.ProductStandardID,
			GrossWeight:                   data.GrossWeight,
			NetWeight:                     data.NetWeight,
			WeightUnit:                    data.WeightUnit,
			ItemCategory:                  data.ItemCategory,
			ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
			CountryOfOrigin:               data.CountryOfOrigin,
			CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrderItemTextKey(length int) []*OrderItemTextKey {
	var res []*OrderItemTextKey

	for i := 0; i < length; i++ {
		pm := &requests.OrderItemTextKey{}

		data := pm
		res = append(res, &OrderItemTextKey{
			Product:  data.Product,
			Language: data.Language,
		})
	}

	return res
}

func (psdc *SDC) ConvertToOrderItemText(rows *sql.Rows) ([]*OrderItemText, error) {
	var res []*OrderItemText

	for i := 0; true; i++ {
		pm := &requests.OrderItemText{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_product_master_product_description_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Product,
			&pm.Language,
			&pm.OrderItemText,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrderItemText{
			Product:       data.Product,
			Language:      data.Language,
			OrderItemText: data.OrderItemText,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToItemCategoryIsINVP() []*ItemCategoryIsINVP {
	var res []*ItemCategoryIsINVP

	for _, v := range psdc.ProductMasterGeneral {
		pm := &requests.ItemCategoryIsINVP{
			Product:            v.Product,
			ItemCategoryIsINVP: false,
		}
		if v.ItemCategory != nil {
			if *v.ItemCategory == "INVP" {
				pm.ItemCategoryIsINVP = true
			}
		}

		data := pm
		res = append(res, &ItemCategoryIsINVP{
			Product:            data.Product,
			ItemCategoryIsINVP: data.ItemCategoryIsINVP,
		})
	}

	return res
}

func (psdc *SDC) ConvertToStockConfPlantRelationProductKey() *StockConfPlantRelationProductKey {
	pm := &requests.StockConfPlantRelationProductKey{}

	data := pm
	res := StockConfPlantRelationProductKey{
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
		Product:                   data.Product,
	}

	return &res
}

func (psdc *SDC) ConvertToStockConfPlantRelationProduct(rows *sql.Rows) ([]*StockConfPlantRelationProduct, error) {
	var res []*StockConfPlantRelationProduct

	for i := 0; true; i++ {
		pm := &requests.StockConfPlantRelationProduct{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_supply_chain_relationship_stock_conf_plant_relation_product_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.Product,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &StockConfPlantRelationProduct{
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipStockConfPlantID: data.SupplyChainRelationshipStockConfPlantID,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			StockConfirmationBusinessPartner:        data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                  data.StockConfirmationPlant,
			Product:                                 data.Product,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToProductMasterBPPlantKey(length int) []*ProductMasterBPPlantKey {
	var res []*ProductMasterBPPlantKey

	for i := 0; i < length; i++ {
		pm := &requests.ProductMasterBPPlantKey{}

		data := pm
		res = append(res, &ProductMasterBPPlantKey{
			Product:         data.Product,
			BusinessPartner: data.BusinessPartner,
			Plant:           data.Plant,
		})
	}

	return res
}

func (psdc *SDC) ConvertToProductMasterBPPlant(rows *sql.Rows) ([]*ProductMasterBPPlant, error) {
	var res []*ProductMasterBPPlant

	for i := 0; true; i++ {
		pm := &requests.ProductMasterBPPlant{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_product_master_bp_plant_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.IsBatchManagementRequired,
			&pm.BatchManagementPolicy,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &ProductMasterBPPlant{
			Product:                   data.Product,
			BusinessPartner:           data.BusinessPartner,
			Plant:                     data.Plant,
			IsBatchManagementRequired: data.IsBatchManagementRequired,
			BatchManagementPolicy:     data.BatchManagementPolicy,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToBusinessPartnerGeneral(rows *sql.Rows) ([]*BusinessPartnerGeneral, error) {
	var res []*BusinessPartnerGeneral

	for i := 0; true; i++ {
		pm := &requests.BusinessPartnerGeneral{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_business_partner_general_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.OrganizationBPName1,
			&pm.Language,
			&pm.AddressID,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &BusinessPartnerGeneral{
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			OrganizationBPName1:     data.OrganizationBPName1,
			Language:                data.Language,
			AddressID:               data.AddressID,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToProductionPlantRelationProductKey() *ProductionPlantRelationProductKey {
	pm := &requests.ProductionPlantRelationProductKey{}

	data := pm
	res := ProductionPlantRelationProductKey{
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
		Product:                   data.Product,
	}

	return &res
}

func (psdc *SDC) ConvertToProductionPlantRelationProduct(rows *sql.Rows) ([]*ProductionPlantRelationProduct, error) {
	var res []*ProductionPlantRelationProduct

	for i := 0; true; i++ {
		pm := &requests.ProductionPlantRelationProduct{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_supply_chain_relationship_production_plant_relation_product_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Product,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &ProductionPlantRelationProduct{
			SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
			SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
			Product:                                  data.Product,
			ProductionPlantBusinessPartner:           data.ProductionPlantBusinessPartner,
			ProductionPlant:                          data.ProductionPlant,
			ProductionPlantStorageLocation:           data.ProductionPlantStorageLocation,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToSupplyChainRelationshipDeliveryPlantRelationProductKey() *SupplyChainRelationshipDeliveryPlantRelationProductKey {
	pm := &requests.SupplyChainRelationshipDeliveryPlantRelationProductKey{
		IsMarkedForDeletion: false,
	}

	data := pm
	res := SupplyChainRelationshipDeliveryPlantRelationProductKey{
		SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
		SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
		DeliverToParty:                         data.DeliverToParty,
		DeliverFromParty:                       data.DeliverFromParty,
		DeliverToPlant:                         data.DeliverToPlant,
		DeliverFromPlant:                       data.DeliverFromPlant,
		Product:                                data.Product,
		IsMarkedForDeletion:                    data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToSupplyChainRelationshipDeliveryPlantRelationProduct(rows *sql.Rows) ([]*SupplyChainRelationshipDeliveryPlantRelationProduct, error) {
	var res []*SupplyChainRelationshipDeliveryPlantRelationProduct

	for i := 0; true; i++ {
		pm := &requests.SupplyChainRelationshipDeliveryPlantRelationProduct{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_supply_chain_relationship_delivery_plant_relation_product_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.Product,
			&pm.DeliverToPlantStorageLocation,
			&pm.DeliverFromPlantStorageLocation,
			&pm.DeliveryUnit,
			&pm.StandardDeliveryDurationInDays,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &SupplyChainRelationshipDeliveryPlantRelationProduct{
			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
			DeliverToParty:                         data.DeliverToParty,
			DeliverFromParty:                       data.DeliverFromParty,
			DeliverToPlant:                         data.DeliverToPlant,
			DeliverFromPlant:                       data.DeliverFromPlant,
			Product:                                data.Product,
			DeliverToPlantStorageLocation:          data.DeliverToPlantStorageLocation,
			DeliverFromPlantStorageLocation:        data.DeliverFromPlantStorageLocation,
			DeliveryUnit:                           data.DeliveryUnit,
			StandardDeliveryDurationInDays:         data.StandardDeliveryDurationInDays,
			IsMarkedForDeletion:                    data.IsMarkedForDeletion,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToTimeZoneKey(length int) []*TimeZoneKey {
	var res []*TimeZoneKey

	for i := 0; i < length; i++ {
		pm := &requests.TimeZoneKey{}

		data := pm
		res = append(res, &TimeZoneKey{
			BusinessPartner: data.BusinessPartner,
			Plant:           data.Plant,
		})
	}

	return res
}

func (psdc *SDC) ConvertToTimeZone(rows *sql.Rows) ([]*TimeZone, error) {
	var res []*TimeZone

	for i := 0; true; i++ {
		pm := &requests.TimeZone{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_plant_general_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.TimeZone,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &TimeZone{
			BusinessPartner: data.BusinessPartner,
			Plant:           data.Plant,
			TimeZone:        data.TimeZone,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToIncoterms(length int, incoterms *string) []*Incoterms {
	var res []*Incoterms

	for i := 0; i < length; i++ {
		pm := &requests.Incoterms{}

		pm.Incoterms = incoterms

		data := pm
		res = append(res, &Incoterms{
			Incoterms: data.Incoterms,
		})
	}

	return res
}

func (psdc *SDC) ConvertToPaymentTerms(length int, paymentTerms *string) []*PaymentTerms {
	var res []*PaymentTerms

	for i := 0; i < length; i++ {
		pm := &requests.PaymentTerms{}

		pm.PaymentTerms = paymentTerms

		data := pm
		res = append(res, &PaymentTerms{
			PaymentTerms: data.PaymentTerms,
		})
	}

	return res
}

func (psdc *SDC) ConvertToPaymentMethod(length int, paymentMethod *string) []*PaymentMethod {
	var res []*PaymentMethod

	for i := 0; i < length; i++ {
		pm := &requests.PaymentMethod{}

		pm.PaymentMethod = paymentMethod

		data := pm
		res = append(res, &PaymentMethod{
			PaymentMethod: data.PaymentMethod,
		})
	}

	return res
}

func (psdc *SDC) ConvertToItemGrossWeight(product string, productGrossWeight, orderQuantityInBaseUnit, itemGrossWeght *float32) *ItemGrossWeight {
	pm := &requests.ItemGrossWeight{}

	pm.Product = product
	pm.ProductGrossWeight = productGrossWeight
	pm.OrderQuantityInBaseUnit = orderQuantityInBaseUnit
	pm.ItemGrossWeight = itemGrossWeght

	data := pm
	res := ItemGrossWeight{
		Product:                 data.Product,
		ProductGrossWeight:      data.ProductGrossWeight,
		OrderQuantityInBaseUnit: data.OrderQuantityInBaseUnit,
		ItemGrossWeight:         data.ItemGrossWeight,
	}

	return &res
}

func (psdc *SDC) ConvertToItemNetWeight(product string, productNetWeight, orderQuantityInBaseUnit, itemNetWeght *float32) *ItemNetWeight {
	pm := &requests.ItemNetWeight{}

	pm.Product = product
	pm.ProductNetWeight = productNetWeight
	pm.OrderQuantityInBaseUnit = orderQuantityInBaseUnit
	pm.ItemNetWeight = itemNetWeght

	data := pm
	res := ItemNetWeight{
		Product:                 data.Product,
		ProductNetWeight:        data.ProductNetWeight,
		OrderQuantityInBaseUnit: data.OrderQuantityInBaseUnit,
		ItemNetWeight:           data.ItemNetWeight,
	}

	return &res
}

func (psdc *SDC) ConvertToCreationDateItem(systemDate string) *CreationDateItem {
	pm := &requests.CreationDateItem{}

	pm.CreationDate = systemDate

	data := pm
	res := CreationDateItem{
		CreationDate: data.CreationDate,
	}

	return &res
}

func (psdc *SDC) ConvertToLastChangeDateItem(systemDate string) *LastChangeDateItem {
	pm := &requests.LastChangeDateItem{}

	pm.LastChangeDate = systemDate

	data := pm
	res := LastChangeDateItem{
		LastChangeDate: data.LastChangeDate,
	}

	return &res
}

func GetBoolPtr(b bool) *bool {
	return &b
}
