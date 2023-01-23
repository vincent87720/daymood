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

func SetupDiscountRouters(router *gin.Engine, db *sql.DB, s settings.Settings) (*gin.Engine, error) {

	router.GET("/deliveryOrders/:id/discounts", GetDiscountsHandler(db))
	router.POST("/discounts", PostDiscountHandler(db, s))
	router.PUT("/discounts/:id", PutDiscountHandler(db))
	router.DELETE("/discounts/:id", DeleteDiscountHandler(db))

	return router, nil
}

func GetDiscountsHandler(db *sql.DB) gin.HandlerFunc {
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
		discountModel := &model.Discount{
			DeliveryOrderID: deliveryOrderIDVal,
		}

		discount := usecases.NewDiscount(discountModel)
		discountXi, modelErr := usecases.Read(discount, db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": discountXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PostDiscountHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		discountModel := &model.Discount{}

		err := context.BindJSON(&discountModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		discount := usecases.NewDiscount(discountModel)
		modelErr := usecases.Create(discount, db)
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

func PutDiscountHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		discountID := context.Param("id")

		if checkEmpty(discountID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		discountIDVal, err := strconv.ParseInt(discountID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		discountModel := &model.Discount{}

		err = context.BindJSON(&discountModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		discountModel.ID = discountIDVal

		discount := usecases.NewDiscount(discountModel)
		modelErr := usecases.Update(discount, db)
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

func DeleteDiscountHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		discountID := context.Param("id")

		if checkEmpty(discountID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		discountIDVal, err := strconv.ParseInt(discountID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		discountModel := &model.Discount{
			ID: discountIDVal,
		}

		discount := usecases.NewDiscount(discountModel)
		modelErr := usecases.Delete(discount, db)
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
