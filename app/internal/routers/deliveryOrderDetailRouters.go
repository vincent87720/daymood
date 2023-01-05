package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/settings"
)

func SetupDeliveryOrderDetailRouters(router *gin.Engine, db *sql.DB, s settings.Settings) (*gin.Engine, error) {

	router.GET("/deliveryOrderDetails", GetAllDeliveryOrderDetailsHandler(db))
	router.POST("/deliveryOrderDetails/multiple", PostDeliveryOrderDetailsHandler(db))
	router.GET("/deliveryOrders/:id/deliveryOrderDetails", GetDeliveryOrderDetailsHandler(db))
	router.POST("/deliveryOrderDetails", PostDeliveryOrderDetailHandler(db))
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

func PostDeliveryOrderDetailsHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		var deliveryOrderDetailXi []model.DeliveryOrderDetail

		err := context.BindJSON(&deliveryOrderDetailXi)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
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

func PostDeliveryOrderDetailHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		deliveryOrderDetail := model.DeliveryOrderDetail{}

		err := context.BindJSON(&deliveryOrderDetail)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		modelErr := deliveryOrderDetail.Create(db)
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
