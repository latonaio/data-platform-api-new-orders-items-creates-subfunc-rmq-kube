package subfunction

import (
	"context"
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"data-platform-api-orders-items-creates-subfunc-rmq-kube/config"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type SubFunction struct {
	ctx  context.Context
	db   *database.Mysql
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
	l    *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, conf *config.Conf, rmq *rabbitmq.RabbitmqClient, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx:  ctx,
		db:   db,
		conf: conf,
		rmq:  rmq,
		l:    l,
	}
}

func (f *SubFunction) MetaData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.MetaData {
	metaData := psdc.ConvertToMetaData(sdc)

	return metaData
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}

	psdc.MetaData = f.MetaData(sdc, psdc)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		// 1-0. サプライチェーンリレーションシップマスタでの取引妥当性確認(一般データ)
		psdc.SupplyChainRelationshipGeneral, e = f.SupplyChainRelationshipGeneral(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-1. サプライチェーンリレーションシップマスタでの取引妥当性確認(入出荷関係データ)  //1-0
		psdc.SupplyChainRelationshipDeliveryRelation, e = f.SupplyChainRelationshipDeliveryRelation(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-2. サプライチェーンリレーションシップマスタ入出荷プラント関係データの取得  //1-1
		psdc.SupplyChainRelationshipDeliveryPlantRelation, e = f.SupplyChainRelationshipDeliveryPlantRelation(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-3. サプライチェーンリレーションシップマスタ取引データの取得  //1-0
		psdc.SupplyChainRelationshipTransaction, e = f.SupplyChainRelationshipTransaction(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 2-0. OrderItem
		psdc.OrderItem = f.OrderItem(sdc, psdc)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 1-7. InvoiceDocumentDate  //1-3
			psdc.HeaderInvoiceDocumentDate, e = f.HeaderInvoiceDocumentDate(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-29. InvoiceDocumentDate  //1-7
			psdc.ItemInvoiceDocumentDate = f.ItemInvoiceDocumentDate(sdc, psdc)
		}(wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 1-4. サプライチェーンリレーションシップマスタ請求関係データの取得  //1-0
			psdc.SupplyChainRelationshipBillingRelation, e = f.SupplyChainRelationshipBillingRelation(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-1 ProductTaxClassificationBillToCountry  //1-4
			psdc.ProductTaxClassificationBillToCountry, e = f.ProductTaxClassificationBillToCountry(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-2. ProductTaxClassificationBillFromCountry  //1-4
			psdc.ProductTaxClassificationBillFromCountry, e = f.ProductTaxClassificationBillFromCountry(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-3. DefinedTaxClassification  //2-1,2-2
			psdc.DefinedTaxClassification, e = f.DefinedTaxClassification(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-20. TaxCode  //1-4,2-3
			psdc.TaxCode, e = f.TaxCode(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			//2-21. TaxRateの計算  //2-20
			psdc.TaxRate, e = f.TaxRate(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 1-12. PricingDate
			psdc.HeaderPricingDate = f.HeaderPricingDate(sdc, psdc)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				// 1-5. サプライチェーンリレーションシップマスタ支払関係データの取得  //1-4
				psdc.SupplyChainRelationshipPaymentRelation, e = f.SupplyChainRelationshipPaymentRelation(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 1-15-1. ビジネスパートナ一般データの取得①(Buyer、Seller、DeliverToParty、DeliverFromParty)  //1-1
				e = f.BusinessPartnerGeneralDeliveryRelation(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 1-15-2. ビジネスパートナ一般データの取得②(BillToParty、BillFromParty)  //1-4
				e = f.BusinessPartnerGeneralBillingRelation(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 1-15-3. ビジネスパートナ一般データの取得③(Payer、Payee)  //1-5
				e = f.BusinessPartnerGeneralPaymentRelation(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				wg.Add(1)
				go func(wg *sync.WaitGroup) {
					defer wg.Done()
					// 2-38. OrderItemTextByBuyer  // 1-15-1
					psdc.OrderItemTextByBuyer, e = f.OrderItemTextByBuyer(sdc, psdc)
					if e != nil {
						err = e
						return
					}

					// 2-39. OrderItemTextBySeller  // 1-15-1
					psdc.OrderItemTextBySeller, e = f.OrderItemTextBySeller(sdc, psdc)
					if e != nil {
						err = e
						return
					}
				}(wg)

				wg.Add(1)
				go func(wg *sync.WaitGroup) {
					defer wg.Done()
					// 10-1. 住所マスタからの住所データの取得([PostalCode, LocalRegion, Country, District, StreetName, CityName, Building, Floor, Room])  //1-15
					psdc.Address, e = f.Address(sdc, psdc)
					if e != nil {
						err = e
						return
					}

					// 10-1. 10-2. AddressIDの登録(ユーザーが任意の住所を入力ファイルで指定した場合)  //1-15
					psdc.Address, e = f.AddressFromInput(sdc, psdc)
					if e != nil {
						err = e
						return
					}
				}(wg)
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				// 2-27. PricingDate  //1-12
				psdc.ItemPricingDate = f.ItemPricingDate(sdc, psdc)
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				// 8-1. 価格マスタデータの取得(入力ファイルの[ConditionAmount]がnullである場合)  //1-0,1-12
				psdc.PriceMaster, e = f.PriceMaster(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 8-2. 価格の計算(入力ファイルの[ConditionAmount]がnullである場合)  //8-1
				psdc.ConditionAmount, e = f.ConditionAmount(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 8-3.”MWST”の価格明細のConditionRateValue(Tax)の計算  //2-20,2-21,8-1,8-2
				psdc.ConditionRateValue, e = f.ConditionRateValue(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 8-4. 入力ファイルの価格の保持(入力ファイルの[ConditionAmount]がnullでない場合)
				psdc.ConditionIsManuallyChanged, e = f.ConditionIsManuallyChanged(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 8-5. PricingProcedureCounter
				psdc.PricingProcedureCounter = f.PricingProcedureCounter(sdc, psdc)

				// 9-1. NetAmount  //8-2
				psdc.NetAmount = f.NetAmount(sdc, psdc)

				// 9-2. TaxAmount  //2-20,2-21,9-1
				psdc.TaxAmount, e = f.TaxAmount(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 9-3. GrossAmount  // 9-1,9-2
				psdc.GrossAmount, e = f.GrossAmount(sdc, psdc)
				if e != nil {
					err = e
					return
				}
			}(wg)
		}(wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 2-12. DeliverToPlantTimeZone  //1-2
			psdc.DeliverToPlantTimeZone, e = f.DeliverToPlantTimeZone(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-13. DeliverFromPlantTimeZone  //1-2
			psdc.DeliverFromPlantTimeZone, e = f.DeliverFromPlantTimeZone(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-15. Incoterms  //1-2
			psdc.Incoterms = f.Incoterms(sdc, psdc)

			// 2-16. PaymentTerms  //1-2
			psdc.ItemPaymentTerms = f.ItemPaymentTerms(sdc, psdc)

			// 2-17. PaymentMethod  //1-2
			psdc.PaymentMethod = f.PaymentMethod(sdc, psdc)
		}(wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 2-4. 品目マスタ一般データの取得
			psdc.ProductMasterGeneral, e = f.ProductMasterGeneral(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-5. OrderItemText  //2-4
			psdc.OrderItemText, e = f.OrderItemText(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-6. ItemCategoryの値のチェックと在庫確認明細への値のセット  //2-4
			psdc.ItemCategoryIsINVP = f.ItemCategoryIsINVP(sdc, psdc)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				// 2-7-1. StockConfirmationBusinessPartner / StockConfirmationPlant  //1-1,2-4,2-6
				psdc.StockConfPlantRelationProduct, e = f.StockConfPlantRelationProduct(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-7-2. ロット管理品目の確認(品目マスタBPプラントデータの取得 - 在庫確認プラント)  //2-7-1
				psdc.StockConfPlantProductMasterBPPlant, e = f.StockConfPlantProductMasterBPPlant(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-7-3. ビジネスパートナ一般データの取得(StockConfirmationBusinessPartner)  //2-7-1
				psdc.StockConfPlantBPGeneral, e = f.StockConfPlantBPGeneral(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-14. StockConfirmationPlantTimeZone  //2-7
				psdc.StockConfirmationPlantTimeZone, e = f.StockConfirmationPlantTimeZone(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-26-3-1,2-26-2-3-2. 在庫確認②(通常の在庫確認)  //2-7
				psdc.OrdinaryStockConfirmation, e = f.OrdinaryStockConfirmation(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-26-3-3. 納入日程行(Orders Item Schedule Line)の生成(通常の在庫確認)  //1-6,2-0,2-26-3-2
				psdc.OrdinaryStockConfirmationOrdersItemScheduleLine, e = f.OrdinaryStockConfirmationOrdersItemScheduleLine(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-28. ConfirmedOrderQuantityInBaseUnit  //2-26-2-3,2-26-3-3
				psdc.ConfirmedOrderQuantityInBaseUnit = f.ConfirmedOrderQuantityInBaseUnit(sdc, psdc)
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				// 2-8-1. ProductionPlantBusinessPartner / ProductionPlant  //1-1,2-4,2-6
				psdc.ProductionPlantRelationProduct, e = f.ProductionPlantRelationProduct(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-8-2. ロット管理品目の確認(品目マスタBPプラントデータの取得 - 製造プラント)  //2-8-1
				psdc.ProductionPlantProductMasterBPPlant, e = f.ProductionPlantProductMasterBPPlant(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-8-3. ビジネスパートナ一般データの取得(ProductionPlantBusinessPartner)
				psdc.ProductionPlantBPGeneral, e = f.ProductionPlantBPGeneral(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-9-1. サプライチェーンリレーションシップマスタ入出荷関係品目データの取得  //1-2,2-6,2-8
				psdc.SupplyChainRelationshipDeliveryPlantRelationProduct, e = f.SupplyChainRelationshipDeliveryPlantRelationProduct(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-9-2. ロット管理品目の確認(品目マスタBPプラントデータの取得 - 入出荷元プラント)  //2-9-1
				psdc.SupplyChainRelationshipProductMasterBPPlantDeliverTo, e = f.SupplyChainRelationshipProductMasterBPPlantDeliverTo(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-9-3. ロット管理品目の確認(品目マスタBPプラントデータの取得 - 入出荷先プラント)  //2-9-1
				psdc.SupplyChainRelationshipProductMasterBPPlantDeliverFrom, e = f.SupplyChainRelationshipProductMasterBPPlantDeliverFrom(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-11. ProductionPlantTimeZone  //2-8
				psdc.ProductionPlantTimeZone, e = f.ProductionPlantTimeZone(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 14-1. 数量単位変換実行の是非の判定  //2-4,2-9
				psdc.QuantityUnitConversion, e = f.QuantityUnitConversion(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 14-2. OrderQuantityInDeliveryUnitへの値のセット  //14-1
				psdc.OrderQuantityInDeliveryUnit, e = f.OrderQuantityInDeliveryUnit(sdc, psdc)
				if e != nil {
					err = e
					return
				}
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				// 2-18. ItemGrossWeight  //2-4
				psdc.ItemGrossWeight, e = f.ItemGrossWeight(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 2-19. ItemNetWeight  //2-4
				psdc.ItemNetWeight, e = f.ItemNetWeight(sdc, psdc)
				if e != nil {
					err = e
					return
				}
			}(wg)
		}(wg)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-13. PriceDetnExchangeRate
		psdc.HeaderPriceDetnExchangeRate = f.HeaderPriceDetnExchangeRate(sdc, psdc)

		// 1-14. AccountingExchangeRate
		psdc.HeaderAccountingExchangeRate = f.HeaderAccountingExchangeRate(sdc, psdc)

		// 2-30. PriceDetnExchangeRate  //1-13
		psdc.ItemPriceDetnExchangeRate = f.ItemPriceDetnExchangeRate(sdc, psdc)

		// 2-31. AccountingExchangeRate  //1-14
		psdc.ItemAccountingExchangeRate = f.ItemAccountingExchangeRate(sdc, psdc)

		// 2-32. ReferenceDocument / ReferenceDocumentItem
		psdc.ItemReferenceDocument = f.ItemReferenceDocument(sdc, psdc)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 99-1-2. CreationDate(Item)、CreationTime(Item)
		psdc.CreationDateItem = f.CreationDateItem(sdc, psdc)
		psdc.CreationTimeItem = f.CreationTimeItem(sdc, psdc)

		// 99-2-2. LastChangeDate(Item)、LastChangedTime(Item)
		psdc.LastChangeDateItem = f.LastChangeDateItem(sdc, psdc)
		psdc.LastChangeTimeItem = f.LastChangeTimeItem(sdc, psdc)
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	// 20. Orders Partner  //1-15-1,1-15-2,1-15-3,2-7-3,2-8-3
	psdc.Partner = f.Partner(sdc, psdc)

	f.l.Info(psdc)
	err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
