package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/settings"
	usecases "github.com/vincent87720/daymood/app/internal/usecases"
)

func SetupDeliveryOrderDetailRouters(router *gin.RouterGroup, db *sql.DB, s settings.Settings) (*gin.RouterGroup, error) {

	router.GET("/api/deliveryOrderDetails", GetAllDeliveryOrderDetailsHandler(db))
	router.POST("/api/deliveryOrderDetails/multiple", PostDeliveryOrderDetailsHandler(db, s))
	router.GET("/api/deliveryOrders/:id/deliveryOrderDetails", GetDeliveryOrderDetailsHandler(db))
	router.POST("/api/deliveryOrderDetails", PostDeliveryOrderDetailHandler(db, s))
	router.PUT("/api/deliveryOrderDetails/:id", PutDeliveryOrderDetailHandler(db))
	router.DELETE("/api/deliveryOrderDetails/:id", DeleteDeliveryOrderDetailHandler(db))

	return router, nil
}

func GetAllDeliveryOrderDetailsHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		deliveryOrderDetailModel := &model.DeliveryOrderDetail{}

		deliveryOrderDetail := usecases.NewDeliveryOrderDetail(deliveryOrderDetailModel)

		deliveryOrderDetailXi, modelErr := usecases.ReadAll(deliveryOrderDetail, db)
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

		deliveryOrderDetail := usecases.NewDeliveryOrderDetail(&model.DeliveryOrderDetail{})
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

		deliveryOrderIDVal, err := strconv.ParseInt(deliveryOrderID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}
		deliveryOrderDetailModel := &model.DeliveryOrderDetail{
			DeliveryOrderID: deliveryOrderIDVal,
		}
		deliveryOrderDetail := usecases.NewDeliveryOrderDetail(deliveryOrderDetailModel)
		deliveryOrderDetailXi, modelErr := usecases.Read(deliveryOrderDetail, db)
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

		deliveryOrderDetailModel := &model.DeliveryOrderDetail{}

		err := context.BindJSON(&deliveryOrderDetailModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		deliveryOrderDetail := usecases.NewDeliveryOrderDetail(deliveryOrderDetailModel)

		modelErr := usecases.Create(deliveryOrderDetail, db)
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

		deliveryOrderDetailModel := &model.DeliveryOrderDetail{}

		err = context.BindJSON(&deliveryOrderDetailModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		deliveryOrderDetailModel.ID = deliveryOrderDetailIDVal

		deliveryOrderDetail := usecases.NewDeliveryOrderDetail(deliveryOrderDetailModel)
		modelErr := usecases.Update(deliveryOrderDetail, db)
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

		deliveryOrderDetailModel := &model.DeliveryOrderDetail{
			ID: deliveryOrderDetailIDVal,
		}

		deliveryOrderDetail := usecases.NewDeliveryOrderDetail(deliveryOrderDetailModel)
		modelErr := usecases.Delete(deliveryOrderDetail, db)
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
