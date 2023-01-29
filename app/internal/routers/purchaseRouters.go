package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/settings"
	"github.com/vincent87720/daymood/app/internal/usecases"
)

func SetupPurchaseRouters(router *gin.RouterGroup, db *sql.DB, s settings.Settings) (*gin.RouterGroup, error) {

	router.GET("/api/purchases", GetPurchasesHandler(db))
	router.GET("/api/purchases/:id", GetPurchaseHandler(db))
	router.POST("/api/purchases", PostPurchaseHandler(db))
	router.PUT("/api/purchases/:id", PutPurchaseHandler(db))
	router.DELETE("/api/purchases/:id", DeletePurchaseHandler(db))
	// router.GET("/suppliers/dumping", DumpFirmHandler(db, s))

	return router, nil
}

func GetPurchasesHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		purchaseModel := &model.Purchase{}

		purchase := usecases.NewPurchase(purchaseModel)
		purchaseXi, modelErr := usecases.ReadAll(purchase, db)
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
		purchaseModel := &model.Purchase{
			ID: purchaseIDVal,
		}

		purchase := usecases.NewPurchase(purchaseModel)
		purchaseXi, modelErr := usecases.Read(purchase, db)
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

		purchaseModel := &model.Purchase{}

		err := context.BindJSON(&purchaseModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(purchaseModel.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		purchase := usecases.NewPurchase(purchaseModel)
		modelErr := usecases.Create(purchase, db)
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

		purchaseModel := &model.Purchase{}

		err = context.BindJSON(&purchaseModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(purchaseModel.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		purchaseModel.ID = supplierIDVal

		purchase := usecases.NewPurchase(purchaseModel)
		modelErr := usecases.Update(purchase, db)
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

		purchaseModel := &model.Purchase{
			ID: purchaseIDVal,
		}

		purchase := usecases.NewPurchase(purchaseModel)
		modelErr := usecases.Delete(purchase, db)
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
