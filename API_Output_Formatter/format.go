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
) (*[]Item, error) {
	var err error

	productMasterGeneralMap := StructArrayToMap(psdc.ProductMasterGeneral, "Product")
	orderItemTextMap := StructArrayToMap(psdc.OrderItemText, "Product")

	res := make([]Item, 0, len(sdc.Header.Item))
	for i, v := range sdc.Header.Item {
		item := Item{}
		inputItem := sdc.Header.Item[i]

		// 入力ファイル
		item, err = jsonTypeConversion[Item](inputItem)
		if err != nil {
			return nil, err
		}

		item.OrderID = psdc.CalculateOrderID.OrderID
		item.OrderItem = psdc.OrderItem[i].OrderItemNumber
		item.OrderItemCategory = *productMasterGeneralMap[*v.Product].ItemCategory
		item.SupplyChainRelationshipID = psdc.SupplyChainRelationshipGeneral[0].SupplyChainRelationshipID
		item.SupplyChainRelationshipDeliveryID = &psdc.SupplyChainRelationshipDeliveryRelation[0].SupplyChainRelationshipDeliveryID
		// item.SupplyChainRelationshipStockConfPlantID =
		// item.SupplyChainRelationshipProductionPlantID =
		item.OrderItemText = *orderItemTextMap[*v.Product].OrderItemText

		item.Product = productMasterGeneralMap[*v.Product].Product
		item.ProductStandardID = *productMasterGeneralMap[*v.Product].ProductStandardID
		item.ProductGroup = productMasterGeneralMap[*v.Product].ProductGroup
		item.BaseUnit = *productMasterGeneralMap[*v.Product].BaseUnit

		item.DeliverToParty = &psdc.SupplyChainRelationshipDeliveryRelation[0].DeliverToParty
		item.DeliverFromParty = &psdc.SupplyChainRelationshipDeliveryRelation[0].DeliverFromParty
		item.CreationDate = psdc.CreationDateItem.CreationDate
		item.LastChangeDate = psdc.LastChangeDateItem.LastChangeDate

		res = append(res, item)
	}

	return &res, nil
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
