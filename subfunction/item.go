package subfunction

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"database/sql"
	"encoding/json"
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
	var args []interface{}
	var rows *sql.Rows
	var err error

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
	rows, err = f.db.Query(
		`SELECT Product, Country, ProductTaxCategory, ProductTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data
		WHERE Product IN ( `+repeat+` )
		AND (Country, ProductTaxCategory) = (?, ?);`, args...,
	)
	if err != nil {
		return nil, err
	}

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
	var args []interface{}
	var rows *sql.Rows
	var err error

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
	rows, err = f.db.Query(
		`SELECT Product, Country, ProductTaxCategory, ProductTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_tax_data
		WHERE Product IN ( `+repeat+` )
		AND (Country, ProductTaxCategory) = (?, ?);`, args...,
	)
	if err != nil {
		return nil, err
	}

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
	var data []*api_processing_data_formatter.DefinedTaxClassification
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
		productTaxClassificationBillFromCountry := productTaxClassificationBillFromCountryMap[v.Product].ProductTaxClassifiication

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
	var args []interface{}

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
		`SELECT Product, BaseUnit, ProductGroup, ProductStandardID, GrossWeight, NetWeight, WeightUnit, ItemCategory, ProductAccountAssignmentGroup, CountryOfOrigin, CountryOfOriginLanguage
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_general_data
		WHERE Product IN ( `+repeat+` )
		AND ValidityStartDate <= ?
		AND IsMarkedForDeletion = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
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
	var args []interface{}

	item := sdc.Header.Item
	dataKey := psdc.ConvertToOrderItemTextKey(len(item))

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
	var args []interface{}

	dataKey := psdc.ConvertToStockConfPlantRelationProductKey()

	supplyChainRelationshipGeneral := psdc.SupplyChainRelationshipGeneral
	dataKey.SupplyChainRelationshipID = supplyChainRelationshipGeneral[0].SupplyChainRelationshipID
	dataKey.Buyer = supplyChainRelationshipGeneral[0].Buyer
	dataKey.Seller = supplyChainRelationshipGeneral[0].Seller

	itemCategoryIsINVPMap := StructArrayToMap(psdc.ItemCategoryIsINVP, "Product")
	for _, v := range psdc.ProductMasterGeneral {
		if itemCategoryIsINVPMap[v.Product].ItemCategoryIsINVP {
			dataKey.Product = append(dataKey.Product, v.Product)
		}
	}

	args = append(args, dataKey.SupplyChainRelationshipID, dataKey.Buyer, dataKey.Seller)
	repeat := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipStockConfPlantID, Buyer, Seller,
		StockConfirmationBusinessPartner, StockConfirmationPlant, Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_stock_conf_plant_rel_pro
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?, ?, ?)
		AND Product IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

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
	var args []interface{}

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
	var args []interface{}

	stockConfPlantRelationProduct := psdc.StockConfPlantRelationProduct

	repeat := strings.Repeat("?,", len(stockConfPlantRelationProduct)-1) + "?"
	for _, v := range stockConfPlantRelationProduct {
		args = append(args, v.StockConfirmationBusinessPartner)
	}

	rows, err := f.db.Query(
		`SELECT BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, OrganizationBPName1, Language, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data
		WHERE BusinessPartner IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

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
	var args []interface{}

	dataKey := psdc.ConvertToProductionPlantRelationProductKey()

	supplyChainRelationshipGeneral := psdc.SupplyChainRelationshipGeneral
	dataKey.SupplyChainRelationshipID = supplyChainRelationshipGeneral[0].SupplyChainRelationshipID
	dataKey.Buyer = supplyChainRelationshipGeneral[0].Buyer
	dataKey.Seller = supplyChainRelationshipGeneral[0].Seller

	itemCategoryIsINVPMap := StructArrayToMap(psdc.ItemCategoryIsINVP, "Product")
	for _, v := range psdc.ProductMasterGeneral {
		if itemCategoryIsINVPMap[v.Product].ItemCategoryIsINVP {
			dataKey.Product = append(dataKey.Product, v.Product)
		}
	}

	args = append(args, dataKey.SupplyChainRelationshipID, dataKey.Buyer, dataKey.Seller)
	repeat := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipProductionPlantID, Product,
		ProductionPlantBusinessPartner, ProductionPlant, ProductionPlantStorageLocation
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_prod_plant_rel_product
		WHERE (SupplyChainRelationshipID, Buyer, Seller) = (?, ?, ?)
		AND Product IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

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
	var args []interface{}

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
	var args []interface{}

	productionPlantRelationProduct := psdc.ProductionPlantRelationProduct

	repeat := strings.Repeat("?,", len(productionPlantRelationProduct)-1) + "?"
	for _, v := range productionPlantRelationProduct {
		args = append(args, v.ProductionPlantBusinessPartner)
	}

	rows, err := f.db.Query(
		`SELECT BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, OrganizationBPName1, Language, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_general_data
		WHERE BusinessPartner IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

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
	var args []interface{}

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

	data, err := psdc.ConvertToSupplyChainRelationshipDeliveryPlantRelationProduct(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipProductMasterBPPlant(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ProductMasterBPPlant, error) {
	var args []interface{}

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
	var args []interface{}

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
	incoterms := psdc.SupplyChainRelationshipTransaction[0].Incoterms

	data := psdc.ConvertToIncoterms(len(sdc.Header.Item), incoterms)

	return data
}

func (f *SubFunction) PaymentTerms(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.PaymentTerms {
	paymentTerms := psdc.SupplyChainRelationshipTransaction[0].PaymentTerms

	data := psdc.ConvertToPaymentTerms(len(sdc.Header.Item), paymentTerms)

	return data
}

func (f *SubFunction) PaymentMethod(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.PaymentMethod {
	paymentMethod := psdc.SupplyChainRelationshipTransaction[0].PaymentMethod

	data := psdc.ConvertToPaymentMethod(len(sdc.Header.Item), paymentMethod)

	return data
}

func (f *SubFunction) ItemGrossWeight(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.ItemGrossWeight {
	var data []*api_processing_data_formatter.ItemGrossWeight

	item := sdc.Header.Item
	itemMap := make(map[string]api_input_reader.Item, len(item))
	for _, v := range item {
		itemMap[*v.Product] = v
	}

	for _, v := range psdc.ProductMasterGeneral {
		product := v.Product
		productGrossWeight := v.GrossWeight
		orderQuantityInBaseUnit := itemMap[product].OrderQuantityInBaseUnit
		itemGrossWeight := parseFloat32Ptr(*productGrossWeight * *orderQuantityInBaseUnit)

		datum := psdc.ConvertToItemGrossWeight(product, productGrossWeight, orderQuantityInBaseUnit, itemGrossWeight)
		data = append(data, datum)
	}

	return data
}

func (f *SubFunction) ItemNetWeight(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.ItemNetWeight {
	var data []*api_processing_data_formatter.ItemNetWeight

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
		value := elem.Field(i).Interface()
		if field == key {
			res[value], _ = jsonTypeConversion[T](elem.Interface())
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
