package subfunction

import (
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"strings"

	"golang.org/x/xerrors"
)

func (f *SubFunction) Address(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.Address, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToAddressKey()

	for _, v := range psdc.BusinessPartnerGeneralBuyer {
		dataKey.AddressID = append(dataKey.AddressID, v.AddressID)
	}

	for _, v := range psdc.BusinessPartnerGeneralSeller {
		dataKey.AddressID = append(dataKey.AddressID, v.AddressID)
	}

	for _, v := range psdc.BusinessPartnerGeneralDeliverToParty {
		dataKey.AddressID = append(dataKey.AddressID, v.AddressID)
	}

	for _, v := range psdc.BusinessPartnerGeneralDeliverFromParty {
		dataKey.AddressID = append(dataKey.AddressID, v.AddressID)
	}

	for _, v := range psdc.BusinessPartnerGeneralBillToParty {
		dataKey.AddressID = append(dataKey.AddressID, v.AddressID)
	}

	for _, v := range psdc.BusinessPartnerGeneralBillFromParty {
		dataKey.AddressID = append(dataKey.AddressID, v.AddressID)
	}

	for _, v := range psdc.BusinessPartnerGeneralPayer {
		dataKey.AddressID = append(dataKey.AddressID, v.AddressID)
	}

	for _, v := range psdc.BusinessPartnerGeneralPayee {
		dataKey.AddressID = append(dataKey.AddressID, v.AddressID)
	}

	dataKey.ValidityEndDate = getSystemDate()

	repeat := strings.Repeat("?,", len(dataKey.AddressID)-1) + "?"
	for _, v := range dataKey.AddressID {
		args = append(args, v)
	}
	args = append(args, dataKey.ValidityEndDate)

	rows, err := f.db.Query(
		`SELECT AddressID, ValidityEndDate, PostalCode, LocalRegion, Country, District, StreetName, CityName, Building, Floor, Room
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_address_address_data
		WHERE AddressID IN ( `+repeat+` )
		AND ValidityEndDate >= ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToAddress(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) AddressFromInput(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.Address, error) {
	processFlag := false

	calculateAddressID, err := f.CalculateAddressID(sdc, psdc)
	if err != nil {
		return nil, err
	}

	addressMasterdata := make([]*api_processing_data_formatter.AddressMaster, 0)
	addressID := calculateAddressID.AddressID
	for i, v := range sdc.Header.Address {
		if v.PostalCode != nil || v.LocalRegion != nil || v.Country != nil || v.District != nil || v.StreetName != nil || v.CityName != nil {
			if len(*v.PostalCode) != 0 || len(*v.LocalRegion) != 0 || len(*v.Country) != 0 || len(*v.District) != 0 || len(*v.StreetName) != 0 || len(*v.CityName) != 0 {
				processFlag = true
				datum := psdc.ConvertToAddressMaster(sdc, i, addressID)
				addressMasterdata = append(addressMasterdata, datum)
				addressID += 1
			}
		}
	}
	psdc.AddressMaster = addressMasterdata

	if !processFlag {
		return psdc.Address, nil
	}

	sessionID := sdc.RuntimeSessionID
	for _, addressData := range addressMasterdata {
		res, err := f.rmq.SessionKeepRequest(f.ctx, f.conf.RMQ.QueueToSQL(), map[string]interface{}{"message": addressData, "function": "Address", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			f.l.Error(err)
			return []*api_processing_data_formatter.Address{}, nil
		}
		res.Success()
	}

	data := psdc.ConvertToAddressFromInput()

	return data, err
}
func (f *SubFunction) CalculateAddressID(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateAddressID, error) {
	dataKey := psdc.ConvertToCalculateAddressIDKey()

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dataQueryGets, err := psdc.ConvertToCalculateAddressIDQueryGets(rows)
	if err != nil {
		return nil, err
	}

	if dataQueryGets.LatestNumber == nil {
		return nil, xerrors.Errorf("LatestNumberがnullです。")
	}

	addressIDLatestNumber := dataQueryGets.LatestNumber
	addressID := *dataQueryGets.LatestNumber + 1

	data := psdc.ConvertToCalculateAddressID(addressIDLatestNumber, addressID)

	return data, err
}
