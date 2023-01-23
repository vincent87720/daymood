package usecases

import (
	"database/sql"
	"fmt"

	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/settings"
)

type DeliveryOrderDetail struct {
	Model    *model.DeliveryOrderDetail
	Settings settings.Settings
}

func NewDeliveryOrderDetail(deliveryOrderDetailModel *model.DeliveryOrderDetail) *DeliveryOrderDetail {
	return &DeliveryOrderDetail{
		Model: deliveryOrderDetailModel,
	}
}

func NewDeliveryOrderDetailWithSettings(deliveryOrderDetailModel *model.DeliveryOrderDetail, settings settings.Settings) *DeliveryOrderDetail {
	return &DeliveryOrderDetail{
		Model:    deliveryOrderDetailModel,
		Settings: settings,
	}
}

func (deliveryOrderDetail *DeliveryOrderDetail) Read(db *sql.DB) ([]interface{}, *model.ModelError) {

	deliveryOrderDetailXi, modelErr := model.Read(deliveryOrderDetail.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return deliveryOrderDetailXi, nil
}

func (deliveryOrderDetail *DeliveryOrderDetail) ReadAll(db *sql.DB) ([]interface{}, *model.ModelError) {

	deliveryOrderDetailXi, modelErr := model.ReadAll(deliveryOrderDetail.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return deliveryOrderDetailXi, nil
}

func (deliveryOrderDetail *DeliveryOrderDetail) Create(db *sql.DB) *model.ModelError {
	avgCosts, modelErr := calcAverageCost(db, deliveryOrderDetail.Settings, *deliveryOrderDetail.Model.ProductID)
	if modelErr != nil {
		return modelErr
	}
	deliveryOrderDetail.Model.Cost = avgCosts

	modelErr = model.Create(deliveryOrderDetail.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (deliveryOrderDetail *DeliveryOrderDetail) CreateMultiple(db *sql.DB, deliveryOrderDetailXi interface{}) *model.ModelError {
	detailXi, ok := deliveryOrderDetailXi.([]model.DeliveryOrderDetail)
	if !ok {
		return nil
	}
	for _, value := range detailXi {
		avgCosts, modelErr := calcAverageCost(db, deliveryOrderDetail.Settings, *value.ProductID)
		if modelErr != nil {
			return modelErr
		}
		value.Cost = float32(int(avgCosts*100)) / 100
	}

	var deliveryOrderDetailModel model.DeliveryOrderDetail

	modelErr := deliveryOrderDetailModel.CreateMultiple(db, detailXi)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (deliveryOrderDetail *DeliveryOrderDetail) Update(db *sql.DB) *model.ModelError {

	modelErr := model.Update(deliveryOrderDetail.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (deliveryOrderDetail *DeliveryOrderDetail) Delete(db *sql.DB) *model.ModelError {

	modelErr := model.Delete(deliveryOrderDetail.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func calcAverageCost(db *sql.DB, s settings.Settings, productID int64) (avgCosts float32, modelErr *model.ModelError) {
	var costs float32 = 0

	historyModel := &model.ProductPurchaseHistory{}
	historyXi, modelErr := historyModel.ReadAll(db, productID)
	if modelErr != nil {
		return costs, modelErr
	}

	for _, value := range historyXi {
		costs += calcImportCost(value)
		costs += calcOtherCost(s)
	}

	if len(historyXi) != 0 {
		avgCosts = costs / float32(len(historyXi))
	}
	avgCosts = float32(int(avgCosts*100)) / 100

	return avgCosts, nil
}

func calcImportCost(history model.ProductPurchaseHistory) (importCost float32) {
	if history.ExchangeRateKrw == nil || history.PurchaseQTY == nil {
		return
	}

	var exchangeRateKrw float32
	//檢查分母不得為0
	if *history.ExchangeRateKrw <= 0 {
		exchangeRateKrw = 1
	} else {
		exchangeRateKrw = *history.ExchangeRateKrw
	}

	var purchaseQTY float32
	//檢查分母不得為0
	if *history.PurchaseQTY <= 0 {
		purchaseQTY = 1
	} else {
		purchaseQTY = float32(*history.PurchaseQTY)
	}

	var wholesalePrice float32
	wholesalePrice = 0
	if history.WholesalePrice != nil {
		wholesalePrice = *history.WholesalePrice
	}

	var shippingAgentCutPercent float32
	shippingAgentCutPercent = 0
	if history.ShippingAgentCutPercent != nil {
		shippingAgentCutPercent = *history.ShippingAgentCutPercent
	}

	var shippingFeeKokusaiKrw float32
	shippingFeeKokusaiKrw = 0
	if history.ShippingFeeKokusaiKrw != nil {
		shippingFeeKokusaiKrw = *history.ShippingFeeKokusaiKrw
	}

	var shippingFeeKr float32
	shippingFeeKr = 0
	if history.ShippingFeeKr != nil {
		shippingFeeKr = *history.ShippingFeeKr
	}

	var tariffTwd float32
	tariffTwd = 0
	if history.TariffTwd != nil {
		tariffTwd = *history.TariffTwd
	}

	var shippingFeeTw float32
	shippingFeeTw = 0
	if history.ShippingFeeTw != nil {
		shippingFeeTw = *history.ShippingFeeTw
	}

	//計算貨運行抽成（商品金額*貨運行抽成百分比）
	shippingAgentCutKrw := wholesalePrice * (shippingAgentCutPercent / 100)

	//計算每個商品的國際運費（國際運費總額/進貨商品數量）
	shippingFeeKokusaiDivideByPurchaseQTY := shippingFeeKokusaiKrw / purchaseQTY

	//計算韓幣開銷總金額
	//商品批價（wholesalePrice）
	//貨運行抽成（shippingAgentCutKrw）
	//每個商品的國際運費（shippingFeeKokusaiDivideByPurchaseQTY）
	//韓國國內運費（shippingFeeKr）
	subtotalKrw := wholesalePrice + shippingAgentCutKrw + shippingFeeKokusaiDivideByPurchaseQTY + shippingFeeKr

	//計算台幣開銷總金額
	//關稅（tariffTwd）
	//台灣國內運費（shippingFeeTw）
	costTwd := tariffTwd + shippingFeeTw

	//計算貨運關稅成本
	//換為台幣後的韓幣開銷（subtotalKrw / 韓圓匯率（exchangeRateKrw)
	//每個商品的台幣開銷（台幣開銷總金額/商品個數）（costTwd / purchaseQTY）
	importCost = (subtotalKrw / exchangeRateKrw) + (costTwd / purchaseQTY)

	return importCost
}

func calcOtherCost(s settings.Settings) (otherCost float32) {
	tradings, err := s.GetTradingSettings()
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range tradings.Costs {
		otherCost += value.Value
	}
	return otherCost
}
