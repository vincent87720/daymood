package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
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
		deliveryOrderXi, modelErr := model.GetAllDeliveryOrders(db)
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
		deliveryOrder := model.DeliveryOrder{
			ID: deliveryOrderIDVal,
		}

		deliveryOrderXi, modelErr := deliveryOrder.GetDeliveryOrder(db)
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

		deliveryOrder := model.DeliveryOrder{}

		err := context.BindJSON(&deliveryOrder)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		modelErr := deliveryOrder.Create(db)
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

		deliveryOrderForm := model.DeliveryOrder{}

		err = context.BindJSON(&deliveryOrderForm)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		deliveryOrder := deliveryOrderForm
		deliveryOrder.ID = supplierIDVal

		modelErr := deliveryOrder.Update(db)
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

		deliveryOrder := model.DeliveryOrder{
			ID: deliveryOrderIDVal,
		}

		modelErr := deliveryOrder.Delete(db)
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
