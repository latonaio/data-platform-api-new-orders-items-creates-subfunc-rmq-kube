package subfunction

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
)

func (f *SubFunction) Partner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) []*api_processing_data_formatter.Partner {
	data := make([]*api_processing_data_formatter.Partner, 0)

	for _, v := range psdc.BusinessPartnerGeneralBuyer {
		datum := psdc.ConvertToPartner("BUYER", v)
		data = append(data, datum)
	}

	for _, v := range psdc.BusinessPartnerGeneralSeller {
		datum := psdc.ConvertToPartner("SELLER", v)
		data = append(data, datum)
	}

	for _, v := range psdc.BusinessPartnerGeneralDeliverToParty {
		datum := psdc.ConvertToPartner("DELIVER_TO", v)
		data = append(data, datum)
	}

	for _, v := range psdc.BusinessPartnerGeneralDeliverFromParty {
		datum := psdc.ConvertToPartner("DELIVER_FROM", v)
		data = append(data, datum)
	}

	for _, v := range psdc.BusinessPartnerGeneralBillToParty {
		datum := psdc.ConvertToPartner("BILL_TO", v)
		data = append(data, datum)
	}

	for _, v := range psdc.BusinessPartnerGeneralBillFromParty {
		datum := psdc.ConvertToPartner("BILL_FROM", v)
		data = append(data, datum)
	}

	for _, v := range psdc.BusinessPartnerGeneralPayer {
		datum := psdc.ConvertToPartner("PAYER", v)
		data = append(data, datum)
	}

	for _, v := range psdc.BusinessPartnerGeneralPayee {
		datum := psdc.ConvertToPartner("PAYEE", v)
		data = append(data, datum)
	}

	for _, v := range psdc.StockConfPlantBPGeneral {
		datum := psdc.ConvertToPartner("STOCK_CONFIRMATION", v)
		data = append(data, datum)
	}

	for _, v := range psdc.ProductionPlantBPGeneral {
		datum := psdc.ConvertToPartner("MANUFACTURER", v)
		data = append(data, datum)
	}

	return data
}
