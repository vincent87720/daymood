package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood.backend/internal/model"
	"github.com/vincent87720/daymood.backend/internal/settings"
)

func SetupPurchaseDetailRouters(router *gin.Engine, db *sql.DB, s settings.Settings) (*gin.Engine, error) {

	router.GET("/purchaseDetails", GetAllPurchaseDetailsHandler(db))
	router.POST("/purchaseDetails/multiple", PostPurchaseDetailsHandler(db))
	router.GET("/purchases/:id/purchaseDetails", GetPurchaseDetailsHandler(db))
	router.POST("/purchaseDetails", PostPurchaseDetailHandler(db))
	router.PUT("/purchaseDetails/:id", PutPurchaseDetailHandler(db))
	router.DELETE("/purchaseDetails/:id", DeletePurchaseDetailHandler(db))
	// router.GET("/suppliers/dumping", DumpFirmHandler(db, s))

	return router, nil
}

func GetAllPurchaseDetailsHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		purchaseDetailXi, modelErr := model.GetAllPurchaseDetails(db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": purchaseDetailXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PostPurchaseDetailsHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		var purchaseDetailXi []model.PurchaseDetail

		err := context.BindJSON(&purchaseDetailXi)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		for _, val := range purchaseDetailXi {
			if checkEmpty(val.Name) == true {
				context.JSON(http.StatusBadRequest, emptyError("name"))
				return
			}
		}

		var purchaseDetail model.PurchaseDetail

		modelErr := purchaseDetail.CreateMultiple(db, purchaseDetailXi)
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

func GetPurchaseDetailsHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		purchaseID := context.Param("id")
		if checkEmpty(purchaseID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		purchaseDetailIDVal, err := strconv.ParseInt(purchaseID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}
		purchaseDetailXi, modelErr := model.GetPurchaseDetails(db, purchaseDetailIDVal)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": purchaseDetailXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PostPurchaseDetailHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		purchaseDetail := model.PurchaseDetail{}

		err := context.BindJSON(&purchaseDetail)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(purchaseDetail.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		modelErr := purchaseDetail.Create(db)
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

func PutPurchaseDetailHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		purchaseDetailID := context.Param("id")

		if checkEmpty(purchaseDetailID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		purchaseDetailIDVal, err := strconv.ParseInt(purchaseDetailID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		purchaseDetailForm := model.PurchaseDetail{}

		err = context.BindJSON(&purchaseDetailForm)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(purchaseDetailForm.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		purchaseDetail := purchaseDetailForm
		purchaseDetail.ID = purchaseDetailIDVal

		modelErr := purchaseDetail.Update(db)
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

func DeletePurchaseDetailHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		purchaseDetailID := context.Param("id")

		if checkEmpty(purchaseDetailID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		purchaseDetailIDVal, err := strconv.ParseInt(purchaseDetailID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		purchaseDetail := model.PurchaseDetail{
			ID: purchaseDetailIDVal,
		}

		modelErr := purchaseDetail.Delete(db)
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
