package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-orders-creates-subfunc-items-rmq-kube/API_Input_Reader"
	"data-platform-api-orders-creates-subfunc-items-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"encoding/json"
	"time"

	"golang.org/x/xerrors"
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
	defer rows.Close()
	res := make([]*SupplyChainRelationshipGeneral, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipGeneral{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_general_data'テーブルに対象のレコードが存在しません。")
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
	defer rows.Close()
	res := make([]*SupplyChainRelationshipDeliveryRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipDeliveryRelation{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_delivery_relation_data'テーブルに対象のレコードが存在しません。")
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
	defer rows.Close()
	res := make([]*SupplyChainRelationshipDeliveryPlantRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipDeliveryPlantRelation{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_delivery_plant_relation_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToSupplyChainRelationshipTransaction(rows *sql.Rows) ([]*SupplyChainRelationshipTransaction, error) {
	defer rows.Close()
	res := make([]*SupplyChainRelationshipTransaction, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipTransaction{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_transaction_data'テーブルに対象のレコードが存在しません。")
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
	defer rows.Close()
	res := make([]*SupplyChainRelationshipBillingRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipBillingRelation{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_delivery_plant_relation_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToSupplyChainRelationshipPaymentRelationKey() *SupplyChainRelationshipPaymentRelationKey {
	pm := &requests.SupplyChainRelationshipPaymentRelationKey{
		DefaultRelation: true,
	}

	data := pm
	res := SupplyChainRelationshipPaymentRelationKey{
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
		BillToParty:               data.BillToParty,
		BillFromParty:             data.BillFromParty,
		DefaultRelation:           data.DefaultRelation,
	}

	return &res
}

func (psdc *SDC) ConvertToSupplyChainRelationshipPaymentRelation(rows *sql.Rows) ([]*SupplyChainRelationshipPaymentRelation, error) {
	defer rows.Close()
	res := make([]*SupplyChainRelationshipPaymentRelation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipPaymentRelation{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.Payer,
			&pm.Payee,
			&pm.DefaultRelation,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &SupplyChainRelationshipPaymentRelation{
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			BillToParty:                      data.BillToParty,
			BillFromParty:                    data.BillFromParty,
			Payer:                            data.Payer,
			Payee:                            data.Payee,
			DefaultRelation:                  data.DefaultRelation,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_payment_relation_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToPaymentTerms(rows *sql.Rows) ([]*PaymentTerms, error) {
	defer rows.Close()
	res := make([]*PaymentTerms, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PaymentTerms{}

		err := rows.Scan(
			&pm.PaymentTerms,
			&pm.BaseDate,
			&pm.BaseDateCalcAddMonth,
			&pm.BaseDateCalcFixedDate,
			&pm.PaymentDueDateCalcAddMonth,
			&pm.PaymentDueDateCalcFixedDate,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &PaymentTerms{
			PaymentTerms:                data.PaymentTerms,
			BaseDate:                    data.BaseDate,
			BaseDateCalcAddMonth:        data.BaseDateCalcAddMonth,
			BaseDateCalcFixedDate:       data.BaseDateCalcFixedDate,
			PaymentDueDateCalcAddMonth:  data.PaymentDueDateCalcAddMonth,
			PaymentDueDateCalcFixedDate: data.PaymentDueDateCalcFixedDate,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_payment_terms_payment_terms_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToHeaderInvoiceDocumentDate(sdc *api_input_reader.SDC) (*HeaderInvoiceDocumentDate, error) {
	pm := &requests.HeaderInvoiceDocumentDate{}

	if sdc.Header.InvoiceDocumentDate == nil {
		return nil, xerrors.Errorf("入力ファイルの'InvoiceDocumentDate'がありません。")
	}
	pm.InvoiceDocumentDate = *sdc.Header.InvoiceDocumentDate
	data := pm

	res := HeaderInvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToRequestedDeliveryDate(sdc *api_input_reader.SDC) (*HeaderInvoiceDocumentDate, error) {
	if sdc.Header.RequestedDeliveryDate == nil {
		return nil, xerrors.Errorf("RequestedDeliveryDateがnullです。")
	}

	if sdc.Header.RequestedDeliveryDate == nil {
		return nil, xerrors.Errorf("入力ファイルの'RequestedDeliveryDate'がありません。")
	}
	pm := &requests.HeaderInvoiceDocumentDate{
		RequestedDeliveryDate: *sdc.Header.RequestedDeliveryDate,
	}

	data := pm
	res := HeaderInvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCaluculateHeaderInvoiceDocumentDate(sdc *api_input_reader.SDC, invoiceDocumentDate string) (*HeaderInvoiceDocumentDate, error) {
	if sdc.Header.RequestedDeliveryDate == nil {
		return nil, xerrors.Errorf("入力ファイルの'RequestedDeliveryDate'がありません。")
	}
	pm := &requests.HeaderInvoiceDocumentDate{
		RequestedDeliveryDate: *sdc.Header.RequestedDeliveryDate,
	}

	pm.InvoiceDocumentDate = invoiceDocumentDate

	data := pm
	res := HeaderInvoiceDocumentDate{
		RequestedDeliveryDate: data.RequestedDeliveryDate,
		InvoiceDocumentDate:   data.InvoiceDocumentDate,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToPricingDate(pricingDate string) *PricingDate {
	pm := &requests.PricingDate{}

	pm.PricingDate = pricingDate

	data := pm
	res := PricingDate{
		PricingDate: data.PricingDate,
	}

	return &res
}

func (psdc *SDC) ConvertToHeaderPriceDetnExchangeRate(sdc *api_input_reader.SDC) *PriceDetnExchangeRate {
	pm := &requests.PriceDetnExchangeRate{
		PriceDetnExchangeRate: sdc.Header.PriceDetnExchangeRate,
	}

	data := pm
	res := PriceDetnExchangeRate{
		PriceDetnExchangeRate: data.PriceDetnExchangeRate,
	}

	return &res
}

func (psdc *SDC) ConvertToHeaderAccountingExchangeRate(sdc *api_input_reader.SDC) *AccountingExchangeRate {
	pm := &requests.AccountingExchangeRate{
		AccountingExchangeRate: sdc.Header.AccountingExchangeRate,
	}

	data := pm
	res := AccountingExchangeRate{
		AccountingExchangeRate: data.AccountingExchangeRate,
	}

	return &res
}

func (psdc *SDC) ConvertToBusinessPartnerGeneralDeliveryRelationKey(length int) []*BusinessPartnerGeneralDeliveryRelationKey {
	res := make([]*BusinessPartnerGeneralDeliveryRelationKey, 0)

	for i := 0; i < length; i++ {
		pm := &requests.BusinessPartnerGeneralDeliveryRelationKey{}

		data := pm
		res = append(res, &BusinessPartnerGeneralDeliveryRelationKey{
			Buyer:            data.Buyer,
			Seller:           data.Seller,
			DeliverToParty:   data.DeliverToParty,
			DeliverFromParty: data.DeliverFromParty,
		})
	}

	return res
}

func (psdc *SDC) ConvertToBusinessPartnerGeneralBillingRelationKey(length int) []*BusinessPartnerGeneralBillingRelationKey {
	res := make([]*BusinessPartnerGeneralBillingRelationKey, 0)

	for i := 0; i < length; i++ {
		pm := &requests.BusinessPartnerGeneralBillingRelationKey{}

		data := pm
		res = append(res, &BusinessPartnerGeneralBillingRelationKey{
			BillToParty:   data.BillToParty,
			BillFromParty: data.BillFromParty,
		})
	}

	return res
}

func (psdc *SDC) ConvertToBusinessPartnerGeneralPaymentRelationKey(length int) []*BusinessPartnerGeneralPaymentRelationKey {
	res := make([]*BusinessPartnerGeneralPaymentRelationKey, 0)

	for i := 0; i < length; i++ {
		pm := &requests.BusinessPartnerGeneralPaymentRelationKey{}

		data := pm
		res = append(res, &BusinessPartnerGeneralPaymentRelationKey{
			Payer: data.Payer,
			Payee: data.Payee,
		})
	}

	return res
}

func (psdc *SDC) ConvertToBusinessPartnerGeneral(rows *sql.Rows) ([]*BusinessPartnerGeneral, error) {
	defer rows.Close()
	res := make([]*BusinessPartnerGeneral, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.BusinessPartnerGeneral{}

		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
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
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			AddressID:               data.AddressID,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_business_partner_general_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

// Item
func (psdc *SDC) ConvertToOrderItem(sdc *api_input_reader.SDC) []*OrderItem {
	res := make([]*OrderItem, 0)

	for i := range sdc.Header.Item {
		pm := &requests.OrderItem{}

		pm.OrderItemNumber = i + 1
		sdc.Header.Item[i].OrderItem = i + 1

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
	defer rows.Close()
	res := make([]*ProductTaxClassificationBillToCountry, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductTaxClassificationBillToCountry{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_master_tax_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToProductTaxClassificationBillFromCountry(rows *sql.Rows) ([]*ProductTaxClassificationBillFromCountry, error) {
	defer rows.Close()
	res := make([]*ProductTaxClassificationBillFromCountry, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductTaxClassificationBillFromCountry{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_master_tax_data'テーブルに対象のレコードが存在しません。")
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
	defer rows.Close()
	res := make([]*ProductMasterGeneral, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductMasterGeneral{}

		err := rows.Scan(
			&pm.Product,
			&pm.BaseUnit,
			&pm.ProductGroup,
			&pm.ProductStandardID,
			&pm.GrossWeight,
			&pm.NetWeight,
			&pm.WeightUnit,
			&pm.InternalCapacityQuantity,
			&pm.InternalCapacityQuantityUnit,
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
			InternalCapacityQuantity:      data.InternalCapacityQuantity,
			InternalCapacityQuantityUnit:  data.InternalCapacityQuantityUnit,
			ItemCategory:                  data.ItemCategory,
			ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
			CountryOfOrigin:               data.CountryOfOrigin,
			CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_master_general_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrderItemTextKey(length int) []*OrderItemTextKey {
	res := make([]*OrderItemTextKey, 0)

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
	defer rows.Close()
	res := make([]*OrderItemText, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrderItemText{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_master_product_description_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToItemCategoryIsINVP() []*ItemCategoryIsINVP {
	res := make([]*ItemCategoryIsINVP, 0)

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
	defer rows.Close()
	res := make([]*StockConfPlantRelationProduct, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.StockConfPlantRelationProduct{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_stock_conf_plant_relation_product_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToProductMasterBPPlantKey(length int) []*ProductMasterBPPlantKey {
	res := make([]*ProductMasterBPPlantKey, 0)

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
	defer rows.Close()
	res := make([]*ProductMasterBPPlant, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductMasterBPPlant{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_master_bp_plant_data'テーブルに対象のレコードが存在しません。")
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
	defer rows.Close()
	res := make([]*ProductionPlantRelationProduct, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductionPlantRelationProduct{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_production_plant_relation_product_data'テーブルに対象のレコードが存在しません。")
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
	defer rows.Close()
	res := make([]*SupplyChainRelationshipDeliveryPlantRelationProduct, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.SupplyChainRelationshipDeliveryPlantRelationProduct{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_supply_chain_relationship_delivery_plant_relation_product_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToTimeZoneKey(length int) []*TimeZoneKey {
	res := make([]*TimeZoneKey, 0)

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
	defer rows.Close()
	res := make([]*TimeZone, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.TimeZone{}

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
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_plant_general_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToIncoterms(incoterms *string) *Incoterms {
	pm := &requests.Incoterms{}

	pm.Incoterms = incoterms

	data := pm
	res := &Incoterms{
		Incoterms: data.Incoterms,
	}

	return res
}

func (psdc *SDC) ConvertToItemPaymentTerms(paymentTerms *string) *ItemPaymentTerms {
	pm := &requests.ItemPaymentTerms{}

	pm.PaymentTerms = paymentTerms

	data := pm
	res := &ItemPaymentTerms{
		PaymentTerms: data.PaymentTerms,
	}

	return res
}

func (psdc *SDC) ConvertToPaymentMethod(paymentMethod *string) *PaymentMethod {
	pm := &requests.PaymentMethod{}

	pm.PaymentMethod = paymentMethod

	data := pm
	res := &PaymentMethod{
		PaymentMethod: data.PaymentMethod,
	}

	return res
}

func (psdc *SDC) ConvertToItemGrossWeight(orderItem int, product string, productGrossWeight, orderQuantityInBaseUnit, itemGrossWeght *float32) *ItemGrossWeight {
	pm := &requests.ItemGrossWeight{}

	pm.OrderItem = orderItem
	pm.Product = product
	pm.ProductGrossWeight = productGrossWeight
	pm.OrderQuantityInBaseUnit = orderQuantityInBaseUnit
	pm.ItemGrossWeight = itemGrossWeght

	data := pm
	res := ItemGrossWeight{
		OrderItem:               data.OrderItem,
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

func (psdc *SDC) ConvertToTaxCode(product, definedTaxClassification string, isExportImport *bool, taxCode *string) *TaxCode {
	pm := &requests.TaxCode{}

	pm.Product = product
	pm.DefinedTaxClassification = definedTaxClassification
	pm.IsExportImport = isExportImport
	pm.TaxCode = taxCode

	data := pm
	res := TaxCode{
		Product:                  data.Product,
		DefinedTaxClassification: data.DefinedTaxClassification,
		IsExportImport:           data.IsExportImport,
		TaxCode:                  data.TaxCode,
	}

	return &res
}

func (psdc *SDC) ConvertToTaxRateKey() *TaxRateKey {
	pm := &requests.TaxRateKey{
		Country: "JP",
	}

	data := pm
	res := TaxRateKey{
		Country:           data.Country,
		TaxCode:           data.TaxCode,
		ValidityEndDate:   data.ValidityEndDate,
		ValidityStartDate: data.ValidityStartDate,
	}

	return &res
}

func (psdc *SDC) ConvertToTaxRate(rows *sql.Rows) ([]*TaxRate, error) {
	defer rows.Close()
	res := make([]*TaxRate, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.TaxRate{}

		err := rows.Scan(
			&pm.Country,
			&pm.TaxCode,
			&pm.ValidityEndDate,
			&pm.ValidityStartDate,
			&pm.TaxRate,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &TaxRate{
			Country:           data.Country,
			TaxCode:           data.TaxCode,
			ValidityEndDate:   data.ValidityEndDate,
			ValidityStartDate: data.ValidityStartDate,
			TaxRate:           data.TaxRate,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_tax_code_tax_rate_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToCheckProductionPlantBatchKey() *ProductionPlantBatchKey {
	pm := &requests.ProductionPlantBatchKey{}

	data := pm
	res := &ProductionPlantBatchKey{
		Product:                        data.Product,
		ProductionPlantBatch:           data.ProductionPlantBatch,
		ProductionPlantBusinessPartner: data.ProductionPlantBusinessPartner,
		ProductionPlant:                data.ProductionPlant,
		ValidityStartDate:              data.ValidityStartDate,
		ValidityEndDate:                data.ValidityEndDate,
	}

	return res
}

func (psdc *SDC) ConvertToProductionPlantBatchExconf(resData map[string]interface{}) (*ProductionPlantBatchMasterdata, error) {
	pm := &requests.ProductionPlantBatchMasterdata{}

	result := resData["result"].(bool)
	if !result {
		return nil, xerrors.Errorf(resData["message"].(string))
	}

	raw, err := json.Marshal(resData["message"].(map[string]interface{})["ProductionPlantBatchCheck"])
	if err != nil {
		return nil, xerrors.Errorf("data marshal error :%#v", err.Error())
	}
	err = json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("input data marshal error :%#v", err.Error())
	}

	data := pm
	res := &ProductionPlantBatchMasterdata{
		Product:                    data.Product,
		BusinessPartner:            data.BusinessPartner,
		Plant:                      data.Plant,
		Batch:                      data.Batch,
		ProductionPlantBatchExConf: data.ProductionPlantBatchExConf,
	}

	return res, nil
}

func (psdc *SDC) ConvertToProductionPlantBatchMaster(i int, countryOfOrigin *string, validityStartDate string) *ProductionPlantBatchMasterdata {
	pm := &requests.ProductionPlantBatchMasterdata{
		Product:             psdc.ProductionPlantBatchMasterdata[i].Product,
		BusinessPartner:     psdc.ProductionPlantBatchMasterdata[i].BusinessPartner,
		Plant:               psdc.ProductionPlantBatchMasterdata[i].Plant,
		Batch:               psdc.ProductionPlantBatchMasterdata[i].Batch,
		CountryOfOrigin:     *countryOfOrigin,
		ValidityStartDate:   &validityStartDate,
		ManufactureDate:     nil,
		CreationDateTime:    getSystemDateTime(),
		LastChangeDateTime:  getSystemDateTime(),
		IsMarkedForDeletion: false,
	}

	data := pm
	res := &ProductionPlantBatchMasterdata{
		Product:             data.Product,
		BusinessPartner:     data.BusinessPartner,
		Plant:               data.Plant,
		Batch:               data.Batch,
		CountryOfOrigin:     data.CountryOfOrigin,
		ValidityStartDate:   data.ValidityStartDate,
		ManufactureDate:     data.ManufactureDate,
		CreationDateTime:    data.CreationDateTime,
		LastChangeDateTime:  data.LastChangeDateTime,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}

	return res
}

func (psdc *SDC) ConvertProductionPlantBatch() []*ProductionPlantBatch {
	res := make([]*ProductionPlantBatch, 0)

	for _, v := range psdc.ProductionPlantBatchMasterdata {
		pm := &requests.ProductionPlantBatch{}

		pm.Product = v.Product
		pm.Batch = v.Batch
		pm.BusinessPartner = *v.BusinessPartner
		pm.Plant = v.Plant
		pm.ProductionPlantBatchExConf = v.ProductionPlantBatchExConf

		data := pm
		res = append(res, &ProductionPlantBatch{
			Product:                    data.Product,
			Batch:                      data.Batch,
			BusinessPartner:            data.BusinessPartner,
			Plant:                      data.Plant,
			ProductionPlantBatchExConf: data.ProductionPlantBatchExConf,
		})

	}

	return res
}

func (psdc *SDC) ConvertToCheckDeliverToPlantBatchKey() *DeliverToPlantBatchKey {
	pm := &requests.DeliverToPlantBatchKey{}

	data := pm
	res := &DeliverToPlantBatchKey{
		Product:             data.Product,
		DeliverToPlantBatch: data.DeliverToPlantBatch,
		DeliverToParty:      data.DeliverToParty,
		DeliverToPlant:      data.DeliverToPlant,
		ValidityStartDate:   data.ValidityStartDate,
		ValidityEndDate:     data.ValidityEndDate,
	}

	return res
}

func (psdc *SDC) ConvertToDeliverToPlantBatchExconf(resData map[string]interface{}) (*DeliverToPlantBatchMasterdata, error) {
	pm := &requests.DeliverToPlantBatchMasterdata{}

	result := resData["result"].(bool)
	if !result {
		return nil, xerrors.Errorf(resData["message"].(string))
	}

	raw, err := json.Marshal(resData["message"].(map[string]interface{})["DeliverToPlantBatchCheck"])
	if err != nil {
		return nil, xerrors.Errorf("data marshal error :%#v", err.Error())
	}
	err = json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("input data marshal error :%#v", err.Error())
	}

	data := pm
	res := &DeliverToPlantBatchMasterdata{
		Product:                   data.Product,
		BusinessPartner:           data.BusinessPartner,
		Plant:                     data.Plant,
		Batch:                     data.Batch,
		DeliverToPlantBatchExConf: data.DeliverToPlantBatchExConf,
	}

	return res, nil
}

func (psdc *SDC) ConvertToDeliverToPlantBatchMaster(i int, countryOfOrigin *string, validityStartDate string) *DeliverToPlantBatchMasterdata {
	pm := &requests.DeliverToPlantBatchMasterdata{
		Product:             psdc.ProductionPlantBatchMasterdata[i].Product,
		BusinessPartner:     psdc.ProductionPlantBatchMasterdata[i].BusinessPartner,
		Plant:               psdc.ProductionPlantBatchMasterdata[i].Plant,
		Batch:               psdc.ProductionPlantBatchMasterdata[i].Batch,
		CountryOfOrigin:     *countryOfOrigin,
		ValidityStartDate:   &validityStartDate,
		ManufactureDate:     nil,
		CreationDateTime:    getSystemDateTime(),
		LastChangeDateTime:  getSystemDateTime(),
		IsMarkedForDeletion: false,
	}

	data := pm
	res := &DeliverToPlantBatchMasterdata{
		Product:             data.Product,
		BusinessPartner:     data.BusinessPartner,
		Plant:               data.Plant,
		Batch:               data.Batch,
		CountryOfOrigin:     data.CountryOfOrigin,
		ValidityStartDate:   data.ValidityStartDate,
		ManufactureDate:     data.ManufactureDate,
		CreationDateTime:    data.CreationDateTime,
		LastChangeDateTime:  data.LastChangeDateTime,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}

	return res

}

func (psdc *SDC) ConvertDeliverToPlantBatch() []*DeliverToPlantBatch {
	res := make([]*DeliverToPlantBatch, 0)

	for _, v := range psdc.DeliverToPlantBatchMasterdata {
		pm := &requests.DeliverToPlantBatch{}

		pm.Product = v.Product
		pm.Batch = v.Batch
		pm.BusinessPartner = *v.BusinessPartner
		pm.Plant = v.Plant
		pm.DeliverToPlantBatchExConf = v.DeliverToPlantBatchExConf

		data := pm
		res = append(res, &DeliverToPlantBatch{
			Product:                   data.Product,
			Batch:                     data.Batch,
			BusinessPartner:           data.BusinessPartner,
			Plant:                     data.Plant,
			DeliverToPlantBatchExConf: data.DeliverToPlantBatchExConf,
		})

	}

	return res
}

func (psdc *SDC) ConvertToCheckDeliverFromPlantBatchKey() *DeliverFromPlantBatchKey {
	pm := &requests.DeliverFromPlantBatchKey{}

	data := pm
	res := &DeliverFromPlantBatchKey{
		Product:               data.Product,
		DeliverFromPlantBatch: data.DeliverFromPlantBatch,
		DeliverFromParty:      data.DeliverFromParty,
		DeliverFromPlant:      data.DeliverFromPlant,
		ValidityStartDate:     data.ValidityStartDate,
		ValidityEndDate:       data.ValidityEndDate,
	}

	return res
}

func (psdc *SDC) ConvertToDeliverFromPlantBatchExconf(resData map[string]interface{}) (*DeliverFromPlantBatchMasterdata, error) {
	pm := &requests.DeliverFromPlantBatchMasterdata{}

	result := resData["result"].(bool)
	if !result {
		return nil, xerrors.Errorf(resData["message"].(string))
	}

	raw, err := json.Marshal(resData["message"].(map[string]interface{})["DeliverFromPlantBatchCheck"])
	if err != nil {
		return nil, xerrors.Errorf("data marshal error :%#v", err.Error())
	}
	err = json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("input data marshal error :%#v", err.Error())
	}

	data := pm
	res := &DeliverFromPlantBatchMasterdata{
		Product:                     data.Product,
		BusinessPartner:             data.BusinessPartner,
		Plant:                       data.Plant,
		Batch:                       data.Batch,
		DeliverFromPlantBatchExConf: data.DeliverFromPlantBatchExConf,
	}

	return res, nil
}

func (psdc *SDC) ConvertToDeliverFromPlantBatchMaster(i int, countryOfOrigin *string, validityStartDate string) *DeliverFromPlantBatchMasterdata {
	pm := &requests.DeliverFromPlantBatchMasterdata{
		Product:             psdc.ProductionPlantBatchMasterdata[i].Product,
		BusinessPartner:     psdc.ProductionPlantBatchMasterdata[i].BusinessPartner,
		Plant:               psdc.ProductionPlantBatchMasterdata[i].Plant,
		Batch:               psdc.ProductionPlantBatchMasterdata[i].Batch,
		CountryOfOrigin:     *countryOfOrigin,
		ValidityStartDate:   &validityStartDate,
		ManufactureDate:     nil,
		CreationDateTime:    getSystemDateTime(),
		LastChangeDateTime:  getSystemDateTime(),
		IsMarkedForDeletion: false,
	}

	data := pm
	res := &DeliverFromPlantBatchMasterdata{
		Product:             data.Product,
		BusinessPartner:     data.BusinessPartner,
		Plant:               data.Plant,
		Batch:               data.Batch,
		CountryOfOrigin:     data.CountryOfOrigin,
		ValidityStartDate:   data.ValidityStartDate,
		ManufactureDate:     data.ManufactureDate,
		CreationDateTime:    data.CreationDateTime,
		LastChangeDateTime:  data.LastChangeDateTime,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}

	return res

}

func (psdc *SDC) ConvertDeliverFromPlantBatch() []*DeliverFromPlantBatch {
	res := make([]*DeliverFromPlantBatch, 0)

	for _, v := range psdc.DeliverFromPlantBatchMasterdata {
		pm := &requests.DeliverFromPlantBatch{}

		pm.Product = v.Product
		pm.Batch = v.Batch
		pm.BusinessPartner = *v.BusinessPartner
		pm.Plant = v.Plant
		pm.DeliverFromPlantBatchExConf = v.DeliverFromPlantBatchExConf

		data := pm
		res = append(res, &DeliverFromPlantBatch{
			Product:                     data.Product,
			Batch:                       data.Batch,
			BusinessPartner:             data.BusinessPartner,
			Plant:                       data.Plant,
			DeliverFromPlantBatchExConf: data.DeliverFromPlantBatchExConf,
		})

	}

	return res
}

func (psdc *SDC) ConvertToStockConfirmationKey() *StockConfirmationKey {
	pm := &requests.StockConfirmationKey{}

	data := pm
	res := &StockConfirmationKey{
		OrderID:                          data.OrderID,
		OrderItem:                        data.OrderItem,
		Product:                          data.Product,
		StockConfirmationBusinessPartner: data.StockConfirmationBusinessPartner,
		StockConfirmationPlant:           data.StockConfirmationPlant,
		ScheduleLineOrderQuantity:        data.ScheduleLineOrderQuantity,
		RequestedDeliveryDate:            data.RequestedDeliveryDate,
		StockConfirmationPlantBatch:      data.StockConfirmationPlantBatch,
		StockConfirmationPlantBatchValidityStartDate: data.StockConfirmationPlantBatchValidityStartDate,
		StockConfirmationPlantBatchValidityEndDate:   data.StockConfirmationPlantBatchValidityEndDate,
		StockConfirmationIsLotUnit:                   data.StockConfirmationIsLotUnit,
		StockConfirmationIsOrdinary:                  data.StockConfirmationIsOrdinary,
	}

	return res
}

func (psdc *SDC) ConvertToStockConfirmation(resData map[string]interface{}, stockConfirmationIsOrdinary, stockConfirmationIsLotUnit bool, orderID, orderItem int) (*StockConfirmation, error) {
	pm := &requests.StockConfirmation{}

	result := resData["result"].(bool)
	if !result {
		return nil, xerrors.Errorf(resData["message"].(string))
	}

	raw, err := json.Marshal(resData["message"].(map[string]interface{})["ProductStockAvailabilityCheck"])
	if err != nil {
		return nil, xerrors.Errorf("data marshal error :%#v", err.Error())
	}
	err = json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("input data marshal error :%#v", err.Error())
	}

	pm.OrderID = orderID
	pm.OrderItem = orderItem
	pm.StockConfirmationIsOrdinary = stockConfirmationIsOrdinary
	pm.StockConfirmationIsLotUnit = stockConfirmationIsLotUnit

	data := pm
	res := &StockConfirmation{
		BusinessPartner:                 data.BusinessPartner,
		Product:                         data.Product,
		Plant:                           data.Plant,
		StorageLocation:                 data.StorageLocation,
		StorageBin:                      data.StorageBin,
		Batch:                           data.Batch,
		RequestedQuantity:               data.RequestedQuantity,
		ProductStockAvailabilityDate:    data.ProductStockAvailabilityDate,
		OrderID:                         data.OrderID,
		OrderItem:                       data.OrderItem,
		InventoryStockType:              data.InventoryStockType,
		InventorySpecialStockType:       data.InventorySpecialStockType,
		AvailableProductStock:           data.AvailableProductStock,
		CheckedQuantity:                 data.CheckedQuantity,
		CheckedDate:                     data.CheckedDate,
		OpenConfirmedQuantityInBaseUnit: data.OpenConfirmedQuantityInBaseUnit,
		StockIsFullyChecked:             data.StockIsFullyChecked,
		Suffix:                          data.Suffix,
		StockConfirmationIsLotUnit:      data.StockConfirmationIsLotUnit,
		StockConfirmationIsOrdinary:     data.StockConfirmationIsOrdinary,
	}

	return res, nil
}

func (psdc *SDC) ConvertToLotUnitStockConfirmationOrdersItemScheduleLine(scheduleLine int, stockConfirmationPlantTimeZone *string, internalCapacityQuantity float32, item api_input_reader.Item, stockConfPlantRelationProduct *StockConfPlantRelationProduct, lotUnitStockConfirmation *StockConfirmation) (*OrdersItemScheduleLine, error) {
	pm := &requests.OrdersItemScheduleLine{}

	if item.RequestedDeliveryDate == nil {
		return nil, xerrors.Errorf("入力ファイルの'RequestedDeliveryDate'がありません。")
	}
	if item.RequestedDeliveryTime == nil {
		return nil, xerrors.Errorf("入力ファイルの'RequestedDeliveryTime'がありません。")
	}
	if item.OrderQuantityInBaseUnit == nil {
		return nil, xerrors.Errorf("入力ファイルの'OrderQuantityInBaseUnit'がありません。")
	}
	if item.ItemScheduleLine[0].ScheduleLineOrderQuantity == nil {
		return nil, xerrors.Errorf("入力ファイルの'ScheduleLineOrderQuantity'がありません。")
	}

	pm.OrderID = lotUnitStockConfirmation.OrderID
	pm.OrderItem = lotUnitStockConfirmation.OrderItem
	pm.ScheduleLine = scheduleLine
	pm.SupplyChainRelationshipID = stockConfPlantRelationProduct.SupplyChainRelationshipID
	pm.SupplyChainRelationshipStockConfPlantID = stockConfPlantRelationProduct.SupplyChainRelationshipStockConfPlantID
	pm.Product = lotUnitStockConfirmation.Product
	pm.StockConfirmationBussinessPartner = lotUnitStockConfirmation.BusinessPartner
	pm.StockConfirmationPlant = lotUnitStockConfirmation.Plant
	pm.StockConfirmationPlantTimeZone = stockConfirmationPlantTimeZone
	pm.StockConfirmationPlantBatch = &lotUnitStockConfirmation.Batch
	pm.StockConfirmationPlantBatchValidityStartDate = item.StockConfirmationPlantBatchValidityStartDate
	pm.StockConfirmationPlantBatchValidityEndDate = item.StockConfirmationPlantBatchValidityEndDate
	pm.RequestedDeliveryDate = *item.RequestedDeliveryDate
	pm.RequestedDeliveryTime = *item.RequestedDeliveryTime
	pm.ConfirmedDeliveryDate = lotUnitStockConfirmation.ProductStockAvailabilityDate
	pm.ScheduleLineOrderQuantity = *item.ItemScheduleLine[0].ScheduleLineOrderQuantity
	pm.OriginalOrderQuantityInBaseUnit = *item.OrderQuantityInBaseUnit
	pm.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit = lotUnitStockConfirmation.CheckedQuantity * internalCapacityQuantity
	pm.ConfirmedOrderQuantityByPDTAvailCheck = lotUnitStockConfirmation.CheckedQuantity
	pm.DeliveredQuantityInBaseUnit = nil
	pm.OpenConfirmedQuantityInBaseUnit = &lotUnitStockConfirmation.OpenConfirmedQuantityInBaseUnit
	pm.StockIsFullyConfirmed = &lotUnitStockConfirmation.StockIsFullyChecked
	pm.PlusMinusFlag = "-"
	pm.ItemScheduleLineDeliveryBlockStatus = getBoolPtr(false)
	pm.IsCancelled = getBoolPtr(false)
	pm.IsMarkedForDeletion = getBoolPtr(false)

	data := pm
	res := &OrdersItemScheduleLine{
		OrderID:                                         data.OrderID,
		OrderItem:                                       data.OrderItem,
		ScheduleLine:                                    data.ScheduleLine,
		SupplyChainRelationshipID:                       data.SupplyChainRelationshipID,
		SupplyChainRelationshipStockConfPlantID:         data.SupplyChainRelationshipStockConfPlantID,
		Product:                                         data.Product,
		StockConfirmationBussinessPartner:               data.StockConfirmationBussinessPartner,
		StockConfirmationPlant:                          data.StockConfirmationPlant,
		StockConfirmationPlantTimeZone:                  data.StockConfirmationPlantTimeZone,
		StockConfirmationPlantBatch:                     data.StockConfirmationPlantBatch,
		StockConfirmationPlantBatchValidityStartDate:    data.StockConfirmationPlantBatchValidityStartDate,
		StockConfirmationPlantBatchValidityEndDate:      data.StockConfirmationPlantBatchValidityEndDate,
		RequestedDeliveryDate:                           data.RequestedDeliveryDate,
		RequestedDeliveryTime:                           data.RequestedDeliveryTime,
		ConfirmedDeliveryDate:                           data.ConfirmedDeliveryDate,
		ScheduleLineOrderQuantity:                       data.ScheduleLineOrderQuantity,
		OriginalOrderQuantityInBaseUnit:                 data.OriginalOrderQuantityInBaseUnit,
		ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit: data.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit,
		ConfirmedOrderQuantityByPDTAvailCheck:           data.ConfirmedOrderQuantityByPDTAvailCheck,
		DeliveredQuantityInBaseUnit:                     data.DeliveredQuantityInBaseUnit,
		OpenConfirmedQuantityInBaseUnit:                 data.OpenConfirmedQuantityInBaseUnit,
		StockIsFullyConfirmed:                           data.StockIsFullyConfirmed,
		PlusMinusFlag:                                   data.PlusMinusFlag,
		ItemScheduleLineDeliveryBlockStatus:             data.ItemScheduleLineDeliveryBlockStatus,
		IsCancelled:                                     data.IsCancelled,
		IsMarkedForDeletion:                             data.IsMarkedForDeletion,
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrdinaryStockConfirmationOrdersItemScheduleLine(scheduleLine int, stockConfirmationPlantTimeZone *string, internalCapacityQuantity float32, item api_input_reader.Item, stockConfPlantRelationProduct *StockConfPlantRelationProduct, ordinaryStockConfirmation *StockConfirmation) (*OrdersItemScheduleLine, error) {
	pm := &requests.OrdersItemScheduleLine{}

	if item.RequestedDeliveryDate == nil {
		return nil, xerrors.Errorf("入力ファイルの'RequestedDeliveryDate'がありません。")
	}
	if item.RequestedDeliveryTime == nil {
		return nil, xerrors.Errorf("入力ファイルの'RequestedDeliveryTime'がありません。")
	}
	if item.OrderQuantityInBaseUnit == nil {
		return nil, xerrors.Errorf("入力ファイルの'OrderQuantityInBaseUnit'がありません。")
	}
	if item.ItemScheduleLine[0].ScheduleLineOrderQuantity == nil {
		return nil, xerrors.Errorf("入力ファイルの'ScheduleLineOrderQuantity'がありません。")
	}

	pm.OrderID = ordinaryStockConfirmation.OrderID
	pm.OrderItem = ordinaryStockConfirmation.OrderItem
	pm.ScheduleLine = scheduleLine
	pm.SupplyChainRelationshipID = stockConfPlantRelationProduct.SupplyChainRelationshipID
	pm.SupplyChainRelationshipStockConfPlantID = stockConfPlantRelationProduct.SupplyChainRelationshipStockConfPlantID
	pm.Product = ordinaryStockConfirmation.Product
	pm.StockConfirmationBussinessPartner = ordinaryStockConfirmation.BusinessPartner
	pm.StockConfirmationPlant = ordinaryStockConfirmation.Plant
	pm.StockConfirmationPlantTimeZone = stockConfirmationPlantTimeZone
	pm.StockConfirmationPlantBatch = getStringPtr("")
	pm.StockConfirmationPlantBatchValidityStartDate = nil
	pm.StockConfirmationPlantBatchValidityEndDate = nil
	pm.RequestedDeliveryDate = *item.RequestedDeliveryDate
	pm.RequestedDeliveryTime = *item.RequestedDeliveryTime
	pm.ConfirmedDeliveryDate = ordinaryStockConfirmation.ProductStockAvailabilityDate
	pm.ScheduleLineOrderQuantity = *item.ItemScheduleLine[0].ScheduleLineOrderQuantity
	pm.OriginalOrderQuantityInBaseUnit = *item.OrderQuantityInBaseUnit
	pm.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit = ordinaryStockConfirmation.CheckedQuantity * internalCapacityQuantity
	pm.ConfirmedOrderQuantityByPDTAvailCheck = ordinaryStockConfirmation.CheckedQuantity
	pm.DeliveredQuantityInBaseUnit = nil
	pm.OpenConfirmedQuantityInBaseUnit = &ordinaryStockConfirmation.OpenConfirmedQuantityInBaseUnit
	pm.StockIsFullyConfirmed = &ordinaryStockConfirmation.StockIsFullyChecked
	pm.PlusMinusFlag = "-"
	pm.ItemScheduleLineDeliveryBlockStatus = getBoolPtr(false)
	pm.IsCancelled = getBoolPtr(false)
	pm.IsMarkedForDeletion = getBoolPtr(false)

	data := pm
	res := &OrdersItemScheduleLine{
		OrderID:                                         data.OrderID,
		OrderItem:                                       data.OrderItem,
		ScheduleLine:                                    data.ScheduleLine,
		SupplyChainRelationshipID:                       data.SupplyChainRelationshipID,
		SupplyChainRelationshipStockConfPlantID:         data.SupplyChainRelationshipStockConfPlantID,
		Product:                                         data.Product,
		StockConfirmationBussinessPartner:               data.StockConfirmationBussinessPartner,
		StockConfirmationPlant:                          data.StockConfirmationPlant,
		StockConfirmationPlantTimeZone:                  data.StockConfirmationPlantTimeZone,
		StockConfirmationPlantBatch:                     data.StockConfirmationPlantBatch,
		StockConfirmationPlantBatchValidityStartDate:    data.StockConfirmationPlantBatchValidityStartDate,
		StockConfirmationPlantBatchValidityEndDate:      data.StockConfirmationPlantBatchValidityEndDate,
		RequestedDeliveryDate:                           data.RequestedDeliveryDate,
		RequestedDeliveryTime:                           data.RequestedDeliveryTime,
		ConfirmedDeliveryDate:                           data.ConfirmedDeliveryDate,
		ScheduleLineOrderQuantity:                       data.ScheduleLineOrderQuantity,
		OriginalOrderQuantityInBaseUnit:                 data.OriginalOrderQuantityInBaseUnit,
		ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit: data.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit,
		ConfirmedOrderQuantityByPDTAvailCheck:           data.ConfirmedOrderQuantityByPDTAvailCheck,
		DeliveredQuantityInBaseUnit:                     data.DeliveredQuantityInBaseUnit,
		OpenConfirmedQuantityInBaseUnit:                 data.OpenConfirmedQuantityInBaseUnit,
		StockIsFullyConfirmed:                           data.StockIsFullyConfirmed,
		PlusMinusFlag:                                   data.PlusMinusFlag,
		ItemScheduleLineDeliveryBlockStatus:             data.ItemScheduleLineDeliveryBlockStatus,
		IsCancelled:                                     data.IsCancelled,
		IsMarkedForDeletion:                             data.IsMarkedForDeletion,
	}

	return res, nil
}

func (psdc *SDC) ConvertToStockConfirmationStatus(ordersItemScheduleLine *OrdersItemScheduleLine, stockConfirmationStatus *string) *StockConfirmationStatus {
	pm := &requests.StockConfirmationStatus{}

	pm.OrderItem = ordersItemScheduleLine.OrderItem
	pm.StockIsFullyConfirmed = ordersItemScheduleLine.StockIsFullyConfirmed
	pm.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit = ordersItemScheduleLine.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit
	pm.StockConfirmationStatus = stockConfirmationStatus

	data := pm
	res := StockConfirmationStatus{
		OrderItem:             data.OrderItem,
		StockIsFullyConfirmed: data.StockIsFullyConfirmed,
		ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit: data.ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit,
		StockConfirmationStatus:                         data.StockConfirmationStatus,
	}

	return &res
}

func (psdc *SDC) ConvertToConfirmedOrderQuantityInBaseUnit(orderItem int, confirmedOrderQuantityInBaseUnit float32) *ConfirmedOrderQuantityInBaseUnit {
	pm := &requests.ConfirmedOrderQuantityInBaseUnit{}

	pm.OrderItem = orderItem
	pm.ConfirmedOrderQuantityInBaseUnit = confirmedOrderQuantityInBaseUnit

	data := pm
	res := ConfirmedOrderQuantityInBaseUnit{
		OrderItem:                        data.OrderItem,
		ConfirmedOrderQuantityInBaseUnit: data.ConfirmedOrderQuantityInBaseUnit,
	}

	return &res
}

func (psdc *SDC) ConvertToItemInvoiceDocumentDate(invoiceDocumentDate string) *ItemInvoiceDocumentDate {
	pm := &requests.ItemInvoiceDocumentDate{}

	pm.InvoiceDocumentDate = invoiceDocumentDate

	data := pm
	res := ItemInvoiceDocumentDate{
		InvoiceDocumentDate: data.InvoiceDocumentDate,
	}

	return &res
}

func (psdc *SDC) ConvertToItemPriceDetnExchangeRate(invoiceDocumentDate *float32) *PriceDetnExchangeRate {
	pm := &requests.PriceDetnExchangeRate{}

	pm.PriceDetnExchangeRate = invoiceDocumentDate

	data := pm
	res := PriceDetnExchangeRate{
		PriceDetnExchangeRate: data.PriceDetnExchangeRate,
	}

	return &res
}

func (psdc *SDC) ConvertToItemAccountingExchangeRate(accountingExchangeRate *float32) *AccountingExchangeRate {
	pm := &requests.AccountingExchangeRate{}

	pm.AccountingExchangeRate = accountingExchangeRate

	data := pm
	res := AccountingExchangeRate{
		AccountingExchangeRate: data.AccountingExchangeRate,
	}

	return &res
}

func (psdc *SDC) ConvertToItemReferenceDocument(referenceDocument, referenceDocumentItem *int) *ItemReferenceDocument {
	pm := &requests.ItemReferenceDocument{}

	pm.ReferenceDocument = referenceDocument
	pm.ReferenceDocumentItem = referenceDocumentItem

	data := pm
	res := ItemReferenceDocument{
		ReferenceDocument:     data.ReferenceDocument,
		ReferenceDocumentItem: data.ReferenceDocumentItem,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderItemTextByBuyerSellerKey() *OrderItemTextByBuyerSellerKey {
	pm := &requests.OrderItemTextByBuyerSellerKey{}

	data := pm
	res := OrderItemTextByBuyerSellerKey{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Language:        data.Language,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderItemTextByBuyerSeller(rows *sql.Rows) ([]*OrderItemTextByBuyerSeller, error) {
	defer rows.Close()
	res := make([]*OrderItemTextByBuyerSeller, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrderItemTextByBuyerSeller{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Language,
			&pm.ProductDescription,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrderItemTextByBuyerSeller{
			Product:            data.Product,
			BusinessPartner:    data.BusinessPartner,
			Language:           data.Language,
			ProductDescription: data.ProductDescription,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_master_product_desc_by_bp_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToProductMasterQualityKey() *ProductMasterQualityKey {
	pm := &requests.ProductMasterQualityKey{
		IsMarkedForDeletion: false,
	}

	data := pm
	res := ProductMasterQualityKey{
		Product:             data.Product,
		BusinessPartner:     data.BusinessPartner,
		Plant:               data.Plant,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToProductMasterQuality(rows *sql.Rows) ([]*ProductMasterQuality, error) {
	defer rows.Close()
	res := make([]*ProductMasterQuality, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductMasterQuality{}

		err := rows.Scan(
			&pm.Product,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &ProductMasterQuality{
			Product:             data.Product,
			BusinessPartner:     data.BusinessPartner,
			Plant:               data.Plant,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_product_master_quality_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToInspectionPlanKey() *InspectionPlanKey {
	pm := &requests.InspectionPlanKey{}

	data := pm
	res := InspectionPlanKey{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
	}

	return &res
}

func (psdc *SDC) ConvertToInspectionPlan(rows *sql.Rows) ([]*InspectionPlan, error) {
	defer rows.Close()
	res := make([]*InspectionPlan, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.InspectionPlan{}

		err := rows.Scan(
			&pm.InspectionPlantBusinessPartner,
			&pm.InspectionPlan,
			&pm.InspectionPlant,
			&pm.Product,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &InspectionPlan{
			InspectionPlantBusinessPartner: data.InspectionPlantBusinessPartner,
			InspectionPlan:                 data.InspectionPlan,
			InspectionPlant:                data.InspectionPlant,
			Product:                        data.Product,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_inspection_plan_header_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToInspectionOrderKey() *InspectionOrderKey {
	pm := &requests.InspectionOrderKey{}

	data := pm
	res := InspectionOrderKey{
		Product:         data.Product,
		BusinessPartner: data.BusinessPartner,
		Plant:           data.Plant,
	}

	return &res
}

func (psdc *SDC) ConvertToInspectionOrder(rows *sql.Rows) ([]*InspectionOrder, error) {
	defer rows.Close()
	res := make([]*InspectionOrder, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.InspectionOrder{}

		err := rows.Scan(
			&pm.InspectionOrder,
			&pm.Product,
			&pm.InspectionPlantBusinessPartner,
			&pm.InspectionPlant,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &InspectionOrder{
			InspectionOrder:                data.InspectionOrder,
			Product:                        data.Product,
			InspectionPlantBusinessPartner: data.InspectionPlantBusinessPartner,
			InspectionPlant:                data.InspectionPlant,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_inspection_order_header_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

// Item Pricing Element
func (psdc *SDC) ConvertToPriceMasterKey() *PriceMasterKey {
	pm := &requests.PriceMasterKey{}

	data := pm
	res := PriceMasterKey{
		Product:                    data.Product,
		SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
		Buyer:                      data.Buyer,
		Seller:                     data.Seller,
		ConditionValidityEndDate:   data.ConditionValidityEndDate,
		ConditionValidityStartDate: data.ConditionValidityStartDate,
	}

	return &res
}

func (psdc *SDC) ConvertToPriceMaster(rows *sql.Rows) ([]*PriceMaster, error) {
	defer rows.Close()
	res := make([]*PriceMaster, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PriceMaster{}

		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionValidityStartDate,
			&pm.ConditionValidityEndDate,
			&pm.Product,
			&pm.ConditionType,
			&pm.ConditionRateValue,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &PriceMaster{
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			Buyer:                      data.Buyer,
			Seller:                     data.Seller,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionValidityStartDate: data.ConditionValidityStartDate,
			ConditionValidityEndDate:   data.ConditionValidityEndDate,
			Product:                    data.Product,
			ConditionType:              data.ConditionType,
			ConditionRateValue:         data.ConditionRateValue,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_price_master_price_master_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToConditionAmount(orderItem int, product string, conditionQuantity *float32, conditionAmount *float32) *ConditionAmount {
	pm := &requests.ConditionAmount{
		ConditionIsManuallyChanged: getBoolPtr(false),
	}

	pm.OrderItem = orderItem
	pm.Product = product
	pm.ConditionQuantity = conditionQuantity
	pm.ConditionAmount = conditionAmount

	data := pm
	res := ConditionAmount{
		OrderItem:                  data.OrderItem,
		Product:                    data.Product,
		ConditionQuantity:          data.ConditionQuantity,
		ConditionAmount:            data.ConditionAmount,
		ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
	}

	return &res
}

func (psdc *SDC) ConvertToConditionRateValue(product string, supplyChainRelationshipID int, taxCode string, priceMasterConditionRateValue, taxRate, conditionRateValue, conditionQuantity, conditionAmount *float32) *ConditionRateValue {
	pm := &requests.ConditionRateValue{
		ConditionIsManuallyChanged: getBoolPtr(false),
	}

	pm.Product = product
	pm.SupplyChainRelationshipID = supplyChainRelationshipID
	pm.TaxCode = taxCode
	pm.PriceMasterConditionRateValue = priceMasterConditionRateValue
	pm.TaxRate = taxRate
	pm.ConditionRateValue = conditionRateValue
	pm.ConditionQuantity = conditionQuantity
	pm.ConditionAmount = conditionAmount

	data := pm
	res := ConditionRateValue{
		Product:                       data.Product,
		SupplyChainRelationshipID:     data.SupplyChainRelationshipID,
		TaxCode:                       data.TaxCode,
		PriceMasterConditionRateValue: data.PriceMasterConditionRateValue,
		TaxRate:                       data.TaxRate,
		ConditionRateValue:            data.ConditionRateValue,
		ConditionQuantity:             data.ConditionQuantity,
		ConditionAmount:               data.ConditionAmount,
		ConditionIsManuallyChanged:    data.ConditionIsManuallyChanged,
	}

	return &res
}

func (psdc *SDC) ConvertToConditionIsManuallyChanged(product string) *ConditionIsManuallyChanged {
	pm := &requests.ConditionIsManuallyChanged{}

	pm.Product = product
	pm.ConditionIsManuallyChanged = getBoolPtr(true)

	data := pm
	res := ConditionIsManuallyChanged{
		Product:                    data.Product,
		ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
	}

	return &res
}

func (psdc *SDC) ConvertToPricingProcedureCounter(product string, supplyChainRelationshipID, buyer, seller, length int) *PricingProcedureCounter {
	pm := &requests.PricingProcedureCounter{}

	counter := make([]int, length)
	for i := 0; i < length; i++ {
		counter[i] = i + 1
	}

	pm.Product = product
	pm.SupplyChainRelationshipID = supplyChainRelationshipID
	pm.Buyer = buyer
	pm.Seller = seller
	pm.PricingProcedureCounter = counter

	data := pm
	res := &PricingProcedureCounter{
		Product:                   data.Product,
		SupplyChainRelationshipID: data.SupplyChainRelationshipID,
		Buyer:                     data.Buyer,
		Seller:                    data.Seller,
		PricingProcedureCounter:   data.PricingProcedureCounter,
	}

	return res
}

// Amount関連の計算
func (psdc *SDC) ConvertToNetAmount(conditionAmount []*ConditionAmount) []*NetAmount {
	res := make([]*NetAmount, 0)

	for _, v := range conditionAmount {
		pm := &requests.NetAmount{}

		pm.OrderItem = v.OrderItem
		pm.Product = v.Product
		pm.NetAmount = v.ConditionAmount

		data := pm
		res = append(res, &NetAmount{
			OrderItem: data.OrderItem,
			Product:   data.Product,
			NetAmount: data.NetAmount,
		})
	}

	return res
}

func (psdc *SDC) ConvertToTaxAmount(orderItem int, product string, taxCode string, taxRate, netAmount, taxAmount *float32) *TaxAmount {
	pm := &requests.TaxAmount{}

	pm.OrderItem = orderItem
	pm.Product = product
	pm.TaxCode = taxCode
	pm.TaxRate = taxRate
	pm.NetAmount = netAmount
	pm.TaxAmount = taxAmount

	data := pm
	res := TaxAmount{
		OrderItem: data.OrderItem,
		Product:   data.Product,
		TaxCode:   data.TaxCode,
		TaxRate:   data.TaxRate,
		NetAmount: data.NetAmount,
		TaxAmount: data.TaxAmount,
	}

	return &res
}

func (psdc *SDC) ConvertToGrossAmount(orderItem int, product string, netAmount, taxAmount, grossAmount *float32) *GrossAmount {
	pm := &requests.GrossAmount{}

	pm.OrderItem = orderItem
	pm.Product = product
	pm.NetAmount = netAmount
	pm.TaxAmount = taxAmount
	pm.GrossAmount = grossAmount

	data := pm
	res := GrossAmount{
		OrderItem:   data.OrderItem,
		Product:     data.Product,
		NetAmount:   data.NetAmount,
		TaxAmount:   data.TaxAmount,
		GrossAmount: data.GrossAmount,
	}

	return &res
}

// Address
func (psdc *SDC) ConvertToAddressKey() *AddressKey {
	pm := &requests.AddressKey{}

	data := pm
	res := AddressKey{
		AddressID:       data.AddressID,
		ValidityEndDate: data.ValidityEndDate,
	}

	return &res
}

func (psdc *SDC) ConvertToAddress(rows *sql.Rows) ([]*Address, error) {
	defer rows.Close()
	res := make([]*Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Address{}

		err := rows.Scan(
			&pm.AddressID,
			&pm.ValidityEndDate,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &Address{
			AddressID:       data.AddressID,
			ValidityEndDate: data.ValidityEndDate,
			PostalCode:      data.PostalCode,
			LocalRegion:     data.LocalRegion,
			Country:         data.Country,
			District:        data.District,
			StreetName:      data.StreetName,
			CityName:        data.CityName,
			Building:        data.Building,
			Floor:           data.Floor,
			Room:            data.Room,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_address_address_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToAddressMaster(sdc *api_input_reader.SDC, idx int) *AddressMaster {
	pm := &requests.AddressMaster{
		AddressID:         sdc.Header.Address[idx].AddressID,
		ValidityStartDate: *sdc.Header.OrderValidityStartDate,
		ValidityEndDate:   *sdc.Header.OrderValidityEndDate,
		PostalCode:        *sdc.Header.Address[idx].PostalCode,
		LocalRegion:       *sdc.Header.Address[idx].LocalRegion,
		Country:           *sdc.Header.Address[idx].Country,
		District:          sdc.Header.Address[idx].District,
		StreetName:        *sdc.Header.Address[idx].StreetName,
		CityName:          *sdc.Header.Address[idx].CityName,
		Building:          sdc.Header.Address[idx].Building,
		Floor:             sdc.Header.Address[idx].Floor,
		Room:              sdc.Header.Address[idx].Room,
	}

	data := pm
	res := &AddressMaster{
		AddressID:         data.AddressID,
		ValidityEndDate:   data.ValidityEndDate,
		ValidityStartDate: data.ValidityStartDate,
		PostalCode:        data.PostalCode,
		LocalRegion:       data.LocalRegion,
		Country:           data.Country,
		District:          data.District,
		StreetName:        data.StreetName,
		CityName:          data.CityName,
		Building:          data.Building,
		Floor:             data.Floor,
		Room:              data.Room,
	}

	return res
}

func (psdc *SDC) ConvertToAddressFromInput() []*Address {
	res := make([]*Address, 0)

	for _, v := range psdc.AddressMaster {
		pm := &requests.Address{}

		pm.AddressID = v.AddressID
		pm.PostalCode = v.PostalCode
		pm.LocalRegion = v.LocalRegion
		pm.Country = v.Country
		pm.District = v.District
		pm.StreetName = v.StreetName
		pm.CityName = v.CityName
		pm.Building = v.Building
		pm.Floor = v.Floor
		pm.Room = v.Room

		data := pm
		res = append(res, &Address{
			AddressID:   data.AddressID,
			PostalCode:  data.PostalCode,
			LocalRegion: data.LocalRegion,
			Country:     data.Country,
			District:    data.District,
			StreetName:  data.StreetName,
			CityName:    data.CityName,
			Building:    data.Building,
			Floor:       data.Floor,
			Room:        data.Room,
		})

	}

	return res
}

// 数量単位変換実行の是非の判定
func (psdc *SDC) ConvertToQuantityUnitConversionKey(product, baseUnit, deliveryUnit string) *QuantityUnitConversionKey {
	pm := &requests.QuantityUnitConversionKey{}

	pm.Product = product
	pm.BaseUnit = baseUnit
	pm.DeliveryUnit = deliveryUnit

	data := pm
	res := QuantityUnitConversionKey{
		Product:      data.Product,
		BaseUnit:     data.BaseUnit,
		DeliveryUnit: data.DeliveryUnit,
	}

	return &res
}

func (psdc *SDC) ConvertToQuantityUnitConversionQueryGets(rows *sql.Rows, dataKey []*QuantityUnitConversionKey) ([]*QuantityUnitConversionQueryGets, error) {
	defer rows.Close()
	res := make([]*QuantityUnitConversionQueryGets, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.QuantityUnitConversionQueryGets{}

		err := rows.Scan(
			&pm.QuantityUnitFrom,
			&pm.QuantityUnitTo,
			&pm.ConversionCoefficient,
		)
		if err != nil {
			return nil, err
		}

		for _, v := range dataKey {
			if v.BaseUnit == pm.QuantityUnitFrom && v.DeliveryUnit == pm.QuantityUnitTo {
				pm.Product = v.Product
				continue
			}
		}

		data := pm
		res = append(res, &QuantityUnitConversionQueryGets{
			Product:               data.Product,
			QuantityUnitFrom:      data.QuantityUnitFrom,
			QuantityUnitTo:        data.QuantityUnitTo,
			ConversionCoefficient: data.ConversionCoefficient,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_quantity_unit_conversion_quantity_unit_conversion_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToQuantityUnitConversion(orderItem int, product string, conversionCoefficient, orderQuantityInDeliveryUnit float32) *QuantityUnitConversion {
	pm := &requests.QuantityUnitConversion{}

	pm.OrderItem = orderItem
	pm.Product = product
	pm.ConversionCoefficient = conversionCoefficient
	pm.OrderQuantityInDeliveryUnit = orderQuantityInDeliveryUnit

	data := pm
	res := QuantityUnitConversion{
		OrderItem:                   data.OrderItem,
		Product:                     data.Product,
		ConversionCoefficient:       data.ConversionCoefficient,
		OrderQuantityInDeliveryUnit: data.OrderQuantityInDeliveryUnit,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderQuantityInDeliveryUnit(orderItem int, orderQuantityInDeliveryUnit float32) *OrderQuantityInDeliveryUnit {
	pm := &requests.OrderQuantityInDeliveryUnit{}

	pm.OrderItem = orderItem
	pm.OrderQuantityInDeliveryUnit = orderQuantityInDeliveryUnit

	data := pm
	res := OrderQuantityInDeliveryUnit{
		OrderItem:                   data.OrderItem,
		OrderQuantityInDeliveryUnit: data.OrderQuantityInDeliveryUnit,
	}

	return &res
}

// Partner
func (psdc *SDC) ConvertToPartner(partnerFunction string, businessPartnerGeneral *BusinessPartnerGeneral) *Partner {
	pm := &requests.Partner{}

	pm.PartnerFunction = partnerFunction
	pm.BusinessPartner = businessPartnerGeneral.BusinessPartner
	pm.BusinessPartnerFullName = businessPartnerGeneral.BusinessPartnerFullName
	pm.BusinessPartnerName = &businessPartnerGeneral.BusinessPartnerName
	pm.Country = &businessPartnerGeneral.Country
	pm.Language = &businessPartnerGeneral.Language
	pm.Currency = &businessPartnerGeneral.Currency
	pm.AddressID = businessPartnerGeneral.AddressID

	data := pm
	res := Partner{
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		AddressID:               data.AddressID,
	}

	return &res
}

// 日付等の処理
func (psdc *SDC) ConvertToCreationDateItem(systemDate string) *CreationDateItem {
	pm := &requests.CreationDateItem{}

	pm.CreationDate = systemDate

	data := pm
	res := CreationDateItem{
		CreationDate: data.CreationDate,
	}

	return &res
}

func (psdc *SDC) ConvertToCreationTimeItem(systemTime string) *CreationTimeItem {
	pm := &requests.CreationTimeItem{}

	pm.CreationTime = systemTime

	data := pm
	res := CreationTimeItem{
		CreationTime: data.CreationTime,
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

func (psdc *SDC) ConvertToLastChangeTimeItem(systemTime string) *LastChangeTimeItem {
	pm := &requests.LastChangeTimeItem{}

	pm.LastChangeTime = systemTime

	data := pm
	res := LastChangeTimeItem{
		LastChangeTime: data.LastChangeTime,
	}

	return &res
}

func getSystemDateTime() string {
	day := time.Now()
	return day.Format("2006-01-02 15:04:05")
}

func getBoolPtr(b bool) *bool {
	return &b
}

func getStringPtr(s string) *string {
	return &s
}

func getFloat32Ptr(f float32) *float32 {
	return &f
}
