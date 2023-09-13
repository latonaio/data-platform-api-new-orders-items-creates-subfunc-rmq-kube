package subfunction

import (
	api_input_reader "data-platform-api-orders-creates-subfunc-items-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-creates-subfunc-items-rmq-kube/API_Processing_Data_Formatter"
	"math"
	"strconv"
	"strings"

	"golang.org/x/xerrors"
)

func (f *SubFunction) PriceMaster(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.PriceMaster, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToPriceMasterKey()

	for _, v := range sdc.Header.Item {
		if v.ItemPricingElement[0].ConditionAmount == nil {
			dataKey.Product = append(dataKey.Product, v.Product)
		}
	}

	if dataKey.Product == nil {
		return nil, nil
	}

	for _, v := range psdc.SupplyChainRelationshipGeneral {
		dataKey.SupplyChainRelationshipID = append(dataKey.SupplyChainRelationshipID, v.SupplyChainRelationshipID)
		dataKey.Buyer = append(dataKey.Buyer, v.Buyer)
		dataKey.Seller = append(dataKey.Seller, v.Seller)
	}

	dataKey.ConditionValidityEndDate = psdc.HeaderPricingDate.PricingDate
	dataKey.ConditionValidityStartDate = psdc.HeaderPricingDate.PricingDate

	if len(dataKey.Product) == 0 {
		return nil, xerrors.Errorf("入力ファイルの'Product'がありません。")
	}
	repeat1 := strings.Repeat("?,", len(dataKey.Product)-1) + "?"
	for _, v := range dataKey.Product {
		args = append(args, v)
	}

	if len(dataKey.SupplyChainRelationshipID) == 0 {
		return nil, xerrors.Errorf("psdc.SupplyChainRelationshipGeneralの'SupplyChainRelationshipID'がありません。")
	}
	repeat2 := strings.Repeat("(?,?,?),", len(dataKey.SupplyChainRelationshipID)-1) + "(?,?,?)"
	for i := range dataKey.SupplyChainRelationshipID {
		args = append(args, dataKey.SupplyChainRelationshipID[i], dataKey.Buyer[i], dataKey.Seller[i])
	}

	args = append(args, dataKey.ConditionValidityEndDate, dataKey.ConditionValidityStartDate)
	rows, err := f.db.Query(
		`SELECT SupplyChainRelationshipID, Buyer, Seller, ConditionRecord, ConditionSequentialNumber,
		ConditionValidityStartDate, ConditionValidityEndDate, Product, ConditionType, ConditionRateValue
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_price_master_price_master_data
		WHERE Product IN ( `+repeat1+` )
		AND (SupplyChainRelationshipID, Buyer, Seller) IN ( `+repeat2+` )
		AND ConditionValidityEndDate >= ?
		AND ConditionValidityStartDate <= ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToPriceMaster(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) ConditionAmount(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ConditionAmount, error) {
	data := make([]*api_processing_data_formatter.ConditionAmount, 0)

	priceMaster := psdc.PriceMaster
	priceMasterMap := make(map[string]*api_processing_data_formatter.PriceMaster, len(priceMaster))
	for _, v := range priceMaster {
		priceMasterMap[v.Product] = v
	}

	for _, v := range sdc.Header.Item {
		orderItem := v.OrderItem
		product := v.Product
		conditionQuantity := v.OrderQuantityInBaseUnit
		if v.ItemPricingElement[0].ConditionAmount == nil && v.Product != nil {
			conditionRateValue := priceMasterMap[*v.Product].ConditionRateValue
			conditionAmount, err := calculateConditionAmount(conditionQuantity, conditionRateValue)
			if err != nil {
				return nil, err
			}

			if product == nil {
				return nil, xerrors.Errorf("入力ファイルの'Product'がありません。")
			}
			datum := psdc.ConvertToConditionAmount(orderItem, *product, conditionQuantity, conditionAmount)
			data = append(data, datum)
		} else if v.ItemPricingElement[0].ConditionAmount != nil && v.Product != nil {
			conditionAmount := v.ItemPricingElement[0].ConditionAmount

			if product == nil {
				return nil, xerrors.Errorf("入力ファイルの'Product'がありません。")
			}
			datum := psdc.ConvertToConditionAmount(orderItem, *product, conditionQuantity, conditionAmount)
			data = append(data, datum)
		}
	}

	return data, nil
}

func calculateConditionAmount(conditionQuantity, conditionRateValue *float32) (*float32, error) {
	if conditionQuantity == nil || conditionRateValue == nil {
		return nil, xerrors.Errorf("ConditionRateValueまたはConditionQuantityがnullです。")
	}

	digit := float32DecimalDigit(*conditionRateValue)
	mul := *conditionRateValue * *conditionQuantity

	s := math.Round(float64(mul)*math.Pow10(digit)) / math.Pow10(digit)
	res := parseFloat32Ptr(float32(s))

	return res, nil
}

func (f *SubFunction) ConditionRateValue(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ConditionRateValue, error) {
	data := make([]*api_processing_data_formatter.ConditionRateValue, 0)

	taxCode := psdc.TaxCode
	taxCodeMap := make(map[string]*api_processing_data_formatter.TaxCode, len(taxCode))
	for _, v := range taxCode {
		taxCodeMap[v.Product] = v
	}

	taxRate := psdc.TaxRate
	taxRateMap := make(map[string]*api_processing_data_formatter.TaxRate, len(taxRate))
	for _, v := range taxRate {
		taxRateMap[v.TaxCode] = v
	}

	conditionAmount := psdc.ConditionAmount
	conditionAmountMap := make(map[string]*api_processing_data_formatter.ConditionAmount, len(conditionAmount))
	for _, v := range conditionAmount {
		conditionAmountMap[v.Product] = v
	}

	for _, v := range psdc.PriceMaster {
		product := v.Product
		supplyChainRelationshipID := v.SupplyChainRelationshipID
		if _, ok := taxCodeMap[product]; !ok {
			continue
		}
		if taxCodeMap[product].TaxCode == nil {
			continue
		}
		taxCode := *taxCodeMap[product].TaxCode
		if taxCode != "1" {
			continue
		}
		if _, ok := taxRateMap[taxCode]; !ok {
			continue
		}

		priceMasterConditionRateValue := v.ConditionRateValue
		taxRate := taxRateMap[taxCode].TaxRate
		conditionRateValue, err := calculateConditionRateValue(priceMasterConditionRateValue, taxRate)
		if err != nil {
			return nil, err
		}
		if _, ok := conditionAmountMap[product]; !ok {
			continue
		}
		conditionQuantity := conditionAmountMap[product].ConditionQuantity
		conditionAmount, err := calculateConditionAmount(conditionQuantity, conditionRateValue)
		if err != nil {
			return nil, err
		}

		datum := psdc.ConvertToConditionRateValue(product, supplyChainRelationshipID, taxCode, priceMasterConditionRateValue, taxRate, conditionRateValue, conditionQuantity, conditionAmount)
		data = append(data, datum)
	}

	return data, nil
}

func calculateConditionRateValue(conditionRateValue *float32, taxRate *float32) (*float32, error) {
	if conditionRateValue == nil || taxRate == nil {
		return nil, xerrors.Errorf("ConditionRateValueまたはTaxRateがnullです。")
	}

	digit := float32DecimalDigit(*conditionRateValue)
	mul := *conditionRateValue * *taxRate / 100

	s := math.Round(float64(mul)*math.Pow10(digit)) / math.Pow10(digit)
	res := parseFloat32Ptr(float32(s))

	return res, nil
}

func (f *SubFunction) ConditionIsManuallyChanged(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ConditionIsManuallyChanged, error) {
	data := make([]*api_processing_data_formatter.ConditionIsManuallyChanged, 0)

	for _, v := range sdc.Header.Item {
		if v.ItemPricingElement[0].ConditionAmount != nil {
			if v.Product == nil {
				return nil, xerrors.Errorf("入力ファイルの'Product'がありません。")
			}
			datum := psdc.ConvertToConditionIsManuallyChanged(*v.Product)
			data = append(data, datum)
		}
	}

	return data, nil
}

func (f *SubFunction) PricingProcedureCounter(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.PricingProcedureCounter {
	data := make([]*api_processing_data_formatter.PricingProcedureCounter, 0)

	searched := make([]int, 0)
	priceMaster := psdc.PriceMaster
	for i, v := range priceMaster {
		if contains(searched, i) {
			continue
		}
		product := v.Product
		supplyChainRelationshipID := v.SupplyChainRelationshipID
		buyer := v.Buyer
		seller := v.Seller
		counter := 1
		for j := i + 1; j < len(priceMaster); j++ {
			if product == priceMaster[j].Product &&
				supplyChainRelationshipID == priceMaster[j].SupplyChainRelationshipID &&
				buyer == priceMaster[j].Buyer &&
				seller == priceMaster[j].Seller {
				counter++
				searched = append(searched, j)
				continue
			}
		}

		datum := psdc.ConvertToPricingProcedureCounter(product, supplyChainRelationshipID, buyer, seller, counter)
		data = append(data, datum)
	}

	return data
}

func float32DecimalDigit(f float32) int {
	s := strconv.FormatFloat(float64(f), 'f', -1, 32)

	i := strings.Index(s, ".")
	if i == -1 {
		return 0
	}

	return len(s[i+1:])
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
