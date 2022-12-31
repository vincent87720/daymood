package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood.backend/internal/model"
	"github.com/vincent87720/daymood.backend/internal/settings"
)

func SetupPurchaseRouters(router *gin.Engine, db *sql.DB, s settings.Settings) (*gin.Engine, error) {

	router.GET("/purchases", GetPurchasesHandler(db))
	router.GET("/purchases/:id", GetPurchaseHandler(db))
	router.POST("/purchases", PostPurchaseHandler(db))
	router.PUT("/purchases/:id", PutPurchaseHandler(db))
	router.DELETE("/purchases/:id", DeletePurchaseHandler(db))
	// router.GET("/suppliers/dumping", DumpFirmHandler(db, s))

	return router, nil
}

func GetPurchasesHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		purchaseXi, modelErr := model.GetAllPurchases(db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": purchaseXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func GetPurchaseHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		purchaseID := context.Param("id")

		if checkEmpty(purchaseID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		purchaseIDVal, err := strconv.ParseInt(purchaseID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}
		purchase := model.Purchase{
			ID: purchaseIDVal,
		}

		purchaseXi, modelErr := purchase.GetPurchase(db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": purchaseXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PostPurchaseHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		purchaseForm := model.Purchase{}

		err := context.BindJSON(&purchaseForm)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(purchaseForm.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		purchase, err := model.NewPurchase(purchaseForm.Name, purchaseForm.Status, purchaseForm.PurchaseType)
		if err != nil {
			context.JSON(http.StatusBadRequest, generalError)
			return
		}

		modelErr := purchase.Create(db)
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

func PutPurchaseHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		supplierID := context.Param("id")

		if checkEmpty(supplierID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		supplierIDVal, err := strconv.ParseInt(supplierID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		purchaseForm := model.Purchase{}

		err = context.BindJSON(&purchaseForm)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(purchaseForm.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		purchase := purchaseForm
		purchase.ID = supplierIDVal

		modelErr := purchase.Update(db)
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

func DeletePurchaseHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		purchaseID := context.Param("id")

		if checkEmpty(purchaseID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		purchaseIDVal, err := strconv.ParseInt(purchaseID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		purchase := model.Purchase{
			ID: purchaseIDVal,
		}

		modelErr := purchase.Delete(db)
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
