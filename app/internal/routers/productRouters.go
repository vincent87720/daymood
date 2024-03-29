package routers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/usecases"
)

func SetupProductRouters(router *gin.RouterGroup, db *sql.DB) *gin.RouterGroup {

	router.GET("/api/products", GetProductsHandler(db))
	router.POST("/api/products", PostProductHandler(db))
	router.POST("/api/products/multiple", PostProductsHandler(db))
	router.PUT("/api/products/:id", PutProductHandler(db))
	router.DELETE("/api/products/:id", DeleteProductHandler(db))
	router.GET("/api/products/:id/purchaseHistories", GetProductPurchaseHistoriesHandler(db))
	router.GET("/api/products/:id/deliveryHistories", GetProductDeliveryHistoriesHandler(db))
	// router.GET("/products/dumping", DumpProductHandler(db, s))
	// router.POST("/stocks/:id", PostStocksHandler(db)) //新增庫存
	// router.GET("/products/images/:id", GetProductImgHandler(db, s))

	return router
}

func GetProductsHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		productModel := &model.Product{}
		product := usecases.NewProduct(productModel)
		productXi, modelErr := usecases.ReadAll(product, db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": productXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PostProductHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		productModel := &model.Product{}

		err := context.BindJSON(&productModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		checkList := []Field{
			{Key: "Name", Val: productModel.Name},
		}
		err = checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		product := usecases.NewProduct(productModel)
		modelErr := usecases.Create(product, db)
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

func PostProductsHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		var productXi []model.Product

		err := context.BindJSON(&productXi)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		for _, val := range productXi {
			checkList := []Field{
				{Key: "Name", Val: val.Name},
			}
			err = checkEmpty(checkList)
			if err != nil {
				context.JSON(http.StatusBadRequest, emptyError(err))
				return
			}
		}

		product := usecases.NewProduct(&model.Product{})
		modelErr := usecases.CreateMultiple(product, db, productXi)
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

func PutProductHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		productID := context.Param("id")

		productModel := &model.Product{}

		err := context.BindJSON(&productModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		checkList := []Field{
			{Key: "id", Val: productID},
			{Key: "Name", Val: productModel.Name},
		}
		err = checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		productModel.ID, err = strconv.ParseInt(productID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		product := usecases.NewProduct(productModel)
		modelErr := usecases.Update(product, db)
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

func DeleteProductHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		productID := context.Param("id")

		productModel := &model.Product{}

		checkList := []Field{
			{Key: "id", Val: productID},
		}
		err := checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		productModel.ID, err = strconv.ParseInt(productID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "PERR2",
				"message": "Invalid input",
			})
			return
		}

		product := usecases.NewProduct(productModel)
		modelErr := usecases.Delete(product, db)
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

func GetProductPurchaseHistoriesHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		productID := context.Param("id")

		productModel := &model.Product{}

		checkList := []Field{
			{Key: "id", Val: productID},
		}
		err := checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		productModel.ID, err = strconv.ParseInt(productID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		product := usecases.NewProduct(productModel)
		historyXi, modelErr := product.ReadPurchaseHistories(db, productModel.ID)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": historyXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func GetProductDeliveryHistoriesHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		productID := context.Param("id")

		productModel := &model.Product{}

		checkList := []Field{
			{Key: "id", Val: productID},
		}
		err := checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		productModel.ID, err = strconv.ParseInt(productID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		product := usecases.NewProduct(productModel)
		historyXi, modelErr := product.ReadDeliveryHistories(db, productModel.ID)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": historyXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

// func DumpProductHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
// 	fn := func(context *gin.Context) {
// 		productXi, err := model.GetAllProducts(db)
// 		if err != nil {
// 			fmt.Println(err)
// 			context.JSON(http.StatusBadRequest, gin.H{
// 				"status": "FAIL",
// 			})
// 			return
// 		}

// 		var prepareCSV [][]string

// 		prepareCSV = append(prepareCSV, []string{"廠商名稱", "商品編號", "採購商品編號", "商品名稱", "商品種類", "庫存", "重量(g)", "批價(KRW)", "批價(TWD)", "定價(TWD)", "售價(TWD)", "加稅運成本(TWD)", "總成本(TWD)", "毛利(售價-總成本)", "毛利率"})

// 		for _, v := range productXi {
// 			tmpXi := []string{}
// 			tmpXi = append(tmpXi, v.FirmInfo.Name, v.ProductSku, v.PurchaseProductID,
// 				v.ProductName, v.ProductType, strconv.Itoa(v.Stocks),
// 				strconv.FormatFloat(float64(v.Weight), 'f', -1, 32),
// 				strconv.FormatFloat(float64(v.KrwWholesalePrice), 'f', -1, 32),
// 				strconv.FormatFloat(float64(v.NtdWholesalePrice), 'f', -1, 32),
// 				strconv.FormatFloat(float64(v.NtdListPrice), 'f', -1, 32),
// 				strconv.FormatFloat(float64(v.NtdSellingPrice), 'f', -1, 32),
// 				strconv.FormatFloat(float64(v.NtdCnf), 'f', -1, 32),
// 				strconv.FormatFloat(float64(v.NtdCost), 'f', -1, 32),
// 				strconv.FormatFloat(float64(v.NtdSellingPrice-v.NtdCost), 'f', -1, 32),
// 				strconv.FormatFloat(float64((v.NtdSellingPrice-v.NtdCost)/v.NtdSellingPrice), 'f', -1, 32))
// 			prepareCSV = append(prepareCSV, tmpXi)
// 		}

// 		err = s.WriteCSV(s.GetExeFilePath()+"/products.csv", prepareCSV)
// 		if err != nil {
// 			fmt.Println(err)
// 			context.JSON(http.StatusBadRequest, gin.H{
// 				"status": "FAIL",
// 			})
// 			return
// 		}

// 		context.FileAttachment(s.GetExeFilePath()+"/products.csv", "products.csv")
// 		return
// 	}

// 	return gin.HandlerFunc(fn)
// }

// func PostStocksHandler(db *sql.DB) gin.HandlerFunc {
// 	fn := func(context *gin.Context) {

// 		productID := context.Param("id")
// 		quantity := context.PostForm("qty")

// 		if checkEmpty(productID) == true {
// 			context.JSON(http.StatusBadRequest, gin.H{
// 				"status":  "FAIL",
// 				"code":    "PERR2",
// 				"message": "Invalid input",
// 			})
// 			return
// 		}

// 		quantityVal, err := strconv.ParseInt(quantity, 10, 64)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			context.JSON(http.StatusBadRequest, gin.H{
// 				"status":  "FAIL",
// 				"code":    "PERR2",
// 				"message": "Invalid input",
// 			})
// 			return
// 		}

// 		err = model.IncreaseStocks(db, productID, int(quantityVal))
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			context.JSON(http.StatusBadRequest, gin.H{
// 				"status": "FAIL",
// 			})
// 			return
// 		}

// 		context.JSON(http.StatusOK, gin.H{
// 			"status": "OK",
// 		})
// 		return
// 	}

// 	return gin.HandlerFunc(fn)
// }

// func GetProductImgHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
// 	fn := func(context *gin.Context) {
// 		productImgID := context.Param("id")
// 		context.File(s.GetExeFilePath() + "/img/" + productImgID)
// 		return
// 	}

// 	return gin.HandlerFunc(fn)
// }

func checkNum(context *gin.Context, s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "FAIL",
			"code":    "PERR2",
			"message": "Invalid input",
		})
		return -1, err
	}
	return float32(f), nil
}
