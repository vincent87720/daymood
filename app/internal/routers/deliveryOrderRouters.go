package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	usecases "github.com/vincent87720/daymood/app/internal/usecases"
)

func SetupDeliveryOrderRouters(router *gin.RouterGroup, db *sql.DB) (*gin.RouterGroup, error) {
	router.GET("/api/deliveryOrders", GetDeliveryOrdersHandler(db))
	router.GET("/api/deliveryOrders/:id", GetDeliveryOrderHandler(db))
	router.POST("/api/deliveryOrders", PostDeliveryOrderHandler(db))
	router.PUT("/api/deliveryOrders/:id", PutDeliveryOrderHandler(db))
	router.DELETE("/api/deliveryOrders/:id", DeleteDeliveryOrderHandler(db))

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

		deliveryOrderModel := &model.DeliveryOrder{}

		checkList := []Field{
			{Key: "id", Val: deliveryOrderID},
		}
		err := checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		deliveryOrderModel.ID, err = strconv.ParseInt(deliveryOrderID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
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
		deliveryOrderID := context.Param("id")

		deliveryOrderModel := &model.DeliveryOrder{}

		err := context.BindJSON(&deliveryOrderModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		checkList := []Field{
			{Key: "id", Val: deliveryOrderID},
		}
		err = checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		deliveryOrderModel.ID, err = strconv.ParseInt(deliveryOrderID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

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

		deliveryOrderModel := &model.DeliveryOrder{}

		checkList := []Field{
			{Key: "id", Val: deliveryOrderID},
		}
		err := checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		deliveryOrderModel.ID, err = strconv.ParseInt(deliveryOrderID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
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
