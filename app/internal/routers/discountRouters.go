package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	usecases "github.com/vincent87720/daymood/app/internal/usecases"
)

func SetupDiscountRouters(router *gin.RouterGroup, db *sql.DB) *gin.RouterGroup {

	router.GET("/api/deliveryOrders/:id/discounts", GetDiscountsHandler(db))
	router.POST("/api/discounts", PostDiscountHandler(db))
	router.PUT("/api/discounts/:id", PutDiscountHandler(db))
	router.DELETE("/api/discounts/:id", DeleteDiscountHandler(db))

	return router
}

func GetDiscountsHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		deliveryOrderID := context.Param("id")

		discountModel := &model.Discount{}

		checkList := []Field{
			{Key: "id", Val: deliveryOrderID},
		}
		err := checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		discountModel.ID, err = strconv.ParseInt(deliveryOrderID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
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

func PostDiscountHandler(db *sql.DB) gin.HandlerFunc {
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

		discountModel := &model.Discount{}

		err := context.BindJSON(&discountModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		checkList := []Field{
			{Key: "id", Val: discountID},
		}
		err = checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		discountModel.ID, err = strconv.ParseInt(discountID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

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

		discountModel := &model.Discount{}

		checkList := []Field{
			{Key: "id", Val: discountID},
		}
		err := checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		discountModel.ID, err = strconv.ParseInt(discountID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
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
