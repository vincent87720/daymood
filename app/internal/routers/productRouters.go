package routers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vincent87720/daymood.backend/internal/model"
	"github.com/vincent87720/daymood.backend/internal/settings"
)

func SetupProductRouters(router *gin.Engine, db *sql.DB, s settings.Settings) (*gin.Engine, error) {

	router.GET("/products", GetProductsHandler(db))
	router.POST("/products", PostProductHandler(db, s))
	router.PUT("/products/:id", PutProductHandler(db, s))
	router.DELETE("/products/:id", DeleteProductHandler(db, s))
	router.GET("/products/dumping", DumpProductHandler(db, s))
	router.POST("/stocks/:id", PostStocksHandler(db)) //新增庫存
	router.GET("/products/images/:id", GetProductImgHandler(db, s))

	return router, nil
}

func GetProductsHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		productXi, err := model.GetAllProducts(db)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":   "OK",
			"products": productXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PostProductHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		productSku := context.PostForm("productSku")
		productName := context.PostForm("productName")
		productType := context.PostForm("productType")
		weight := context.PostForm("weight")
		krwWholesalePrice := context.PostForm("krwWholesalePrice")
		ntdWholesalePrice := context.PostForm("ntdWholesalePrice")
		ntdListPrice := context.PostForm("ntdListPrice")
		ntdSellingPrice := context.PostForm("ntdSellingPrice")
		ntdCost := context.PostForm("ntdCost")
		ntdCnf := context.PostForm("ntdCnf")
		firmID := context.PostForm("firmID")
		purchaseProductID := context.PostForm("purchaseProductID")
		imgFile, err := context.FormFile("imgFile")
		imgName := ""
		imgID := ""

		if imgFile != nil {
			imgName = imgFile.Filename

			re := regexp.MustCompile("\\.\\w+$")
			match := re.FindStringSubmatch(imgFile.Filename)
			imgID = uuid.New().String()[:6] + match[0]

			if _, err := os.Stat(s.GetExeFilePath() + "/img"); os.IsNotExist(err) {
				_ = os.Mkdir(s.GetExeFilePath()+"/img", os.ModePerm)
			}
			err = context.SaveUploadedFile(imgFile, s.GetExeFilePath()+"/img/"+imgID)
			if err != nil {
				fmt.Println(err)
				context.JSON(http.StatusBadRequest, gin.H{
					"status":  "FAIL",
					"message": "Upload failed.",
				})
				return
			}
		}

		if checkEmpty(productSku) == true ||
			checkEmpty(productName) == true ||
			checkEmpty(productType) == true ||
			checkEmpty(weight) == true ||
			checkEmpty(krwWholesalePrice) == true ||
			checkEmpty(ntdWholesalePrice) == true ||
			checkEmpty(ntdListPrice) == true ||
			checkEmpty(ntdSellingPrice) == true ||
			checkEmpty(ntdCnf) == true ||
			checkEmpty(ntdCost) == true ||
			checkEmpty(firmID) == true ||
			checkEmpty(purchaseProductID) == true {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "PERR2",
				"message": "Invalid input",
			})
			return
		}

		weightVal, err := checkNum(context, weight)
		if err != nil {
			return
		}

		krwWholesalePriceVal, err := checkNum(context, krwWholesalePrice)
		if err != nil {
			return
		}

		ntdWholesalePriceVal, err := checkNum(context, ntdWholesalePrice)
		if err != nil {
			return
		}

		ntdListPriceVal, err := checkNum(context, ntdListPrice)
		if err != nil {
			return
		}

		ntdSellingPriceVal, err := checkNum(context, ntdSellingPrice)
		if err != nil {
			return
		}

		ntdCnfVal, err := checkNum(context, ntdCnf)
		if err != nil {
			return
		}

		ntdCostVal, err := checkNum(context, ntdCost)
		if err != nil {
			return
		}

		firmIDVal, err := strconv.ParseInt(firmID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "PERR2",
				"message": "Invalid input",
			})
			return
		}

		product, err := model.NewProduct(
			productSku, productName, productType,
			imgName, imgID,
			weightVal,
			krwWholesalePriceVal, ntdWholesalePriceVal,
			ntdListPriceVal, ntdSellingPriceVal, ntdCnfVal, ntdCostVal, purchaseProductID, firmIDVal)
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		errInfo, err := product.Create(db)
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    errInfo.Code,
				"message": errInfo.Message,
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

func PutProductHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		productID := context.Param("id")
		productSku := context.PostForm("productSku")
		productName := context.PostForm("productName")
		productType := context.PostForm("productType")
		weight := context.PostForm("weight")
		krwWholesalePrice := context.PostForm("krwWholesalePrice")
		ntdWholesalePrice := context.PostForm("ntdWholesalePrice")
		ntdListPrice := context.PostForm("ntdListPrice")
		ntdSellingPrice := context.PostForm("ntdSellingPrice")
		ntdCost := context.PostForm("ntdCost")
		ntdCnf := context.PostForm("ntdCnf")
		firmID := context.PostForm("firmID")
		purchaseProductID := context.PostForm("purchaseProductID")
		productImgName := context.PostForm("productImgName")
		productImgID := context.PostForm("productImgID")
		delImgID := context.PostForm("delImgID")
		imgFile, err := context.FormFile("imgFile")
		//舊檔案存在，刪除檔案
		if checkEmpty(delImgID) == false {
			err = os.Remove(s.GetExeFilePath() + "/img/" + delImgID)
			if err != nil {
				fmt.Println(err)
			}
		}

		//新檔案存在，儲存檔案
		if imgFile != nil {
			productImgName = imgFile.Filename

			re := regexp.MustCompile("\\.\\w+$")
			match := re.FindStringSubmatch(imgFile.Filename)
			productImgID = uuid.New().String()[:6] + match[0]

			if _, err := os.Stat(s.GetExeFilePath() + "/img"); os.IsNotExist(err) {
				_ = os.Mkdir(s.GetExeFilePath()+"/img", os.ModePerm)
			}
			err = context.SaveUploadedFile(imgFile, s.GetExeFilePath()+"/img/"+productImgID)
			if err != nil {
				fmt.Println(err)
				context.JSON(http.StatusBadRequest, gin.H{
					"status":  "FAIL",
					"message": "Upload failed.",
				})
				return
			}
		}

		if checkEmpty(productID) == true ||
			checkEmpty(productSku) == true ||
			checkEmpty(productName) == true ||
			checkEmpty(productType) == true ||
			checkEmpty(weight) == true ||
			checkEmpty(krwWholesalePrice) == true ||
			checkEmpty(ntdWholesalePrice) == true ||
			checkEmpty(ntdListPrice) == true ||
			checkEmpty(ntdSellingPrice) == true ||
			checkEmpty(ntdCnf) == true ||
			checkEmpty(ntdCost) == true ||
			checkEmpty(firmID) == true ||
			checkEmpty(purchaseProductID) == true {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "PERR2",
				"message": "Invalid input",
			})
			return
		}

		productIDVal, err := strconv.ParseInt(productID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "PERR2",
				"message": "Invalid input",
			})
			return
		}

		weightVal, err := checkNum(context, weight)
		if err != nil {
			return
		}

		krwWholesalePriceVal, err := checkNum(context, krwWholesalePrice)
		if err != nil {
			return
		}

		ntdWholesalePriceVal, err := checkNum(context, ntdWholesalePrice)
		if err != nil {
			return
		}

		ntdListPriceVal, err := checkNum(context, ntdListPrice)
		if err != nil {
			return
		}

		ntdSellingPriceVal, err := checkNum(context, ntdSellingPrice)
		if err != nil {
			return
		}

		ntdCnfVal, err := checkNum(context, ntdCnf)
		if err != nil {
			return
		}

		ntdCostVal, err := checkNum(context, ntdCost)
		if err != nil {
			return
		}

		firmIDVal, err := strconv.ParseInt(firmID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "PERR2",
				"message": "Invalid input",
			})
			return
		}

		product, err := model.NewProduct(
			productSku, productName, productType,
			productImgName, productImgID,
			weightVal,
			krwWholesalePriceVal, ntdWholesalePriceVal,
			ntdListPriceVal, ntdSellingPriceVal,
			ntdCnfVal, ntdCostVal,
			purchaseProductID, firmIDVal)
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		product.ProductID = productIDVal

		errInfo, err := product.Update(db)
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    errInfo.Code,
				"message": errInfo.Message,
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

func DeleteProductHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		// context.Header("Access-Control-Allow-Origin", "*") //allow CORS actions

		productID := context.Param("id")

		if checkEmpty(productID) == true {
			err := fmt.Errorf("Invalid input")
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		productIDVal, err := strconv.ParseInt(productID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "PERR2",
				"message": "Invalid input",
			})
			return
		}

		product := model.Product{
			ProductID: productIDVal,
		}

		//刪除照片
		productMap, err := model.GetProductsByID(db, []string{productID})
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}
		if len(productMap) > 0 {
			imgID := productMap[productID].ProductImgID
			if checkEmpty(imgID) == false {
				os.Remove(s.GetExeFilePath() + "/img/" + imgID)
			}
		}

		err = product.Delete(db)
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
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

func DumpProductHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		productXi, err := model.GetAllProducts(db)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		var prepareCSV [][]string

		prepareCSV = append(prepareCSV, []string{"廠商名稱", "商品編號", "採購商品編號", "商品名稱", "商品種類", "庫存", "重量(g)", "批價(KRW)", "批價(TWD)", "定價(TWD)", "售價(TWD)", "加稅運成本(TWD)", "總成本(TWD)", "毛利(售價-總成本)", "毛利率"})

		for _, v := range productXi {
			tmpXi := []string{}
			tmpXi = append(tmpXi, v.FirmInfo.Name, v.ProductSku, v.PurchaseProductID,
				v.ProductName, v.ProductType, strconv.Itoa(v.Stocks),
				strconv.FormatFloat(float64(v.Weight), 'f', -1, 32),
				strconv.FormatFloat(float64(v.KrwWholesalePrice), 'f', -1, 32),
				strconv.FormatFloat(float64(v.NtdWholesalePrice), 'f', -1, 32),
				strconv.FormatFloat(float64(v.NtdListPrice), 'f', -1, 32),
				strconv.FormatFloat(float64(v.NtdSellingPrice), 'f', -1, 32),
				strconv.FormatFloat(float64(v.NtdCnf), 'f', -1, 32),
				strconv.FormatFloat(float64(v.NtdCost), 'f', -1, 32),
				strconv.FormatFloat(float64(v.NtdSellingPrice-v.NtdCost), 'f', -1, 32),
				strconv.FormatFloat(float64((v.NtdSellingPrice-v.NtdCost)/v.NtdSellingPrice), 'f', -1, 32))
			prepareCSV = append(prepareCSV, tmpXi)
		}

		err = s.WriteCSV(s.GetExeFilePath()+"/products.csv", prepareCSV)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		context.FileAttachment(s.GetExeFilePath()+"/products.csv", "products.csv")
		return
	}

	return gin.HandlerFunc(fn)
}

func PostStocksHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		productID := context.Param("id")
		quantity := context.PostForm("qty")

		if checkEmpty(productID) == true {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "PERR2",
				"message": "Invalid input",
			})
			return
		}

		quantityVal, err := strconv.ParseInt(quantity, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "PERR2",
				"message": "Invalid input",
			})
			return
		}

		err = model.IncreaseStocks(db, productID, int(quantityVal))
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
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

func GetProductImgHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		productImgID := context.Param("id")
		context.File(s.GetExeFilePath() + "/img/" + productImgID)
		return
	}

	return gin.HandlerFunc(fn)
}

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
