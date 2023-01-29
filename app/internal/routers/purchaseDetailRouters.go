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

func SetupPurchaseDetailRouters(router *gin.RouterGroup, db *sql.DB, s settings.Settings) (*gin.RouterGroup, error) {

	router.GET("/api/purchaseDetails", GetAllPurchaseDetailsHandler(db))
	router.POST("/api/purchaseDetails/multiple", PostPurchaseDetailsHandler(db))
	router.GET("/api/purchases/:id/purchaseDetails", GetPurchaseDetailsHandler(db))
	router.POST("/api/purchaseDetails", PostPurchaseDetailHandler(db))
	router.PUT("/api/purchaseDetails/:id", PutPurchaseDetailHandler(db))
	router.DELETE("/api/purchaseDetails/:id", DeletePurchaseDetailHandler(db))
	// router.GET("/suppliers/dumping", DumpFirmHandler(db, s))

	return router, nil
}

func GetAllPurchaseDetailsHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		purchaseDetailModel := &model.PurchaseDetail{}
		purchaseDetail := usecases.NewPurchaseDetail(purchaseDetailModel)
		purchaseDetailXi, modelErr := model.ReadAll(purchaseDetail, db)
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

		purchaseDetail := usecases.NewPurchaseDetail(&model.PurchaseDetail{})
		modelErr := usecases.CreateMultiple(purchaseDetail, db, purchaseDetailXi)
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

		purchaseIDVal, err := strconv.ParseInt(purchaseID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}
		purchaseDetailModel := &model.PurchaseDetail{
			PurchaseID: purchaseIDVal,
		}
		purchaseDetail := usecases.NewPurchaseDetail(purchaseDetailModel)
		purchaseDetailXi, modelErr := usecases.Read(purchaseDetail, db)
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

		purchaseDetailModel := &model.PurchaseDetail{}

		err := context.BindJSON(&purchaseDetailModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(purchaseDetailModel.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		purchaseDetail := usecases.NewPurchaseDetail(purchaseDetailModel)
		modelErr := usecases.Create(purchaseDetail, db)
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

		purchaseDetailModel := &model.PurchaseDetail{}

		err = context.BindJSON(&purchaseDetailModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(purchaseDetailModel.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		purchaseDetailModel.ID = purchaseDetailIDVal

		purchaseDetail := usecases.NewPurchaseDetail(purchaseDetailModel)
		modelErr := usecases.Update(purchaseDetail, db)
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

		purchaseDetailModel := &model.PurchaseDetail{
			ID: purchaseDetailIDVal,
		}

		purchaseDetail := usecases.NewPurchaseDetail(purchaseDetailModel)
		modelErr := usecases.Delete(purchaseDetail, db)
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
