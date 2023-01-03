package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/settings"
)

func SetupSupplierRouters(router *gin.Engine, db *sql.DB, s settings.Settings) (*gin.Engine, error) {

	router.GET("/suppliers", GetSuppliersHandler(db))
	router.POST("/suppliers", PostSupplierHandler(db))
	router.PUT("/suppliers/:id", PutSupplierHandler(db))
	router.DELETE("/suppliers/:id", DeleteSupplierHandler(db))
	// router.GET("/suppliers/dumping", DumpFirmHandler(db, s))

	return router, nil
}

func GetSuppliersHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		supplierXi, modelErr := model.GetAllSuppliers(db)
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

		supplier := model.Supplier{}

		err := context.BindJSON(&supplier)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(supplier.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		//DataStatus預設為使用中
		supplier.DataStatus = 1

		modelErr := supplier.Create(db)
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

		supplier := model.Supplier{}

		err = context.BindJSON(&supplier)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(supplier.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		supplier.ID = supplierIDVal

		modelErr := supplier.Update(db)
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

		supplier := model.Supplier{
			ID: supplierIDVal,
		}

		modelErr := supplier.Delete(db)
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
