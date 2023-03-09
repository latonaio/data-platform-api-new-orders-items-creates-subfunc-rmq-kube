package subfunction

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"sort"
	"strings"
	"time"

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
	defer rows.Close()

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
	args := make([]interface{}, 0)

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

	if len(dataKey.SupplyChainRelationshipID) == 0 {
		return nil, xerrors.Errorf("psdc.SupplyChainRelationshipGeneralの'SupplyChainRelationshipID'がありません。")
	}
	repeat1 := strings.Repeat("(?,?,?),", len(dataKey.SupplyChainRelationshipID)-1) + "(?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(args, dataKey.SupplyChainRelationshipID[i], dataKey.Buyer[i], dataKey.Seller[i])
	}

	if len(dataKey.DeliverToParty) == 0 {
		return nil, xerrors.Errorf("入力ファイルの'DeliverToParty'がありません。")
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
	defer rows.Close()

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
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToSupplyChainRelationshipDeliveryPlantRelationKey()

	for _, v := range psdc.SupplyChainRelationshipDeliveryRelation {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.SupplyChainRelationshipDeliveryID = append(dataKey.SupplyChainRelationshipDeliveryID, v.SupplyChainRelationshipDeliveryID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
		dataKey.DeliverToParty = append(dataKey.DeliverToParty, v.DeliverToParty)
		dataKey.DeliverFromParty = append(dataKey.DeliverFromParty, v.DeliverFromParty)
	}

	if len(dataKey.DeliverToParty) == 0 {
		return nil, xerrors.Errorf("入力ファイルの'DeliverToParty'がありません。")
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
	defer rows.Close()

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
	args := make([]interface{}, 0)

	supplyChainRelationshipGeneral := psdc.SupplyChainRelationshipGeneral

	if len(supplyChainRelationshipGeneral) == 0 {
		return nil, xerrors.Errorf("'psdc.SupplyChainRelationshipGeneral'がありません。")
	}
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
	defer rows.Close()

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
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToSupplyChainRelationshipBillingRelationKey()

	for _, v := range psdc.SupplyChainRelationshipGeneral {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
	}

	if len(dataKey.SupplyChainRelationshipID) == 0 {
		return nil, xerrors.Errorf("psdc.SupplyChainRelationshipGeneralの'SupplyChainRelationshipID'がありません。")
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
	defer rows.Close()

	data, err := psdc.ConvertToSupplyChainRelationshipBillingRelation(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) SupplyChainRelationshipPaymentRelation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.SupplyChainRelationshipPaymentRelation, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToSupplyChainRelationshipPaymentRelationKey()

	for _, v := range psdc.SupplyChainRelationshipBillingRelation {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
		dataKey.BillToParty = append(dataKey.BillToParty, v.BillToParty)
		dataKey.BillFromParty = append(dataKey.BillFromParty, v.BillFromParty)
	}

	if len(dataKey.SupplyChainRelationshipID) == 0 {
		return nil, xerrors.Errorf("psdc.SupplyChainRelationshipBillingRelation'SupplyChainRelationshipID'がありません。")
	}
	repeat := strings.Repeat("(?,?,?,?,?),", len(dataKey.SupplyChainRelationshipID)-1) + "(?,?,?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(
			args,
			dataKey.SupplyChainRelationshipID[i],
			dataKey.Buyer[i],
			dataKey.Seller[i],
			dataKey.BillToParty[i],
			dataKey.BillFromParty[i],
		)
	}

	args = append(args, dataKey.DefaultRelation)

	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, SupplyChainRelationshipBillingID, SupplyChainRelationshipPaymentID, Buyer, Seller, BillToParty, BillFromParty, Payer, Payee, DefaultRelation
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_supply_chain_relationship_payment_relation_data
		WHERE (SupplyChainRelationshipID, Buyer, Seller, BillToParty, BillFromParty) IN ( `+repeat+` )
		AND DefaultRelation = ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToSupplyChainRelationshipPaymentRelation(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) HeaderInvoiceDocumentDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.HeaderInvoiceDocumentDate, error) {
	rows, err := f.db.Query(
		`SELECT PaymentTerms, BaseDate, BaseDateCalcAddMonth, BaseDateCalcFixedDate, PaymentDueDateCalcAddMonth, PaymentDueDateCalcFixedDate
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_data
			WHERE PaymentTerms = ?;`, psdc.SupplyChainRelationshipTransaction[0].PaymentTerms,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	psdc.PaymentTerms, err = psdc.ConvertToPaymentTerms(rows)
	if err != nil {
		return nil, err
	}

	if sdc.Header.InvoiceDocumentDate != nil {
		if *sdc.Header.InvoiceDocumentDate != "" {
			data, err := psdc.ConvertToHeaderInvoiceDocumentDate(sdc)
			return data, err
		}
	}

	requestedDeliveryDate, err := psdc.ConvertToRequestedDeliveryDate(sdc)
	if err != nil {
		return nil, err
	}

	invoiceDocumentDate, err := calculateHeaderInvoiceDocumentDate(psdc, requestedDeliveryDate.RequestedDeliveryDate, psdc.PaymentTerms)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToCaluculateHeaderInvoiceDocumentDate(sdc, invoiceDocumentDate)

	return data, err
}

func calculateHeaderInvoiceDocumentDate(
	psdc *api_processing_data_formatter.SDC,
	requestedDeliveryDate string,
	paymentTerms []*api_processing_data_formatter.PaymentTerms,
) (string, error) {
	format := "2006-01-02"
	t, err := time.Parse(format, requestedDeliveryDate)
	if err != nil {
		return "", err
	}

	sort.Slice(paymentTerms, func(i, j int) bool {
		return paymentTerms[i].BaseDate < paymentTerms[j].BaseDate
	})

	day := t.Day()
	for i, v := range paymentTerms {
		if day <= v.BaseDate {
			if v.BaseDateCalcAddMonth == nil {
				return "", xerrors.Errorf("paymentTermsの'BaseDateCalcAddMonth'がありません。")
			}
			t = time.Date(t.Year(), t.Month()+time.Month(*v.BaseDateCalcAddMonth)+1, 0, 0, 0, 0, 0, time.UTC)
			if v.BaseDateCalcFixedDate == nil {
				return "", xerrors.Errorf("paymentTermsの'BaseDateCalcFixedDate'がありません。")
			}
			if *v.BaseDateCalcFixedDate == 31 {
				t = time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC)
			} else {
				t = time.Date(t.Year(), t.Month(), *v.BaseDateCalcFixedDate, 0, 0, 0, 0, time.UTC)
			}
			break
		}
		if len(paymentTerms) == 0 {
			return "", xerrors.Errorf("'paymentTerms'がありません。")
		}
		if i == len(paymentTerms)-1 {
			return "", xerrors.Errorf("'data_platform_payment_terms_payment_terms_data'テーブルが不適切です。")
		}
	}

	res := t.Format(format)

	return res, nil
}

func (f *SubFunction) HeaderPricingDate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.PricingDate {
	pricingDate := getSystemDate()

	if sdc.Header.PricingDate != nil {
		if *sdc.Header.PricingDate != "" {
			pricingDate = *sdc.Header.PricingDate
		}
	}

	data := psdc.ConvertToPricingDate(pricingDate)

	return data
}

func (f *SubFunction) HeaderPriceDetnExchangeRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.PriceDetnExchangeRate {
	data := new(api_processing_data_formatter.PriceDetnExchangeRate)

	if sdc.Header.PriceDetnExchangeRate != nil {
		data = psdc.ConvertToHeaderPriceDetnExchangeRate(sdc)
	}

	return data
}

func (f *SubFunction) HeaderAccountingExchangeRate(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.AccountingExchangeRate {
	data := new(api_processing_data_formatter.AccountingExchangeRate)

	if sdc.Header.AccountingExchangeRate != nil {
		data = psdc.ConvertToHeaderAccountingExchangeRate(sdc)
	}

	return data
}

func (f *SubFunction) BusinessPartnerGeneralDeliveryRelation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) error {
	args := make([]interface{}, 0)
	var err error

	supplyChainRelationshipDeliveryRelation := psdc.SupplyChainRelationshipDeliveryRelation

	dataKey := psdc.ConvertToBusinessPartnerGeneralDeliveryRelationKey(len(supplyChainRelationshipDeliveryRelation))

	for i := range dataKey {
		dataKey[i].Buyer = supplyChainRelationshipDeliveryRelation[i].Buyer
		dataKey[i].Seller = supplyChainRelationshipDeliveryRelation[i].Seller
		dataKey[i].DeliverToParty = supplyChainRelationshipDeliveryRelation[i].DeliverToParty
		dataKey[i].DeliverFromParty = supplyChainRelationshipDeliveryRelation[i].DeliverFromParty
	}

	if len(dataKey) == 0 {
		return xerrors.Errorf("BusinessPartnerGeneralDeliveryRelation'dataKey'がありません。")
	}
	repeat := strings.Repeat("?,", len(dataKey)-1) + "?"

	for _, v := range dataKey {
		args = append(args, v.Buyer)
	}

	psdc.BusinessPartnerGeneralBuyer, err = f.businessPartnerGeneral(sdc, psdc, repeat, args)
	if err != nil {
		return err
	}

	args = nil
	for _, v := range dataKey {
		args = append(args, v.Seller)
	}

	psdc.BusinessPartnerGeneralSeller, err = f.businessPartnerGeneral(sdc, psdc, repeat, args)
	if err != nil {
		return err
	}

	args = nil
	for _, v := range dataKey {
		args = append(args, v.DeliverToParty)
	}

	psdc.BusinessPartnerGeneralDeliverToParty, err = f.businessPartnerGeneral(sdc, psdc, repeat, args)
	if err != nil {
		return err
	}

	args = nil
	for _, v := range dataKey {
		args = append(args, v.DeliverFromParty)
	}

	psdc.BusinessPartnerGeneralDeliverFromParty, err = f.businessPartnerGeneral(sdc, psdc, repeat, args)
	if err != nil {
		return err
	}

	return err
}

func (f *SubFunction) BusinessPartnerGeneralBillingRelation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) error {
	args := make([]interface{}, 0)
	var err error

	supplyChainRelationshipBillingRelation := psdc.SupplyChainRelationshipBillingRelation

	dataKey := psdc.ConvertToBusinessPartnerGeneralBillingRelationKey(len(supplyChainRelationshipBillingRelation))

	for i := range dataKey {
		dataKey[i].BillToParty = supplyChainRelationshipBillingRelation[i].BillToParty
		dataKey[i].BillFromParty = supplyChainRelationshipBillingRelation[i].BillFromParty
	}

	if len(dataKey) == 0 {
		return xerrors.Errorf("BusinessPartnerGeneralBillingRelation'dataKey'がありません。")
	}
	repeat := strings.Repeat("?,", len(dataKey)-1) + "?"

	for _, v := range dataKey {
		args = append(args, v.BillToParty)
	}

	psdc.BusinessPartnerGeneralBillToParty, err = f.businessPartnerGeneral(sdc, psdc, repeat, args)
	if err != nil {
		return err
	}

	args = nil
	for _, v := range dataKey {
		args = append(args, v.BillFromParty)
	}

	psdc.BusinessPartnerGeneralBillFromParty, err = f.businessPartnerGeneral(sdc, psdc, repeat, args)
	if err != nil {
		return err
	}

	return err
}

