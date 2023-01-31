package subfunction

import (
	"context"
	api_input_reader "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-orders-items-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"sync"

	database "github.com/latonaio/golang-mysql-network-connector"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type SubFunction struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx: ctx,
		db:  db,
		l:   l,
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
	wg.Add(6)

	psdc.MetaData = f.MetaData(sdc, psdc)

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
		psdc.PaymentTerms = f.PaymentTerms(sdc, psdc)

		// 2-17. PaymentMethod  //1-2
		psdc.PaymentMethod = f.PaymentMethod(sdc, psdc)

	}(&wg)

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
		psdc.SupplyChainRelationshipProductMasterBPPlant, e = f.SupplyChainRelationshipProductMasterBPPlant(sdc, psdc)
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

		// 2-14. StockConfirmationPlantTimeZone  //2-7
		psdc.StockConfirmationPlantTimeZone, e = f.StockConfirmationPlantTimeZone(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 2-18. ItemGrossWeight  //2-4
		psdc.ItemGrossWeight = f.ItemGrossWeight(sdc, psdc)

		// 2-19. ItemNetWeight  //2-4
		psdc.ItemNetWeight = f.ItemNetWeight(sdc, psdc)

	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-6. OrderID
		psdc.CalculateOrderID, e = f.CalculateOrderID(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 2-0. OrderItem
		psdc.OrderItem = f.OrderItem(sdc, psdc)
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 99-1-2. CreationDate(Item)
		psdc.CreationDateItem = f.CreationDateItem(sdc, psdc)

		// 99-2-2. LastChangeDate(Item)
		psdc.LastChangeDateItem = f.LastChangeDateItem(sdc, psdc)
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// // 2-1. ビジネスパートナマスタの取引先機能データの取得
		// psdc.HeaderPartnerFunction, e = f.HeaderPartnerFunction(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// // 2-2. ビジネスパートナの一般データの取得  // 2-1
		// psdc.HeaderPartnerBPGeneral, e = f.HeaderPartnerBPGeneral(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// // 5-1-1. BPTaxClassification  // 2-2
		// psdc.ItemBPTaxClassification, e = f.ItemBPTaxClassification(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// // 5-1-3 TaxCode  // 5-1-1, 5-1-2
		// psdc.TaxCode, e = f.TaxCode(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// // 4-1. ビジネスパートナマスタの取引先プラントデータの取得  // 2-1
		// psdc.HeaderPartnerPlant, e = f.HeaderPartnerPlant(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// // 5-4-1 StockConfirmationPlant  // 4-1
		// psdc.StockConfirmationPlant = f.StockConfirmationPlant(sdc, psdc)

		// // TODO: 仕様変更に対応する必要あり
		// // 5-5-1 ProductMasterBPPlant  // 4-1
		// psdc.ProductMasterBPPlant, e = f.ProductMasterBPPlant(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// // 5-6-1 GetPartnerPlantData  // 4-1
		// psdc.ProductionPlant = f.ProductionPlant(sdc, psdc)

		// // 1-8. PricingDate
		// psdc.PricingDate = f.PricingDate(sdc, psdc)

		// // 8-1. 価格マスタデータの取得(入力ファイルの[ConditionAmount]がnullである場合)  // 1-8
		// psdc.PriceMaster, e = f.PriceMaster(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// // 8-2. 価格の計算(入力ファイルの[ConditionAmount]がnullである場合)  // 8-1
		// psdc.ConditionAmount, e = f.ConditionAmount(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// // 5-21. NetAmount  // 8-2
		// psdc.NetAmount = f.NetAmount(sdc, psdc)

		// // 5-20. TaxRateの計算  // 5-1-3
		// psdc.TaxRate, e = f.TaxRate(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// // 5-22. TaxAmount  // 5-1-3, 5-20, 5-21
		// psdc.TaxAmount, e = f.TaxAmount(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// // 5-23. GrossAmount  // 5-21, 5-22
		// psdc.GrossAmount, e = f.GrossAmount(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	f.l.Info(psdc)
	err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
