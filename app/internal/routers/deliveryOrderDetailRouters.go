package routers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/settings"
)

func SetupDeliveryOrderDetailRouters(router *gin.Engine, db *sql.DB, s settings.Settings) (*gin.Engine, error) {

	router.GET("/deliveryOrderDetails", GetAllDeliveryOrderDetailsHandler(db))
	router.POST("/deliveryOrderDetails/multiple", PostDeliveryOrderDetailsHandler(db, s))
	router.GET("/deliveryOrders/:id/deliveryOrderDetails", GetDeliveryOrderDetailsHandler(db))
	router.POST("/deliveryOrderDetails", PostDeliveryOrderDetailHandler(db, s))
	router.PUT("/deliveryOrderDetails/:id", PutDeliveryOrderDetailHandler(db))
	router.DELETE("/deliveryOrderDetails/:id", DeleteDeliveryOrderDetailHandler(db))

	return router, nil
}

func GetAllDeliveryOrderDetailsHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		deliveryOrderDetailXi, modelErr := model.GetAllDeliveryOrderDetails(db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": deliveryOrderDetailXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PostDeliveryOrderDetailsHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		var deliveryOrderDetailXi []model.DeliveryOrderDetail

		err := context.BindJSON(&deliveryOrderDetailXi)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		for _, value := range deliveryOrderDetailXi {
			avgCosts, modelErr := calcAverageCost(db, s, *value.ProductID)
			if modelErr != nil {
				context.JSON(http.StatusBadRequest, modelError(modelErr))
				return
			}
			value.Cost = float32(int(avgCosts*100)) / 100

		}

		var deliveryOrderDetail model.DeliveryOrderDetail

		modelErr := deliveryOrderDetail.CreateMultiple(db, deliveryOrderDetailXi)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func GetDeliveryOrderDetailsHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		deliveryOrderID := context.Param("id")
		if checkEmpty(deliveryOrderID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		deliveryOrderDetailIDVal, err := strconv.ParseInt(deliveryOrderID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}
		deliveryOrderDetailXi, modelErr := model.GetDeliveryOrderDetails(db, deliveryOrderDetailIDVal)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": deliveryOrderDetailXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PostDeliveryOrderDetailHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		deliveryOrderDetail := model.DeliveryOrderDetail{}

		err := context.BindJSON(&deliveryOrderDetail)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		avgCosts, modelErr := calcAverageCost(db, s, *deliveryOrderDetail.ProductID)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}
		deliveryOrderDetail.Cost = float32(int(avgCosts*100)) / 100

		modelErr = deliveryOrderDetail.Create(db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func calcAverageCost(db *sql.DB, s settings.Settings, productID int64) (avgCosts float32, modelErr *model.ModelError) {
	var costs float32 = 0
	historyXi, modelErr := model.GetProductPurchaseHistories(db, productID)
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

func PutDeliveryOrderDetailHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		deliveryOrderDetailID := context.Param("id")

		if checkEmpty(deliveryOrderDetailID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		deliveryOrderDetailIDVal, err := strconv.ParseInt(deliveryOrderDetailID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		deliveryOrderDetailForm := model.DeliveryOrderDetail{}

		err = context.BindJSON(&deliveryOrderDetailForm)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		deliveryOrderDetail := deliveryOrderDetailForm
		deliveryOrderDetail.ID = deliveryOrderDetailIDVal

		modelErr := deliveryOrderDetail.Update(db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func DeleteDeliveryOrderDetailHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		deliveryOrderDetailID := context.Param("id")

		if checkEmpty(deliveryOrderDetailID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		deliveryOrderDetailIDVal, err := strconv.ParseInt(deliveryOrderDetailID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		deliveryOrderDetail := model.DeliveryOrderDetail{
			ID: deliveryOrderDetailIDVal,
		}

		modelErr := deliveryOrderDetail.Delete(db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
		return
	}

	return gin.HandlerFunc(fn)
}