func (f *SubFunction) BusinessPartnerGeneralPaymentRelation(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) error {
	args := make([]interface{}, 0)
	var err error

	supplyChainRelationshipPaymentRelation := psdc.SupplyChainRelationshipPaymentRelation

	dataKey := psdc.ConvertToBusinessPartnerGeneralPaymentRelationKey(len(supplyChainRelationshipPaymentRelation))

	for i := range dataKey {
		dataKey[i].Payer = supplyChainRelationshipPaymentRelation[i].Payer
		dataKey[i].Payee = supplyChainRelationshipPaymentRelation[i].Payee
	}

	if len(dataKey) == 0 {
		return xerrors.Errorf("BusinessPartnerGeneralPaymentRelation'dataKey'がありません。")
	}
	repeat := strings.Repeat("?,", len(dataKey)-1) + "?"

	for _, v := range dataKey {
		args = append(args, v.Payer)
	}

	psdc.BusinessPartnerGeneralPayer, err = f.businessPartnerGeneral(sdc, psdc, repeat, args)
	if err != nil {
		return err
	}

	args = nil
	for _, v := range dataKey {
		args = append(args, v.Payee)
	}

	psdc.BusinessPartnerGeneralPayee, err = f.businessPartnerGeneral(sdc, psdc, repeat, args)
	if err != nil {
		return err
	}

	return err
}

func (f *SubFunction) businessPartnerGeneral(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	repeat string,
	args []interface{},
) ([]*api_processing_data_formatter.BusinessPartnerGeneral, error) {

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
