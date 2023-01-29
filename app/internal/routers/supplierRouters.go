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

func SetupSupplierRouters(router *gin.RouterGroup, db *sql.DB, s settings.Settings) (*gin.RouterGroup, error) {

	router.GET("/api/suppliers", GetSuppliersHandler(db))
	router.POST("/api/suppliers", PostSupplierHandler(db))
	router.PUT("/api/suppliers/:id", PutSupplierHandler(db))
	router.DELETE("/api/suppliers/:id", DeleteSupplierHandler(db))
	// router.GET("/suppliers/dumping", DumpFirmHandler(db, s))

	return router, nil
}

func GetSuppliersHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		supplierModel := &model.Supplier{}

		supplier := usecases.NewSupplier(supplierModel)
		supplierXi, modelErr := model.ReadAll(supplier, db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": supplierXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PostSupplierHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		supplierModel := &model.Supplier{}

		err := context.BindJSON(&supplierModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(supplierModel.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		supplier := usecases.NewSupplier(supplierModel)
		modelErr := usecases.Create(supplier, db)
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

func PutSupplierHandler(db *sql.DB) gin.HandlerFunc {
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

		supplierModel := &model.Supplier{}

		err = context.BindJSON(&supplierModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(supplierModel.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		supplierModel.ID = supplierIDVal

		supplier := usecases.NewSupplier(supplierModel)
		modelErr := usecases.Update(supplier, db)
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

func DeleteSupplierHandler(db *sql.DB) gin.HandlerFunc {
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

		supplierModel := &model.Supplier{
			ID: supplierIDVal,
		}

		supplier := usecases.NewSupplier(supplierModel)
		modelErr := usecases.Delete(supplier, db)
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

// func DumpFirmHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
// 	fn := func(context *gin.Context) {
// 		firmXi, err := model.GetAllFirm(db)
// 		if err != nil {
// 			fmt.Println(err)
// 			context.JSON(http.StatusBadRequest, gin.H{
// 				"status": "FAIL",
// 			})
// 			return
// 		}

// 		var prepareCSV [][]string

// 		prepareCSV = append(prepareCSV, []string{"廠商名稱", "廠商地址", "備註"})

// 		for _, v := range firmXi {
// 			tmpXi := []string{}
// 			tmpXi = append(tmpXi, v.Name, v.Address, v.Remark)
// 			prepareCSV = append(prepareCSV, tmpXi)
// 		}

// 		err = s.WriteCSV(s.GetExeFilePath()+"/firms.csv", prepareCSV)
// 		if err != nil {
// 			fmt.Println(err)
// 			context.JSON(http.StatusBadRequest, gin.H{
// 				"status": "FAIL",
// 			})
// 			return
// 		}

// 		context.FileAttachment(s.GetExeFilePath()+"/firms.csv", "firms.csv")
// 		return
// 	}

// 	return gin.HandlerFunc(fn)
// }
