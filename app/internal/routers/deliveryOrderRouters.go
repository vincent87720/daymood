package routers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood.backend/internal/model"
)

func SetupDeliveryOrderRouters(router *gin.Engine, db *sql.DB) (*gin.Engine, error) {

	//從EasyStore及資料庫取得訂單資訊
	router.GET("/deliveryorders/:id", GetEasyPickOrderHandler(db))
	router.POST("/deliveryorders", PostDeliveryOrderHandler(db)) //減少庫存
	// router.GET("/products", GetProductHandler(db))
	// router.DELETE("/products/:id", DeleteProductHandler(db))

	return router, nil
}

type JSONItems struct {
	DeliveryProductName string `json:"product_name"`
	NtdRetailPrice      string `json:"price"`
	ProductSku          string `json:"sku"`
	DeliveryProductQty  int    `json:"quantity"`
}

type JSONOrder struct {
	DeliveryOrderID    int         `json:"id"`
	DeliveryOrderTotal string      `json:"total_price"`
	Products           []JSONItems `json:"items"`
}

type JSONData struct {
	Order JSONOrder `json:"order"`
}

type JSONRAW struct {
	Data JSONData `json:"data"`
}

func GetEasyPickOrderHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		deliveryOrderID := context.Param("id")
		resp, err := http.Get("http://127.0.0.1:5000/easystoreorders/" + deliveryOrderID)
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "DERR0",
				"message": "Can't connect to EasyStore.",
			})
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		var data JSONRAW

		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		deliveryOrder, errInfo, err := organizeData(db, data)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    errInfo.Code,
				"message": errInfo.Message,
			})
			return
		}

		//檢查productIDXi陣列長度，小於等於零則不能進入GetProducts
		if len(deliveryOrder.Products) <= 0 {
			context.JSON(http.StatusOK, gin.H{
				"status": "OK",
				"order":  "[]",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"order":  deliveryOrder,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func organizeData(db *sql.DB, easyStoreData JSONRAW) (model.DeliveryOrder, *model.ModelError, error) {

	var deliveryOrder model.DeliveryOrder

	//將從EasyStore取得的價格轉型後放入deliveryOrder.EasyStoreOrderTotal
	totalPriceFloat, err := strconv.ParseFloat(easyStoreData.Data.Order.DeliveryOrderTotal, 32)
	if err != nil {
		fmt.Println(err)
	}
	deliveryOrder.EasyStoreOrderTotal = float32(totalPriceFloat)

	//將從EasyStore取得的訂單編號轉型後放入deliveryOrder.DeliveryOrderID
	deliveryOrder.DeliveryOrderID = strconv.Itoa(easyStoreData.Data.Order.DeliveryOrderID)

	//將從EasyStore取得的訂單商品資訊放入deliveryOrder.Products
	for _, val := range easyStoreData.Data.Order.Products {
		priceFloat, err := strconv.ParseFloat(val.NtdRetailPrice, 32)
		if err != nil {
			fmt.Println(err)
		}
		p := model.DeliveryProduct{
			ProductSku:          val.ProductSku,
			DeliveryProductName: val.DeliveryProductName,
			NtdSellingPrice:     float32(priceFloat),
			DeliveryProductQty:  val.DeliveryProductQty,
		}
		deliveryOrder.Products = append(deliveryOrder.Products, p)
	}

	var productSkuXi []string

	//檢查productID是否只能是[A-Za-z0-9_-]，不滿足規則便將其移除，滿足規則將其加入productStrXi
	for idx, val := range deliveryOrder.Products {
		r, _ := regexp.Compile(`[\w\-]+`)
		if r.MatchString(val.ProductSku) == false {
			deliveryOrder.Products = append(deliveryOrder.Products[:idx], deliveryOrder.Products[idx+1:]...)
		} else {
			productSkuXi = append(productSkuXi, val.ProductSku)
		}
	}

	//依照productStrXi裡面的字串在資料庫內查詢對應資訊
	productMap, err := model.GetProducts(db, productSkuXi)
	if err != nil {
		fmt.Println(err)
		return deliveryOrder, &model.ModelError{Code: 0, Message: "?"}, err
	}
	if len(productSkuXi) != len(productMap) {

		return deliveryOrder, &model.ModelError{Code: 1, Message: "Some products not found in database."}, errors.New("Some products not found in database.")
	}

	for idx, val := range deliveryOrder.Products {
		//將查詢到的NtdRetailPrice放入deliveryOrder中
		deliveryOrder.Products[idx].NtdSellingPrice = productMap[val.ProductSku].NtdSellingPrice

		//計算並記錄商品每個商品的小計
		retailPrice := deliveryOrder.Products[idx].NtdSellingPrice
		qty := deliveryOrder.Products[idx].DeliveryProductQty
		deliveryOrder.Products[idx].DeliveryProductSubtotal = retailPrice * float32(qty)

		//累加每個商品的小計，計算總價
		deliveryOrder.DeliveryOrderTotal += deliveryOrder.Products[idx].DeliveryProductSubtotal
	}

	//將總價減去從EasyStore取得的價格，用於計算折扣
	deliveryOrder.Discount = deliveryOrder.DeliveryOrderTotal - deliveryOrder.EasyStoreOrderTotal

	return deliveryOrder, nil, nil
}

func PostDeliveryOrderHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		type JsonRequest struct {
			Order model.DeliveryOrder `json:"order"`
		}

		var json JsonRequest
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := json.Order.DecreaseStocks(db)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "DERR2",
				"message": "Transaction fail.",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
		return
	}

	return gin.HandlerFunc(fn)
}
