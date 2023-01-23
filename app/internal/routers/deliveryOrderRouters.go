package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	usecases "github.com/vincent87720/daymood/app/internal/usecases"
)

func SetupDeliveryOrderRouters(router *gin.Engine, db *sql.DB) (*gin.Engine, error) {
	router.GET("/deliveryOrders", GetDeliveryOrdersHandler(db))
	router.GET("/deliveryOrders/:id", GetDeliveryOrderHandler(db))
	router.POST("/deliveryOrders", PostDeliveryOrderHandler(db))
	router.PUT("/deliveryOrders/:id", PutDeliveryOrderHandler(db))
	router.DELETE("/deliveryOrders/:id", DeleteDeliveryOrderHandler(db))

	return router, nil
}
func GetDeliveryOrdersHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		deliveryOrderModel := &model.DeliveryOrder{}
		deliveryOrder := usecases.NewDeliveryOrder(deliveryOrderModel)
		deliveryOrderXi, modelErr := usecases.ReadAll(deliveryOrder, db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": deliveryOrderXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func GetDeliveryOrderHandler(db *sql.DB) gin.HandlerFunc {
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
		deliveryOrderModel := &model.DeliveryOrder{
			ID: deliveryOrderIDVal,
		}

		deliveryOrder := usecases.NewDeliveryOrder(deliveryOrderModel)
		deliveryOrderXi, modelErr := usecases.Read(deliveryOrder, db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": deliveryOrderXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PostDeliveryOrderHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		deliveryOrderModel := &model.DeliveryOrder{}

		err := context.BindJSON(&deliveryOrderModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		deliveryOrder := usecases.NewDeliveryOrder(deliveryOrderModel)
		modelErr := usecases.Create(deliveryOrder, db)
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

func PutDeliveryOrderHandler(db *sql.DB) gin.HandlerFunc {
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

		deliveryOrderModel := &model.DeliveryOrder{}

		err = context.BindJSON(&deliveryOrderModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		deliveryOrderModel.ID = supplierIDVal

		deliveryOrder := usecases.NewDeliveryOrder(deliveryOrderModel)
		modelErr := model.Update(deliveryOrder, db)
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

func DeleteDeliveryOrderHandler(db *sql.DB) gin.HandlerFunc {
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

		deliveryOrderModel := &model.DeliveryOrder{
			ID: deliveryOrderIDVal,
		}

		deliveryOrder := usecases.NewDeliveryOrder(deliveryOrderModel)
		modelErr := model.Delete(deliveryOrder, db)
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
