package subfunction

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"strings"

	"golang.org/x/xerrors"
)

func (f *SubFunction) SupplyChainRelationshipGeneral(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipGeneral, error) {
	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, Buyer, Seller
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_general_data
		WHERE (Buyer, Seller) = (?, ?);`, sdc.Header.Buyer, sdc.Header.Seller,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToSupplyChainRelationshipGeneral(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipDeliveryRelation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipDeliveryRelation, error) {
	var args []interface{}

	dataKey := psdc.ConvertToSupplyChainRelationshipDeliveryRelationKey()

	for _, v := range psdc.SupplyChainRelationshipGeneral {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
	}

	for _, v := range sdc.Header.Item {
		if v.DeliverToParty == nil {
			return nil, xerrors.Errorf("入力ファイルの'DeliverToParty'がnullです。")
		} else if v.DeliverFromParty == nil {
			return nil, xerrors.Errorf("入力ファイルの'DeliverToParty'がnullです。")
		}
		dataKey.DeliverToParty = append(dataKey.DeliverToParty, *v.DeliverToParty)
		dataKey.DeliverFromParty = append(dataKey.DeliverFromParty, *v.DeliverFromParty)
	}

	repeat1 := strings.Repeat("(?,?,?),", len(dataKey.SupplyChainRelationshipID)-1) + "(?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(args, dataKey.SupplyChainRelationshipID[i], dataKey.Buyer[i], dataKey.Seller[i])
	}

	repeat2 := strings.Repeat("(?,?),", len(dataKey.DeliverToParty)-1) + "(?,?)"
	for i := range dataKey.DeliverToParty {
		args = append(args, dataKey.DeliverToParty[i], dataKey.DeliverFromParty[i])
	}

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, Buyer, Seller, DeliverToParty, DeliverFromParty
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_delivery_relation_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller) IN ( `+repeat1+` )
		AND (DeliverToParty, DeliverFromParty) IN ( `+repeat2+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToSupplyChainRelationshipDeliveryRelation(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipDeliveryPlantRelation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipDeliveryPlantRelation, error) {
	var args []interface{}

	dataKey := psdc.ConvertToSupplyChainRelationshipDeliveryPlantRelationKey()

	for _, v := range psdc.SupplyChainRelationshipDeliveryRelation {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.SupplyChainRelationshipDeliveryID = append(dataKey.SupplyChainRelationshipDeliveryID, v.SupplyChainRelationshipDeliveryID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
		dataKey.DeliverToParty = append(dataKey.DeliverToParty, v.DeliverToParty)
		dataKey.DeliverFromParty = append(dataKey.DeliverFromParty, v.DeliverFromParty)
	}

	repeat := strings.Repeat("(?,?,?,?,?,?),", len(dataKey.DeliverToParty)-1) + "(?,?,?,?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(
			args,
			dataKey.SupplyChainRelationshipID[i],
			dataKey.SupplyChainRelationshipDeliveryID[i],
			dataKey.Buyer[i],
			dataKey.Seller[i],
			dataKey.DeliverToParty[i],
			dataKey.DeliverFromParty[i],
		)
	}

	args = append(args, dataKey.DefaultRelation)

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, SupplyChainRelationshipDeliveryPlantID,
		Buyer, Seller, DeliverToParty, DeliverFromParty, DeliverToPlant, DeliverFromPlant, DefaultRelation
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_delivery_plant_rel_data
		WHERE (SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, Buyer, Seller, DeliverToParty, DeliverFromParty) IN ( `+repeat+` )
		AND DefaultRelation = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToSupplyChainRelationshipDeliveryPlantRelation(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipTransaction(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipTransaction, error) {
	var args []interface{}

	supplyChainRelationshipGeneral := psdc.SupplyChainRelationshipGeneral

	repeat := strings.Repeat("(?,?,?),", len(supplyChainRelationshipGeneral)-1) + "(?,?,?)"
	for _, v := range supplyChainRelationshipGeneral {
		args = append(args, v.SupplyChainRelationshipID, v.Buyer, v.Seller)
	}

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, Buyer, Seller, TransactionCurrency, Incoterms, PaymentTerms, PaymentMethod, AccountAssignmentGroup
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_transaction_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToSupplyChainRelationshipTransaction(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipBillingRelation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipBillingRelation, error) {
	var args []interface{}

	dataKey := psdc.ConvertToSupplyChainRelationshipBillingRelationKey()

	for _, v := range psdc.SupplyChainRelationshipGeneral {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
	}

	repeat := strings.Repeat("(?,?,?),", len(dataKey.SupplyChainRelationshipID)-1) + "(?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(
			args,
			dataKey.SupplyChainRelationshipID[i],
			dataKey.Buyer[i],
			dataKey.Seller[i],
		)
	}

	args = append(args, dataKey.DefaultRelation)

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipBillingID, Buyer, Seller, BillToParty, BillFromParty, DefaultRelation, BillToCountry, BillFromCountry, IsExportImport, TransactionTaxCategory, TransactionTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_billing_relation_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller) IN ( `+repeat+` )
		AND DefaultRelation = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToSupplyChainRelationshipBillingRelation(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CalculateOrderID(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateOrderID, error) {
	dataKey := psdc.ConvertToCalculateOrderIDKey()

	dataKey.ServiceLabel = psdc.MetaData.ServiceLabel

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}

	dataQueryGets, err := psdc.ConvertToCalculateOrderIDQueryGets(rows)
	if err != nil {
		return nil, err
	}

	if dataQueryGets.LatestNumber == nil {
		return nil, xerrors.Errorf("LatestNumberがnullです。")
	}

	orderIDLatestNumber := dataQueryGets.LatestNumber
	orderID := *dataQueryGets.LatestNumber + 1

	data := psdc.ConvertToCalculateOrderID(orderIDLatestNumber, orderID)

	return data, err
}
