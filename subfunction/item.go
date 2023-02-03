package subfunction

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"encoding/json"
	"math"
	"reflect"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

func (f *SubFunction) OrderItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.OrderItem {
	data := psdc.ConvertToOrderItem(sdc)

	return data
}

func (f *SubFunction) ProductTaxClassificationBillToCountry(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductTaxClassificationBillToCountry, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductTaxClassificationKey()

	for _, v := range sdc.Header.Item {
		dataKey.Product = append(dataKey.Product, v.Product)
	}

	dataKey.Country = psdc.SupplyChainRelationshipBillingRelation[0].BillToCountry

	repeat := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	args = append(args, dataKey.Country, dataKey.ProductTaxCategory)
	rows, err := f.db.Query(
		`SELECT Product, Country, ProductTaxCategory, ProductTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data
		WHERE Product IN ( `+repeat+` )
		AND (Country, ProductTaxCategory) = (?, ?);`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductTaxClassificationBillToCountry(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ProductTaxClassificationBillFromCountry(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductTaxClassificationBillFromCountry, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductTaxClassificationKey()

	for _, v := range sdc.Header.Item {
		dataKey.Product = append(dataKey.Product, v.Product)
	}

	dataKey.Country = psdc.SupplyChainRelationshipBillingRelation[0].BillFromCountry

	repeat := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	args = append(args, dataKey.Country, dataKey.ProductTaxCategory)
	rows, err := f.db.Query(
		`SELECT Product, Country, ProductTaxCategory, ProductTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data
		WHERE Product IN ( `+repeat+` )
		AND (Country, ProductTaxCategory) = (?, ?);`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductTaxClassificationBillFromCountry(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DefinedTaxClassification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.DefinedTaxClassification, error) {
	data := make([]*api_processing_data_formatter.DefinedTaxClassification, 0)
	var err error

	transactionTaxClassification := psdc.SupplyChainRelationshipBillingRelation[0].TransactionTaxClassification

	productTaxClassificationBillFromCountry := psdc.ProductTaxClassificationBillFromCountry
	productTaxClassificationBillFromCountryMap := make(map[string]*api_processing_data_formatter.ProductTaxClassificationBillFromCountry, len(productTaxClassificationBillFromCountry))
	for _, v := range productTaxClassificationBillFromCountry {
		productTaxClassificationBillFromCountryMap[v.Product] = v
	}

	for _, v := range psdc.ProductTaxClassificationBillToCountry {
		var definedTaxClassification string

		product := v.Product
		productTaxClassificationBillToCountry := v.ProductTaxClassifiication

		if _, ok := productTaxClassificationBillFromCountryMap[product]; !ok {
			continue
		}
		productTaxClassificationBillFromCountry := productTaxClassificationBillFromCountryMap[product].ProductTaxClassifiication

		if transactionTaxClassification == nil || productTaxClassificationBillToCountry == nil || productTaxClassificationBillFromCountry == nil {
			return nil, xerrors.Errorf("TransactionTaxClassificationまたはProductTaxClassificationBillToCountryまたはProductTaxClassificationBillFromCountryがnullです。")
		}
		if *transactionTaxClassification == "1" && *productTaxClassificationBillToCountry == "1" && *productTaxClassificationBillFromCountry == "1" {
			definedTaxClassification = "1"
		} else {
			definedTaxClassification = "0"
		}

		datum := psdc.ConvertToDefinedTaxClassification(product, transactionTaxClassification, productTaxClassificationBillToCountry, productTaxClassificationBillFromCountry, definedTaxClassification)
		data = append(data, datum)
	}

	return data, err
}

func (f *SubFunction) ProductMasterGeneral(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductMasterGeneral, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductMasterGeneralKey()

	for _, v := range sdc.Header.Item {
		dataKey.Product = append(dataKey.Product, v.Product)
	}
	dataKey.ValidityStartDate = getSystemDate()

	repeat := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	args = append(args, dataKey.ValidityStartDate, dataKey.IsMarkedForDeletion)
	rows, err := f.db.Query(
		`SELECT Product, BaseUnit, ProductGroup, ProductStandardID, GrossWeight, NetWeight, WeightUnit, InternalCapacityQuantity, InternalCapacityQuantityUnit, ItemCategory, ProductAccountAssignmentGroup, CountryOfOrigin, CountryOfOriginLanguage
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_general_data
		WHERE Product IN ( `+repeat+` )
		AND ValidityStartDate <= ?
		AND IsMarkedForDeletion = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductMasterGeneral(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderItemText(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderItemText, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToOrderItemTextKey(len(sdc.Header.Item))

	for i := range psdc.ProductMasterGeneral {
		dataKey[i].Product = psdc.ProductMasterGeneral[i].Product
		dataKey[i].Language = *psdc.ProductMasterGeneral[i].CountryOfOriginLanguage
	}

	repeat := strings.Repeat("(?, ?),", len(dataKey)-1) + "(?, ?)"
	for _, v := range dataKey {
		args = append(args, v.Product, v.Language)
	}

	rows, err := f.db.Query(
		`SELECT Product, Language, ProductDescription
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_description_data
		WHERE (Product, Language) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrderItemText(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ItemCategoryIsINVP(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.ItemCategoryIsINVP {
	data := psdc.ConvertToItemCategoryIsINVP()

	return data
}

func (f *SubFunction) StockConfPlantRelationProduct(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.StockConfPlantRelationProduct, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToStockConfPlantRelationProductKey()

	for _, v := range psdc.SupplyChainRelationshipGeneral {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
	}

	itemCategoryIsINVPMap := StructArrayToMap(psdc.ItemCategoryIsINVP, "Product")
	for _, v := range psdc.ProductMasterGeneral {
		if itemCategoryIsINVPMap[v.Product].ItemCategoryIsINVP {
			dataKey.Product = append(dataKey.Product, v.Product)
		}
	}

	repeat1 := strings.Repeat("(?,?,?),", len(dataKey.SupplyChainRelationshipID)-1) + "(?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(args, dataKey.SupplyChainRelationshipID[i], dataKey.Buyer[i], dataKey.Seller[i])
	}

	repeat2 := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipStockConfPlantID, Buyer, Seller,
		StockConfirmationBusinessPartner, StockConfirmationPlant, Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_stock_conf_plant_rel_pro
		WHERE (SupplyChainRelationshipID, Buyer, Seller) IN ( `+repeat1+` )
		AND Product IN ( `+repeat2+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToStockConfPlantRelationProduct(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) StockConfPlantProductMasterBPPlant(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductMasterBPPlant, error) {
	args := make([]interface{}, 0)

	stockConfPlantRelationProduct := psdc.StockConfPlantRelationProduct

	dataKey := psdc.ConvertToProductMasterBPPlantKey(len(stockConfPlantRelationProduct))

	for i := range dataKey {
		dataKey[i].Product = stockConfPlantRelationProduct[i].Product
		dataKey[i].BusinessPartner = stockConfPlantRelationProduct[i].StockConfirmationBusinessPartner
		dataKey[i].Plant = stockConfPlantRelationProduct[i].StockConfirmationPlant
	}

	repeat := strings.Repeat("(?,?,?),", len(dataKey)-1) + "(?,?,?)"
	for _, v := range dataKey {
		args = append(args, v.Product, v.BusinessPartner, v.Plant)
	}

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Plant, IsBatchManagementRequired, BatchManagementPolicy
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductMasterBPPlant(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) StockConfPlantBPGeneral(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.BusinessPartnerGeneral, error) {
	args := make([]interface{}, 0)

	stockConfPlantRelationProduct := psdc.StockConfPlantRelationProduct

	repeat := strings.Repeat("?,", len(stockConfPlantRelationProduct)-1) + "?"
	for _, v := range stockConfPlantRelationProduct {
		args = append(args, v.StockConfirmationBusinessPartner)
	}

	rows, err := f.db.Query(
		`SELECT BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, Country, Language, Currency, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data
		WHERE BusinessPartner IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToBusinessPartnerGeneral(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ProductionPlantRelationProduct(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductionPlantRelationProduct, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToProductionPlantRelationProductKey()

	for _, v := range psdc.SupplyChainRelationshipGeneral {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
	}

	itemCategoryIsINVPMap := StructArrayToMap(psdc.ItemCategoryIsINVP, "Product")
	for _, v := range psdc.ProductMasterGeneral {
		if itemCategoryIsINVPMap[v.Product].ItemCategoryIsINVP {
			dataKey.Product = append(dataKey.Product, v.Product)
		}
	}

	repeat1 := strings.Repeat("(?,?,?),", len(dataKey.SupplyChainRelationshipID)-1) + "(?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(args, dataKey.SupplyChainRelationshipID[i], dataKey.Buyer[i], dataKey.Seller[i])
	}

	repeat2 := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipProductionPlantID, Product,
		ProductionPlantBusinessPartner, ProductionPlant, ProductionPlantStorageLocation
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_prod_plant_rel_product
		WHERE (SupplyChainRelationshipID, Buyer, Seller) IN ( `+repeat1+` )
		AND Product IN ( `+repeat2+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductionPlantRelationProduct(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ProductionPlantProductMasterBPPlant(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductMasterBPPlant, error) {
	args := make([]interface{}, 0)

	productionPlantRelationProduct := psdc.ProductionPlantRelationProduct

	dataKey := psdc.ConvertToProductMasterBPPlantKey(len(productionPlantRelationProduct))

	for i := range dataKey {
		dataKey[i].Product = productionPlantRelationProduct[i].Product
		dataKey[i].BusinessPartner = productionPlantRelationProduct[i].ProductionPlantBusinessPartner
		dataKey[i].Plant = productionPlantRelationProduct[i].ProductionPlant
	}

	repeat := strings.Repeat("(?,?,?),", len(dataKey)-1) + "(?,?,?)"
	for _, v := range dataKey {
		args = append(args, v.Product, v.BusinessPartner, v.Plant)
	}

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Plant, IsBatchManagementRequired, BatchManagementPolicy
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductMasterBPPlant(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ProductionPlantBPGeneral(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.BusinessPartnerGeneral, error) {
	args := make([]interface{}, 0)

	productionPlantRelationProduct := psdc.ProductionPlantRelationProduct

	repeat := strings.Repeat("?,", len(productionPlantRelationProduct)-1) + "?"
	for _, v := range productionPlantRelationProduct {
		args = append(args, v.ProductionPlantBusinessPartner)
	}

	rows, err := f.db.Query(
		`SELECT BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, Country, Language, Currency, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data
		WHERE BusinessPartner IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToBusinessPartnerGeneral(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipDeliveryPlantRelationProduct(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipDeliveryPlantRelationProduct, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToSupplyChainRelationshipDeliveryPlantRelationProductKey()

	for _, v := range psdc.SupplyChainRelationshipDeliveryPlantRelation {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.SupplyChainRelationshipDeliveryID = append(dataKey.SupplyChainRelationshipDeliveryID, v.SupplyChainRelationshipDeliveryID)
		dataKey.SupplyChainRelationshipDeliveryPlantID = append(dataKey.SupplyChainRelationshipDeliveryPlantID, v.SupplyChainRelationshipDeliveryPlantID)
		dataKey.DeliverToParty = append(dataKey.DeliverToParty, v.DeliverToParty)
		dataKey.DeliverFromParty = append(dataKey.DeliverFromParty, v.DeliverFromParty)
		dataKey.DeliverToPlant = append(dataKey.DeliverToPlant, v.DeliverToPlant)
		dataKey.DeliverFromPlant = append(dataKey.DeliverFromPlant, v.DeliverFromPlant)
	}

	repeat1 := strings.Repeat("(?,?,?,?,?,?,?),", len(dataKey.DeliverToParty)-1) + "(?,?,?,?,?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(
			args,
			dataKey.SupplyChainRelationshipID[i],
			dataKey.SupplyChainRelationshipDeliveryID[i],
			dataKey.SupplyChainRelationshipDeliveryPlantID[i],
			dataKey.DeliverToParty[i],
			dataKey.DeliverFromParty[i],
			dataKey.DeliverToPlant[i],
			dataKey.DeliverFromPlant[i],
		)
	}

	itemCategoryIsINVPMap := StructArrayToMap(psdc.ItemCategoryIsINVP, "Product")
	for _, v := range psdc.ProductionPlantRelationProduct {
		if itemCategoryIsINVPMap[v.Product].ItemCategoryIsINVP {
			dataKey.Product = append(dataKey.Product, v.Product)
		}
	}

	repeat2 := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	args = append(args, dataKey.IsMarkedForDeletion)

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, SupplyChainRelationshipDeliveryPlantID,
		DeliverToParty, DeliverFromParty, DeliverToPlant, DeliverFromPlant, Product, DeliverToPlantStorageLocation,
		DeliverFromPlantStorageLocation, DeliveryUnit, StandardDeliveryDurationInDays, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_delivery_plant_rel_prod
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, SupplyChainRelationshipDeliveryPlantID, DeliverToParty, DeliverFromParty, DeliverToPlant, DeliverFromPlant) IN ( `+repeat1+` )
		AND Product IN ( `+repeat2+` )
		AND IsMarkedForDeletion = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToSupplyChainRelationshipDeliveryPlantRelationProduct(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipProductMasterBPPlantDeliverTo(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductMasterBPPlant, error) {
	args := make([]interface{}, 0)

	supplyChainRelationshipDeliveryPlantRelationProduct := psdc.SupplyChainRelationshipDeliveryPlantRelationProduct

	dataKey := psdc.ConvertToProductMasterBPPlantKey(len(supplyChainRelationshipDeliveryPlantRelationProduct))

	for i := range dataKey {
		dataKey[i].Product = supplyChainRelationshipDeliveryPlantRelationProduct[i].Product
		dataKey[i].BusinessPartner = supplyChainRelationshipDeliveryPlantRelationProduct[i].DeliverToParty
		dataKey[i].Plant = supplyChainRelationshipDeliveryPlantRelationProduct[i].DeliverToPlant
	}

	repeat := strings.Repeat("(?,?,?),", len(dataKey)-1) + "(?,?,?)"
	for _, v := range dataKey {
		args = append(args, v.Product, v.BusinessPartner, v.Plant)
	}

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Plant, IsBatchManagementRequired, BatchManagementPolicy
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductMasterBPPlant(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipProductMasterBPPlantDeliverFrom(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductMasterBPPlant, error) {
	args := make([]interface{}, 0)

	supplyChainRelationshipDeliveryPlantRelationProduct := psdc.SupplyChainRelationshipDeliveryPlantRelationProduct

	dataKey := psdc.ConvertToProductMasterBPPlantKey(len(supplyChainRelationshipDeliveryPlantRelationProduct))

	for i := range dataKey {
		dataKey[i].Product = supplyChainRelationshipDeliveryPlantRelationProduct[i].Product
		dataKey[i].BusinessPartner = supplyChainRelationshipDeliveryPlantRelationProduct[i].DeliverFromParty
		dataKey[i].Plant = supplyChainRelationshipDeliveryPlantRelationProduct[i].DeliverFromPlant
	}

	repeat := strings.Repeat("(?,?,?),", len(dataKey)-1) + "(?,?,?)"
	for _, v := range dataKey {
		args = append(args, v.Product, v.BusinessPartner, v.Plant)
	}

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Plant, IsBatchManagementRequired, BatchManagementPolicy
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_bp_plant_data
		WHERE (Product, BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToProductMasterBPPlant(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ProductionPlantTimeZone(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.TimeZone, error) {
	productionPlantRelationProduct := psdc.ProductionPlantRelationProduct

	dataKey := psdc.ConvertToTimeZoneKey(len(productionPlantRelationProduct))

	for i := range dataKey {
		dataKey[i].BusinessPartner = productionPlantRelationProduct[i].ProductionPlantBusinessPartner
		dataKey[i].Plant = productionPlantRelationProduct[i].ProductionPlant
	}

	data, err := f.timeZone(sdc, psdc, dataKey)

	return data, err
}

func (f *SubFunction) DeliverToPlantTimeZone(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.TimeZone, error) {
	supplyChainRelationshipDeliveryPlantRelation := psdc.SupplyChainRelationshipDeliveryPlantRelation

	dataKey := psdc.ConvertToTimeZoneKey(len(supplyChainRelationshipDeliveryPlantRelation))

	for i := range dataKey {
		dataKey[i].BusinessPartner = supplyChainRelationshipDeliveryPlantRelation[i].DeliverToParty
		dataKey[i].Plant = supplyChainRelationshipDeliveryPlantRelation[i].DeliverToPlant
	}

	data, err := f.timeZone(sdc, psdc, dataKey)

	return data, err
}

func (f *SubFunction) DeliverFromPlantTimeZone(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.TimeZone, error) {
	supplyChainRelationshipDeliveryPlantRelation := psdc.SupplyChainRelationshipDeliveryPlantRelation

	dataKey := psdc.ConvertToTimeZoneKey(len(supplyChainRelationshipDeliveryPlantRelation))

	for i := range dataKey {
		dataKey[i].BusinessPartner = supplyChainRelationshipDeliveryPlantRelation[i].DeliverFromParty
		dataKey[i].Plant = supplyChainRelationshipDeliveryPlantRelation[i].DeliverFromPlant
	}

	data, err := f.timeZone(sdc, psdc, dataKey)

	return data, err
}

func (f *SubFunction) StockConfirmationPlantTimeZone(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.TimeZone, error) {
	stockConfPlantRelationProduct := psdc.StockConfPlantRelationProduct

	dataKey := psdc.ConvertToTimeZoneKey(len(stockConfPlantRelationProduct))

	for i := range dataKey {
		dataKey[i].BusinessPartner = stockConfPlantRelationProduct[i].StockConfirmationBusinessPartner
		dataKey[i].Plant = stockConfPlantRelationProduct[i].StockConfirmationPlant
	}

	data, err := f.timeZone(sdc, psdc, dataKey)

	return data, err
}

func (f *SubFunction) timeZone(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	dataKey []*api_processing_data_formatter.TimeZoneKey,
) ([]*api_processing_data_formatter.TimeZone, error) {
	args := make([]interface{}, 0)

	repeat := strings.Repeat("(?,?),", len(dataKey)-1) + "(?,?)"
	for _, v := range dataKey {
		args = append(args, v.BusinessPartner, v.Plant)
	}

	rows, err := f.db.Query(
		`SELECT BusinessPartner, Plant, TimeZone
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_plant_general_data
		WHERE (BusinessPartner, Plant) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToTimeZone(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) Incoterms(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.Incoterms {
	data := make([]*api_processing_data_formatter.Incoterms, 0)

	incoterms := psdc.SupplyChainRelationshipTransaction[0].Incoterms

	for range sdc.Header.Item {
		datum := psdc.ConvertToIncoterms(incoterms)
		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) ItemPaymentTerms(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.ItemPaymentTerms {
	data := make([]*api_processing_data_formatter.ItemPaymentTerms, 0)

	paymentTerms := psdc.SupplyChainRelationshipTransaction[0].PaymentTerms

	for range sdc.Header.Item {
		datum := psdc.ConvertToItemPaymentTerms(paymentTerms)
		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) PaymentMethod(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.PaymentMethod {
	data := make([]*api_processing_data_formatter.PaymentMethod, 0)

	paymentMethod := psdc.SupplyChainRelationshipTransaction[0].PaymentMethod

	for range sdc.Header.Item {
		datum := psdc.ConvertToPaymentMethod(paymentMethod)
		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) ItemGrossWeight(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.ItemGrossWeight {
	data := make([]*api_processing_data_formatter.ItemGrossWeight, 0)

	productMasterGeneral := psdc.ProductMasterGeneral
	productMasterGeneralMap := make(map[string]*api_processing_data_formatter.ProductMasterGeneral, len(productMasterGeneral))
	for _, v := range productMasterGeneral {
		productMasterGeneralMap[v.Product] = v
	}

	for i, v := range sdc.Header.Item {
		orderItem := psdc.OrderItem[i].OrderItemNumber
		if v.Product == nil {
			continue
		}
		product := *v.Product
		if _, ok := productMasterGeneralMap[product]; !ok {
			continue
		}
		productGrossWeight := productMasterGeneralMap[product].GrossWeight
		orderQuantityInBaseUnit := v.OrderQuantityInBaseUnit
		itemGrossWeight := parseFloat32Ptr(*productGrossWeight * *orderQuantityInBaseUnit)

		datum := psdc.ConvertToItemGrossWeight(orderItem, product, productGrossWeight, orderQuantityInBaseUnit, itemGrossWeight)
		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) ItemNetWeight(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.ItemNetWeight {
	data := make([]*api_processing_data_formatter.ItemNetWeight, 0)

	item := sdc.Header.Item
	itemMap := make(map[string]api_input_reader.Item, len(item))
	for _, v := range item {
		itemMap[*v.Product] = v
	}

	for _, v := range psdc.ProductMasterGeneral {
		product := v.Product
		productNetWeight := v.GrossWeight
		orderQuantityInBaseUnit := itemMap[product].OrderQuantityInBaseUnit
		itemNetWeight := parseFloat32Ptr(*productNetWeight * *orderQuantityInBaseUnit)

		datum := psdc.ConvertToItemNetWeight(product, productNetWeight, orderQuantityInBaseUnit, itemNetWeight)
		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) TaxCode(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.TaxCode, error) {
	data := make([]*api_processing_data_formatter.TaxCode, 0)
	var err error

	for _, v := range psdc.DefinedTaxClassification {
		product := v.Product
		definedTaxClassification := v.DefinedTaxClassification
		for _, w := range psdc.SupplyChainRelationshipBillingRelation {
			isExportImport := w.IsExportImport
			taxCode := new(string)

			if isExportImport == nil {
				return nil, xerrors.Errorf("IsExportImportがnullです。")
			}
			if definedTaxClassification == "1" && !*isExportImport {
				taxCode = getStringPtr("1")
			} else if definedTaxClassification == "0" && !*isExportImport {
				taxCode = getStringPtr("8")
			} else if definedTaxClassification == "0" && *isExportImport {
				taxCode = getStringPtr("9")
			} else {
				taxCode = getStringPtr("0")
			}

			datum := psdc.ConvertToTaxCode(product, definedTaxClassification, isExportImport, taxCode)
			data = append(data, datum)
		}
	}

	return data, err

}

func (f *SubFunction) TaxRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.TaxRate, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToTaxRateKey()

	for _, v := range psdc.TaxCode {
		dataKey.TaxCode = append(dataKey.TaxCode, v.TaxCode)
	}

	dataKey.ValidityEndDate = getSystemDate()
	dataKey.ValidityStartDate = getSystemDate()

	repeat := strings.Repeat("?,", len(dataKey.TaxCode)-1) + "?"
	args = append(args, dataKey.Country)
	for _, v := range dataKey.TaxCode {
		args = append(args, v)
	}

	args = append(args, dataKey.ValidityEndDate, dataKey.ValidityStartDate)
	rows, err := f.db.Query(
		`SELECT Country, TaxCode, ValidityEndDate, ValidityStartDate, TaxRate
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_tax_code_tax_rate_data
		WHERE Country = ?
		AND TaxCode IN ( `+repeat+` )
		AND ValidityEndDate >= ?
		AND ValidityStartDate <= ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToTaxRate(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrdinaryStockConfirmation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdinaryStockConfirmation, error) {
	var err error
	data := make([]*api_processing_data_formatter.OrdinaryStockConfirmation, 0)

	length := 0
	for _, v := range psdc.StockConfPlantProductMasterBPPlant {
		if v.IsBatchManagementRequired == nil {
			continue
		}
		if !*v.IsBatchManagementRequired {
			length++
		}
	}

	dataKey := psdc.ConvertToOrdinaryStockConfirmationKey(length)

	stockConfPlantProductMasterBPPlantMap := StructArrayToMap(psdc.StockConfPlantProductMasterBPPlant, "Product")

	idx := 0
	for _, v := range sdc.Header.Item {
		if v.Product == nil || v.RequestedDeliveryDate == nil {
			continue
		}
		if _, ok := stockConfPlantProductMasterBPPlantMap[*v.Product]; !ok {
			continue
		}
		if stockConfPlantProductMasterBPPlantMap[*v.Product].IsBatchManagementRequired == nil {
			continue
		}
		if !*stockConfPlantProductMasterBPPlantMap[*v.Product].IsBatchManagementRequired {
			dataKey[idx].Product = *v.Product
			dataKey[idx].StockConfirmationBusinessPartner = stockConfPlantProductMasterBPPlantMap[*v.Product].BusinessPartner
			dataKey[idx].StockConfirmationPlant = stockConfPlantProductMasterBPPlantMap[*v.Product].Plant
			dataKey[idx].RequestedDeliveryDate = *v.RequestedDeliveryDate
			idx++
		}
	}

	for _, v := range dataKey {
		req, err := jsonTypeConversion[api_processing_data_formatter.ProductAvailabilityCheck](sdc)
		if err != nil {
			err = xerrors.Errorf("request create error: %w", err)
			return nil, err
		}
		req.ProductStock.BusinessPartner = v.StockConfirmationBusinessPartner
		req.ProductStock.Product = v.Product
		req.ProductStock.Plant = v.StockConfirmationPlant
		req.ProductStock.Availability.ProductStockAvailabilityDate = v.RequestedDeliveryDate

		res, err := f.rmq.SessionKeepRequest(f.ctx, "data-platform-function-product-stock-availability-check-queue", req)
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil, err
		}
		res.Success()

		datum, err := psdc.ConvertToOrdinaryStockConfirmation(res.Data())
		if err != nil {
			return nil, err
		}

		data = append(data, datum)
	}

	return data, err
}

func (f *SubFunction) OrdinaryStockConfirmationOrdersItemScheduleLine(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdersItemScheduleLine, error) {
	var err error
	data := make([]*api_processing_data_formatter.OrdersItemScheduleLine, 0)

	stockConfPlantRelationProductMap := StructArrayToMap(psdc.StockConfPlantRelationProduct, "Product")
	ordinaryStockConfirmationMap := StructArrayToMap(psdc.OrdinaryStockConfirmation, "Product")

	idx := 1
	for i, item := range sdc.Header.Item {
		if item.Product == nil || item.RequestedDeliveryDate == nil || item.OrderQuantityInBaseUnit == nil {
			continue
		}
		if _, ok := stockConfPlantRelationProductMap[*item.Product]; !ok {
			continue
		}
		if _, ok := ordinaryStockConfirmationMap[*item.Product]; !ok {
			continue
		}

		orderID := psdc.CalculateOrderID.OrderID
		orderItem := psdc.OrderItem[i].OrderItemNumber
		scheduleLine := idx
		stockConfirmationPlantTimeZone := new(string)
		for _, v := range psdc.StockConfirmationPlantTimeZone {
			if v.BusinessPartner == ordinaryStockConfirmationMap[*item.Product].BusinessPartner && v.Plant == ordinaryStockConfirmationMap[*item.Product].Plant {
				stockConfirmationPlantTimeZone = v.TimeZone
			}
		}
		stockConfPlantRelationProduct := stockConfPlantRelationProductMap[*item.Product]
		ordinaryStockConfirmation := ordinaryStockConfirmationMap[*item.Product]
		datum, err := psdc.ConvertToOrdinaryStockConfirmationOrdersItemScheduleLine(orderID, orderItem, scheduleLine, stockConfirmationPlantTimeZone, item, stockConfPlantRelationProduct, ordinaryStockConfirmation)
		if err != nil {
			return nil, err
		}

		data = append(data, datum)
		idx++
	}

	return data, err
}

func (f *SubFunction) ItemPricingDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.PricingDate {
	data := make([]*api_processing_data_formatter.PricingDate, 0)

	pricingDate := psdc.HeaderPricingDate.PricingDate

	for range sdc.Header.Item {
		datum := psdc.ConvertToPricingDate(pricingDate)
		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) ConfirmedOrderQuantityInBaseUnit(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.ConfirmedOrderQuantityInBaseUnit {
	data := make([]*api_processing_data_formatter.ConfirmedOrderQuantityInBaseUnit, 0)

	itemCategoryIsINVPMap := StructArrayToMap(psdc.ItemCategoryIsINVP, "Product")
	for _, v := range psdc.OrdinaryStockConfirmationOrdersItemScheduleLine {
		if itemCategoryIsINVPMap[v.Product].ItemCategoryIsINVP {
			datum := psdc.ConvertToConfirmedOrderQuantityInBaseUnit(v.OrderItem, v.ConfirmedOrderQuantityByPDTAvailCheck)

			data = append(data, datum)
		}
	}

	return data
}

func (f *SubFunction) ItemInvoiceDocumentDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.ItemInvoiceDocumentDate {
	data := make([]*api_processing_data_formatter.ItemInvoiceDocumentDate, 0)

	invoiceDocumentDate := psdc.HeaderInvoiceDocumentDate.InvoiceDocumentDate

	for _, v := range sdc.Header.Item {
		if v.InvoiceDocumentDate != nil {
			if *v.InvoiceDocumentDate != "" {
				invoiceDocumentDate = *v.InvoiceDocumentDate
			}
		}
		datum := psdc.ConvertToItemInvoiceDocumentDate(invoiceDocumentDate)
		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) ItemPriceDetnExchangeRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.PriceDetnExchangeRate {
	data := make([]*api_processing_data_formatter.PriceDetnExchangeRate, 0)

	priceDetnExchangeRate := psdc.HeaderPriceDetnExchangeRate.PriceDetnExchangeRate

	for range sdc.Header.Item {
		datum := psdc.ConvertToItemPriceDetnExchangeRate(priceDetnExchangeRate)
		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) ItemAccountingExchangeRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.AccountingExchangeRate {
	data := make([]*api_processing_data_formatter.AccountingExchangeRate, 0)

	accountingExchangeRate := psdc.HeaderAccountingExchangeRate.AccountingExchangeRate

	for range sdc.Header.Item {
		datum := psdc.ConvertToItemAccountingExchangeRate(accountingExchangeRate)
		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) ItemReferenceDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.ItemReferenceDocument {
	data := make([]*api_processing_data_formatter.ItemReferenceDocument, 0)

	for range sdc.Header.Item {
		datum := psdc.ConvertToItemReferenceDocument(sdc.Header.ReferenceDocument, sdc.Header.ReferenceDocumentItem)
		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) OrderItemTextByBuyer(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderItemTextByBuyerSeller, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToOrderItemTextByBuyerSellerKey()

	for _, v := range sdc.Header.Item {
		dataKey.Product = append(dataKey.Product, v.Product)
	}

	for _, v := range psdc.BusinessPartnerGeneralBuyer {
		dataKey.BusinessPartner = append(dataKey.BusinessPartner, v.BusinessPartner)
		dataKey.Language = append(dataKey.Language, v.Language)
	}

	repeat1 := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	repeat2 := strings.Repeat("(?,?),", len(dataKey.BusinessPartner)-1) + "(?,?)"
	for i := range dataKey.BusinessPartner {
		args = append(args, dataKey.BusinessPartner[i], dataKey.Language[i])
	}

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Language, ProductDescription
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_desc_by_bp_data
		WHERE Product IN ( `+repeat1+` )
		AND (BusinessPartner, Language) IN ( `+repeat2+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrderItemTextByBuyerSeller(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderItemTextBySeller(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderItemTextByBuyerSeller, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToOrderItemTextByBuyerSellerKey()

	for _, v := range sdc.Header.Item {
		dataKey.Product = append(dataKey.Product, v.Product)
	}

	for _, v := range psdc.BusinessPartnerGeneralSeller {
		dataKey.BusinessPartner = append(dataKey.BusinessPartner, v.BusinessPartner)
	}

	for _, v := range psdc.BusinessPartnerGeneralBuyer {
		dataKey.Language = append(dataKey.Language, v.Language)
	}

	repeat1 := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	repeat2 := strings.Repeat("(?,?),", len(dataKey.BusinessPartner)-1) + "(?,?)"
	for i := range dataKey.BusinessPartner {
		args = append(args, dataKey.BusinessPartner[i], dataKey.Language[i])
	}

	rows, err := f.db.Query(
		`SELECT Product, BusinessPartner, Language, ProductDescription
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_product_desc_by_bp_data
		WHERE Product IN ( `+repeat1+` )
		AND (BusinessPartner, Language) IN ( `+repeat2+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrderItemTextByBuyerSeller(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

// Amount関連の計算
func (f *SubFunction) NetAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.NetAmount {
	data := psdc.ConvertToNetAmount(psdc.ConditionAmount)

	return data
}

func (f *SubFunction) TaxAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.TaxAmount, error) {
	data := make([]*api_processing_data_formatter.TaxAmount, 0)

	item := sdc.Header.Item
	itemMap := make(map[string]api_input_reader.Item, len(item))
	for _, v := range item {
		itemMap[*v.Product] = v
	}

	taxRate := psdc.TaxRate
	taxRateMap := make(map[string]*api_processing_data_formatter.TaxRate, len(taxRate))
	for _, v := range taxRate {
		taxRateMap[v.TaxCode] = v
	}

	netAmount := psdc.NetAmount
	netAmountMap := make(map[string]*api_processing_data_formatter.NetAmount, len(netAmount))
	for _, v := range netAmount {
		netAmountMap[v.Product] = v
	}

	for _, v := range psdc.TaxCode {
		taxAmount := new(float32)
		if *v.TaxCode == "1" {
			taxAmount, _ = calculateTaxAmount(taxRateMap[*v.TaxCode].TaxRate, netAmountMap[v.Product].NetAmount)
		} else {
			taxAmount = parseFloat32Ptr(0)
		}

		if itemMap[v.Product].TaxAmount == nil {
			datum := psdc.ConvertToTaxAmount(v.Product, v.TaxCode, taxRateMap[*v.TaxCode].TaxRate, netAmountMap[v.Product].NetAmount, taxAmount)
			data = append(data, datum)
		} else {
			datum := psdc.ConvertToTaxAmount(v.Product, v.TaxCode, taxRateMap[*v.TaxCode].TaxRate, netAmountMap[v.Product].NetAmount, itemMap[v.Product].TaxAmount)
			data = append(data, datum)
			if math.Abs(float64(*taxAmount-*itemMap[v.Product].TaxAmount)) >= 2 {
				return nil, xerrors.Errorf("TaxAmountについて入力ファイルの値と計算結果の差の絶対値が2以上の明細が一つ以上存在します。")
			}
		}
	}

	return data, nil
}

func (f *SubFunction) GrossAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.GrossAmount, error) {
	data := make([]*api_processing_data_formatter.GrossAmount, 0)

	item := sdc.Header.Item
	itemMap := make(map[string]api_input_reader.Item, len(item))
	for _, v := range item {
		itemMap[*v.Product] = v
	}

	for _, v := range psdc.TaxAmount {
		grossAmount := parseFloat32Ptr(*v.NetAmount + *v.TaxAmount)

		if itemMap[v.Product].GrossAmount == nil {
			datum := psdc.ConvertToGrossAmount(v.Product, v.NetAmount, v.TaxAmount, grossAmount)
			data = append(data, datum)
		} else {
			datum := psdc.ConvertToGrossAmount(v.Product, v.NetAmount, v.TaxAmount, itemMap[v.Product].GrossAmount)
			data = append(data, datum)
			if math.Abs(float64(*grossAmount-*itemMap[v.Product].GrossAmount)) >= 2 {
				return nil, xerrors.Errorf("GrossAmountについて入力ファイルの値と計算結果の差の絶対値が2以上の明細が一つ以上存在します。")
			}
		}
	}

	return data, nil
}

func calculateTaxAmount(taxRate *float32, netAmount *float32) (*float32, error) {
	if taxRate == nil || netAmount == nil {
		return nil, xerrors.Errorf("TaxRateまたはNetAmountがnullです。")
	}

	digit := float32DecimalDigit(*netAmount)
	mul := *netAmount * *taxRate / 100

	s := math.Round(float64(mul)*math.Pow10(digit)) / math.Pow10(digit)
	res := parseFloat32Ptr(float32(s))

	return res, nil
}

// 数量単位変換実行の是非の判定
func (f *SubFunction) QuantityUnitConversion(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.QuantityUnitConversion, error) {
	args := make([]interface{}, 0)
	data := make([]*api_processing_data_formatter.QuantityUnitConversion, 0)
	dataKey := make([]*api_processing_data_formatter.QuantityUnitConversionKey, 0)

	supplyChainRelationshipDeliveryPlantRelationProductMap := StructArrayToMap(psdc.SupplyChainRelationshipDeliveryPlantRelationProduct, "Product")

	for _, v := range psdc.ProductMasterGeneral {
		datumKey := new(api_processing_data_formatter.QuantityUnitConversionKey)

		if v.BaseUnit == nil {
			continue
		}
		product := v.Product
		baseUnit := *v.BaseUnit
		_, ok := supplyChainRelationshipDeliveryPlantRelationProductMap[product]
		if !ok {
			continue
		}
		deliveryUnit := supplyChainRelationshipDeliveryPlantRelationProductMap[product].DeliveryUnit

		if baseUnit == deliveryUnit {
			continue
		}

		datumKey = psdc.ConvertToQuantityUnitConversionKey(product, baseUnit, deliveryUnit)
		dataKey = append(dataKey, datumKey)
	}

	if len(dataKey) == 0 {
		return nil, nil
	}

	repeat := strings.Repeat("(?,?),", len(dataKey)-1) + "(?,?)"
	for _, v := range dataKey {
		args = append(args, v.BaseUnit, v.DeliveryUnit)
	}

	rows, err := f.db.Query(
		`SELECT QuantityUnitFrom, QuantityUnitTo, ConversionCoefficient
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_quantity_unit_conversion_quantity_unit_conv_data
		WHERE (QuantityUnitFrom, QuantityUnitTo) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dataQueryGets, err := psdc.ConvertToQuantityUnitConversionQueryGets(rows, dataKey)
	if err != nil {
		return nil, err
	}

	dataQueryGetsMap := StructArrayToMap(dataQueryGets, "Product")

	for i, v := range sdc.Header.Item {
		_, ok := dataQueryGetsMap[*v.Product]
		if !ok {
			continue
		}

		orderItem := psdc.OrderItem[i].OrderItemNumber
		product := *v.Product
		orderQuantityInBaseUnit := v.OrderQuantityInBaseUnit
		conversionCoefficient := dataQueryGetsMap[*v.Product].ConversionCoefficient

		orderQuantityInDeliveryUnit, err := calculateOrderQuantityInDeliveryUnit(orderQuantityInBaseUnit, conversionCoefficient)
		if err != nil {
			return nil, err
		}

		datum := psdc.ConvertToQuantityUnitConversion(orderItem, product, conversionCoefficient, orderQuantityInDeliveryUnit)
		data = append(data, datum)
	}

	return data, err
}

func calculateOrderQuantityInDeliveryUnit(orderQuantityInBaseUnit *float32, conversionCoefficient float32) (float32, error) {
	if orderQuantityInBaseUnit == nil {
		return 0, xerrors.Errorf("OrderQuantityInBaseUnitがnullです。")
	}

	res := *orderQuantityInBaseUnit * conversionCoefficient

	return res, nil
}

func (f *SubFunction) OrderQuantityInDeliveryUnit(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderQuantityInDeliveryUnit, error) {
	data := make([]*api_processing_data_formatter.OrderQuantityInDeliveryUnit, 0)

	quantityUnitConversionMap := StructArrayToMap(psdc.QuantityUnitConversion, "OrderItem")

	for i, v := range sdc.Header.Item {
		var orderQuantityInDeliveryUnit float32
		orderItem := i + 1

		_, ok := quantityUnitConversionMap[orderItem]
		if ok {
			orderQuantityInDeliveryUnit = quantityUnitConversionMap[orderItem].OrderQuantityInDeliveryUnit
		} else {
			if v.OrderQuantityInDeliveryUnit != nil {
				orderQuantityInDeliveryUnit = *v.OrderQuantityInDeliveryUnit
			}
		}

		datum := psdc.ConvertToOrderQuantityInDeliveryUnit(orderItem, orderQuantityInDeliveryUnit)
		data = append(data, datum)
	}

	return data, nil
}

// 日付等の処理
func (f *SubFunction) CreationDateItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.CreationDateItem {
	data := psdc.ConvertToCreationDateItem(getSystemDate())

	return data
}

func (f *SubFunction) LastChangeDateItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.LastChangeDateItem {
	data := psdc.ConvertToLastChangeDateItem(getSystemDate())

	return data
}

func getSystemDate() string {
	day := time.Now()
	return day.Format("2006-01-02")
}

func getStringPtr(s string) *string {
	return &s
}

func parseFloat32Ptr(f float32) *float32 {
	return &f
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
			res[value], _ = jsonTypeConversion[T](elem.Interface())
			break
		}
	}

	return res
}

func jsonTypeConversion[T any](data interface{}) (T, error) {
	var dist T
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
