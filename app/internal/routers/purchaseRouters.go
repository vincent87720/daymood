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

		purchaseModel := &model.Purchase{}

		checkList := []Field{
			{Key: "id", Val: purchaseID},
		}
		err := checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		purchaseModel.ID, err = strconv.ParseInt(purchaseID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
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

		checkList := []Field{
			{Key: "Name", Val: purchaseModel.Name},
		}
		err = checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
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
		purchaseID := context.Param("id")

		purchaseModel := &model.Purchase{}

		err := context.BindJSON(&purchaseModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		checkList := []Field{
			{Key: "id", Val: purchaseID},
			{Key: "Name", Val: purchaseModel.Name},
		}
		err = checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		purchaseModel.ID, err = strconv.ParseInt(purchaseID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

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

		purchaseModel := &model.Purchase{}

		checkList := []Field{
			{Key: "id", Val: purchaseID},
		}
		err := checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		purchaseModel.ID, err = strconv.ParseInt(purchaseID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
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
